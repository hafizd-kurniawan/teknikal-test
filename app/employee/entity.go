package employee

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	"employee-management/infra/middleware"
)

type Employee struct {
	EmployeeId   int        `json:"id" db:"employee_id" gorm:"primaryKey;autoIncrement"`
	EmployeeCode string     `json:"employee_code" db:"employee_code"`
	EmployeeName string     `json:"employee_name" db:"employee_name"`
	Password     string     `json:"password" db:"password"`
	DepartmentID int        `json:"department_id" db:"department_id"`
	PositionID   int        `json:"position_id" db:"position_id"`
	Superior     int        `json:"superior" db:"superior"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	CreatedBy    string     `json:"created_by" db:"created_by"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	UpdatedBy    string     `json:"updated_by" db:"updated_by"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at" gorm:"index"`
}

func NewCreateEmployee(req CreateEmployeeRequest) Employee {
	emp := Employee{
		EmployeeCode: req.EmployeeCode,
		EmployeeName: req.EmployeeName,
		Password:     req.Password,
		DepartmentID: req.DepartmentID,
		PositionID:   req.PositionID,
		Superior:     req.Superior,
		CreatedAt:    time.Now(),
		CreatedBy:    req.CreatedBy,
		UpdatedAt:    time.Now(),
		UpdatedBy:    req.CreatedBy,
	}
	return emp

}

func NewUpdateRequest(req UpdateEmployeeRequest) Employee {
	return Employee{
		EmployeeId:   req.EmployeeID,
		EmployeeCode: req.EmployeeCode,
		EmployeeName: req.EmployeeName,
		Password:     req.Password,
		DepartmentID: req.DepartmentID,
		PositionID:   req.PositionID,
		Superior:     req.Superior,
		UpdatedAt:    time.Now(),
		UpdatedBy:    req.UpdatedBy,
	}
}

func NewLoginRequest(req LoginRequest) Employee {
	return Employee{
		EmployeeCode: req.EmployeeCode,
		Password:     req.Password,
	}
}
func (e Employee) Validate() error {
	if e.EmployeeCode == "" {
		return errors.New("employee code is required")
	}
	if e.EmployeeName == "" {
		return errors.New("employee name is required")
	}
	if e.Password == "" {
		return errors.New("password is required")
	}
	if len(e.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}

func (e *Employee) EncryptPassword() error {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	e.Password = string(encryptedPass)
	return nil
}

func (e Employee) ValidatePasswordFromPlainText(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (e Employee) GenerateAccessToken() (token string, err error) {
	// token, err = middleware.GenereateJWT2(e.EmployeeName)
	token, err = middleware.GenerateNewJWT(&middleware.Claims{
		EmployeeID:   e.EmployeeId,
		EmployeeCode: e.EmployeeCode,
		EmployeeName: e.EmployeeName,
	})

	if err != nil {
		return "", err
	}

	return token, nil
}
