// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package riza_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/riza-api-go"
	"github.com/stainless-sdks/riza-api-go/internal/testutil"
	"github.com/stainless-sdks/riza-api-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := riza.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	v1ExecuteResponse, err := client.V1.Execute(context.TODO(), riza.V1ExecuteParams{})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", v1ExecuteResponse.ExitCode)
}
