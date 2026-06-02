# MockExamLab Frontend

Next.js 15 App Router frontend for the MockExamLab platform. TypeScript, Zustand state management, Firebase Auth, and SCSS modules.

---

## Page Structure

```mermaid
flowchart TD
    Root["app/layout.tsx\n(global SCSS · Providers · ToastContainer)"]

    Root --> Public["app/(public)/\nno extra layout"]
    Root --> Dashboard["app/dashboard/\nlayout.tsx → DashboardLayout"]

    Public --> Home["page.tsx\n/"]
    Public --> SignIn["signin/page.tsx\n/signin"]
    Public --> SignUp["signup/page.tsx\n/signup"]

    Dashboard --> DashPage["page.tsx\n/dashboard\n(Profile)"]
    Dashboard --> Exams["exams/page.tsx\n/dashboard/exams"]
    Dashboard --> Exam["exams/exam/page.tsx\n/dashboard/exams/exam\n?section=&level="]
    Dashboard --> Chat["chat/page.tsx\n/dashboard/chat"]
    Dashboard --> Financial["financial/page.tsx\n/dashboard/financial"]
    Dashboard --> Stats["stats/page.tsx\n/dashboard/stats"]
    Dashboard --> Social["social/page.tsx\n/dashboard/social"]
    Dashboard --> Post["social/post/[id]/page.tsx\n/dashboard/social/post/:id"]
    Dashboard --> Shop["shop/page.tsx\n/dashboard/shop"]
    Dashboard --> Books["shop/books/page.tsx\n/dashboard/shop/books"]
    Dashboard --> ProfileExams["profile/exams/page.tsx\n/dashboard/profile/exams"]
    Dashboard --> Inventory["profile/inventory/page.tsx\n/dashboard/profile/inventory"]

    Exam -->|"?section=reading"| Reading["Reading.tsx"]
    Exam -->|"?section=listening"| Listening["Listening.tsx"]
    Exam -->|"?section=writing"| Writing["Writing.tsx"]
    Exam -->|"?section=speaking"| Speaking["Speaking.tsx"]
```

---

## Auth Flow

```mermaid
sequenceDiagram
    participant U as User
    participant Page as /signin page
    participant AS as authStore (Zustand)
    participant FB as Firebase Auth
    participant API as Go API /user/auth
    participant MW as middleware.ts (Edge)
    participant D as /dashboard/exams

    U->>Page: submit email + password
    Page->>AS: login(email, password)
    AS->>FB: signInWithEmailAndPassword()
    FB-->>AS: credential + ID token
    AS->>API: POST /api/v1/user/auth\nAuthorization: <id-token>
    API-->>AS: { userId }
    AS->>AS: set token, userId, email, userName\nsetCookie("token") · setCookie("userId")
    AS-->>Page: token updated in store
    Page->>Page: useEffect sees token → router.push("/dashboard/exams")
    Note over MW: on any subsequent navigation
    MW->>MW: read "token" cookie
    MW-->>D: pass through (authenticated)
```

---

## State Management

Two Zustand stores handle all application state. Both run entirely client-side.

```mermaid
flowchart LR
    subgraph authStore["store/authStore.ts\n(persisted to localStorage)"]
        direction TB
        AS_State["email · token · userName · userId\nisLoading"]
        AS_Actions["login(email, password)\nsignUp(name, family, email, password)\nreset()"]
    end

    subgraph testsStore["store/testsStore.ts"]
        direction TB
        TS_State["tests · createdTest · test"]
        TS_Actions["getAllExams(token)\ncreateTest(testId, type, date, onSuccess?)\ngetTest(testId)\nanswerTest(answer, questionId, userTestId)"]
    end

    FB["Firebase Auth SDK"]
    API["Go REST API\nmel-api.go7.ir/api/v1"]
    Cookies["Browser Cookies\ntoken · userId · testId"]

    AS_Actions -->|"signIn / createUser"| FB
    FB -->|"ID token"| AS_Actions
    AS_Actions -->|"POST /user/auth"| API
    AS_Actions -->|"setCookie"| Cookies
    TS_Actions -->|"GET/POST with Bearer token"| API
    TS_Actions -->|"setCookie testId"| Cookies
    Cookies -->|"read by middleware.ts"| MW["Edge Middleware\nredirects if no token"]
```

### Store Persistence

`authStore` uses `zustand/middleware persist` — state survives page refresh via `localStorage` key `auth-storage`. `testsStore` is session-only (no persist middleware).

---

## Component Hierarchy

```mermaid
flowchart TD
    subgraph PublicUI["Public pages"]
        HomeLayout --> Navbar
        HomeLayout --> Footer
        HomeLanding --> HomeHeader & HomeExam & HomeExamDetails & HomeState & HomeWorks
    end

    subgraph DashboardUI["Dashboard pages"]
        DashboardLayout --> Sidebar
        DashboardLayout --> MobileMenu
        DashboardLayout --> Children["page content"]
    end

    subgraph ExamUI["Active exam (/dashboard/exams/exam)"]
        QuestionContainer --> TestHeader
        QuestionContainer --> ExamsFooter
        QuestionContainer --> ExamSection["Reading | Listening\nWriting | Speaking"]
        ExamSection --> Instruction["intro screen"]
        ExamSection --> Questions["question components\n(ReadingQuestion · WritingQuestion\nCheckQuestion · NoteCompletion\nSpeakingQuestion)"]
    end
```

---

## Exam Session Flow

```mermaid
stateDiagram-v2
    [*] --> ExamsPage: /dashboard/exams
    ExamsPage --> ExamResults: getAllExams() loads test list
    ExamResults --> Creating: user clicks a test → getTest(id)
    Creating --> IntroScreen: navigate to\n/exams/exam?section=reading&level=intro
    IntroScreen --> Questions: click "Start test"\n?level=questions
    Questions --> NextSection: click "Finish"\n→ next section intro
    NextSection --> IntroScreen
    Questions --> Dashboard: last section finished\n→ /dashboard/exams

    note right of Questions
        answerTest() called on each answer
        useSearchParams drives which
        component renders
    end note
```

---

## Directory Reference

```
frontend/
├── app/
│   ├── layout.tsx              # Root layout — global SCSS, Providers
│   ├── (public)/               # Route group, no extra layout
│   │   ├── page.tsx            # Home landing
│   │   ├── signin/page.tsx     # Firebase sign-in form
│   │   └── signup/page.tsx     # Firebase sign-up form
│   └── dashboard/
│       ├── layout.tsx          # Empty (DashboardLayout lives in container)
│       ├── page.tsx            # Profile
│       ├── exams/
│       │   ├── page.tsx        # Exam list (calls getAllExams)
│       │   └── exam/page.tsx   # Active session (Suspense + useSearchParams)
│       ├── chat/page.tsx
│       ├── financial/page.tsx
│       ├── profile/
│       │   ├── exams/page.tsx
│       │   └── inventory/page.tsx
│       ├── shop/
│       │   ├── page.tsx
│       │   └── books/page.tsx
│       ├── social/
│       │   ├── page.tsx
│       │   └── post/[id]/page.tsx
│       └── stats/page.tsx
│
├── components/
│   ├── Providers.tsx           # 'use client' — ToastContainer
│   ├── ui/                     # Navbar, Footer, MobileMenu, Spinner ...
│   ├── dashboard/              # Sidebar, DashboardLayout, exam/profile/shop/social/stats
│   ├── exams/                  # Reading, Listening, Writing, Speaking + sub-components
│   └── home/                   # HomeHeader, HomeExam, HomeExamDetails ...
│
├── container/                  # Page-level containers (data fetching + layout wiring)
│   ├── ExamContainer.tsx
│   ├── HomeLanding.tsx
│   ├── QuessionContainer.tsx   # Suspense wrapper for useSearchParams
│   └── dashboard/
│       ├── ProfileContainer.tsx
│       ├── FinancialContainer.tsx
│       ├── ProfileInventoryContainer.tsx
│       ├── ShopContainer.tsx
│       ├── SocialContainer.tsx
│       └── StatsContainer.tsx
│
├── store/
│   ├── authStore.ts            # Auth state (persisted)
│   └── testsStore.ts           # Exam/test state
│
├── lib/
│   └── firebase.ts             # Firebase app + getAuth()
│
├── middleware.ts               # Edge middleware — reads "token" cookie
├── next.config.ts              # sassOptions only
└── styles/                     # SCSS modules per page/feature
```

---

## Running Locally

```bash
# from repo root
pnpm install          # installs all workspace deps
pnpm --dir frontend dev       # or: task ts:dev
```

App runs at `http://localhost:3000`.

### Available Scripts

```bash
pnpm --dir frontend dev          # dev server (Turbopack)
pnpm --dir frontend build        # production build
pnpm --dir frontend lint         # ESLint
pnpm --dir frontend exec tsc --noEmit   # type check
pnpm --dir frontend test         # Jest
```

---

## Edge Middleware

`middleware.ts` runs on every request before rendering. It reads the `token` cookie and redirects unauthenticated users to `/signin` for any non-public path.

```
Public paths:  /  /signin  /signup
Protected:     everything else (except _next/static, images, icons)
```

The matcher excludes static assets via a regex to avoid unnecessary edge invocations.
