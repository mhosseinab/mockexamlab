# MockExamLab

A polyglot monorepo for an IELTS mock-exam platform — Go REST API backed by PostgreSQL and Firebase Auth, Next.js 15 frontend with App Router.

---

## Architecture

```mermaid
flowchart TD
    Browser["Browser / Mobile"]

    subgraph Frontend["frontend/ — Next.js 15"]
        MW["middleware.ts\n(cookie auth guard)"]
        Pages["App Router\n(public) + dashboard"]
        Stores["Zustand Stores\nauthStore · testsStore"]
        FBClient["Firebase Auth SDK\n(client)"]
    end

    subgraph API["services/api — Go 1.24 / Gin"]
        Router["Router\n/api/v1/*"]
        Auth["Auth Middleware\nFirebase token verify"]
        Handlers["Handlers\nTest · Section · Question\nUserTest · UserAnswer\nUser · Paddle"]
        DI["Wire DI"]
    end

    subgraph Infra["Infrastructure"]
        PG[("PostgreSQL 13")]
        FBAdmin["Firebase Admin SDK\n(server)"]
        Paddle["Paddle\nWebhook"]
        Swagger["Swagger UI\n/api/v1/swagger"]
    end

    Browser -->|"HTTPS requests"| MW
    MW -->|"passes or\nredirects /signin"| Pages
    Pages <-->|"Zustand actions"| Stores
    Stores -->|"Firebase signIn /\ncreateUser"| FBClient
    FBClient -->|"ID token"| Stores
    Stores -->|"POST /api/v1/user/auth\n+ Bearer token"| Router
    Stores -->|"REST calls\n+ Bearer token"| Router
    Router --> Auth
    Auth -->|"verifyIDToken()"| FBAdmin
    Auth --> Handlers
    Handlers <--> DI
    DI --> PG
    DI --> FBAdmin
    Paddle -->|"POST /webhook"| Router
```

---

## Repository Layout

```
mockexamlab/
├── services/
│   └── api/                  # Go 1.24 REST API
│       ├── internal/
│       │   ├── api/          # Gin handlers + router + wire
│       │   ├── cmd/          # Cobra CLI (runner, migrations)
│       │   ├── mapper/       # DTO ↔ model conversions
│       │   ├── models/       # GORM models + enums + DTOs
│       │   ├── repository/   # DB access layer
│       │   ├── service/      # Business logic
│       │   └── tools/        # env helpers
│       ├── docs/             # Swagger generated output
│       ├── docker-compose.yml
│       ├── DockerFile
│       └── go.mod
│
├── frontend/                 # Next.js 15 App Router
│   ├── app/
│   │   ├── (public)/         # Unauthenticated routes
│   │   │   ├── page.tsx      # Home landing
│   │   │   ├── signin/
│   │   │   └── signup/
│   │   ├── dashboard/        # Authenticated routes
│   │   │   ├── exams/
│   │   │   │   ├── page.tsx
│   │   │   │   └── exam/     # Active exam session
│   │   │   ├── profile/
│   │   │   ├── shop/
│   │   │   ├── social/
│   │   │   ├── stats/
│   │   │   ├── financial/
│   │   │   └── chat/
│   │   └── layout.tsx        # Root layout (global SCSS)
│   ├── components/           # Reusable UI components
│   ├── container/            # Page-level data containers
│   ├── store/                # Zustand stores
│   │   ├── authStore.ts
│   │   └── testsStore.ts
│   ├── lib/
│   │   └── firebase.ts       # Firebase app + Auth init
│   ├── middleware.ts          # Edge auth guard
│   └── styles/               # SCSS modules
│
├── Taskfile.yml              # Root orchestration
├── Taskfile.go.yml           # Go tasks (build / test / lint)
├── Taskfile.ts.yml           # TS tasks (dev / build / typecheck)
├── lefthook.yml              # Git hooks (fmt, lint, tidy)
├── pnpm-workspace.yaml
└── .github/workflows/ci.yml
```

---

## Quick Start

### Prerequisites

| Tool | Version |
|------|---------|
| Go | ≥ 1.24 |
| Node.js | ≥ 20 |
| pnpm | ≥ 9 |
| Task | ≥ 3 |
| Docker + Compose | any recent |

### Setup

```bash
# 1. Clone
git clone <repo-url> && cd mockexamlab

# 2. Install all deps + initialise go.work
task setup

# 3. Start the database
cd services/api && docker compose up -d app-db

# 4. Copy and fill env vars
cp services/api/.env.prod services/api/.env
# Edit DB_DSN, ADMIN_EMAIL, ADMIN_PASSWORD, Firebase credentials

# 5. Run the API (auto-migrates on start)
cd services/api && go run main.go serve

# 6. Run the frontend (separate terminal)
task ts:dev
```

### Useful Tasks

```bash
task doctor        # verify all tools present
task build:all     # build API + frontend
task test:all      # run all tests
task lint:all      # lint all code

task go:build      # GOWORK=off go build ./...
task go:test       # GOWORK=off go test ./... -race
task ts:dev        # pnpm --dir frontend dev
task ts:typecheck  # pnpm --dir frontend exec tsc --noEmit
```

---

## Environment Variables

### API (`services/api`)

| Variable | Required | Description |
|----------|----------|-------------|
| `DB_DSN` | yes | PostgreSQL DSN |
| `ADMIN_EMAIL` | no | Seeds admin Firebase user on start |
| `ADMIN_PASSWORD` | no | Seeds admin Firebase user on start |

### Frontend (`frontend/`)

Firebase config is embedded in `lib/firebase.ts`. No `.env` needed for local dev unless you override it.

---

## CI/CD

GitHub Actions (`.github/workflows/ci.yml`) runs two independent jobs using path filters:

```mermaid
flowchart LR
    Push["git push"] --> Detect["detect job\ndorny/paths-filter"]
    Detect -->|"frontend/**"| TS["typescript job\npnpm install → lint → tsc → jest"]
    Detect -->|"services/api/**"| GO["golang job\ngo test -race → golangci-lint"]
```

---

## Contributing

- Go: `gofmt` enforced on pre-commit via lefthook; `go mod tidy` enforced on pre-push
- TypeScript: `next lint` on pre-commit; `tsc --noEmit` on pre-push
- Commits are not required to build both services — CI skips unchanged paths
