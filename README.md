# Employee_Data_Management

### Requirement :- 
Design and implement (with tests) a _employee database management_ using Go programming
language with Gin framework and Mysql as Data Base.

### Tools or frameworks used:-
    Mysql
    Go
    Gin

### Details :- 

#### Schema :- 
```mysql
+---------------------+
| Database            |
+---------------------+
| employee_management |
| information_schema  |
| mysql               |
| performance_schema  |
| sys                 |
| zocket_db           |
+---------------------+
```

#### Employee Table :- 

```mysql
+----------+---------------+------+-----+---------+----------------+
| Field    | Type          | Null | Key | Default | Extra          |
+----------+---------------+------+-----+---------+----------------+
| id       | bigint        | NO   | PRI | NULL    | auto_increment |
| name     | varchar(255)  | NO   |     | NULL    |                |
| position | varchar(255)  | NO   |     | NULL    |                |
| salary   | decimal(10,2) | NO   |     | NULL    |                |
+----------+---------------+------+-----+---------+----------------+
```

#### How to Run?

>To run Api use  **go run cmd/main.go** and to terminate use **control + c**

> To run test cases use **go test ./internal/models/.**

>To check coverage use **go test -cover ./internal/models/.**

>To generate coverage profile file after run test with coverage use **go test -coverprofile=coverage.out ./internal/models/.**

>To analyze the coverage profile file and to and generate a detailed HTML report use **go tool cover -html=coverage.out**



**Note :-** 
    It is just a simplefied version developed and there can be a lot of modifications or enhancements possible as per the requirment.