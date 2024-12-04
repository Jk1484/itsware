package models

type Device struct {
	ID      int    `json:"id"`
	Serial  string `json:"serial"`
	Profile string `json:"profile"`
	Cabinet string `json:"cabinet"`
	Team    string `json:"team"`
}
