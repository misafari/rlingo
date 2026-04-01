package domain

import (
	"time"

	"github.com/google/uuid"
)

type MemberRole string

const (
	RoleViewer MemberRole = "VIEWER"
	RoleEditor MemberRole = "EDITOR"
	RoleAdmin  MemberRole = "ADMIN"
)

type ProjectMember struct {
	ProjectID uuid.UUID
	UserID    uuid.UUID
	Role      MemberRole
	InvitedBy *uuid.UUID
	JoinedAt  time.Time
}

func NewAdminProjectMembership(projectID, userID uuid.UUID) *ProjectMember {
	return &ProjectMember{
		ProjectID: projectID,
		UserID:    userID,
		Role:      RoleAdmin,
		JoinedAt:  time.Now().UTC(),
	}
}
