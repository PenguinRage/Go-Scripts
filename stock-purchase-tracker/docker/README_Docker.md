# Stock Purchase Tracker - Docker Setup

This project includes a Docker setup for PostgreSQL database to store stock purchase data.

## Prerequisites

- Docker
- Docker Compose

## Quick Start

1. **Start the PostgreSQL database:**
   ```bash
   docker-compose up -d
   ```

2. **Check if the database is running:**
   ```bash
   docker-compose ps
   ```

3. **View database logs:**
   ```bash
   docker-compose logs postgres
   ```

## Database Details

- **Database Name:** stocks
- **Username:** stockuser
- **Password:** stockpass
- **Port:** 5432 (mapped to localhost:5432)
- **Container Name:** stock-tracker-db

## Table Structure

The `stocks` table includes the following columns:

| Column | Type | Description |
|--------|------|-------------|
| id | SERIAL | Primary key, auto-incrementing |
| exchange | VARCHAR(10) | Stock exchange (e.g., ASX, NYSE) |
| code | VARCHAR(20) | Stock symbol/code |
| state | VARCHAR(10) | Transaction state (BUY/SELL) |
| quantity | INTEGER | Number of shares |
| currency | VARCHAR(3) | Currency code (AUD, USD, etc.) |
| price | DECIMAL(10,2) | Price per share |
| created_at | TIMESTAMP | Record creation timestamp |
| updated_at | TIMESTAMP | Last update timestamp |

## Sample Data

The database is initialized with sample stock data for testing purposes.

## Connecting to the Database

### Using psql (if you have PostgreSQL client installed):
```bash
psql -h localhost -p 5432 -U stockuser -d stocks
```

### Using Docker:
```bash
docker exec -it stock-tracker-db psql -U stockuser -d stocks
```

## Useful Commands

### Stop the database:
```bash
docker-compose down
```

### Stop and remove all data:
```bash
docker-compose down -v
```

### Restart the database:
```bash
docker-compose restart
```

### View database logs:
```bash
docker-compose logs -f postgres
```

## Application Connection

Your Go application is configured to connect to the database using:
- Host: localhost
- Port: 5432
- Database: stocks
- Username: stockuser
- Password: stockpass

The connection string is already updated in `db/db.go`.

## Troubleshooting

1. **Port already in use:** If port 5432 is already in use, modify the port mapping in `docker-compose.yml`
2. **Permission issues:** Make sure Docker has the necessary permissions to create volumes
3. **Database not starting:** Check the logs with `docker-compose logs postgres` 