package handler

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// AddVehicle is a method that returns a handler for the route POST /vehicles
func (h *VehicleDefault) AddVehicle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
					Respuestas:
			201 Created: Vehículo creado exitosamente.
			400 Bad Request: Datos del vehículo mal formados o incompletos.
			409 Conflict: Identificador del vehículo ya existente.

		*/
		// request
		// read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body 1"})
			return
		}

		// - validate request body parse into map[string]any
		bodyMap := make(map[string]any)
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body Unmarshal"})
			return
		}

		// validate requerid fields
		/* if err := tools.CheckFieldExistance(bodyMap, "name", "quantity", "code_value", "expiration"); err != nil {
			var FieldError *tools.FieldError
			if errors.As(err, &FieldError) {
				response.JSON(w, http.StatusInternalServerError, map[string]any{"error": fmt.Sprintf("%s is required", FieldError.Field)})
				return
			}
		} */

		var bodyMapRequest VehicleJSON
		if err := json.Unmarshal(bytes, &bodyMapRequest); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body Unmarshal"})
			return
		}

		// process
		// - create a new vehicle
		vehicle := internal.Vehicle{
			// Id: bodyMapRequest.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           bodyMapRequest.Brand,
				Model:           bodyMapRequest.Model,
				Registration:    bodyMapRequest.Registration,
				Color:           bodyMapRequest.Color,
				FabricationYear: bodyMapRequest.FabricationYear,
				Capacity:        bodyMapRequest.Capacity,
				MaxSpeed:        bodyMapRequest.MaxSpeed,
				FuelType:        bodyMapRequest.FuelType,
				Transmission:    bodyMapRequest.Transmission,
				Weight:          bodyMapRequest.Weight,
				Dimensions: internal.Dimensions{
					Height: bodyMapRequest.Height,
					Length: bodyMapRequest.Length,
					Width:  bodyMapRequest.Width,
				},
			},
		}

		// - add the vehicle
		if err := h.sv.AddVehicle(&vehicle); err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "success",
			"data":    vehicle,
		})
	}
}

// GetByColorAndYear is a method that returns a handler for the route GET /vehicles/color/:color/year/:year
func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - get color and year from the URL
		// two disctint ways to get the parameters
		color := chi.URLParam(r, "color")
		// color := r.URL.Query().Get("color")

		// ------------------------------------------------------------------------------------
		year, err := strconv.Atoi(chi.URLParam(r, "year"))
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"error": "invalid year"})
			return
		}
		// ------------------------------------------------------------------------------------
		// process
		// - get vehicles by color and year
		v, err := h.sv.GetByColorAndYear(color, year)
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}
		fmt.Println("color:", color, "year:", year, "v:", v)
		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
