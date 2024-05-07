package controllers

import (
	"employee_data_management/internal/models"
	"employee_data_management/internal/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct{}

var EmployeeModel = new(models.EmployeeModel)

func (ectrl EmployeeController) CreateEmployee(c *gin.Context) {
	var emp structs.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if emp.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid name"})
		return
	}
	if emp.Position == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid position"})
		return
	}
	if emp.Salary <= 0.0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid salary"})
		return
	}
	if emp.Id != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not required"})
		return
	}
	id, err := EmployeeModel.CreateEmployee(emp.Name, emp.Position, emp.Salary)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (ectrl EmployeeController) GetEmployeeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}
	status, emp, err := EmployeeModel.GetEmployeeByID(int64(id))
	if status == "404" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid employee ID"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"employee_details": emp})
}

func (ectrl EmployeeController) UpdateEmployee(c *gin.Context) {
	var emp structs.Employee
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if emp.Id != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not required"})
		return
	}
	if emp.Name == "" && emp.Position == "" && emp.Salary == 0.0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "inadequate parameters"})
		return
	}
	if emp.Salary < 0.0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid salary"})
		return
	}
	emp.Id = int64(id)
	status, err := EmployeeModel.UpdateEmployee(emp)
	if status == "400" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (ectrl EmployeeController) DeleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}
	status, err := EmployeeModel.DeleteEmployee(int64(id))
	if status == "404" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid employee ID"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (ectrl EmployeeController) ListEmployees(c *gin.Context) {
	// Pagination parameters
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Page Number"})
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Size"})
		return
	}
	offset := (page - 1) * limit

	employees, err := EmployeeModel.ListEmployees(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"employee_details": employees})
}
