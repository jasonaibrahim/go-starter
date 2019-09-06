# Go Starter

## Overview

Gin
- routing and html rendering
Auth0
- client cookies for session persistence
Deployment
- creates a docker image
- can deploy to the environment of your choosing. set up your own CI/CD

## Getting Started

Add `.env`

```text
SESSION_KEY=*****
SESSION_NAME=session
AUTH0_CLIENT_ID=**********
AUTH0_DOMAIN=*********
AUTH0_CLIENT_SECRET=********
AUTH0_CALLBACK_URL=http://localhost:3000/callback
PORT=3000
GIN_MODE=debug
```

Run the app in development
```text
go run src/main.go
```

Run tests
```text
go test test/*
```