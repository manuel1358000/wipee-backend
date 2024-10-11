package dtos

type RatingDTO struct {
	UserProfileId string  `json:"userId"`    // ID del usuario que califica
	Score         float64 `json:"score"`     // Puntuación (ej. 1 a 5)
	Comment       string  `json:"comment"`   // Comentario del usuario
	CreatedAt     string  `json:"createdAt"` // Fecha de creación de la calificación
}
