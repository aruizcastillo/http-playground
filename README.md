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

The POST response is:

```json
{
  "received": "Hello from frontend"
}
```

This makes every frontend and backend interchangeable.

---

## Requirements

Install the following tools before running the playground:

* Node.js
* pnpm
* Go

Verify the installation:

```bash
node --version
pnpm --version
go version
```

---

## Setup

Install all Node.js dependencies from the repository root:

```bash
pnpm install
```

The pnpm workspace includes:

```text
frontend/vue
frontend/angular
backend/node-express
```

The Go backends are grouped using `go.work`, so they can also be run directly from the repository root.

Go dependencies are downloaded automatically when running a Go backend for the first time.

---

## Running the playground

Start one frontend and one backend in separate terminals.

```text
Terminal 1 → one frontend
Terminal 2 → one backend
```

All frontends run at:

```text
http://localhost:3000
```

All backends run at:

```text
http://localhost:8080
```

Only one frontend and one backend can run at a time because implementations in the same group share the same port.

You can switch frontend or backend implementations without changing the HTTP contract.

### Frontends

Vue:

```bash
pnpm dev:vue
```

Angular:

```bash
pnpm dev:angular
```

### Backends

Node + Express:

```bash
pnpm dev:express
```

Go `net/http`:

```bash
pnpm dev:go-nethttp
```

Go + Gin:

```bash
pnpm dev:go-gin
```

Go + Fiber:

```bash
pnpm dev:go-fiber
```

All commands are executed from the repository root.

---

# Frontends

## Vue

Location:

```text
frontend/vue
```

Runs at:

```text
http://localhost:3000
```

The backend URL is configured in:

```text
frontend/vue/.env
```

```env
VITE_API_URL=http://localhost:8080
```

### HTTP clients

Vue includes two implementations of the same API calls:

* Native Fetch API
* Axios

Both implementations use the same HTTP contract.

---

## Angular

Location:

```text
frontend/angular
```

Runs at:

```text
http://localhost:3000
```

The backend URL is configured using Angular environments:

```text
frontend/angular/src/environments/environment.ts
```

```ts
export const environment = {
  apiUrl: 'http://localhost:8080',
}
```

### HTTP client

Angular uses its built-in `HttpClient` with the same HTTP contract.

---

# Backends

All backends expose the same API at:

```text
http://localhost:8080
```

Their `.env` files contain local development configuration:

```env
PORT=8080
FRONTEND_ORIGIN=http://localhost:3000
```

These `.env` files are intentionally tracked because they only contain local playground configuration and no secrets.

---

## Node + Express

Location:

```text
backend/node-express
```

Uses Express to implement the shared HTTP API.

---

## Go `net/http`

Location:

```text
backend/go-nethttp
```

Uses Go's standard `net/http` package without an external HTTP framework.

---

## Go + Gin

Location:

```text
backend/go-gin
```

Uses Gin to implement the shared HTTP API.

---

## Go + Fiber

Location:

```text
backend/go-fiber
```

Uses Fiber to implement the shared HTTP API.

---

All four backends implement the same GET and POST endpoints.

Their differences are intentionally limited to the HTTP library or framework used, making the implementations easy to compare directly.
