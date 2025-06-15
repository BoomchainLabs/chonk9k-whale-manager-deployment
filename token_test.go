// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package earnapp_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/BoomchainLabs/release-please"
	"github.com/BoomchainLabs/release-please/internal/testutil"
	"github.com/BoomchainLabs/release-please/option"
)

func TestTokenGetInfo(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Token.GetInfo(context.TODO())
	if err != nil {
		var apierr *earnapp.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
