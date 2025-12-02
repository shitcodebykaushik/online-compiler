# 🚀 Online Compiler - Docker Deployment Guide

This guide will help you deploy the entire Online Compiler application using Docker in just a few steps.

## 📋 Prerequisites

- **Docker** (version 20.10 or higher)
- **Docker Compose** (version 2.0 or higher)
- At least **4GB RAM** available
- At least **10GB disk space**

## 🎯 Quick Start (Recommended)

### Option 1: Automated Setup Script

1. **Clone/Extract the project:**
```bash
cd online-compiler
```

2. **Run the setup script:**
```bash
chmod +x setup.sh
./setup.sh
```

3. **Access the application:**
```
http://localhost
```

That's it! The script will:
- ✅ Build all Docker images
- ✅ Start all services (Piston, Backend, Frontend)
- ✅ Install all language runtimes
- ✅ Verify everything is working

---

## 🔧 Manual Setup

### Step 1: Build and Start Services

```bash
docker-compose up -d --build
```

This will start:
- **Frontend** on port `80`
- **Backend** on port `8080`
- **Piston** on port `2000`

### Step 2: Install Language Runtimes

Wait about 30 seconds for Piston to start, then install the runtimes:

```bash
# Python 3.10.0
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"python","version":"3.10.0"}'

# JavaScript (Node.js 18.15.0)
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"node","version":"18.15.0"}'

# Java 15.0.2
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"java","version":"15.0.2"}'

# C/C++ (GCC 10.2.0)
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"gcc","version":"10.2.0"}'

# Go 1.16.2
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"go","version":"1.16.2"}'
```

### Step 3: Verify Installation

```bash
# Check if all services are running
docker-compose ps

# Verify installed runtimes
curl -s http://localhost:2000/api/v2/runtimes
```

### Step 4: Access the Application

Open your browser and visit:
```
http://localhost
```

---

## 🐳 Docker Commands Reference

### View Logs
```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f frontend
docker-compose logs -f backend
docker-compose logs -f piston
```

### Stop Services
```bash
docker-compose down
```

### Restart Services
```bash
docker-compose restart
```

### Rebuild and Restart
```bash
docker-compose up -d --build
```

### Remove Everything (including volumes)
```bash
docker-compose down -v
```

### Check Service Status
```bash
docker-compose ps
```

---

## 🌐 Service Ports

| Service  | Internal Port | External Port | URL                    |
|----------|--------------|---------------|------------------------|
| Frontend | 80           | 80            | http://localhost       |
| Backend  | 8080         | 8080          | http://localhost:8080  |
| Piston   | 2000         | 2000          | http://localhost:2000  |

---

## 🔒 Production Deployment

### 1. Update Environment Variables

Create a `.env` file:

```env
# Backend Configuration
GIN_MODE=release
PORT=8080
JUDGE0_URL=http://piston:2000
DATABASE_PATH=/app/data/compiler.db
ALLOWED_ORIGINS=https://yourdomain.com

# Frontend Configuration
VITE_API_URL=https://api.yourdomain.com
```

### 2. Use Production Docker Compose

Create `docker-compose.prod.yml`:

```yaml
version: '3.8'

services:
  piston:
    image: ghcr.io/engineer-man/piston
    container_name: piston_api
    privileged: true
    volumes:
      - piston_packages:/piston/packages
    tmpfs:
      - /piston/jobs:exec,uid=1000,gid=1000,mode=711
    networks:
      - app-network
    restart: always

  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    container_name: online_compiler_backend
    env_file: .env
    volumes:
      - backend_data:/app/data
    depends_on:
      - piston
    networks:
      - app-network
    restart: always

  frontend:
    build:
      context: .
      dockerfile: Dockerfile
      args:
        - VITE_API_URL=${VITE_API_URL}
    container_name: online_compiler_frontend
    depends_on:
      - backend
    networks:
      - app-network
    restart: always

  nginx:
    image: nginx:alpine
    container_name: nginx_proxy
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx-prod.conf:/etc/nginx/nginx.conf
      - ./ssl:/etc/nginx/ssl
    depends_on:
      - frontend
      - backend
    networks:
      - app-network
    restart: always

volumes:
  piston_packages:
  backend_data:

networks:
  app-network:
    driver: bridge
```

### 3. Deploy
```bash
docker-compose -f docker-compose.prod.yml up -d --build
```

---

## 🐛 Troubleshooting

### Services Not Starting

**Check logs:**
```bash
docker-compose logs
```

**Check if ports are in use:**
```bash
lsof -i :80
lsof -i :8080
lsof -i :2000
```

### Language Runtime Not Working

**Reinstall the runtime:**
```bash
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"<language>","version":"<version>"}'
```

**Check available packages:**
```bash
curl -s http://localhost:2000/api/v2/packages
```

### Frontend Can't Connect to Backend

1. Check if backend is running:
```bash
curl http://localhost:8080/api/v1/health
```

2. Check backend logs:
```bash
docker-compose logs backend
```

3. Verify CORS settings in backend configuration

### Database Issues

**Reset database:**
```bash
docker-compose down -v
docker-compose up -d
```

### Out of Memory

**Increase Docker memory limit:**
- Docker Desktop → Settings → Resources → Memory
- Allocate at least 4GB

---

## 📦 What's Included

```
online-compiler/
├── docker-compose.yml       # Main Docker Compose configuration
├── setup.sh                 # Automated setup script
├── Dockerfile               # Frontend Docker image
├── nginx.conf               # Nginx configuration
├── .dockerignore            # Docker ignore file
├── backend/
│   ├── Dockerfile           # Backend Docker image
│   └── .dockerignore        # Backend Docker ignore
└── README.md                # Main documentation
```

---

## 🎯 Health Checks

### Frontend
```bash
curl http://localhost
```

### Backend
```bash
curl http://localhost:8080/api/v1/health
```

### Piston
```bash
curl http://localhost:2000/api/v2/runtimes
```

---

## 🚀 Scaling

To run multiple instances of the backend:

```bash
docker-compose up -d --scale backend=3
```

Add a load balancer (nginx) in front for distribution.

---

## 💾 Data Persistence

Data is persisted in Docker volumes:
- `piston_packages`: Installed language runtimes
- `backend_data`: SQLite database with code snippets

To backup:
```bash
docker run --rm -v online-compiler_backend_data:/data -v $(pwd):/backup alpine tar czf /backup/backup.tar.gz /data
```

To restore:
```bash
docker run --rm -v online-compiler_backend_data:/data -v $(pwd):/backup alpine tar xzf /backup/backup.tar.gz -C /
```

---

## 🔐 Security Notes

- Change default ports in production
- Use HTTPS with SSL certificates
- Implement rate limiting (Redis recommended)
- Set up firewall rules
- Regular security updates
- Use secrets for sensitive data

---

## 📞 Support

For issues or questions:
- Check the logs: `docker-compose logs -f`
- Verify all services are running: `docker-compose ps`
- Review the main README.md for detailed documentation

---

**Built with ❤️ using Docker, React, Go, and Piston**
