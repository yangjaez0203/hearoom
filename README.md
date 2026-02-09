# hearoom

우리 함께, 지금 여기서, 같이 들어요.

플레이리스트를 같이 만들고, 같은 곡을 동시에 들으며 의견을 나눌 수 있는 서비스.

## Core Features

### 함께 듣기
- 실시간 동기화 스트리밍 (같은 지점을 같이 들음)
- 이전곡/다음곡 동기화
- 백그라운드 재생

### 공유 플레이리스트
- 함께 플레이리스트 생성 및 편집
- 간편한 초대 플로우

### 소셜
- 노래 추천 및 공유
- 노래별 코멘트 및 이모지 리액션

### 커스터마이징
- 플레이리스트 썸네일로 직접 찍은 현장 사진 업로드

### 비회원 우선 설계
- 회원가입 없이 서비스 이용 가능
- 기록이 필요한 시점에 가입 시 기존 데이터 자동 연계

## Tech Stack

| Layer | Stack |
|-------|-------|
| Backend | Go, Fiber v2 |
| Frontend | TBD |

## Getting Started

```bash
# 백엔드 실행
cd backend && go run cmd/server/main.go

# Health check
curl http://localhost:8080/health

# API 문서 미리보기 (포트 4000)
npx @redocly/cli preview -p 4000
```
