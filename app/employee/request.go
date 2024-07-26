package employee

// type CreateEmployeeRequest struct {
// 	EmployeeName string `json:"employee_name" validate:"required"`
// 	Password     string `json:"password" validate:"required"`
// 	DepartmentId int    `json:"department_id" validate:"required"`
// 	PositionId   int    `json:"position_id" validate:"required"`
// 	Superior     int    `json:"superior"`
// }

// type UpdateEmployeeRequest struct {
// 	EmployeeId   int64  `json:"employee_id" validate:"required"`
// 	DepartmentId int64  `json:"department_id" validate:"required"`
// 	PositionId   int64  `json:"position_id" validate:"required"`
// 	EmployeeName string `json:"employee_name" validate:"required"`
// }

// type ForgotPasswordRequest struct {
// 	EmployeeCode string `json:"employee_code" validate:"required"`
// 	Password     string `json:"password" validate:"required"`
// }

type LoginRequest struct {
	EmployeeCode string `json:"employee_code"`
	Password     string `json:"password"`
}

type CreateEmployeeRequest struct {
	EmployeeCode string `json:"employee_code"`
	EmployeeName string `json:"employee_name"`
	Password     string `json:"password"`
	DepartmentID int    `json:"department_id"`
	PositionID   int    `json:"position_id"`
	Superior     int    `json:"superior"`
	CreatedBy    string `json:"created_by"`
}

type UpdateEmployeeRequest struct {
	EmployeeID   int    `json:"id"`
	EmployeeCode string `json:"employee_code"`
	EmployeeName string `json:"employee_name"`
	Password     string `json:"password"`
	DepartmentID int    `json:"department_id"`
	PositionID   int    `json:"position_id"`
	Superior     int    `json:"superior"`
	UpdatedBy    string `json:"updated_by"`
}
