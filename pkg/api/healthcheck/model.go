package healthcheck

type pingResponse struct {
	// Ping response message
	Ping string `json:"ping" example:"pong"`
}
