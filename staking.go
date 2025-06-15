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

// StakingService contains methods and other services that help with interacting
// with the earn-app API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStakingService] method instead.
type StakingService struct {
	Options []option.RequestOption
}

// NewStakingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStakingService(opts ...option.RequestOption) (r StakingService) {
	r = StakingService{}
	r.Options = opts
	return
}

// Retrieve all available staking vaults
func (r *StakingService) ListVaults(ctx context.Context, opts ...option.RequestOption) (res *[]StakingVault, err error) {
	opts = append(r.Options[:], opts...)
	path := "staking/vaults"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type StakingVault struct {
	// Unique vault identifier
	ID int64 `json:"id"`
	// Whether the vault is active
	Active bool `json:"active"`
	// Annual percentage rate
	Apr float64 `json:"apr"`
	// Vault description
	Description string `json:"description"`
	// Minimum lock period in days
	MinLockPeriod int64 `json:"minLockPeriod"`
	// Vault name
	Name string `json:"name"`
	// Rewards offered by this vault
	Rewards []StakingVaultReward `json:"rewards"`
	// Total amount staked in this vault
	TotalStaked float64 `json:"totalStaked"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Active        respjson.Field
		Apr           respjson.Field
		Description   respjson.Field
		MinLockPeriod respjson.Field
		Name          respjson.Field
		Rewards       respjson.Field
		TotalStaked   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StakingVault) RawJSON() string { return r.JSON.raw }
func (r *StakingVault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StakingVaultReward struct {
	// Unique reward identifier
	ID int64 `json:"id"`
	// Reward description
	Description string `json:"description"`
	// When the reward is distributed
	//
	// Any of "daily", "weekly", "monthly", "end".
	Distribution string `json:"distribution"`
	// Reward rate
	Rate float64 `json:"rate"`
	// Type of reward
	//
	// Any of "token", "nft", "boost", "other".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Description  respjson.Field
		Distribution respjson.Field
		Rate         respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StakingVaultReward) RawJSON() string { return r.JSON.raw }
func (r *StakingVaultReward) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
