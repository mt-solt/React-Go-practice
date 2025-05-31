package models

type Random struct {
	UUID   string `json:"uuid"`
	Value  int64  `json:"value"`
	UserID string `json:"user_id"`
}
