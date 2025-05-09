This project uses the dependencies:
1. Gin: REST API ENDPOINTS
2. godotenv: .ENV FILES
3. sqlx: DATABASE INTERACTION
4. golang-migrate (CLI TOOL and Dependency): DATABASE MIGRATION
5. go-sql-driver/mysql: MySQL DRIVER
6. golang/x/crpyto: For bcrypt encryption

# Features to be added
1. Reactions
2. Comments
3. Upload files

# For improvement
1. Add logs in repository and service layers
2. Add tests for all modules

# How to run
1. Install golang-migrate for database migration
```go
-- windows
irm get.scoop.sh | iex
scoop install migrate

-- linux
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
2. Create the database
```sql
CREATE DATABASE sma_db;
```

3. Run the database migration file on root project folder
```go
-- syntax
migrate -path migrations -database "mysql://<username>:<password>@tcp(<host>:<port>)/<database_name>" up

-- usage
migrate -path migrations -database "mysql://root:root@tcp(localhost:3307)/sma_db" up
```

4. Download the dependencies and missing dependencies
```go
go mod download
go mod tidy
```

5. Supply the proper environment variables in .env file present in root project
```
PORT=:8000

DB_USERNAME=root
DB_PASSWORD=root
DB_HOST=localhost
DB_PORT=3307
DB_NAME=sma_db

# Generate this with `openssl rand -base64 32`
JWT_SECRET_KEY=7nnoasdFD58zhjVO+GjQLhpRl6ps2x9+bZVfolJJlpI=
```