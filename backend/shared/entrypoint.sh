#!/bin/sh
set -e

# Run migrations if directory exists
if [ -d "/app/migrations" ]; then
    echo "Running migrations..."
    /app/migrate -path /app/migrations -database "${DB_URL}" up
fi

# Start the service
exec /app/${SERVICE_NAME}