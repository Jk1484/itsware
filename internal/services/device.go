package services

import (
	"itsware/internal/models"
	"itsware/internal/repositories"
)

type Device struct {
	Repository *repositories.Device
}

func (s *Device) Create(device models.Device) error {
	return s.Repository.Create(device)
}

func (s *Device) Get(id int) (models.Device, error) {
	return s.Repository.Get(id)
}

func (s *Device) Update(device models.Device) error {
	return s.Repository.Update(device)
}

func (s *Device) Delete(id int) error {
	return s.Repository.Delete(id)
}

func (s *Device) GetAll() ([]models.Device, error) {
	return s.Repository.GetAll()
}
