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

func TestSandboxExecuteWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := riza.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sandbox.Execute(context.TODO(), riza.SandboxExecuteParams{
		Code:     riza.F("print(\"Hello world!\")"),
		Language: riza.F(riza.SandboxExecuteParamsLanguagePython),
		Args:     riza.F([]string{"string", "string", "string"}),
		Env: riza.F(map[string]string{
			"foo": "string",
		}),
		Stdin: riza.F("string"),
	})
	if err != nil {
		var apierr *riza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
