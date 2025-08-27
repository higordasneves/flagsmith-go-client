package identities

import (
	"github.com/higordasneves/flagsmith-go-client/flagengine/features"
	"github.com/higordasneves/flagsmith-go-client/flagengine/identities/traits"
	"github.com/higordasneves/flagsmith-go-client/flagengine/utils"
)

type IdentityModel struct {
	DjangoID          int                           `json:"django_id"`
	Identifier        string                        `json:"identifier"`
	EnvironmentAPIKey string                        `json:"environment_api_key"`
	CreatedDate       utils.ISOTime                 `json:"created_date"`
	IdentityUUID      string                        `json:"identity_uuid"`
	IdentityTraits    []*traits.TraitModel          `json:"identity_traits"`
	IdentityFeatures  []*features.FeatureStateModel `json:"identity_features"`
	CompositeKey_     string                        `json:"composite_key"`
}

func (i *IdentityModel) CompositeKey() string {
	if i.CompositeKey_ == "" {
		i.CompositeKey_ = i.EnvironmentAPIKey + "_" + i.Identifier
	}
	return i.CompositeKey_
}
