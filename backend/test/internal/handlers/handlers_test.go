package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yangjaez0203/hearoom/backend/internal/auth"
	"github.com/yangjaez0203/hearoom/backend/internal/handlers"
	"github.com/yangjaez0203/hearoom/backend/internal/middleware"
	"github.com/yangjaez0203/hearoom/backend/internal/models"
	"github.com/yangjaez0203/hearoom/backend/test/helper"
)

func setupApp() *fiber.App {
	cfg := helper.TestConfig()
	app := fiber.New()

	app.Use(middleware.Identity(cfg))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
	app.Post("/auth/token", handlers.CreateToken(cfg))
	app.Get("/me", handlers.GetMe)

	return app
}

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

func TestHealthEndpoint(t *testing.T) {
	app := setupApp()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	status, body := helper.DoRequest(t, app, req)
	if status != 200 {
		t.Fatalf("expected 200, got %d", status)
	}

	m := helper.ParseJSON(t, body)
	if m["status"] != "ok" {
		t.Fatalf("expected status=ok, got %v", m["status"])
	}
}

func TestCreateToken(t *testing.T) {
	t.Run("returns 200 with token and user", func(t *testing.T) {
		app := setupApp()
		req := httptest.NewRequest(http.MethodPost, "/auth/token", nil)

		status, body := helper.DoRequest(t, app, req)
		if status != 200 {
			t.Fatalf("expected 200, got %d", status)
		}

		m := helper.ParseJSON(t, body)

		token, ok := m["token"].(string)
		if !ok || token == "" {
			t.Fatal("token is missing or empty")
		}

		user, ok := m["user"].(map[string]interface{})
		if !ok {
			t.Fatal("user is missing")
		}

		id, _ := user["id"].(string)
		if !helper.UUIDRegexp.MatchString(id) {
			t.Fatalf("user.id is not a valid UUID: %q", id)
		}

		anon, ok := user["anonymous"].(bool)
		if !ok || !anon {
			t.Fatalf("expected user.anonymous=true, got %v", user["anonymous"])
		}
	})

	t.Run("returned token is a valid JWT", func(t *testing.T) {
		cfg := helper.TestConfig()
		app := setupApp()
		req := httptest.NewRequest(http.MethodPost, "/auth/token", nil)

		_, body := helper.DoRequest(t, app, req)
		m := helper.ParseJSON(t, body)

		token := m["token"].(string)
		user := m["user"].(map[string]interface{})

		validated, err := auth.ValidateToken(cfg.JWTSecret, token)
		if err != nil {
			t.Fatalf("ValidateToken failed: %v", err)
		}

		if validated.ID != user["id"].(string) {
			t.Fatalf("ID mismatch: validated=%q, response=%q", validated.ID, user["id"])
		}
		if validated.Username != user["username"].(string) {
			t.Fatalf("Username mismatch: validated=%q, response=%q", validated.Username, user["username"])
		}
		if validated.Anonymous != user["anonymous"].(bool) {
			t.Fatalf("Anonymous mismatch: validated=%v, response=%v", validated.Anonymous, user["anonymous"])
		}
	})
}

func TestGetMe(t *testing.T) {
	t.Run("200 with valid token", func(t *testing.T) {
		app := setupApp()

		// Issue a token first.
		tokenReq := httptest.NewRequest(http.MethodPost, "/auth/token", nil)
		_, tokenBody := helper.DoRequest(t, app, tokenReq)
		tokenResp := helper.ParseJSON(t, tokenBody)
		token := tokenResp["token"].(string)

		// GET /me with the token.
		meReq := httptest.NewRequest(http.MethodGet, "/me", nil)
		meReq.Header.Set("Authorization", "Bearer "+token)

		status, meBody := helper.DoRequest(t, app, meReq)
		if status != 200 {
			t.Fatalf("expected 200, got %d", status)
		}

		user := helper.ParseJSON(t, meBody)
		if _, ok := user["id"].(string); !ok {
			t.Fatal("user.id missing from /me response")
		}
	})

	t.Run("401 without authorization header", func(t *testing.T) {
		app := setupApp()
		req := httptest.NewRequest(http.MethodGet, "/me", nil)

		status, _ := helper.DoRequest(t, app, req)
		if status != 401 {
			t.Fatalf("expected 401, got %d", status)
		}
	})

	t.Run("401 with invalid token", func(t *testing.T) {
		app := setupApp()
		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		req.Header.Set("Authorization", "Bearer invalid.token.here")

		status, _ := helper.DoRequest(t, app, req)
		if status != 401 {
			t.Fatalf("expected 401, got %d", status)
		}
	})

	t.Run("401 with expired token", func(t *testing.T) {
		cfg := helper.TestConfig()
		app := setupApp()

		user := &models.User{ID: "test-id", Username: "test", Anonymous: true}
		token, err := auth.GenerateToken(cfg.JWTSecret, -1*time.Hour, user)
		if err != nil {
			t.Fatalf("GenerateToken failed: %v", err)
		}

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		status, _ := helper.DoRequest(t, app, req)
		if status != 401 {
			t.Fatalf("expected 401, got %d", status)
		}
	})

	t.Run("401 with Basic auth scheme", func(t *testing.T) {
		app := setupApp()
		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		req.Header.Set("Authorization", "Basic dXNlcjpwYXNz")

		status, _ := helper.DoRequest(t, app, req)
		if status != 401 {
			t.Fatalf("expected 401, got %d", status)
		}
	})

	t.Run("401 with wrong signing secret", func(t *testing.T) {
		app := setupApp()

		user := &models.User{ID: "test-id", Username: "test", Anonymous: true}
		token, err := auth.GenerateToken("wrong-secret", time.Hour, user)
		if err != nil {
			t.Fatalf("GenerateToken failed: %v", err)
		}

		req := httptest.NewRequest(http.MethodGet, "/me", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		status, _ := helper.DoRequest(t, app, req)
		if status != 401 {
			t.Fatalf("expected 401, got %d", status)
		}
	})
}

func TestE2EFlow(t *testing.T) {
	app := setupApp()

	// Step 1: POST /auth/token → get token + user
	tokenReq := httptest.NewRequest(http.MethodPost, "/auth/token", nil)
	status, tokenBody := helper.DoRequest(t, app, tokenReq)
	if status != 200 {
		t.Fatalf("POST /auth/token: expected 200, got %d", status)
	}

	tokenResp := helper.ParseJSON(t, tokenBody)
	token := tokenResp["token"].(string)
	issuedUser := tokenResp["user"].(map[string]interface{})

	// Step 2: GET /me with Bearer token
	meReq := httptest.NewRequest(http.MethodGet, "/me", nil)
	meReq.Header.Set("Authorization", "Bearer "+token)

	status, meBody := helper.DoRequest(t, app, meReq)
	if status != 200 {
		t.Fatalf("GET /me: expected 200, got %d", status)
	}

	meUser := helper.ParseJSON(t, meBody)

	// Step 3: Verify issued user == /me user
	if meUser["id"] != issuedUser["id"] {
		t.Fatalf("id mismatch: issued=%v, me=%v", issuedUser["id"], meUser["id"])
	}
	if meUser["username"] != issuedUser["username"] {
		t.Fatalf("username mismatch: issued=%v, me=%v", issuedUser["username"], meUser["username"])
	}
	if meUser["anonymous"] != issuedUser["anonymous"] {
		t.Fatalf("anonymous mismatch: issued=%v, me=%v", issuedUser["anonymous"], meUser["anonymous"])
	}
}
