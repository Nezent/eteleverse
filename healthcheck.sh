#!/bin/bash

HEALTH_ENDPOINT="http://localhost:8080/api/v1/health"
MAX_RETRIES=5
RETRY_DELAY=2

echo "Checking health endpoint: $HEALTH_ENDPOINT"
echo ""

for i in $(seq 1 $MAX_RETRIES); do
    echo "Attempt $i/$MAX_RETRIES..."
    
    HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" $HEALTH_ENDPOINT)
    
    if [ "$HTTP_CODE" -eq 200 ]; then
        echo "✅ Health check passed!"
        echo ""
        echo "Response:"
        curl -s $HEALTH_ENDPOINT | jq '.' 2>/dev/null || curl -s $HEALTH_ENDPOINT
        echo ""
        exit 0
    else
        echo "❌ Health check failed with HTTP status: $HTTP_CODE"
        if [ $i -lt $MAX_RETRIES ]; then
            echo "Retrying in $RETRY_DELAY seconds..."
            sleep $RETRY_DELAY
        fi
    fi
done

echo ""
echo "❌ Health check failed after $MAX_RETRIES attempts"
exit 1
