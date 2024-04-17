// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package temprizaapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/TEMP_riza-api-go"
	"github.com/stainless-sdks/TEMP_riza-api-go/internal/testutil"
	"github.com/stainless-sdks/TEMP_riza-api-go/option"
)

func TestV1ExecuteWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Execute(context.TODO(), temprizaapi.V1ExecuteParams{
		Args: temprizaapi.F([]string{"string", "string", "string"}),
		Code: temprizaapi.F("string"),
		Env: temprizaapi.F(map[string]string{
			"foo": "string",
		}),
		Language: temprizaapi.F(temprizaapi.V1ExecuteParamsLanguageUnspecified),
		Stdin:    temprizaapi.F("string"),
	})
	if err != nil {
		var apierr *temprizaapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
