package models

import (
	"errors"
	"rest_mysql/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type Employee struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email" validate:"required,email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func validate(e *Employee) error {
	if e.ID == 0 {
		return errors.New("id is required")
	}
	if e.Name == "" {
		return errors.New("name is required")
	}
	if e.Email == "" {
		return errors.New("email is required")
	}
	return nil
}
func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Employee{})
}

// GetEmployees returns a list of employees from the database
func GetEmployees() ([]Employee, error) {
	employees := []Employee{}
	result := db.Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}
	return employees, nil
}

func GetEmployee(id int) (*Employee, error) {
	employee := &Employee{}
	result := db.First(&employee, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return employee, nil

}

// CreateEmployee creates an employee in the database
func CreateEmployee(employee *Employee) (*Employee, error) {
	err := validate(employee)
	if err != nil {
		return nil, err
	}
	result := db.Create(&employee)
	if result.Error != nil {
		return nil, result.Error
	}
	return employee, nil
}

// UpdateEmployee updates an existing employee in the database
func UpdateEmployee(id int, updatedEmployee *Employee) (*Employee, error) {
	employee := Employee{}
	result := db.First(&employee, id)
	if result.Error != nil {
		return nil, result.Error
	}
	//updatedEmployee.ID=employee.ID

	updatedEmployee.ID = employee.ID
	result = db.Save(&updatedEmployee)
	if result.Error != nil {
		return nil, result.Error
	}
	return updatedEmployee, nil

}

// DeleteEmployee deletes an employee with given ID from the database
func DeleteEmployee(id int) error {
	employee := Employee{}
	result := db.First(&employee, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return result.Error
		}
	}
	result = db.Delete(&employee)
	if result != nil {
		return result.Error
	}
	return nil
}
