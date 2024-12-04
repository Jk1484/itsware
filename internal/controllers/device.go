package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"itsware/internal/models"
	"itsware/internal/services"
)

type Device struct {
	Service *services.Device
}

func (c *Device) Create(w http.ResponseWriter, r *http.Request) {
	var device models.Device

	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.Service.Create(device); err != nil {
		http.Error(w, "Failed to create device", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *Device) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}

	device, err := c.Service.Get(id)
	if err != nil {
		http.Error(w, "Device not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(device); err != nil {
		http.Error(w, "Failed to encode device", http.StatusInternalServerError)
		return
	}
}

func (c *Device) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}

	var device models.Device
	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	device.ID = id

	if err := c.Service.Update(device); err != nil {
		http.Error(w, "Failed to update device", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Device) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid device ID", http.StatusBadRequest)
		return
	}

	if err := c.Service.Delete(id); err != nil {
		http.Error(w, "Failed to delete device", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Device) GetAll(w http.ResponseWriter, r *http.Request) {
	devices, err := c.Service.GetAll()
	if err != nil {
		http.Error(w, "Failed to retrieve devices", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(devices); err != nil {
		http.Error(w, "Failed to encode devices", http.StatusInternalServerError)
	}
}
