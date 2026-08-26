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

Every frontend talks to every backend using the same HTTP API:

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

Response:

```json
{
  "received": "Hello from frontend"
}
```

The important idea is:

```text
Frontend
   ↓
HTTP
   ↓
Backend
```

The frontend only knows the HTTP contract. It does not need to know whether the backend is implemented with Express, `net/http`, Gin, Fiber, or something else.

---

## Running the playground

Two terminals are required:

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

Only one frontend and one backend should be running at the same time.

Example:

```text
Vue :3000
   ↓
HTTP
   ↓
Fiber :8080
```

You can stop Fiber, start Express, and use the same Vue application without changing the HTTP calls.

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

The backend URL is configured in:

```text
frontend/vue/.env
```

```env
VITE_API_URL=http://localhost:8080
```

### HTTP clients

The Vue frontend intentionally contains two implementations of the same API calls.

#### Native Fetch API

```text
messageFetch.service.js
```

Uses the browser's native `fetch()` API.

With `fetch`, JSON handling is relatively explicit:

```js
const response = await fetch(url)
const data = await response.json()
```

For POST requests, the JavaScript object must also be serialized:

```js
body: JSON.stringify({ message })
```

#### Axios

```text
messageAxios.service.js
```

Axios adds abstractions on top of HTTP requests.

For example:

```js
axios.post(url, { message })
```

automatically serializes the object as JSON.

Responses are also parsed automatically and exposed through:

```js
response.data
```

This allows the frontend to compare:

```text
Browser Fetch API
       vs
Axios
```

while sending the same HTTP requests.

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

### HttpClient

Angular uses its built-in `HttpClient`:

```ts
this.http.get(...)
this.http.post(...)
```

Unlike the Promise-based Fetch API and Axios usage in the Vue example, Angular `HttpClient` returns RxJS `Observable`s.

Example:

```ts
this.messageService.fetchMessage().subscribe((data) => {
  this.getResponse = data.message
})
```

The HTTP contract is still exactly the same.

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

### Model

Express handlers use:

```js
(req, res)
```

Conceptually:

```text
request  → req
response → res
```

Example:

```js
app.get('/api/message', (req, res) => {
  res.json({
    message: 'Hello from node-express backend',
  })
})
```

For JSON request bodies:

```js
app.use(express.json())
```

parses the incoming HTTP body and makes the resulting JavaScript object available as:

```js
req.body
```

Express also provides helpers such as:

```js
req.params
req.query
req.body

res.status(...)
res.json(...)
res.send(...)
```

Architecture:

```text
Application
    ↓
Express
    ↓
Node HTTP
    ↓
Network
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

### Model

`net/http` is part of the Go standard library.

Handlers receive:

```go
func(w http.ResponseWriter, r *http.Request)
```

Conceptually:

```text
r → request
w → response
```

Unlike Express, JSON processing is explicit.

Incoming JSON:

```go
json.NewDecoder(r.Body).Decode(&body)
```

Outgoing JSON:

```go
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(response)
```

This makes `net/http` useful for understanding what higher-level frameworks are doing automatically.

Architecture:

```text
Application
    ↓
net/http
    ↓
Network
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

### Model

Gin handlers receive:

```go
func(c *gin.Context)
```

The context provides helpers around the underlying HTTP request and response.

For example:

```go
c.ShouldBindJSON(&body)
```

parses an incoming JSON body.

And:

```go
c.JSON(http.StatusOK, gin.H{
    "received": body.Message,
})
```

creates a JSON response.

Gin still uses Go's standard `net/http` underneath.

Architecture:

```text
Application
    ↓
Gin
    ↓
net/http
    ↓
Network
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

### Model

Fiber handlers use:

```go
func(c fiber.Ctx) error
```

Fiber centralizes request, response, parameters, headers, body parsing, and response helpers in `fiber.Ctx`.

Incoming JSON:

```go
c.Bind().Body(&body)
```

Outgoing JSON:

```go
return c.JSON(fiber.Map{
    "received": body.Message,
})
```

Unlike Gin, Fiber is not built directly on `net/http`. Its HTTP engine is `fasthttp`.

Architecture:

```text
Application
    ↓
Fiber
    ↓
fasthttp
    ↓
Network
```

---

# Backend comparison

The same POST endpoint looks roughly like this in each backend.

### Express

```js
app.post('/api/message', (req, res) => {
  const { message } = req.body

  res.json({
    received: message,
  })
})
```

### `net/http`

```go
func handlePostMessage(w http.ResponseWriter, r *http.Request) {
	var body MessageRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(MessageResponse{
		Received: body.Message,
	})
}
```

### Gin

```go
router.POST("/api/message", func(c *gin.Context) {
	var body MessageRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"received": body.Message,
	})
})
```

### Fiber

```go
app.Post("/api/message", func(c fiber.Ctx) error {
	var body MessageRequest

	if err := c.Bind().Body(&body); err != nil {
		return fiber.ErrBadRequest
	}

	return c.JSON(fiber.Map{
		"received": body.Message,
	})
})
```

All four handlers implement the same HTTP operation:

```text
HTTP request
    ↓
read JSON body
    ↓
application value
    ↓
create response
    ↓
serialize JSON
    ↓
HTTP response
```

The difference is how much of that process each framework exposes or abstracts.
