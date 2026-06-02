# MockExamLab API

Go 1.24 REST API for the MockExamLab platform. Built with Gin, GORM, Firebase Auth, and PostgreSQL.

---

## Layer Architecture

```mermaid
flowchart TD
    Client["HTTP Client"] -->|"Bearer token"| Router

    subgraph Gin["Gin Router"]
        Router["router.go\nCORS · routes · swagger"]
        MW["authMiddleware\nFirebase verifyIDToken"]
    end

    Router --> MW
    MW -->|"injects UserClaims\ninto context"| H

    subgraph Handlers["internal/api — handlers"]
        H["TestAPI · SectionAPI\nQuestionGroupAPI · QuestionAPI\nUserTestAPI · UserAnswerAPI\nUserAPI · PaddleAPI"]
    end

    subgraph Services["internal/service — business logic"]
        S["TestService · SectionService\nQuestionGroupService · QuestionService\nUserTestService · UserAnswerService\nUserService · PaddleService"]
    end

    subgraph Repos["internal/repository — data access"]
        R["TestRepository · SectionRepository\nQuestionGroupRepository\nQuestionRepository\nUserTestRepository · UserAnswerRepository\nUserRepository · FirebaseRepository\nSubscriptionEventRepository · TestStatRepository"]
    end

    H -->|"calls"| S
    S -->|"calls"| R
    R -->|"GORM"| PG[("PostgreSQL")]
    R -->|"Firebase Admin SDK"| FB["Firebase Auth"]

    Wire["wire_gen.go\n(compile-time DI)"] -.->|"wires"| H
```

---

## Data Model

```mermaid
erDiagram
    User {
        uuid ID PK
        string FirebaseId
        string DisplayName
        string Email
        int64 LastLogin
        string Role
        string InvitationCode
        int64 SubscriptionExpirationMs
        int SubscriptionState
    }

    Test {
        uuid ID PK
        string Name
        string Description
        int64 TestDate
        string Module
        int TestTime
        uuid CreatorId
        jsonb ExtraProperty
    }

    Section {
        uuid ID PK
        uuid TestID FK
        string Title
        string ComponentType
        int Time
        int QStart
        int QEnd
    }

    QuestionGroup {
        uuid ID PK
        uuid SectionID FK
        string Title
        string QuestionType
        int QStart
        int QEnd
    }

    Question {
        uuid ID PK
        uuid GroupID FK
        string CorrectAnswer
        string Component
        int QNumber
    }

    UserTest {
        uuid ID PK
        uuid TestID FK
        uuid UserID FK
        int64 TestDate
        int Duration
        float32 OverallScore
        float32 ListeningScore
        float32 ReadingScore
        float32 WritingScore
        float32 SpeakingScore
    }

    UserAnswer {
        uuid ID PK
        uuid UserTestID FK
        uuid QuestionID FK
        uuid MarkerID FK
        string Answer
        float32 MarkerScore
        string MarkerComment
    }

    QuestionStatistic {
        uuid ID PK
        uuid QuestionID FK
        int TotalAnswer
        int TotalCorrectAnswer
    }

    TestStat {
        uuid ID PK
        uuid TestId FK
        jsonb Listening
        jsonb Reading
        jsonb Writing
        jsonb Speaking
        jsonb Overall
    }

    SubscriptionEvent {
        uuid ID PK
        uuid UserID FK
        string AlertName
        string SubscriptionId
        string Status
        string Currency
    }

    User ||--o{ UserTest : "takes"
    Test ||--o{ UserTest : "attempted via"
    Test ||--o{ Section : "has"
    Test ||--|| TestStat : "has"
    Section ||--o{ QuestionGroup : "has"
    QuestionGroup ||--o{ Question : "has"
    Question ||--|| QuestionStatistic : "has"
    UserTest ||--o{ UserAnswer : "contains"
    Question ||--o{ UserAnswer : "answered by"
    User ||--o{ SubscriptionEvent : "has"
```

---

## Request Auth Flow

```mermaid
sequenceDiagram
    participant C as Client
    participant G as Gin Router
    participant MW as authMiddleware
    participant FB as Firebase Admin
    participant H as Handler

    C->>G: POST /api/v1/user-test\nAuthorization: Bearer <id-token>
    G->>MW: route matched
    MW->>FB: VerifyIDToken(token)
    alt token valid
        FB-->>MW: DecodedToken (uid, email, role, ...)
        MW->>MW: build UserClaims
        MW->>G: set "user" in context
        G->>H: handle request
        H->>H: read UserClaims from ctx
        H-->>C: 200 OK + body
    else token invalid / expired
        FB-->>MW: error
        MW-->>C: 401 Unauthorized
    end
```

---

## API Routes

Base path: `/api/v1` — all routes except `/user/auth` and `/webhook` require a valid Firebase ID token.

### Tests

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/test/all` | List all tests |
| `GET` | `/test/:id` | Get test by ID |
| `GET` | `/test/full/:id` | Get test with all nested data (sections → groups → questions) |
| `POST` | `/test` | Create test |
| `PUT` | `/test` | Update test |
| `DELETE` | `/test/:id` | Delete test |

### Sections

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/section/all` | List all sections |
| `GET` | `/section/:id` | Get section by ID |
| `POST` | `/section` | Create section |
| `PUT` | `/section` | Update section |
| `DELETE` | `/section/:id` | Delete section |

### Question Groups

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/question-group/all` | List groups by section ID |
| `GET` | `/question-group/:id` | Get group by ID |
| `POST` | `/question-group` | Create group |
| `PUT` | `/question-group` | Update group |
| `DELETE` | `/question-group/:id` | Delete group |

### Questions

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/question/all` | List questions by group ID |
| `GET` | `/question/test/all` | List questions by test ID |
| `GET` | `/question/:id` | Get question by ID |
| `POST` | `/question` | Create question |
| `PUT` | `/question` | Update question |
| `DELETE` | `/question/:id` | Delete question |

### User Tests (Attempts)

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/user-test/all` | List user's test attempts |
| `GET` | `/user-test/:id` | Get attempt by ID |
| `POST` | `/user-test` | Start a new test attempt |
| `PUT` | `/user-test/submit` | Submit completed test |
| `DELETE` | `/user-test/:id` | Delete attempt |

### User Answers

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/user-answer/all` | List answers by user-test ID |
| `GET` | `/user-answer/:id` | Get answer by ID |
| `POST` | `/user-answer` | Submit single answer |
| `POST` | `/user-answer/speaking` | Submit speaking answer (audio) |
| `POST` | `/user-answer/batch` | Batch submit answers |
| `PUT` | `/user-answer/answer` | Update answer |
| `PUT` | `/user-answer/answer/speaking` | Update speaking answer |
| `PUT` | `/user-answer/marker` | Marker scores an answer |
| `DELETE` | `/user-answer/:id` | Delete answer |

### Users

| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/user/auth` | Login or sign-up — no auth middleware |
| `GET` | `/user` | List users with filter + pagination |

### Public

| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/webhook` | Paddle subscription webhook |
| `GET` | `/swagger/*any` | Swagger UI |

---

## Dependency Injection

Wire generates the full constructor chain at compile time (`wire_gen.go`). Each API group is wired independently — no global state.

```mermaid
flowchart LR
    DB[("*gorm.DB")] --> UR[UserRepo]
    DB --> TR[TestRepo]
    DB --> SR[SectionRepo]
    DB --> QGR[QuestionGroupRepo]
    DB --> QR[QuestionRepo]
    DB --> UTR[UserTestRepo]
    DB --> UAR[UserAnswerRepo]
    DB --> TSR[TestStatRepo]
    DB --> SER[SubscriptionEventRepo]
    FB["*auth.Client"] --> FBR[FirebaseRepo]

    UR --> US[UserService]
    FBR --> US
    TR --> TS[TestService]
    SR --> SS[SectionService]
    QGR --> QGS[QuestionGroupService]
    QR --> QS[QuestionService]
    UTR --> UTS[UserTestService]
    QR --> UTS
    TSR --> UTS
    UR --> UTS
    UAR --> UAS[UserAnswerService]
    SER --> PS[PaddleService]
    UR --> PS

    US --> UserAPI
    TS --> TestAPI
    SS --> SectionAPI
    QGS --> QuestionGroupAPI
    QS --> QuestionAPI
    UTS --> UserTestAPI
    UAS --> UserAnswerAPI
    PS --> PaddleAPI
```

---

## Running Locally

### Database only (Docker)

```bash
docker compose up -d app-db
```

### API directly

```bash
export DB_DSN="host=localhost user=postgres password=... dbname=MockExam port=5432 sslmode=disable"
export ADMIN_EMAIL="admin@example.com"   # optional — seeds admin user
export ADMIN_PASSWORD="secret"           # optional

go run main.go serve
# API on :8080
```

### Full stack (Docker Compose)

```bash
docker compose up
# copies vendor/ into container
```

---

## Swagger Docs

Available at `http://localhost:8080/api/v1/swagger/index.html`.

Regenerate after changing handler annotations:

```bash
# from services/api/
swag init -g main.go
```

---

## Conventions

- All primary keys are UUIDs generated in `BeforeCreate` GORM hooks
- Soft deletes via GORM `DeletedAt` on all models
- `CreatedAt` / `UpdatedAt` are Unix milliseconds (`int64`), not `time.Time`
- `ExtraProperty` and score arrays are stored as PostgreSQL `JSONB` via `jackc/pgtype`
- The auth middleware validates the raw Firebase ID token; the `Authorization` header value is the token itself
