package domain

import (
	"testing"

	"github.com/google/uuid"
	_const "github.com/misafari/rlingo/internal/project/const"
	"github.com/misafari/rlingo/internal/project/domain"
	"github.com/misafari/rlingo/tests/project/utils"
	"github.com/stretchr/testify/assert"
)

func TestProjectValidation(t *testing.T) {
	tests := []struct {
		name          string
		input         *domain.Project
		expectedError error
	}{
		{
			"returns nil error when project is valid",
			utils.GenerateProject(),
			nil,
		},
		{
			"returns error when project id is not valid",
			utils.GenerateProject(utils.WithID(uuid.Nil)),
			_const.ErrMissingProjectID,
		},
		{
			"returns error when project tenant id is not valid",
			utils.GenerateProject(utils.WithTenantID(uuid.Nil)),
			_const.ErrMissingTenantID,
		},
		{
			"returns error when project name is not valid (empty string)",
			utils.GenerateProject(utils.WithName("")),
			_const.ErrInvalidName,
		},
		{
			"returns error when project name is not valid (blank string)",
			utils.GenerateProject(utils.WithName("    ")),
			_const.ErrInvalidName,
		},
		{
			"returns error when project created by id is not valid",
			utils.GenerateProject(utils.WithCreatedById(uuid.Nil)),
			_const.ErrMissingCreatedBy,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Validate()
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
