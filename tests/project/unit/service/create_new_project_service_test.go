package service

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/misafari/rlingo/internal/project"
	_const "github.com/misafari/rlingo/internal/project/const"
	"github.com/misafari/rlingo/internal/project/domain"
	"github.com/misafari/rlingo/tests/project/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewProject(t *testing.T) {
	tests := []struct {
		name           string
		input          *domain.Project
		expectedEntity *domain.Project
		expectedError  error
	}{
		{
			"returns created when project is created",
			&domain.Project{
				ID:          uuid.MustParse("4f7af2c7-6b95-4721-9cc5-c880027f34e5"),
				Name:        "Sample Project",
				Description: "This is a sample project description",
			},
			&domain.Project{
				ID:          uuid.MustParse("4f7af2c7-6b95-4721-9cc5-c880027f34e5"),
				Name:        "Sample Project",
				Description: "This is a sample project description",
			},
			nil,
		},
		{
			"returns error when given project is nil",
			nil,
			nil,
			_const.ErrProjectIsNil,
		},
		{
			"returns error when given project is not valid",
			&domain.Project{
				ID:          uuid.MustParse("4f7af2c7-6b95-4721-9cc5-c880027f34e5"),
				TenantID:    uuid.Nil,
				Name:        "Sample Project",
				Description: "This is a sample project description",
			},
			nil,
			_const.ErrMissingTenantID,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mr := utils.NewMockProjectRepositoryWithCreateFn(tt.expectedEntity, tt.expectedError)
			ps := project.NewProjectService(mr)

			createdProject, err := ps.Create(context.Background(), tt.input)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedEntity, createdProject)
		})
	}
}
