package repositories

import (
	"database/sql"
	"itsware/internal/models"
)

type Device struct {
	DB *sql.DB
}

func (r *Device) Create(device models.Device) error {
	query := `INSERT INTO devices (serial, profile, cabinet, team) VALUES ($1, $2, $3, $4)`
	_, err := r.DB.Exec(query, device.Serial, device.Profile, device.Cabinet, device.Team)
	return err
}

func (r *Device) Get(id int) (models.Device, error) {
	var device models.Device
	query := `SELECT id, serial, profile, cabinet, team FROM devices WHERE id = $1`
	row := r.DB.QueryRow(query, id)
	err := row.Scan(&device.ID, device.Serial, device.Profile, device.Cabinet, device.Team)
	return device, err
}

func (r *Device) Update(device models.Device) error {
	query := `UPDATE devices SET serial = $1, profile = $2, cabinet = $3, team = $4 WHERE id = $5`
	_, err := r.DB.Exec(query, device.Serial, device.Profile, device.Cabinet, device.Team, device.ID)
	return err
}

func (r *Device) Delete(id int) error {
	query := `DELETE FROM devices WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *Device) GetAll() ([]models.Device, error) {
	var devices []models.Device
	rows, err := r.DB.Query("SELECT id, serial, profile, cabinet, team FROM devices FROM devices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var device models.Device
		if err := rows.Scan(&device.ID, &device.Serial, &device.Profile, &device.Cabinet, &device.Team); err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}
