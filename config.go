package flagsmith

import (
	"time"
)

const (
	// Number of seconds to wait for a request to
	// complete before terminating the request.
	DefaultTimeout = 10 * time.Second

	// Default base URL for the API.
	DefaultBaseURL = "https://edge.api.flagsmith.com/api/v1/"

	BulkIdentifyMaxCount   = 100
	DefaultRealtimeBaseUrl = "https://realtime.flagsmith.com/"
)

// config contains all configurable Client settings.
type config struct {
	baseURL            string
	timeout            time.Duration
	localEvaluation    bool
	envRefreshInterval time.Duration
	enableAnalytics    bool
	offlineMode        bool
	realtimeBaseUrl    string
	useRealtime        bool
	polling            bool
	userProvidedClient bool

	// localEvaluationPageLimit caps how many pages of the environment-document
	// are fetched per update cycle. 0 means unlimited. Default is 1 (first page only)
	// to preserve existing behaviour.
	localEvaluationPageLimit int

	// localEvaluationMemoryAllocBytes is the maximum total size (in bytes, measured
	// as the marshaled JSON of identity_overrides) that may be accumulated across
	// pages. 0 means no limit. Use multiples of 1024*1024 for MB-based limits.
	localEvaluationMemoryAllocBytes int
}

// defaultConfig returns default configuration.
func defaultConfig() config {
	return config{
		baseURL:                  DefaultBaseURL,
		timeout:                  DefaultTimeout,
		envRefreshInterval:       time.Second * 60,
		realtimeBaseUrl:          DefaultRealtimeBaseUrl,
		userProvidedClient:       false,
		localEvaluationPageLimit: 1,
	}
}
