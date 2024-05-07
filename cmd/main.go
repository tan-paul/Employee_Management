package main

import (
	controllers "employee_data_management/internal/controllers"
	"employee_data_management/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	//Connecting with  mysql
	err := database.ConnectMysqlDB()
	if err != nil {
		panic(err)
	}
	defer database.DisConnectMysqlDB(database.Database_Conn)

	// Creating a gin router with default middleware
	r := gin.Default()

	employee := new(controllers.EmployeeController)
	r.POST("/employees", employee.CreateEmployee)
	r.GET("/employees/:id", employee.GetEmployeeByID)
	r.PATCH("/employees/:id", employee.UpdateEmployee)
	r.DELETE("/employees/:id", employee.DeleteEmployee)
	r.GET("/employees", employee.ListEmployees)

	r.Run(":8080") // listen and serve on port 8080
}
