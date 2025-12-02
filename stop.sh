#!/bin/bash

# Stop script for Online Compiler

echo "🛑 Stopping Online Compiler..."

# Stop all services
docker-compose down

echo ""
echo "✅ All services stopped!"
echo ""
echo "💡 To start again, run: ./start.sh"
echo "💡 To remove all data, run: docker-compose down -v"
echo ""
