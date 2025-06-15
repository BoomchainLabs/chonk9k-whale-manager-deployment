// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package earnapp_test

import (
	"context"
	"os"
	"testing"

	"github.com/BoomchainLabs/release-please"
	"github.com/BoomchainLabs/release-please/internal/testutil"
	"github.com/BoomchainLabs/release-please/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := earnapp.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	user, err := client.Users.New(context.TODO(), earnapp.UserNewParams{
		WalletAddress: "walletAddress",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", user.ID)
}
