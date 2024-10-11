package dtos

import "time"

type CommentDto struct {
	CommentId       string    `json:"commentId"`                 // ID único del comentario
	LocationId      string    `json:"locationId"`                // ID de la ubicación asociada
	UserId          string    `json:"userId"`                    // ID del usuario que hizo el comentario
	CommentText     string    `json:"commentText"`               // Texto del comentario
	CreatedAt       time.Time `json:"createdAt"`                 // Fecha de creación
	ParentCommentId *string   `json:"parentCommentId,omitempty"` // ID del comentario padre, si es un comentario en respuesta
}
