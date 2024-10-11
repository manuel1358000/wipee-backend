package models

import (
	"fmt"
	"time"
	"wipee/lib/dtos"

	"github.com/segmentio/ksuid"
)

// Location representa la estructura de la ubicación para DynamoDB
type Location struct {
	PK                string        `dynamodbav:"PK"`                // Clave de partición
	SK                string        `dynamodbav:"SK"`                // Clave de ordenación
	LocationId        string        `dynamodbav:"locationId"`        // ID de la ubicación
	Name              string        `dynamodbav:"name"`              // Nombre de la ubicación
	Address           string        `dynamodbav:"address"`           // Dirección
	Coordinates       Coordinates   `dynamodbav:"coordinates"`       // Coordenadas
	Photos            []string      `dynamodbav:"photos"`            // Fotos
	Description       string        `dynamodbav:"description"`       // Descripción
	AdditionalNote    string        `dynamodbav:"additionalNote"`    // Notas adicionales
	Price             float64       `dynamodbav:"price"`             // Precio
	Schedule          []DaySchedule `dynamodbav:"schedule"`          // Horarios
	Status            string        `dynamodbav:"status"`            // Estado
	CreatedBy         string        `dynamodbav:"createdBy"`         // ID del usuario que creó la ubicación
	IsDelete          bool          `dynamodbav:"isDelete"`          // Indica si está borrado lógicamente
	CreateDateTime    *time.Time    `dynamodbav:"createDateTime"`    // Fecha de creación
	DeleteDateTime    *time.Time    `dynamodbav:"deleteDateTime"`    // Fecha de eliminación
	ModifyDateTime    *time.Time    `dynamodbav:"modifyDateTime"`    // Fecha de modificación
	Category          string        `dynamodbav:"category"`          // Categoría
	ChangingTableType string        `dynamodbav:"changingTableType"` // Tipo de cambiador
	Amenities         []string      `dynamodbav:"amenities"`         // Amenidades
	PrivacyLevel      string        `dynamodbav:"privacyLevel"`      // Nivel de privacidad
}

// Coordinates representa las coordenadas de una ubicación
type Coordinates struct {
	Latitude  float64 `dynamodbav:"latitude"`  // Latitud
	Longitude float64 `dynamodbav:"longitude"` // Longitud
}

// DaySchedule representa el horario de apertura para cada día de la semana
type DaySchedule struct {
	Day       string    `dynamodbav:"day"`       // Día de la semana
	IsActive  bool      `dynamodbav:"isActive"`  // Indica si está activo ese día
	Morning   TimeRange `dynamodbav:"morning"`   // Horario de la mañana
	Afternoon TimeRange `dynamodbav:"afternoon"` // Horario de la tarde
}

// TimeRange representa un rango de tiempo
type TimeRange struct {
	Start string `dynamodbav:"start"` // Hora de inicio
	End   string `dynamodbav:"end"`   // Hora de fin
}

// newLocation crea una nueva instancia de Location.
func newLocation(locationDto *dtos.LocationsDto) *Location {
	now := time.Now()
	locationId := ksuid.New().String() // Generar el LocationId usando KSUID

	return &Location{
		PK:                NewLocationPK(locationDto.CreatedBy),
		SK:                NewLocationSK(locationId),
		LocationId:        locationId,
		Name:              locationDto.Name,
		Address:           locationDto.Address,
		Coordinates:       convertCoordinatesDtoToModel(locationDto.Coordinates),
		Photos:            locationDto.Photos,
		Description:       locationDto.Description,
		AdditionalNote:    locationDto.AdditionalNote,
		Price:             locationDto.Price,
		Schedule:          convertDaySchedulesDtoToModel(locationDto.Schedule),
		Status:            locationDto.Status,
		CreatedBy:         locationDto.CreatedBy,
		IsDelete:          locationDto.IsDelete,
		CreateDateTime:    &now,
		DeleteDateTime:    locationDto.DeleteDateTime,
		ModifyDateTime:    locationDto.ModifyDateTime,
		Category:          locationDto.Category,
		ChangingTableType: locationDto.ChangingTableType,
		Amenities:         locationDto.Amenities,
		PrivacyLevel:      locationDto.PrivacyLevel,
	}
}

// Convertir un slice de dtos.DaySchedule a un slice de DaySchedule
func convertDaySchedulesDtoToModel(dtoDaySchedules []dtos.DaySchedule) []DaySchedule {
	daySchedules := make([]DaySchedule, len(dtoDaySchedules))
	for i, dto := range dtoDaySchedules {
		daySchedules[i] = DaySchedule{
			Day:       dto.Day,
			IsActive:  dto.IsActive,
			Morning:   convertTimeRangeDtoToModel(dto.Morning),
			Afternoon: convertTimeRangeDtoToModel(dto.Afternoon),
		}
	}
	return daySchedules
}

// convertTimeRangeDtoToModel convierte un DTO TimeRange a un modelo TimeRange
func convertTimeRangeDtoToModel(dtoTimeRange dtos.TimeRange) TimeRange {
	return TimeRange{
		Start: dtoTimeRange.Start,
		End:   dtoTimeRange.End,
	}
}

// convertDtoToModelCoordinates convierte un DTO Coordinates a un modelo Coordinates
func convertCoordinatesDtoToModel(dtoCoords dtos.Coordinates) Coordinates {
	return Coordinates{
		Latitude:  dtoCoords.Latitude,
		Longitude: dtoCoords.Longitude,
	}
}

// LocationsToDto convierte un modelo Location a un DTO LocationsDto.
func (l *Location) LocationsToDto() *dtos.LocationsDto {
	return &dtos.LocationsDto{
		LocationId:        l.LocationId,
		Name:              l.Name,
		Address:           l.Address,
		Coordinates:       convertModelToDtoCoordinates(l.Coordinates),
		Photos:            l.Photos,
		Description:       l.Description,
		AdditionalNote:    l.AdditionalNote,
		Price:             l.Price,
		Schedule:          convertDaySchedulesModelToDto(l.Schedule),
		Status:            l.Status,
		CreatedBy:         l.CreatedBy,
		IsDelete:          l.IsDelete,
		CreateDateTime:    l.CreateDateTime,
		DeleteDateTime:    l.DeleteDateTime,
		ModifyDateTime:    l.ModifyDateTime,
		Category:          l.Category,
		ChangingTableType: l.ChangingTableType,
		Amenities:         l.Amenities,
		PrivacyLevel:      l.PrivacyLevel,
	}
}

// convertModelToDtoCoordinates convierte un modelo Coordinates a un DTO Coordinates
func convertModelToDtoCoordinates(coords Coordinates) dtos.Coordinates {
	return dtos.Coordinates{
		Latitude:  coords.Latitude,
		Longitude: coords.Longitude,
	}
}

// convertDaySchedulesModelToDto convierte un slice de DaySchedule a un slice de dtos.DaySchedule
func convertDaySchedulesModelToDto(daySchedules []DaySchedule) []dtos.DaySchedule {
	dtosDaySchedules := make([]dtos.DaySchedule, len(daySchedules))
	for i, daySchedule := range daySchedules {
		dtosDaySchedules[i] = dtos.DaySchedule{
			Day:       daySchedule.Day,
			IsActive:  daySchedule.IsActive,
			Morning:   convertTimeRange(daySchedule.Morning),
			Afternoon: convertTimeRange(daySchedule.Afternoon),
		}
	}
	return dtosDaySchedules
}

// convertTimeRange convierte un modelo TimeRange a un DTO TimeRange
func convertTimeRange(timeRange TimeRange) dtos.TimeRange {
	return dtos.TimeRange{
		Start: timeRange.Start,
		End:   timeRange.End,
	}
}

// NewLocationPK genera la clave de partición para una ubicación.
func NewLocationPK(userID string) string {
	return fmt.Sprintf("USER#%s", userID)
}

// NewLocationSK genera la clave de ordenación para una ubicación.
func NewLocationSK(locationID string) string {
	return fmt.Sprintf("LOCATION#%s", locationID)
}
