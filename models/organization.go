package models

type Organization struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	City     string `json:"city"`
	State    string `json:"state"`
	Postal   string `json:"postal"`
	Category string `json:"category"`
}
