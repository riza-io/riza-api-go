// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package riza_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/riza-io/riza-api-go"
	"github.com/riza-io/riza-api-go/internal/testutil"
	"github.com/riza-io/riza-api-go/option"
)

func TestCodeExecuteWithOptionalParams(t *testing.T) {
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
	_, err := client.Code.Execute(context.TODO(), riza.CodeExecuteParams{
		Args: riza.F([]string{"string", "string", "string"}),
		Code: riza.F("string"),
		Env: riza.F(map[string]string{
			"foo": "string",
		}),
		Language: riza.F(riza.CodeExecuteParamsLanguageUnspecified),
		Stdin:    riza.F("string"),
	})
	if err != nil {
		var apierr *riza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
