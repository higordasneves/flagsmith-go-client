package projects

import (
	"github.com/higordasneves/flagsmith-go-client/flagengine/organisations"
	"github.com/higordasneves/flagsmith-go-client/flagengine/segments"
)

type ProjectModel struct {
	ID                int                              `json:"id"`
	Name              string                           `json:"name"`
	HideDisabledFlags bool                             `json:"hide_disabled_flags"`
	Organisation      *organisations.OrganisationModel `json:"organisation"`
	Segments          []*segments.SegmentModel         `json:"segments"`
}
