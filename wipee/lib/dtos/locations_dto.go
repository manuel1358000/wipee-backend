package dtos

import (
	"time"
)

type LocationsDto struct {
	LocationId        string        `json:"locationId"`                  // ID único de la ubicación (Auto Generated)
	Name              string        `json:"name" validate:"required"`    // Nombre de la ubicación
	Address           string        `json:"address" validate:"required"` // Dirección
	Coordinates       Coordinates   `json:"coordinates,omitempty"`       // Coordenadas
	Photos            []string      `json:"photos,omitempty"`            // Fotos (puedes usar un array de URLs)
	Description       string        `json:"description,omitempty"`       // Descripción
	AdditionalNote    string        `json:"additionalNote,omitempty"`    // Notas adicionales
	Price             float64       `json:"price,omitempty"`             // Precio
	Schedule          []DaySchedule `json:"schedule,omitempty"`          // Horarios
	Status            string        `json:"status,omitempty"`            // Estado (e.g., activo, inactivo)
	CreatedBy         string        `json:"createdBy"`                   // ID del usuario que creó la ubicación
	IsDelete          bool          `json:"isDelete,omitempty"`          // Indica si está borrado lógicamente
	CreateDateTime    *time.Time    `json:"createDateTime,omitempty"`    // Fecha de creación
	DeleteDateTime    *time.Time    `json:"deleteDateTime,omitempty"`    // Fecha de eliminación
	ModifyDateTime    *time.Time    `json:"modifyDateTime,omitempty"`    // Fecha de modificación
	Category          string        `json:"category,omitempty"`          // Categoría
	ChangingTableType string        `json:"changingTableType,omitempty"` // Tipo de cambiador
	Amenities         []string      `json:"amenities,omitempty"`         // Amenidades
	PrivacyLevel      string        `json:"privacyLevel,omitempty"`      // Nivel de privacidad
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type DaySchedule struct {
	Day       string    `json:"day"`       // Día de la semana (e.g., "Monday")
	IsActive  bool      `json:"isActive"`  // Indica si está activo ese día
	Morning   TimeRange `json:"morning"`   // Horario de la mañana
	Afternoon TimeRange `json:"afternoon"` // Horario de la tarde
}

// Estructura para el rango de tiempo
type TimeRange struct {
	Start string `json:"start"` // Hora de inicio (formato HH:MM)
	End   string `json:"end"`   // Hora de fin (formato HH:MM)
}
