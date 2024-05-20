# Web App

## Installation

```bash
npm i
npm run build
go build -o web-server src/main.go
docker build -t andreasatle/go-fiber .
docker run -p 80:8080 -p 443:8443 -d andreasatle/go-fiber
```
