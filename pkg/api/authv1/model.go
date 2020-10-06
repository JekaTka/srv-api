package authv1

type registerQueryParams struct {
	Limit         int    `json:"limit" query:"limit" default:"100" example:"100"`
	Offset        int    `json:"offset" query:"offset" default:"0" example:"200"`
	Query         string `query:"q" validate:"required"`
}

//easyjson:json
type registrationResponse struct {
	Status string `json:"status" example:"success"`
}
