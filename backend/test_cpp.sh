#!/bin/bash

CODE=$(cat test_fibonacci.cpp)

curl -X POST http://localhost:8080/api/v1/execute \
  -H "Content-Type: application/json" \
  -d @- << EOF
{
  "language_id": 54,
  "code": $(printf '%s' "$CODE" | python3 -c 'import json,sys; print(json.dumps(sys.stdin.read()))'),
  "stdin": ""
}
EOF
