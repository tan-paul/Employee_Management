package models

import (
	"database/sql"
	db "employee_data_management/internal/database"
	"employee_data_management/internal/structs"
	"fmt"
)

type EmployeeModel struct{}

func (emod EmployeeModel) CreateEmployee(name, position string, salary float64) (int64, error) {
	res, err := db.Database_Conn.Exec("INSERT INTO employees (name, position, salary) VALUES (?, ?, ?)", name, position, salary)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (emod EmployeeModel) GetEmployeeByID(id int64) (string, structs.Employee, error) {
	var emp structs.Employee
	err := db.Database_Conn.QueryRow("SELECT id, name, position, salary FROM employees WHERE id = ?", id).Scan(&emp.Id, &emp.Name, &emp.Position, &emp.Salary)
	if err != nil {
		if err == sql.ErrNoRows {
			return "404", structs.Employee{}, err
		}
		return "", structs.Employee{}, err
	}
	return "", emp, nil
}
func (emod EmployeeModel) UpdateEmployee(emp structs.Employee) (string, error) {
	var id int
	var flag bool = false
	if err := db.Database_Conn.QueryRow("SELECT id FROM employees WHERE id = ?", emp.Id).Scan(&id); err != nil {
		return "400", nil
	}
	query := "UPDATE employees SET "
	if emp.Name != "" {
		query += "name = '" + emp.Name + "'"
		flag = true
	}
	if emp.Position != "" {
		if flag {
			query += ", position = '" + emp.Position + "'"
		} else {
			query += "position = '" + emp.Position + "'"
			flag = true
		}
	}
	if emp.Salary > 0.0 {
		if flag {
			query += ", salary = " + fmt.Sprintf("%f", emp.Salary)
		} else {
			query += "salary = " + fmt.Sprintf("%f", emp.Salary)
		}
	}
	query += " WHERE id = ?"
	_, err := db.Database_Conn.Exec(query, emp.Id)
	if err != nil {
		return "", err
	}
	return "", nil
}

func (emod EmployeeModel) DeleteEmployee(id int64) (string, error) {
	var db_id int
	if err := db.Database_Conn.QueryRow("SELECT id FROM employees WHERE id = ?", id).Scan(&db_id); err != nil {
		return "404", nil
	}
	if _, err := db.Database_Conn.Exec("DELETE FROM employees WHERE id = ?", id); err != nil {
		return "", err
	}
	return "", nil
}

func (emod EmployeeModel) ListEmployees(limit, offset int) ([]structs.Employee, error) {
	rows, err := db.Database_Conn.Query("SELECT id, name, position, salary FROM employees LIMIT ?, ?", offset, limit)
	if err != nil {
		return []structs.Employee{}, err
	}
	defer rows.Close()

	var employees []structs.Employee
	for rows.Next() {
		var emp structs.Employee
		err := rows.Scan(&emp.Id, &emp.Name, &emp.Position, &emp.Salary)
		if err != nil {
			return []structs.Employee{}, err
		}
		employees = append(employees, emp)
	}
	if err := rows.Err(); err != nil {
		return []structs.Employee{}, err
	}
	return employees, nil
}
