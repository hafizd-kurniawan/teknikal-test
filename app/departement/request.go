package departement

type CreateDepartmentRequest struct {
	DepartmentName string `json:"department_name"`
	CreatedBy      string `json:"created_by"`
}

type UpdateDepartmentRequest struct {
	DepartmentID   int    `json:"id"`
	DepartmentName string `json:"department_name"`
	UpdatedBy      string `json:"updated_by"`
}
