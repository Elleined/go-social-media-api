# System Features
## Main features
1. CRUD of posts
2. CRUD of post reactions
3. CRUD of comments
4. CRUD of comment reactions
5. CRUD of emoji
6. CRUD of provider type
7. CRUD of users

## Special features
1. Robust pagination
2. Social login (Google, Microsoft, and Facebook) and local login [Documented here](https://github.com/Elleined/go-social-media-api/issues/2)
3. Refresh token for 1 week [Refresh token feature](https://github.com/Elleined/security-project?tab=readme-ov-file#refresh-token)
4. Applied access token for 15 minutes as middleware in protected routes
5. Upload, Delete, and Reading attachments using [go-file-server-api](https://github.com/Elleined/go-file-server-api)

# How to run
## dev
1. CD to deployment > dev
2. Supply the correct environment variables in (./deployment/dev/.env) and local project .env
3. Run these command it will run the following:
   - file-server
   - redis-server
   - mysql-server
```
docker compose up -d dev-migration
```
4. Create post folder for post attachments
```
http://localhost:8090/folders/post
```
5. Create comment folder for comment attachments
```
http://localhost:8090/folders/comment
```
6. Create comment folder for comment attachments
```
http://localhost:8090/folders/user
```
7. Add GIN_MODE=debug to IDE environment variable (important!)
8. Run the local project

## prod
1. CD to deployment > prod
2. Supply the correct environment variables
3. Run these command it will run the following:
   - file-server
   - redis-server
   - mysql-server
   - backend
```
docker compose up -d migration
```
4. Create post folder for post attachments
```
http://localhost:8090/folders/post
```
5. Create comment folder for comment attachments
```
http://localhost:8090/folders/comment
```
6. Create comment folder for comment attachments
```
http://localhost:8090/folders/user
```