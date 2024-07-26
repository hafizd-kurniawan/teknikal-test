package position

type CreatePositionRequest struct {
	DepartmentID int    `json:"department_id"`
	PositionName string `json:"position_name"`
	CreatedBy    string `json:"created_by"`
}

type UpdatePositionRequest struct {
	PositionID   int    `json:"id"`
	DepartmentID int    `json:"department_id"`
	PositionName string `json:"position_name"`
	UpdatedBy    string `json:"updated_by"`
}
