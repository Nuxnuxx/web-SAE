services:
- type: web
  name: app-backend
  runtime: go
  rootDir: backend
  buildCommand: go build -o app .
  startCommand: ./app
- type: web
  name: app-community
  runtime: docker
  rootDir: community/docker
  dockerfilePath: ./Dockerfile
  dockerContext: .
