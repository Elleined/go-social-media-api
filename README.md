This project uses the dependencies:
1. Gin: REST API ENDPOINTS
2. godotenv: .ENV FILES
3. sqlx: DATABASE INTERACTION
4. golang-migrate (CLI TOOL and Dependency): DATABASE MIGRATION
5. go-sql-driver/mysql: MySQL DRIVER
6. golang/x/crpyto: For bcrypt encryption

# Features to be added
1. Comment Reactions
3. Upload files with go-file-server-api project

# For improvement
1. Add logs in repository and service layers
2. Unit testing for all modules
3. Integration testing for all modules
4. Proper error messages
5. Implement robust pagination
7. Add pessimistic/ preventive validations
   - has()
   - isAlreadyExists()
   - exists()

# Validations
## User
   - save
     - prevents duplicate email
     - strong password (8 length, 1 uppercase, 1 lowercase, 1 special character)
   - getAll
     - page is required defaults to 1
     - pageSize is required defaults to 10
     - isActive is required defaults to true
   - updateStatus
     - user id is required
     - new status is required (true or false)
   - updatePassword
     - user id is required
     - new password is required
   - login
     - username is required
     - password is required
## Post
   - save
     - must have a logged in user
     - subject is required
     - content is required
   - get all
     - page size is required defaults to 10
     - page is required defaults to 1
     - isDeleted is required defaults to false (not deleted)
     - must have a logged in user
   - updateSubject
     - must have a logged in user
     - new subject is required
     - checks if post author is the currentuser else cannot be updated
   - updateContent
     - must have logged in usre
     - new content is required
     - checks if post author is the currentuser else cannot be updated
   - updateAttachment
     - must have logged in user
     - new attachment is required
     - checks if post author is the currentuser else cannot be updated
   - deleteById
     - must have logged in user
     - post id is required
     - checks if post author is the currentuser else cannot be deleted
   - getAllByUser
     - must have a logged in user
     - page is required defaults to 1
     - pageSize is required defaults to 10
     - isDeleted is required defaults to false (not deleted)
## Post reaction
   - save
     - must have a logged in user
     - post id is required
     - emoji id is required
     - returns error if logged in user already reacted to the post
   - getAll
      - must have a logged in user
      - post id is required
   - getAllByEmoji
      - must have a logged in user
      - post id is required
      - emoji id is required
   - Update
      - must have a logged in user
      - post id is required
      - new emoji id is required
      - returns error if logged in user doesn't already reacted to the post
   - Delete
      - must have a logged in user
      - post id is required
      - returns error if logged in user doesn't already reacted to the post
## Comment
   - save
     - must have a logged in user
     - post id is required
     - content is required
   - getAll
     - must have a logged in user
     - post id is required
     - isDeleted is required defaults to false (not deleted)
   - updateContent
      - must have a logged in user
      - post id is required
      - new content is required
      - checks if post has the comment
      - checks if logged in user is the author of the comment
   - updateAttachment
      - must have a logged in user
      - post id is required
      - new attachment is required
      - checks if post has the comment
      - checks if logged in user is the author of the comment
   - deleteById
      - must have a logged in user
      - post id is required
      - comment id is required
      - checks if post has the comment
      - checks if logged in user the author of the comment
## Emoji
   - save
     - must have a logged in user
     - prevent duplicate name
   - update
     - must have a logged in user
     - prevents duplicate name
     - cannot update if already been used as FK
   - delete
     - must have a logged in user
     - cannot delete if already been used as FK
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

4. Supply the proper environment variables in .env file present in root project
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

5. Add `GIN_MODE=VALUE` release or debug in IDE environment or host machine.
    - The default mode is release.
    - When GIN_MODE value is not supplied or set to release it will be release mode
    - Should explicitly declare in debug mode when in development

6. Run the project