#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}================================${NC}"
echo -e "${BLUE}  Online Compiler Setup Script${NC}"
echo -e "${BLUE}================================${NC}\n"

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo -e "${RED}❌ Docker is not installed. Please install Docker first.${NC}"
    exit 1
fi

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo -e "${RED}❌ Docker Compose is not installed. Please install Docker Compose first.${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Docker and Docker Compose are installed${NC}\n"

# Build and start all services
echo -e "${BLUE}🚀 Building and starting all services...${NC}\n"
docker-compose up -d --build

# Wait for services to be ready
echo -e "\n${YELLOW}⏳ Waiting for services to be ready...${NC}"
sleep 10

# Check if services are running
echo -e "\n${BLUE}📊 Checking service status...${NC}\n"
docker-compose ps

# Install language runtimes in Piston
echo -e "\n${BLUE}📦 Installing language runtimes in Piston...${NC}\n"

echo -e "${YELLOW}Installing Python 3.10.0...${NC}"
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"python","version":"3.10.0"}' 2>/dev/null
echo ""

echo -e "${YELLOW}Installing JavaScript (Node.js 18.15.0)...${NC}"
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"node","version":"18.15.0"}' 2>/dev/null
echo ""

echo -e "${YELLOW}Installing Java 15.0.2...${NC}"
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"java","version":"15.0.2"}' 2>/dev/null
echo ""

echo -e "${YELLOW}Installing C/C++ (GCC 10.2.0)...${NC}"
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"gcc","version":"10.2.0"}' 2>/dev/null
echo ""

echo -e "${YELLOW}Installing Go 1.16.2...${NC}"
curl -X POST http://localhost:2000/api/v2/packages \
  -H "Content-Type: application/json" \
  -d '{"language":"go","version":"1.16.2"}' 2>/dev/null
echo ""

# Verify installations
echo -e "\n${BLUE}✅ Verifying installed runtimes...${NC}"
curl -s http://localhost:2000/api/v2/runtimes | grep -o '"language":"[^"]*"' | sed 's/"language":"//g' | sed 's/"//g' | sort | uniq

echo -e "\n${GREEN}================================${NC}"
echo -e "${GREEN}  🎉 Setup Complete!${NC}"
echo -e "${GREEN}================================${NC}\n"

echo -e "${BLUE}📍 Application URLs:${NC}"
echo -e "   Frontend:  ${GREEN}http://localhost${NC}"
echo -e "   Backend:   ${GREEN}http://localhost:8080${NC}"
echo -e "   Piston:    ${GREEN}http://localhost:2000${NC}\n"

echo -e "${YELLOW}📝 Useful Commands:${NC}"
echo -e "   View logs:        ${BLUE}docker-compose logs -f${NC}"
echo -e "   Stop services:    ${BLUE}docker-compose down${NC}"
echo -e "   Restart services: ${BLUE}docker-compose restart${NC}"
echo -e "   View status:      ${BLUE}docker-compose ps${NC}\n"

echo -e "${GREEN}🌐 Open your browser and visit: http://localhost${NC}\n"
