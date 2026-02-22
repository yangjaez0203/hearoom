package helper

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/yangjaez0203/hearoom/backend/internal/config"
)

var UUIDRegexp = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)

func TestConfig() *config.Config {
	return &config.Config{
		JWTSecret:  "test-secret-key",
		JWTExpiry:  time.Hour,
		ServerPort: "8080",
	}
}

func DoRequest(t *testing.T, app *fiber.App, req *http.Request) (int, []byte) {
	t.Helper()
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("app.Test failed: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response body: %v", err)
	}
	return resp.StatusCode, body
}

func ParseJSON(t *testing.T, body []byte) map[string]interface{} {
	t.Helper()
	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		t.Fatalf("JSON parse error: %v\nbody: %s", err, body)
	}
	return m
}
