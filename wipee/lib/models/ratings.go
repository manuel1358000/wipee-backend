package models

import (
	"fmt"
	"wipee/lib/dtos"
)

type Rating struct {
	UserProfileId string  `dynamodbav:"userId"`    // ID del usuario que califica
	Score         float64 `dynamodbav:"score"`     // Puntuación (ej. 1 a 5)
	Comment       string  `dynamodbav:"comment"`   // Comentario del usuario
	CreatedAt     string  `dynamodbav:"createdAt"` // Fecha de creación de la calificación
	PK            string  `dynamodbav:"pk"`        // Clave de partición
	SK            string  `dynamodbav:"sk"`        // Clave de ordenación
}

func NewRating(dto *dtos.RatingDTO) *Rating {
	return &Rating{
		UserProfileId: dto.UserProfileId,
		Score:         dto.Score,
		Comment:       dto.Comment,
		CreatedAt:     dto.CreatedAt,
		PK:            NewRatingPK(dto.UserProfileId),              // Asigna la clave de partición
		SK:            NewRatingSK(fmt.Sprintf("%.2f", dto.Score)), // Asigna la clave de ordenación
	}
}

func (r *Rating) RatingToDto() *dtos.RatingDTO {
	return &dtos.RatingDTO{
		UserProfileId: r.UserProfileId,
		Score:         r.Score,
		Comment:       r.Comment,
		CreatedAt:     r.CreatedAt,
	}
}

// NewRatingPK genera la clave de partición para una calificación.
func NewRatingPK(userID string) string {
	return fmt.Sprintf("USER#%s", userID) // La clave de partición se basa en el ID del usuario.
}

// NewRatingSK genera la clave de ordenación para una calificación.
func NewRatingSK(score string) string {
	return fmt.Sprintf("RATING#%s", score) // La clave de ordenación se basa en el puntaje.
}
