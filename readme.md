# Mybiiz Backend

### Compiling
1. Install dependencies
```
go get
```

2. Create `.env` file
```
# .env
DB_HOST=localhost
DB_USERNAME=myusername
DB_PASSWORD=mypassword
DB_NAME=mybiiz
JWT_SECRET=
```

3. Create mysql database named `mybiiz`
4. Generate JWT secret using `GET /generate` and put the token to `.env` `JWT_SECRET`

5. Run
```
go build
```

### Running
```
./mybiiz-backend
```
