package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// AddVehicle is a method that adds a vehicle to the service
	AddVehicle(v *Vehicle) (err error)
}
