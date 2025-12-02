# 🚀 Quick Start Guide - What to Do Now

## Current Status
✅ Frontend is running on http://localhost:5173  
⏳ Docker containers are being built and started  
⏳ Backend is ready but waiting for services  

---

## Option 1: Wait for Docker (Recommended for Full Production Setup)

**What's Happening:**
- Judge0 images downloaded (10.5GB)
- Docker is building the Go backend container
- Containers will start automatically

**Wait Time:** ~5-10 more minutes

**When Ready, Check:**
```bash
cd /home/shiiit/.gemini/antigravity/scratch/online-compiler/backend
docker-compose ps
```

**You should see:**
- judge0-server (running)
- judge0-worker (running)
- app-redis (running)
- compiler-backend (running)

**Then Test:**
```bash
# Test health
curl http://localhost:8080/api/v1/health

# Test code execution
curl -X POST http://localhost:8080/api/v1/execute \
  -H "Content-Type: application/json" \
  -d '{"language_id": 71, "code": "print(\"Hello!\")"}'
```

---

## Option 2: Run Backend Without Docker (Quick Test Now)

If you want to test immediately without waiting:

**Terminal 1 - Start Redis Only:**
```bash
cd /home/shiiit/.gemini/antigravity/scratch/online-compiler/backend
docker run -d --name temp-redis -p 6379:6379 redis:7-alpine
```

**Terminal 2 - Run Go Backend:**
```bash
cd /home/shiiit/.gemini/antigravity/scratch/online-compiler/backend
go run cmd/server/main.go
```

**Note:** This won't have Judge0, so code execution will fail, but you can:
- Test the API is responding
- Test rate limiting
- Test snippet creation/retrieval
- Test the frontend connection

---

## Option 3: Use Public Judge0 Instance (Temporary)

Edit `.env` to use a public Judge0 instance:

```bash
JUDGE0_URL=https://judge0-ce.p.rapidapi.com
```

Then run:
```bash
go run cmd/server/main.go
```

**Limitation:** Public instances may have rate limits

---

## ✅ What You Can Do RIGHT NOW

### 1. Test the Frontend
Open http://localhost:5173 and try:
- ✅ Change languages
- ✅ Edit code
- ✅ Toggle dark/light theme
- ✅ Test UI responsiveness

### 2. Explore the Code
Check out:
- `src/components/` - React components
- `backend/internal/api/handlers/` - API handlers
- `backend/internal/services/` - Business logic
- `backend/docker-compose.yml` - Infrastructure

### 3. Review Documentation
- `/backend/README.md` - Backend docs
- `README.md` - Project overview
- Walkthrough artifact - Complete guide

---

## 🎯 Recommended Path

**For Learning/Testing:**
→ **Option 2** (Run without Docker now, add Judge0 later)

**For Production:**
→ **Option 1** (Wait for full Docker stack)

**For Quick Demo:**
→ **Option 3** (Use public Judge0)

---

## 📝 Commands Summary

**Check Docker Status:**
```bash
cd backend
docker-compose ps
docker-compose logs -f
```

**Start Backend Manually:**
```bash
cd backend
go run cmd/server/main.go
```

**Test Backend API:**
```bash
curl http://localhost:8080/api/v1/health
```

**Stop Everything:**
```bash
# Stop frontend
Ctrl+C in frontend terminal

# Stop backend
Ctrl+C in backend terminal

# Stop Docker
cd backend
docker-compose down
```

---

## 🐛 Troubleshooting

**Port Already in Use:**
```bash
# Find what's using port 8080
lsof -i :8080

# Kill it or change PORT in .env
```

**Docker Issues:**
```bash
# Restart Docker
sudo systemctl restart docker

# Clean up
docker-compose down -v
docker system prune -f
```

**Go Module Issues:**
```bash
cd backend
go mod tidy
go mod download
```

---

## 💡 What I Recommend

**Do this NOW:**

1. ✅ **Keep frontend running** - It's already working perfectly!

2. ✅ **Start a simple Redis container:**
   ```bash
   docker run -d --name temp-redis -p 6379:6379 redis:7-alpine
   ```

3. ✅ **Start the Go backend:**
   ```bash
   cd /home/shiiit/.gemini/antigravity/scratch/online-compiler/backend
   go run cmd/server/main.go
   ```

4. ✅ **Test in the frontend:**
   - Click "Run Code" 
   - You'll see an error about Judge0 (expected)
   - But API connection works!

5. ⏳ **Later, when Docker finishes:**
   - Stop the manual backend (Ctrl+C)
   - Stop temp Redis: `docker stop temp-redis && docker rm temp-redis`
   - Start full stack: `docker-compose up -d`
   - Test again with full Judge0 support!

---

**Current Time:** You've been waiting ~10 minutes already.  
**Docker should be done in:** ~5 minutes  

**Choose your path!** 🚀
