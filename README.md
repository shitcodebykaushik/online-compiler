# 🚀 Online Code Compiler

A modern, full-stack online code compiler similar to Programiz, built with **React.js**, **Go**, **Piston**, and **Docker**. Supports multiple programming languages with real-time code execution in a secure, isolated environment.

![Status](https://img.shields.io/badge/Status-Production_Ready-success)
![React](https://img.shields.io/badge/React-18.2.0-blue)
![Go](https://img.shields.io/badge/Go-1.16+-00ADD8)
![Piston](https://img.shields.io/badge/Piston-Latest-green)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED)

---

## 🎯 Quick Deploy with Docker (Recommended)

**One-command deployment for the entire application!**

```bash
# Run the setup script
chmod +x setup.sh
./setup.sh
```

Then open: **http://localhost**

📖 **For detailed deployment instructions, see [DEPLOYMENT.md](DEPLOYMENT.md)**

---

## ✨ Features

### 🎨 **Modern UI/UX**
- Clean and intuitive interface
- Monaco Editor (VS Code engine) integration
- Syntax highlighting for multiple languages
- Real-time code execution
- Console output display
- Responsive design

### 💻 **Code Editor**
- Monaco Editor with IntelliSense
- Auto-completion
- Line numbering
- Code folding
- Multi-cursor support
- Customizable themes

### 🌍 **Multi-Language Support**
- 🐍 Python 3.10.0
- 🟨 JavaScript (Node.js 18.15.0)
- ☕ Java 15.0.2
- 🔷 C (GCC 10.2.0)
- 🔶 C++ (GCC 10.2.0)
- 🔵 Go 1.16.2

### 🎯 **Code Execution**
- Real-time code execution via Piston
- Display output and errors
- Show execution time
- Secure sandboxed environment
- Support for stdin input

---

## 🛠️ Tech Stack

**Frontend:**
- React 18.2.0
- Vite 5.0.0
- Monaco Editor (@monaco-editor/react)
- Lucide React (Icons)
- Vanilla CSS with CSS Variables

**Backend:**
- Go 1.16+
- Gin framework
- Piston API for code execution
- Docker containers
- SQLite (code snippets storage)
- Redis (optional - rate limiting & caching)

**Code Execution:**
- Piston (Docker-based execution engine)
- Supports 40+ programming languages
- Secure and isolated execution

---

## 📋 Prerequisites

- **Node.js** 16+ and npm
- **Go** 1.16+
- **Docker** and Docker Compose
- **Git**

---

## 🚀 Quick Start

### 1. Install Dependencies

**Frontend:**
```bash
npm install
```

**Backend:**
```bash
cd backend
go mod download
cd ..
```

### 2. Start Piston (Code Execution Engine)

```bash
cd backend
docker-compose -f piston-compose.yml up -d
```

### 3. Install Language Runtimes (First Time Only)

```bash
# Install Python
curl -X POST http://localhost:2000/api/v2/packages -H "Content-Type: application/json" -d '{"language":"python","version":"3.10.0"}'

# Install JavaScript
curl -X POST http://localhost:2000/api/v2/packages -H "Content-Type: application/json" -d '{"language":"node","version":"18.15.0"}'

# Install Java
curl -X POST http://localhost:2000/api/v2/packages -H "Content-Type: application/json" -d '{"language":"java","version":"15.0.2"}'

# Install C/C++
curl -X POST http://localhost:2000/api/v2/packages -H "Content-Type: application/json" -d '{"language":"gcc","version":"10.2.0"}'

# Install Go
curl -X POST http://localhost:2000/api/v2/packages -H "Content-Type: application/json" -d '{"language":"go","version":"1.16.2"}'
```

**Verify installations:**
```bash
curl -s http://localhost:2000/api/v2/runtimes
```

### 4. Start Backend Server

```bash
cd backend
go run cmd/server/main.go
```

Backend will run on: **http://localhost:8080**

### 5. Start Frontend Development Server

In a new terminal:
```bash
npm run dev
```

Frontend will run on: **http://localhost:5173**

### 6. Access the Application

Open your browser and navigate to:
```
http://localhost:5173
```

---

## 📁 Project Structure

```
online-compiler/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go              # Backend entry point
│   ├── configs/
│   │   └── config.go                # Configuration
│   ├── internal/
│   │   ├── api/
│   │   │   ├── router.go            # API routes
│   │   │   ├── handlers/            # Request handlers
│   │   │   └── middleware/          # Middleware
│   │   ├── database/
│   │   │   └── db.go                # Database setup
│   │   ├── models/
│   │   │   └── models.go            # Data models
│   │   └── services/
│   │       ├── piston.go            # Piston integration
│   │       └── snippet.go           # Snippet service
│   ├── piston-compose.yml           # Piston Docker Compose
│   └── go.mod                       # Go dependencies
├── src/
│   ├── components/                  # React components
│   │   ├── Header.jsx
│   │   ├── CodeEditor.jsx
│   │   ├── Console.jsx
│   │   ├── LanguageSelector.jsx
│   │   └── ControlPanel.jsx
│   ├── hooks/                       # Custom React hooks
│   │   └── useCodeExecution.js
│   ├── utils/                       # Utility functions
│   │   ├── codeTemplates.js
│   │   └── languageConfig.js
│   ├── styles/
│   │   └── variables.css
│   ├── App.jsx                      # Main app component
│   ├── main.jsx                     # App entry point
│   └── index.css                    # Global styles
├── public/                          # Static assets
├── package.json
├── vite.config.js
└── README.md
```

---

## 🎮 Usage

1. **Select Language**: Choose from the language dropdown
2. **Write Code**: Use Monaco editor (with IntelliSense)
3. **Run Code**: Click "Run Code" button
4. **View Output**: See results in console panel
5. **Reset**: Restore default template
6. **Save Snippets**: Save your code for later

---

## 🔌 API Endpoints

### Health Check
```
GET /api/v1/health
```

### Execute Code
```
POST /api/v1/execute
Content-Type: application/json

{
  "language_id": 71,
  "code": "print('Hello World')",
  "stdin": ""
}
```

**Language IDs:**
- Python: `71`
- JavaScript: `63`
- Java: `62`
- C++: `54`
- C: `50`
- Go: `60`

**Response:**
```json
{
  "stdout": "Hello World\n",
  "stderr": "",
  "output": "Hello World\n",
  "run": {
    "stdout": "Hello World\n",
    "stderr": "",
    "code": 0,
    "signal": null,
    "output": "Hello World\n"
  }
}
```

### Save Code Snippet
```
POST /api/v1/snippets
Content-Type: application/json

{
  "language": "python",
  "code": "print('Hello World')",
  "title": "My Snippet"
}
```

### Get Code Snippet
```
GET /api/v1/snippets/:id
```

---

## 🐳 Docker Commands

### Start Piston
```bash
cd backend
docker-compose -f piston-compose.yml up -d
```

### Stop Piston
```bash
docker-compose -f piston-compose.yml down
```

### View Logs
```bash
docker logs piston_api
```

### Restart Piston
```bash
docker-compose -f piston-compose.yml restart
```

### Check Status
```bash
docker ps | grep piston
```

---

## 🐛 Troubleshooting

### Piston Container Issues

**Check status:**
```bash
docker ps | grep piston
```

**View logs:**
```bash
docker logs piston_api --tail 50
```

**Restart:**
```bash
cd backend
docker-compose -f piston-compose.yml restart
```

### Backend Port Already in Use

Kill the process using port 8080:
```bash
lsof -ti:8080 | xargs kill -9
```

Or change the port in `backend/configs/config.go`.

### Frontend Port Already in Use

Kill the process using port 5173:
```bash
lsof -ti:5173 | xargs kill -9
```

Or change the port in `vite.config.js`.

### Language Runtime Not Available

Install the required runtime:
```bash
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"<language>","version":"<version>"}'
```

Check available packages:
```bash
curl -s http://localhost:2000/api/v2/packages
```

Check installed runtimes:
```bash
curl -s http://localhost:2000/api/v2/runtimes
```

### Frontend Not Connecting to Backend

1. Ensure backend is running on port 8080
2. Check CORS settings in `backend/configs/config.go`
3. Verify frontend API endpoint in `src/hooks/useCodeExecution.js`
4. Check browser console for errors

### Code Execution Errors

1. Verify Piston is running: `docker ps | grep piston`
2. Check language runtime is installed: `curl -s http://localhost:2000/api/v2/runtimes`
3. Test Piston directly:
```bash
curl -X POST http://localhost:2000/api/v2/execute \
  -H "Content-Type: application/json" \
  -d '{"language":"python","version":"3.10.0","files":[{"content":"print(\"test\")"}]}'
```

---

## 📝 Adding New Languages

### 1. Check Available Languages
```bash
curl -s http://localhost:2000/api/v2/packages
```

### 2. Install the Runtime
```bash
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"<language>","version":"<version>"}'
```

### 3. Update Backend

Edit `backend/internal/services/piston.go`:

```go
var languageMap = map[int]struct {
    Language string
    Version  string
    FileName string
}{
    // ... existing mappings
    XX: {"<language>", "<version>", "<filename>"},
}
```

### 4. Update Frontend

Edit `src/utils/languageConfig.js`:

```javascript
export const LANGUAGES = [
  // ... existing languages
  { id: XX, name: "<Language>", extension: "<ext>", icon: "<emoji>" }
];
```

Add template in `src/utils/codeTemplates.js`:

```javascript
export const CODE_TEMPLATES = {
  // ... existing templates
  XX: `// Your template code here`
};
```

### 5. Restart Backend
```bash
cd backend
go run cmd/server/main.go
```

---

## 🔒 Security

- **Rate limiting**: 30 requests per 15 minutes (configurable)
- **Sandboxed execution**: All code runs in isolated Piston containers
- **Input validation**: All API endpoints validate input
- **CORS configuration**: Restricted origins
- **Environment-based config**: Sensitive data in environment variables

### For Production:

1. **Set production mode:**
```bash
export GIN_MODE=release
```

2. **Configure CORS** in `backend/configs/config.go`:
```go
AllowedOrigins: []string{"https://yourdomain.com"}
```

3. **Use HTTPS** with a reverse proxy (nginx, Caddy)

4. **Implement authentication** for API endpoints

5. **Set up monitoring** and logging

6. **Use environment variables** for all sensitive configuration

7. **Enable rate limiting** aggressively

8. **Regular security updates** for all dependencies

---

## 🔧 Configuration

### Backend Configuration

Edit `backend/configs/config.go`:

```go
type Config struct {
    ServerPort      string
    PistonURL       string
    DatabasePath    string
    AllowedOrigins  []string
    RateLimitWindow time.Duration
    RateLimitMax    int
}
```

### Frontend Configuration

Edit `src/hooks/useCodeExecution.js`:

```javascript
const API_BASE_URL = 'http://localhost:8080/api/v1';
```

### Piston Configuration

Edit `backend/piston-compose.yml`:

```yaml
services:
  piston_api:
    container_name: piston_api
    image: ghcr.io/engineer-man/piston
    # ... configuration
```

---

## 🚀 Deployment

### Docker Deployment (Recommended)

1. **Build backend:**
```bash
cd backend
docker build -t online-compiler-backend .
```

2. **Create production docker-compose.yml:**
```yaml
version: '3.8'
services:
  piston:
    image: ghcr.io/engineer-man/piston
    # ... Piston config
  
  backend:
    image: online-compiler-backend
    ports:
      - "8080:8080"
    environment:
      - GIN_MODE=release
    depends_on:
      - piston
  
  frontend:
    build: .
    ports:
      - "80:80"
```

3. **Deploy:**
```bash
docker-compose up -d
```

### Manual Deployment

1. **Build frontend:**
```bash
npm run build
```

2. **Deploy static files** to your web server

3. **Run backend** with process manager (PM2, systemd)

4. **Set up reverse proxy** (nginx, Caddy)

---

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📄 License

This project is licensed under the MIT License.

---

## 🙏 Acknowledgments

- **Monaco Editor** - VS Code's powerful code editor
- **Piston** - Open-source code execution engine
- **Gin** - High-performance Go web framework
- **Lucide React** - Beautiful icon library
- **Programiz** - Design inspiration

---

## 💡 Why Piston Over Judge0?

This project initially used Judge0 but switched to Piston for the following reasons:

- **Better cgroup v2 support**: Piston works natively with modern Linux systems (Kali, Ubuntu 22.04+)
- **Simpler setup**: No complex configuration needed
- **More languages**: 40+ languages supported out of the box
- **Active development**: Regular updates and community support
- **Docker-first design**: Better container isolation
- **No sandbox issues**: Works on all modern Linux distributions

---

## 📞 Support

For issues and questions:
- Create an issue on GitHub
- Check the troubleshooting section
- Review the API documentation

---

**🎯 Status:** Production Ready ✅

**📍 Project Location:** `/home/shiiit/.gemini/antigravity/scratch/online-compiler`

**🌐 Development URLs:**
- Frontend: http://localhost:5173
- Backend: http://localhost:8080
- Piston: http://localhost:2000

---

**Built with ❤️ using React, Go, and Piston**
