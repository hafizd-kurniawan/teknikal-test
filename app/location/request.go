package location

type CreateLocationRequest struct {
	LocationName string `json:"location_name"`
	CreatedBy    string `json:"created_by"`
}

type UpdateLocationRequest struct {
	LocationId   int    `json:"id"`
	LocationName string `json:"location_name"`
	UpdatedBy    string `json:"updated_by"`
}
