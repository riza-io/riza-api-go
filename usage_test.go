// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package riza_test

import (
	"context"
	"os"
	"testing"

	"github.com/riza-io/riza-api-go"
	"github.com/riza-io/riza-api-go/internal/testutil"
	"github.com/riza-io/riza-api-go/option"
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
	codeExecuteResponse, err := client.Code.Execute(context.TODO(), riza.CodeExecuteParams{})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", codeExecuteResponse.ExitCode)
}
