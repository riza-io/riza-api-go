// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package temprizaapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/TEMP_riza-api-go"
	"github.com/stainless-sdks/TEMP_riza-api-go/internal/testutil"
	"github.com/stainless-sdks/TEMP_riza-api-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := temprizaapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	v1ExecuteResponse, err := client.V1.Execute(context.TODO(), temprizaapi.V1ExecuteParams{})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", v1ExecuteResponse.ExitCode)
}
