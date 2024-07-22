## To migrate database:
``` migrate -path=database/sql_migrations -database "postgresql://postgres:admin@localhost:5432/practice?sslmode=disable" -verbose up```

## To Run server:
```go run main.go```