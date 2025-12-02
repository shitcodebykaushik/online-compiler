# 📦 Docker Package - Complete Application

This package contains everything needed to run the Online Code Compiler application using Docker.

## 📋 What's Included

```
online-compiler/
├── 🚀 setup.sh              # Automated setup (RECOMMENDED)
├── ▶️  start.sh              # Start all services
├── ⏹️  stop.sh               # Stop all services
├── 🐳 docker-compose.yml    # Docker orchestration
├── 📖 DEPLOYMENT.md         # Detailed deployment guide
├── 📘 README.md             # Full documentation
├── 🌐 Dockerfile            # Frontend container
├── ⚙️  nginx.conf            # Web server config
├── 🚫 .dockerignore         # Docker ignore rules
├── backend/
│   ├── 🐳 Dockerfile        # Backend container
│   └── 🚫 .dockerignore     # Backend ignore rules
└── [source code...]
```

## 🎯 Quick Start (3 Steps)

### 1. Prerequisites
- Install Docker: https://docs.docker.com/get-docker/
- Install Docker Compose: https://docs.docker.com/compose/install/

### 2. Run Setup
```bash
chmod +x setup.sh
./setup.sh
```

### 3. Access Application
Open your browser: **http://localhost**

**That's it! 🎉**

## 🎮 Common Commands

```bash
# First time setup (installs everything)
./setup.sh

# Start services (after initial setup)
./start.sh

# Stop services
./stop.sh

# View logs
docker-compose logs -f

# Check status
docker-compose ps

# Restart
docker-compose restart

# Complete removal
docker-compose down -v
```

## 🌐 Service URLs

| Service  | URL                   | Description           |
|----------|-----------------------|-----------------------|
| Frontend | http://localhost      | Web Interface         |
| Backend  | http://localhost:8080 | API Server            |
| Piston   | http://localhost:2000 | Code Execution Engine |

## 📦 Supported Languages

- 🐍 Python 3.10.0
- 🟨 JavaScript (Node.js 18.15.0)
- ☕ Java 15.0.2
- 🔷 C (GCC 10.2.0)
- 🔶 C++ (GCC 10.2.0)
- 🔵 Go 1.16.2

## 📤 Sharing This Package

To share with others:

### Option 1: ZIP Archive
```bash
zip -r online-compiler.zip online-compiler/
```

### Option 2: Git Repository
```bash
git init
git add .
git commit -m "Initial commit"
git remote add origin <your-repo-url>
git push -u origin main
```

### Option 3: Docker Hub (Images Only)
```bash
# Tag images
docker tag online-compiler-frontend:latest yourusername/online-compiler-frontend:latest
docker tag online-compiler-backend:latest yourusername/online-compiler-backend:latest

# Push to Docker Hub
docker push yourusername/online-compiler-frontend:latest
docker push yourusername/online-compiler-backend:latest
```

## 🔧 System Requirements

| Resource | Minimum | Recommended |
|----------|---------|-------------|
| CPU      | 2 cores | 4+ cores    |
| RAM      | 4 GB    | 8+ GB       |
| Disk     | 10 GB   | 20+ GB      |
| OS       | Any with Docker | Linux/Mac |

## 🐛 Troubleshooting

### Port Already in Use
```bash
# Check what's using the port
lsof -i :80
lsof -i :8080
lsof -i :2000

# Stop conflicting services or change ports in docker-compose.yml
```

### Services Not Starting
```bash
# View logs
docker-compose logs

# Restart services
docker-compose restart

# Complete rebuild
docker-compose down -v
docker-compose up -d --build
```

### Language Not Working
```bash
# Reinstall runtime (example: Python)
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"python","version":"3.10.0"}'
```

## 📚 Documentation

- **DEPLOYMENT.md** - Complete deployment guide with production tips
- **README.md** - Full application documentation
- **setup.sh** - Automated setup script (view source for details)

## 🎓 For Recipients

If someone sent you this package:

1. **Extract** the archive (if zipped)
2. **Install** Docker and Docker Compose
3. **Run** `./setup.sh`
4. **Open** http://localhost in your browser
5. **Enjoy** coding online!

## ⚡ Performance Tips

- Allocate at least 4GB RAM to Docker
- Use SSD storage for better performance
- Keep Docker updated to latest version
- Close unnecessary applications

## 🔒 Security Notes

**Development Use:**
- Default setup is for development/testing
- All services run on localhost

**Production Use:**
- See DEPLOYMENT.md for production setup
- Configure HTTPS/SSL
- Set up proper firewall rules
- Use environment variables for secrets
- Enable rate limiting with Redis

## 🆘 Getting Help

1. Check **DEPLOYMENT.md** for detailed guides
2. View logs: `docker-compose logs -f`
3. Check service status: `docker-compose ps`
4. Verify Docker: `docker --version`

## 📝 Quick Reference Card

```bash
# First Time
./setup.sh              # Complete setup with all languages

# Daily Use
./start.sh              # Start application
./stop.sh               # Stop application

# Monitoring
docker-compose ps       # Check status
docker-compose logs -f  # View logs

# Maintenance
docker-compose restart  # Restart all services
docker-compose down -v  # Complete cleanup
```

## ✅ Verification Checklist

After running setup.sh, verify:

- [ ] http://localhost loads the frontend
- [ ] http://localhost:8080/api/v1/health returns `{"status":"ok"}`
- [ ] Can select a language from dropdown
- [ ] Can write and execute Python code
- [ ] Can write and execute JavaScript code
- [ ] Output appears in console
- [ ] No errors in browser console

## 🎉 Success!

If all services are running and you can execute code, you're all set!

---

**Built with ❤️ using Docker, React, Go, and Piston**

**Version:** 1.0.0  
**Last Updated:** December 2025  
**License:** MIT
