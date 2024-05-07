package models

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"employee_data_management/internal/database"
	"employee_data_management/internal/structs"
)

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	database.Database_Conn = mockDB
	return mockDB, mock
}

func TestCreateEmployee(t *testing.T) {
	mockDB, mock := setupMockDB(t)
	defer mockDB.Close()

	em := EmployeeModel{}

	// Mocking database expectations
	mock.ExpectExec("INSERT INTO employees").WithArgs("John Jana", "Developer", 50000.0).WillReturnResult(sqlmock.NewResult(123, 1))
	mock.ExpectExec("INSERT INTO employees").WithArgs("", "Position", 0.0).WillReturnError(errors.New("Some database error"))

	// Test case 1: Valid employee creation
	id, err := em.CreateEmployee("John Jana", "Developer", 50000.0)
	assert.NoError(t, err)
	assert.Equal(t, int64(123), id)

	// Test case 2: Invalid employee creation
	_, err = em.CreateEmployee("Invalid", "Position", 0.0)
	assert.Error(t, err)
}

func TestGetEmployeeByID(t *testing.T) {
	mockDB, mock := setupMockDB(t)
	defer mockDB.Close()

	em := EmployeeModel{}

	// Test case 1: Valid employee ID
	mock.ExpectQuery("SELECT id, name, position, salary FROM employees").WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "position", "salary"}).AddRow(123, "John Doe", "Developer", 50000.0))
	status, emp, err := em.GetEmployeeByID(123)
	assert.Equal(t, "", status)
	assert.NoError(t, err)
	assert.Equal(t, structs.Employee{Id: 123, Name: "John Doe", Position: "Developer", Salary: 50000.0}, emp)

	// Test case 2: Invalid employee ID
	mock.ExpectQuery("SELECT id, name, position, salary FROM employees").WithArgs(456).WillReturnError(sql.ErrNoRows)
	status, _, err = em.GetEmployeeByID(456)
	assert.Equal(t, "404", status)
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestUpdateEmployee(t *testing.T) {
	mockDB, mock := setupMockDB(t)
	defer mockDB.Close()

	em := EmployeeModel{}

	// Test case 1: Valid employee update
	mock.ExpectQuery("SELECT id FROM employees").WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(123))
	mock.ExpectExec("UPDATE employees").WithArgs(sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(0, 1))
	status, err := em.UpdateEmployee(structs.Employee{Id: 123, Name: "Jane Doe", Position: "Manager", Salary: 60000.0})
	assert.Equal(t, "", status)
	assert.NoError(t, err)

	// Test case 2: Employee ID not found
	mock.ExpectQuery("SELECT id FROM employees").WithArgs(456).WillReturnError(sql.ErrNoRows)
	status, err = em.UpdateEmployee(structs.Employee{Id: 456, Name: "Invalid", Position: "Invalid", Salary: 0})
	assert.Equal(t, "400", status)
	assert.NoError(t, err)

	// Test case 3: Database error
	mock.ExpectQuery("SELECT id FROM employees").WithArgs(789).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(789))
	mock.ExpectExec("UPDATE employees").WithArgs(structs.Employee{Id: 789, Name: "Invalid", Position: "Invalid", Salary: 0}).WillReturnError(errors.New("Some database error"))
	status, err = em.UpdateEmployee(structs.Employee{Id: 789, Name: "Invalid", Position: "Invalid", Salary: 0})
	assert.Error(t, err)
	assert.Equal(t, "", status)
}

func TestDeleteEmployee(t *testing.T) {
	mockDB, mock := setupMockDB(t)
	defer mockDB.Close()

	em := EmployeeModel{}

	// Test case 1: Valid employee deletion
	mock.ExpectQuery("SELECT id FROM employees").WithArgs(123).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(123))
	mock.ExpectExec("DELETE FROM employees").WithArgs(123).WillReturnResult(sqlmock.NewResult(0, 1))
	status, err := em.DeleteEmployee(123)
	assert.NoError(t, err)
	assert.Equal(t, "", status)

	// Test case 2: Employee ID not found
	mock.ExpectQuery("SELECT id FROM employees").WithArgs(456).WillReturnError(sql.ErrNoRows)
	status, err = em.DeleteEmployee(456)
	assert.Equal(t, "404", status)
	assert.NoError(t, err)

	// Test case 3: Database error
	mock.ExpectQuery("SELECT id FROM employees").WithArgs(789).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(789))
	mock.ExpectExec("DELETE FROM employees").WithArgs(789).WillReturnError(errors.New("Some database error"))
	status, err = em.DeleteEmployee(789)
	assert.Error(t, err)
	assert.Equal(t, "", status)
}

func TestListEmployees(t *testing.T) {
	mockDB, mock := setupMockDB(t)
	defer mockDB.Close()

	em := EmployeeModel{}

	// Test case 1: Valid list of employees
	rows := sqlmock.NewRows([]string{"id", "name", "position", "salary"}).
		AddRow(1, "Test Tanmoy", "Developer", 50000.0).
		AddRow(2, "Nitish Rana", "Manager", 60000.0)
	mock.ExpectQuery("SELECT id, name, position, salary FROM employees").WithArgs(0, 10).WillReturnRows(rows)
	employees, err := em.ListEmployees(10, 0)
	assert.NoError(t, err)
	assert.Len(t, employees, 2)
	assert.Equal(t, "Test Tanmoy", employees[0].Name)
	assert.Equal(t, "Developer", employees[0].Position)
	assert.Equal(t, 50000.0, employees[0].Salary)
	assert.Equal(t, "Nitish Rana", employees[1].Name)
	assert.Equal(t, "Manager", employees[1].Position)
	assert.Equal(t, 60000.0, employees[1].Salary)

	// Test case 2: Error querying the database
	mock.ExpectQuery("SELECT id, name, position, salary FROM employees").WithArgs(0, 10).WillReturnError(errors.New("Database error"))
	_, err = em.ListEmployees(10, 0)
	assert.Error(t, err)
}
