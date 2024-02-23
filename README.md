
#Go JWT Authentication

external libraries used:
   1. GORM - https://gorm.io/
   2. Gin - https://gin-gonic.com/
   3. Dotenv - https://github.com/joho/godotenv
   4. Bcrypt - https://pkg.go.dev/golang.org/x/crypt...
   5. JWT-Go - https://github.com/golang-jwt/jwt
   6. Compile Daemon - https://github.com/githubnemo/Compile...

```
go get ./
```
```
go install github.com/githubnemo/CompileDaemon
```
To run using compile daemon:
```
compiledaemon --command="./go-jwt"
```

SIGNUP
{
"email":"",
"password":""
}
```
http://localhost:3000/signup
http://localhost:3000/login
http://localhost:3000/validate
```
