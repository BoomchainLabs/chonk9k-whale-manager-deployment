// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package earnapp

import (
	"context"
	"net/http"

	"github.com/BoomchainLabs/release-please/internal/apijson"
	"github.com/BoomchainLabs/release-please/internal/requestconfig"
	"github.com/BoomchainLabs/release-please/option"
	"github.com/BoomchainLabs/release-please/packages/respjson"
)

// StatService contains methods and other services that help with interacting with
// the earn-app API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStatService] method instead.
type StatService struct {
	Options []option.RequestOption
}

// NewStatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStatService(opts ...option.RequestOption) (r StatService) {
	r = StatService{}
	r.Options = opts
	return
}

// Retrieve overall platform statistics and metrics
func (r *StatService) Get(ctx context.Context, opts ...option.RequestOption) (res *StatGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "stats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StatGetResponse struct {
	// Number of active missions
	ActiveMissions int64 `json:"activeMissions"`
	// Number of active quests
	ActiveQuests int64 `json:"activeQuests"`
	// Daily active users
	DailyActiveUsers int64 `json:"dailyActiveUsers"`
	// Monthly active users
	MonthlyActiveUsers int64 `json:"monthlyActiveUsers"`
	// Total amount staked across all vaults
	TotalStaked float64 `json:"totalStaked"`
	// Total tokens distributed as rewards
	TotalTokensDistributed float64 `json:"totalTokensDistributed"`
	// Total number of users
	TotalUsers int64 `json:"totalUsers"`
	// Weekly active users
	WeeklyActiveUsers int64 `json:"weeklyActiveUsers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveMissions         respjson.Field
		ActiveQuests           respjson.Field
		DailyActiveUsers       respjson.Field
		MonthlyActiveUsers     respjson.Field
		TotalStaked            respjson.Field
		TotalTokensDistributed respjson.Field
		TotalUsers             respjson.Field
		WeeklyActiveUsers      respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StatGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StatGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
