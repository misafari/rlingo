package _const

import "errors"

var (
	ErrInvalidName      = errors.New("project: name must not be empty")
	ErrMissingTenantID  = errors.New("project: tenant ID is required")
	ErrMissingUserID    = errors.New("project: user ID is required")
	ErrMissingProjectID = errors.New("project: ID is required")
	ErrMissingCreatedBy = errors.New("project: created by is required")
	ErrNothingToUpdate  = errors.New("project: at least one field must be provided")
	ErrProjectIsNil     = errors.New("project: project is nil")

	ErrProjectNotFound = errors.New("project: project not found")
	ErrMemberNotFound  = errors.New("project: member not found")

	ErrAlreadyArchived = errors.New("project: project is already archived")
	ErrAlreadyMember   = errors.New("project: user is already a member of this project")

	ErrForbidden = errors.New("project: insufficient permissions for this action")

	ErrUnauthenticated       = errors.New("project: request has no authenticated user")
	ErrProjectCreationFailed = errors.New("project: failed to create project")
	ErrProjectDeletionFailed = errors.New("project: failed to delete project")
	ErrProjectFetchingFailed = errors.New("project: failed to fetch projects")
	ErrProjectUpdateFailed   = errors.New("project: failed to update project")
)
