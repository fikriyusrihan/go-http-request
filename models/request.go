package models

type StatusRequest struct {
	Water uint8 `json:"water"`
	Wind  uint8 `json:"wind"`
}

type StatusResponse struct {
	WaterStatus string `json:"water"`
	WindStatus  string `json:"wind"`
}
