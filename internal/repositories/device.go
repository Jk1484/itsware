package repositories

import (
	"context"
	"itsware/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Device struct {
	DB *pgxpool.Pool
}

func (r *Device) Create(device models.Device) error {
	query := `SELECT create_device($1, $2, $3, $4)`
	_, err := r.DB.Exec(context.Background(), query, device.Serial, device.CabinetID, device.TeamID, device.Profile)
	return err
}

func (r *Device) Get(id int) (models.Device, error) {
	var device models.Device
	query := `SELECT id, serial, profile, cabinet, team FROM devices WHERE id = $1`
	row := r.DB.QueryRow(context.Background(), query, id)
	err := row.Scan(&device.ID, device.Serial, device.Profile, device.CabinetID, device.TeamID)
	return device, err
}

func (r *Device) Update(device models.Device) error {
	query := `UPDATE devices SET serial = $1, profile = $2, cabinet = $3, team = $4 WHERE id = $5`
	_, err := r.DB.Exec(context.Background(), query, device.Serial, device.Profile, device.CabinetID, device.TeamID, device.ID)
	return err
}

func (r *Device) Delete(id int) error {
	query := `DELETE FROM devices WHERE id = $1`
	_, err := r.DB.Exec(context.Background(), query, id)
	return err
}

func (r *Device) GetAll() ([]models.Device, error) {
	var devices []models.Device
	rows, err := r.DB.Query(context.Background(), "SELECT id, serial, profile, cabinet, team FROM devices FROM devices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var device models.Device
		if err := rows.Scan(&device.ID, &device.Serial, &device.Profile, &device.CabinetID, &device.TeamID); err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}
