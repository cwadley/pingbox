# pingbox
A simple ping webserver

A webserver that listens on port 8080 and always returns 200 for requests. Useful as a ping endpoint for network uptime probes.

## Usage
### Docker Compose
`docker-compose up -d`

### Docker
`docker run -d -p 8080:8080 cwadley/pingbox:latest`

### Non-Docker
`go run pingbox.go`
