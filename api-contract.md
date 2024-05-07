# api contract for Employee Data Management

## sample curl for create api

```
curl --location 'http://localhost:8080/employees' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Rajo singh",
    "position": "Recruiter",
    "salary": 90000
}'
```

## sample o/p

```json
{
    "id": 5
}
```

## sample curl for update api

```
curl --location --request PATCH 'http://localhost:8080/employees/5' \
--header 'Content-Type: application/json' \
--data '{
    "name":"natoar lal",
	"position": "Senior Developer",
	"salary": 190000
}'
```
## sample curl for delete api

```
curl --location --request DELETE 'http://localhost:8080/employees/1'
```

## sample curl for get-id api

```
curl --location 'http://localhost:8080/employees/5'
```

## sample o/p

```json
{
    "employee_details": {
        "Id": 5,
        "Name": "natoar lal",
        "Position": "Senior Developer",
        "Salary": 190000
    }
}
```
## sample curl for get-all api

```
curl --location 'http://localhost:8080/employees?page=1&limit=2'
```

## sample o/p

```json
{
    "employee_details": [
        {
            "Id": 2,
            "Name": "Ronak Patel",
            "Position": "Recruiter",
            "Salary": 90000
        },
        {
            "Id": 3,
            "Name": "Tanmoy Paul",
            "Position": "Recruiter",
            "Salary": 90000
        }
    ]
}
```