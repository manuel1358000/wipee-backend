package models

import (
	"fmt"
	"time"
	"wipee/lib/dtos"

	"github.com/segmentio/ksuid"
)

// Comment representa un comentario en la aplicación.
type Comment struct {
	CommentId       string    `dynamodbav:"commentId"`                 // ID único del comentario
	LocationId      string    `dynamodbav:"locationId"`                // ID de la ubicación asociada
	UserId          string    `dynamodbav:"userId"`                    // ID del usuario que hizo el comentario
	CommentText     string    `dynamodbav:"commentText"`               // Texto del comentario
	CreatedAt       time.Time `dynamodbav:"createdAt"`                 // Fecha de creación
	ParentCommentId *string   `dynamodbav:"parentCommentId,omitempty"` // ID del comentario padre
	PK              string    `dynamodbav:"pk"`                        // Clave de partición
	SK              string    `dynamodbav:"sk"`                        // Clave de clasificación
}

func createComment(commentDto dtos.CommentDto) Comment {
	commentId := ksuid.New().String()
	return Comment{
		CommentId:       commentId,
		LocationId:      commentDto.LocationId,
		UserId:          commentDto.UserId,
		CommentText:     commentDto.CommentText,
		CreatedAt:       time.Now(),
		ParentCommentId: commentDto.ParentCommentId,
		PK:              "LOCATION#" + commentDto.LocationId, // Clave de partición
		SK:              "COMMENT#" + commentId,              // Clave de clasificación
	}
}
func (c *Comment) CommentsToDto() *dtos.CommentDto {
	return &dtos.CommentDto{
		CommentId:       c.CommentId,
		LocationId:      c.LocationId,
		UserId:          c.UserId,
		CommentText:     c.CommentText,
		CreatedAt:       c.CreatedAt,
		ParentCommentId: c.ParentCommentId,
	}
}

// NewCommentPK genera la clave de partición para un comentario.
func NewCommentPK(userID string) string {
	return fmt.Sprintf("USER#%s", userID) // Asumiendo que el comentario está asociado a un usuario.
}

// NewCommentSK genera la clave de ordenación para un comentario.
func NewCommentSK(commentID string) string {
	return fmt.Sprintf("COMMENT#%s", commentID) // La clave de ordenación se basa en el ID del comentario.
}
