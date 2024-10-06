package dtos

import (
	"time"
)

type UserProfileDto struct {
	UserProfileId    string     `json:"userProfileId"`                       // Auto Generated
	UserProfileName  string     `json:"userProfileName" validate:"required"` // e.g., Elderberry Syrup
	UserProfileEmail string     `json:"userProfileEmail" validate:"required"`
	Status           string     `json:"status,omitempty"`
	IsDelete         bool       `json:"isDelete,omitempty"`
	Country          string     `json:"country,omitempty"`
	PathProfilePhoto string     `json:"pathProfilePhoto,omitempty"`
	BirthDateTime    *time.Time `json:"birthDateTime,omitempty" validate:"required"`
	CreateDateTime   *time.Time `json:"createDateTime,omitempty"`
	DeleteDateTime   *time.Time `json:"deleteDateTime,omitempty"`
	ModifyDateTime   *time.Time `json:"modifyDateTime,omitempty"`
}
