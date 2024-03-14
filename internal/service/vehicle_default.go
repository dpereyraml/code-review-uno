package service

import "app/internal"

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// AddVehicle is a method that adds a vehicle to the service
func (s *VehicleDefault) AddVehicle(v *internal.Vehicle) (err error) {
	err = s.rp.AddVehicle(v)
	return
}

// GetByColorAndYear is a method that returns vehicles by color and year
func (s *VehicleDefault) GetByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.GetByColorAndYear(color, year)
	return
}
