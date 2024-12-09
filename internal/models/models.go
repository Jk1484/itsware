package models

type Device struct {
	ID        int           `json:"id"`
	CabinetID int           `json:"cabinet"`
	TeamID    int           `json:"team"`
	Status    string        `json:"status"`
	Serial    string        `json:"serial"`
	Profile   DeviceProfile `json:"profile"`
}

type DeviceProfile struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Cabinet struct {
	ID       string `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Location string `json:"location" db:"location"`
}

type Team struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type User struct {
	ID     int    `json:"id"`
	TeamID int    `json:"teamID"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
