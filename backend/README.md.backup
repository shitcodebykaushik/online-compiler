# Online Compiler Backend - README

A robust Go backend for the online code compiler platform, integrated with Judge0, Redis, and SQLite.

---

## 🏗️ Architecture

- **Go API Server** (Gin framework)
- **Judge0 CE** for code execution
- **Redis** for caching and rate limiting
- **SQLite** for snippet storage

---

## 🚀 Quick Start

### Prerequisites

- Go 1.21+
- Docker & Docker Compose
- Port 8080 available (backend)
- Port 2358 available (Judge0)
- Port 6379 available (Redis)

### Method 1: Docker Compose (Recommended)

```bash
# Start all services
docker-compose up -d

# Check status
docker-compose ps

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Method 2: Local Development

**1. Start Judge0 and Redis:**
```bash
docker-compose up -d judge0-db judge0-redis judge0-server judge0-worker app-redis
```

**2. Wait for Judge0 to be ready (30-60 seconds):**
```bash
curl http://localhost:2358/about
```

**3. Run Go backend:**
```bash
cd backend
go run cmd/server/main.go
```

---

## 📡 API Endpoints

### Health Check
```bash
curl http://localhost:8080/api/v1/health
```

### Execute Code
```bash
curl -X POST http://localhost:8080/api/v1/execute \
  -H "Content-Type: application/json" \
  -d '{
    "language_id": 71,
    "code": "print(\"Hello, World!\")",
    "stdin": ""
  }'
```

**Response:**
```json
{
  "success": true,
  "output": "Hello, World!\n",
  "execution_time": 42.5,
  "memory_kb": 256,
  "status": "Accepted"
}
```

### Create Snippet
```bash
curl -X POST http://localhost:8080/api/v1/snippets \
  -H "Content-Type: application/json" \
  -d '{
    "language": "python",
    "code": "print(\"My snippet\")",
    "title": "Test Snippet"
  }'
```

### Get Snippet
```bash
curl http://localhost:8080/api/v1/snippets/{snippet_id}
```

---

## 🔧 Configuration

Edit `.env` file:

```bash
# Server
PORT=8080
GIN_MODE=debug  # or 'release' for production

# Judge0
JUDGE0_URL=http://localhost:2358
JUDGE0_TIMEOUT=10

# Redis
REDIS_URL=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# Database
DATABASE_PATH=./data/compiler.db

# Rate Limiting
RATE_LIMIT_REQUESTS=30
RATE_LIMIT_WINDOW=900  # 15 minutes

# CORS
ALLOWED_ORIGINS=http://localhost:5173,http://localhost:3000
```

---

## 📊 Language IDs (Judge0)

| Language   | ID |
|------------|-----|
| C          | 50  |
| C++        | 54  |
| Python     | 71  |
| Java       | 62  |
| JavaScript | 63  |
| Go         | 60  |
| Rust       | 73  |
| PHP        | 68  |

---

## 🛠️ Project Structure

```
backend/
├── cmd/server/main.go              # Entry point
├── internal/
│   ├── api/
│   │   ├── handlers/              # HTTP handlers
│   │   ├── middleware/            # Middleware
│   │   └── router.go              # Routes
│   ├── models/                    # Data models
│   ├── services/                  # Business logic
│   │   ├── judge0.go             # Judge0 integration
│   │   ├── cache.go              # Redis caching
│   │   └── snippet.go            # Snippet management
│   └── database/                 # Database setup
├── configs/                       # Configuration
├── docker-compose.yml            # Docker orchestration
└── .env                          # Environment variables
```

---

## 🔒 Security Features

✅ **Input Validation** - Max 64KB code size, language ID validation  
✅ **Rate Limiting** - 30 requests per 15 minutes per IP  
✅ **CORS Protection** - Whitelist allowed origins  
✅ **Isolated Execution** - Judge0 runs in containers  
✅ **Result Caching** - Reduces load on Judge0

---

## 🧪 Testing

### Test Health Endpoint
```bash
curl http://localhost:8080/api/v1/health
```

Expected:
```json
{
  "status": "healthy",
  "redis": "connected",
  "database": "connected",
  "judge0": "available"
}
```

### Test Code Execution (All Languages)

**Python:**
```bash
curl -X POST http://localhost:8080/api/v1/execute \
  -H "Content-Type: application/json" \
  -d '{"language_id": 71, "code": "print(2 + 2)"}'
```

**C:**
```bash
curl -X POST http://localhost:8080/api/v1/execute \
  -H "Content-Type: application/json" \
  -d '{"language_id": 50, "code": "#include <stdio.h>\nint main() { printf(\"Hello\"); return 0; }"}'
```

**JavaScript:**
```bash
curl -X POST http://localhost:8080/api/v1/execute \
  -H "Content-Type: application/json" \
  -d '{"language_id": 63, "code": "console.log(\"Hello\")"}'
```

### Test Rate Limiting
```bash
# Run this 31 times quickly
for i in {1..31}; do
  curl -X POST http://localhost:8080/api/v1/execute \
    -H "Content-Type: application/json" \
    -d '{"language_id": 71, "code": "print(1)"}' &
done
```

Expected on 31st request:
```json
{
  "success": false,
  "error": "Rate limit exceeded. Please try again later.",
  "code": "RATE_LIMIT_EXCEEDED"
}
```

---

## 📝 Troubleshooting

### Judge0 Not Available

```bash
# Check if Judge0 server is running
docker ps | grep judge0

# View Judge0 logs
docker logs judge0-server

# Restart Judge0
docker-compose restart judge0-server
```

### Redis Connection Failed

```bash
# Check Redis status
docker ps | grep redis

# Test Redis connection
docker exec -it app-redis redis-cli ping
```

### Database Issues

```bash
# Check if database file exists
ls -lh data/compiler.db

# Remove and recreate (WARNING: deletes all data)
rm data/compiler.db
# Restart backend to recreate
```

---

## 🚢 Deployment

### Production Build

```bash
# Build binary
go build -o compiler-server cmd/server/main.go

# Run
GIN_MODE=release ./compiler-server
```

### Docker Production

```bash
# Build and start all services
docker-compose up -d --build

# View status
docker-compose ps
```

---

## 📦 Dependencies

```
github.com/gin-gonic/gin          # Web framework
github.com/go-redis/redis/v8      # Redis client
github.com/google/uuid            # UUID generation
github.com/joho/godotenv          # Environment variables
gorm.io/driver/sqlite             # SQLite driver
gorm.io/gorm                      # ORM
```

---

## 🔗 Integration with Frontend

The frontend is already configured to connect to the backend at `http://localhost:8080`.

**Frontend API Call:**
```javascript
const response = await fetch('http://localhost:8080/api/v1/execute', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    language_id: 71,
    code: code,
    stdin: ''
  })
});
```

---

## 📈 Performance

- **Response Time:** < 100ms (cached results)
- **Judge0 Execution:** 1-5 seconds (varies by language)
- **Rate Limit:** 30 requests per 15 min per IP
- **Cache TTL:** 1 hour for results

---

## 🎯 Next Steps

1. **Start Services:** `docker-compose up -d`
2. **Test Health:** `curl http://localhost:8080/api/v1/health`
3. **Test Execution:** Run code via API
4. **Start Frontend:** Frontend will connect automatically

---

**Status:** ✅ Backend Complete | ⏳ Waiting for Judge0 Docker Images to Download

**Project Location:** `/home/shiiit/.gemini/antigravity/scratch/online-compiler/backend`
