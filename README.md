## Quick start
```bash
# Local
go run cmd/api/main.go --packs configs/pack.json

# Docker
docker build -t pack-calculator .
docker run -p 8000:8000 pack-calculator

# Compose
docker compose up --build