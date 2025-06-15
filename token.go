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

// TokenService contains methods and other services that help with interacting with
// the earn-app API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenService] method instead.
type TokenService struct {
	Options []option.RequestOption
}

// NewTokenService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTokenService(opts ...option.RequestOption) (r TokenService) {
	r = TokenService{}
	r.Options = opts
	return
}

// Retrieve current information about the $LERF token
func (r *TokenService) GetInfo(ctx context.Context, opts ...option.RequestOption) (res *TokenGetInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "token/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TokenGetInfoResponse struct {
	// Token contract address
	Address string `json:"address"`
	// Circulating token supply
	CirculatingSupply float64 `json:"circulatingSupply"`
	// Token decimals
	Decimals int64 `json:"decimals"`
	// Exchanges where the token is listed
	Listings []TokenGetInfoResponseListing `json:"listings"`
	// Current market capitalization
	MarketCap float64 `json:"marketCap"`
	// Token name
	Name string `json:"name"`
	// Current token price in USD
	Price float64 `json:"price"`
	// 24-hour price change percentage
	PriceChange24h float64 `json:"priceChange24h"`
	// Token symbol
	Symbol string `json:"symbol"`
	// Total token supply
	TotalSupply float64 `json:"totalSupply"`
	// 24-hour trading volume
	Volume24h float64 `json:"volume24h"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address           respjson.Field
		CirculatingSupply respjson.Field
		Decimals          respjson.Field
		Listings          respjson.Field
		MarketCap         respjson.Field
		Name              respjson.Field
		Price             respjson.Field
		PriceChange24h    respjson.Field
		Symbol            respjson.Field
		TotalSupply       respjson.Field
		Volume24h         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenGetInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *TokenGetInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TokenGetInfoResponseListing struct {
	// Exchange name
	Exchange string `json:"exchange"`
	// Trading pair
	Pair string `json:"pair"`
	// Link to exchange
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exchange    respjson.Field
		Pair        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenGetInfoResponseListing) RawJSON() string { return r.JSON.raw }
func (r *TokenGetInfoResponseListing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
