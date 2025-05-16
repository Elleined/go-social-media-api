# This project uses the dependencies:
1. Gin: REST API ENDPOINTS
2. godotenv: .ENV FILES
3. sqlx: DATABASE INTERACTION
4. golang-migrate (CLI TOOL and Dependency): DATABASE MIGRATION
5. go-sql-driver/mysql: MySQL DRIVER
6. golang/x/crpyto: For bcrypt encryption
7. golang-jwt/jwt/v5: For JWT

# Features to be added
1. automatic creation of uploadpath + post and comment
2. rename all pageRequest to request

# For improvement
1. API testing for all modules
2. Unit testing for all modules
3. Integration testing for all modules
4. Add redis caching

# Validations
## User
   - save
     - prevents duplicate email
     - strong password (8 length, 1 uppercase, 1 lowercase, 1 special character, 1 digit)
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
     - checks if post author is the currentuser else cannot be updated (only checks for affected rows)
   - updateContent
     - must have logged in usre
     - new content is required
     - checks if post author is the currentuser else cannot be updated (only checks for affected rows)
   - updateAttachment
     - must have logged in user
     - new attachment is required
     - checks if post author is the currentuser else cannot be updated (only checks for affected rows)
   - deleteById
     - must have logged in user
     - post id is required
     - checks if post author is the currentuser else cannot be deleted (only checks for affected rows)
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
      - checks if logged in user is the author of the comment (only checks for affected rows)
   - updateAttachment
      - must have a logged in user
      - post id is required
      - new attachment is required
      - checks if logged in user is the author of the comment (only checks for affected rows)
   - deleteById
      - must have a logged in user
      - post id is required
      - comment id is required
      - checks if logged in user the author of the comment (only checks for affected rows)
## Comment reaction
- save
     - must have a logged in user
     - post id is required
     - comment id is required
     - emoji id is required
     - returns error if logged in user already reacted to the comment
- getAll
     - must have a logged in user
     - post id is required
     - comment id is required
- getAllByEmoji
     - must have a logged in user
     - post id is required
     - comment id is required
     - emoji id is required
- Update
     - must have a logged in user
     - post id is required
     - comment id is required
     - new emoji id is required
     - returns error if logged in user doesn't already reacted to the comment
- Delete
     - must have a logged in user
     - post id is required
     - comment id is required
     - returns error if logged in user doesn't already reacted to the comment
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
## dev
1. CD to deployment > prod or dev
2. Supply the correct environment variables
3. Run database migration
```
docker compose up -d backend-migration
```
4. Run the file-server-api
```
docker compose up -d file-server
```

## prod
1. CD to deployment > prod or dev
2. Supply the correct environment variables
3. Run database migration
```
docker compose up -d backend-migration
```

4. Run the project
```
docker compose up -d backend
```