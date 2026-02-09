# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**hearoom** — 모노레포 구조의 웹 애플리케이션. Go 백엔드 + 프론트엔드(미정).

## Architecture

```
backend/           # Go (Fiber v2) 서버 — 포트 8080
├── cmd/server/    # 엔트리포인트 (main.go)
├── internal/      # 비공개 패키지 (config, handlers, models)
└── pkg/           # 공개 패키지
frontend/          # 프론트엔드 (미설정)
```

## Build & Run

```bash
# 백엔드 실행
cd backend && go run cmd/server/main.go

# 백엔드 테스트
cd backend && go test ./...

# 단일 패키지 테스트
cd backend && go test ./internal/handlers/...

# Go 포맷팅
cd backend && gofmt -w .
```

## Commit Convention

**형식:** `[{노션아이디}] {타입}: {설명}`

- 브랜치명이 노션 태스크 ID (예: `TSK-17`)
- 타입: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`
- 예시: `[TSK-17] feat: 로그인 기능 추가`
- 기능별 커밋 분리를 위해 `/commit-by-feature` 커스텀 명령어 사용 가능

## Tech Stack

- **Backend:** Go 1.25, Fiber v2, UUID
- **Frontend:** TBD

## CI

- PR 오픈 시 변경 경로 기반 조건부 빌드 (`.github/workflows/pr-check.yml`)
- `backend/**` → Go 빌드 + 테스트
- `docs/**` → Redocly lint + 문서 빌드
- `frontend/**` → 미설정 (프레임워크 선택 후 추가)
