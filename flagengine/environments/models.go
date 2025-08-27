package environments

import (
	"time"

	"github.com/higordasneves/flagsmith-go-client/flagengine/features"
	"github.com/higordasneves/flagsmith-go-client/flagengine/identities"
	"github.com/higordasneves/flagsmith-go-client/flagengine/projects"
)

type EnvironmentModel struct {
	ID                int                           `json:"id"`
	APIKey            string                        `json:"api_key"`
	Project           *projects.ProjectModel        `json:"project"`
	FeatureStates     []*features.FeatureStateModel `json:"feature_states"`
	IdentityOverrides []*identities.IdentityModel   `json:"identity_overrides"`
	UpdatedAt         time.Time                     `json:"updated_at"`
}
