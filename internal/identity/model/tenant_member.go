package model

import (
	"time"

	"github.com/google/uuid"
)

type TenantMemberRole string

const (
	RoleTenantAdmin TenantMemberRole = "TENANT_ADMIN"
	RoleMember      TenantMemberRole = "MEMBER"
)

type TenantMember struct {
	TenantID  uuid.UUID
	UserID    uuid.UUID
	Role      TenantMemberRole
	InvitedBy *uuid.UUID
	JoinedAt  time.Time
}

func NewTenantOwner(tenantID, userID uuid.UUID) *TenantMember {
	return &TenantMember{
		TenantID: tenantID,
		UserID:   userID,
		Role:     RoleTenantAdmin,
		JoinedAt: time.Now().UTC(),
	}
}
