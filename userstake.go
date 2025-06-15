// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package earnapp

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/BoomchainLabs/release-please/internal/apijson"
	"github.com/BoomchainLabs/release-please/internal/requestconfig"
	"github.com/BoomchainLabs/release-please/option"
	"github.com/BoomchainLabs/release-please/packages/param"
	"github.com/BoomchainLabs/release-please/packages/respjson"
)

// UserStakeService contains methods and other services that help with interacting
// with the earn-app API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserStakeService] method instead.
type UserStakeService struct {
	Options []option.RequestOption
}

// NewUserStakeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserStakeService(opts ...option.RequestOption) (r UserStakeService) {
	r = UserStakeService{}
	r.Options = opts
	return
}

// Stake tokens in a vault
func (r *UserStakeService) New(ctx context.Context, userID int64, body UserStakeNewParams, opts ...option.RequestOption) (res *UserStake, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("users/%v/stakes", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve all staking positions for a specific user
func (r *UserStakeService) List(ctx context.Context, userID int64, opts ...option.RequestOption) (res *[]UserStake, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("users/%v/stakes", userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserStake struct {
	// Unique stake identifier
	ID int64 `json:"id"`
	// Staked amount
	Amount float64 `json:"amount"`
	// Whether rewards are auto-compounded
	AutoCompound bool `json:"autoCompound"`
	// End timestamp
	EndDate time.Time `json:"endDate" format:"date-time"`
	// Lock period in days
	LockPeriod int64 `json:"lockPeriod"`
	// Start timestamp
	StartDate time.Time `json:"startDate" format:"date-time"`
	// Stake status
	//
	// Any of "active", "completed", "claimed".
	Status UserStakeStatus `json:"status"`
	// Total rewards earned
	TotalRewards float64 `json:"totalRewards"`
	// User identifier
	UserID int64        `json:"userId"`
	Vault  StakingVault `json:"vault"`
	// Vault identifier
	VaultID int64 `json:"vaultId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Amount       respjson.Field
		AutoCompound respjson.Field
		EndDate      respjson.Field
		LockPeriod   respjson.Field
		StartDate    respjson.Field
		Status       respjson.Field
		TotalRewards respjson.Field
		UserID       respjson.Field
		Vault        respjson.Field
		VaultID      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserStake) RawJSON() string { return r.JSON.raw }
func (r *UserStake) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Stake status
type UserStakeStatus string

const (
	UserStakeStatusActive    UserStakeStatus = "active"
	UserStakeStatusCompleted UserStakeStatus = "completed"
	UserStakeStatusClaimed   UserStakeStatus = "claimed"
)

type UserStakeNewParams struct {
	// Amount to stake
	Amount float64 `json:"amount,required"`
	// Vault identifier
	VaultID int64 `json:"vaultId,required"`
	// Whether to auto-compound rewards
	AutoCompound param.Opt[bool] `json:"autoCompound,omitzero"`
	// Lock period in days
	LockPeriod param.Opt[int64] `json:"lockPeriod,omitzero"`
	paramObj
}

func (r UserStakeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UserStakeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserStakeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
