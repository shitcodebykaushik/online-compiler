#!/bin/bash

# Simple start script for Online Compiler

echo "🚀 Starting Online Compiler..."

# Start all services
docker-compose up -d

echo ""
echo "✅ Services started!"
echo ""
echo "📍 Application URLs:"
echo "   Frontend:  http://localhost"
echo "   Backend:   http://localhost:8080"
echo "   Piston:    http://localhost:2000"
echo ""
echo "💡 Tip: Run './setup.sh' first if this is your first time!"
echo ""
