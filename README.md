# Fullstack react + Go http server

> Basic single page react application with an http backend server implemented in Go.

## How to run

Must have go >= v1.18 installed.

**Development server**

Concurrently runs the react client with hot reloading and the go server. Requests from the client in development are proxied to the server port.

- `Client port`: 3002
- `Go server port`: 8080

```bash
yarn start
```

**Production build**

```bash
yarn build
```

**Start server**

```bash
yarn serve

# or

go run server.go
```

- GET `/` route serves static react build
- GET `/hello` route returns a json payload
