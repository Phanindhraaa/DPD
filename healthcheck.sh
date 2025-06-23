#!/bin/bash

echo "🔍 Checking Service 1..."
curl -s -f http://localhost:8080/service1/ping > /dev/null
SERVICE1_STATUS=$?

echo "🔍 Checking Service 2..."
curl -s -f http://localhost:8080/service2/ping > /dev/null
SERVICE2_STATUS=$?

if [ $SERVICE1_STATUS -eq 0 ] && [ $SERVICE2_STATUS -eq 0 ]; then
  echo "✅ All services are healthy!"
  exit 0
else
  echo "❌ One or more services failed the health check."
  if [ $SERVICE1_STATUS -ne 0 ]; then
    echo "❌ Service 1 failed"
  fi
  if [ $SERVICE2_STATUS -ne 0 ]; then
    echo "❌ Service 2 failed"
  fi
  exit 1
fi
