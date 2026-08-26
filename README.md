# HTTP Playground

Minimal playground for comparing how different frontend and backend technologies communicate over HTTP.

The goal is to implement the same HTTP contract using different clients and servers while keeping the application deliberately small.

## Architecture

The repository contains multiple interchangeable frontends and backends:

```text
http-playground/
├─ frontend/
│  ├─ vue/
│  └─ angular/
│
└─ backend/
   ├─ node-express/
   ├─ go-nethttp/
   ├─ go-gin/
   └─ go-fiber/
```

Every frontend talks to every backend through the same HTTP API:

```http
GET /api/message
```

```json
{
  "message": "Hello from backend"
}
```

```http
POST /api/message
Content-Type: application/json
```

```json
{
  "message": "Hello from frontend"
}
```

The POST response is `{ "received": "Hello from frontend" }`. This makes the frontends and backends interchangeable.

---

## Running the playground

Start one frontend and one backend in separate terminals:

```text
Terminal 1 → one frontend
Terminal 2 → one backend
```

All frontends use:

```text
http://localhost:3000
```

All backends use:

```text
http://localhost:8080
```

Only one frontend and one backend can run at a time because they share ports. You can switch backends without changing the frontend.

---

# Frontends

## Vue

Location:

```text
frontend/vue
```

Install dependencies:

```bash
pnpm install
```

Start:

```bash
pnpm dev
```

Runs at:

```text
http://localhost:3000
```

Set the backend URL in `frontend/vue/.env` to:

```env
VITE_API_URL=http://localhost:8080
```

### HTTP clients

Vue includes two implementations of the same API calls: one using the native Fetch API and another using Axios.

---

## Angular

Location:

```text
frontend/angular
```

Install dependencies:

```bash
pnpm install
```

Start:

```bash
pnpm start
```

Runs at:

```text
http://localhost:3000
```

The backend URL is configured using Angular environments:

```text
src/environments/environment.ts
```

```ts
export const environment = {
  apiUrl: 'http://localhost:8080',
}
```

### HTTP client

Angular uses its built-in `HttpClient` with the same API contract.

---

# Backends

All backends expose the same API on:

```text
http://localhost:8080
```

Their `.env` files contain:

```env
PORT=8080
FRONTEND_ORIGIN=http://localhost:3000
```

---

## Node + Express

Location:

```text
backend/node-express
```

Install dependencies:

```bash
pnpm install
```

Start:

```bash
pnpm dev
```

---

## Go `net/http`

Location:

```text
backend/go-nethttp
```

Start:

```bash
go run .
```

---

## Go + Gin

Location:

```text
backend/go-gin
```

Start:

```bash
go run .
```

---

## Go + Fiber

Location:

```text
backend/go-fiber
```

Start:

```bash
go run .
```

All four backends implement the same GET and POST endpoints. Their differences are limited to the HTTP library or framework used.
