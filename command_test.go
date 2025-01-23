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

func TestCommandExecWithOptionalParams(t *testing.T) {
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
	_, err := client.Command.Exec(context.TODO(), riza.CommandExecParams{
		Code:           riza.F("print(\"Hello world!\")"),
		Language:       riza.F(riza.CommandExecParamsLanguagePython),
		AllowHTTPHosts: riza.F([]string{"string"}),
		Args:           riza.F([]string{"string"}),
		Env: riza.F(map[string]string{
			"foo": "string",
		}),
		Files: riza.F([]riza.CommandExecParamsFile{{
			Contents: riza.F("contents"),
			Path:     riza.F("path"),
		}}),
		HTTP: riza.F(riza.CommandExecParamsHTTP{
			Allow: riza.F([]riza.CommandExecParamsHTTPAllow{{
				Auth: riza.F(riza.CommandExecParamsHTTPAllowAuth{
					Basic: riza.F(riza.CommandExecParamsHTTPAllowAuthBasic{
						Password: riza.F("password"),
						UserID:   riza.F("user_id"),
					}),
					Bearer: riza.F(riza.CommandExecParamsHTTPAllowAuthBearer{
						Token: riza.F("token"),
					}),
					Header: riza.F(riza.CommandExecParamsHTTPAllowAuthHeader{
						Name:  riza.F("name"),
						Value: riza.F("value"),
					}),
					Query: riza.F(riza.CommandExecParamsHTTPAllowAuthQuery{
						Key:   riza.F("key"),
						Value: riza.F("value"),
					}),
				}),
				Host: riza.F("host"),
			}}),
		}),
		Limits: riza.F(riza.CommandExecParamsLimits{
			ExecutionTimeout: riza.F(int64(0)),
			MemorySize:       riza.F(int64(0)),
		}),
		RuntimeRevisionID: riza.F("runtime_revision_id"),
		Stdin:             riza.F("stdin"),
	})
	if err != nil {
		var apierr *riza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCommandExecFuncWithOptionalParams(t *testing.T) {
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
	_, err := client.Command.ExecFunc(context.TODO(), riza.CommandExecFuncParams{
		Code:     riza.F("def execute(input): return { \"name\": \"John\", \"executed\": True }"),
		Language: riza.F(riza.CommandExecFuncParamsLanguagePython),
		Env: riza.F(map[string]string{
			"foo": "string",
		}),
		Files: riza.F([]riza.CommandExecFuncParamsFile{{
			Contents: riza.F("contents"),
			Path:     riza.F("path"),
		}}),
		HTTP: riza.F(riza.CommandExecFuncParamsHTTP{
			Allow: riza.F([]riza.CommandExecFuncParamsHTTPAllow{{
				Auth: riza.F(riza.CommandExecFuncParamsHTTPAllowAuth{
					Basic: riza.F(riza.CommandExecFuncParamsHTTPAllowAuthBasic{
						Password: riza.F("password"),
						UserID:   riza.F("user_id"),
					}),
					Bearer: riza.F(riza.CommandExecFuncParamsHTTPAllowAuthBearer{
						Token: riza.F("token"),
					}),
					Header: riza.F(riza.CommandExecFuncParamsHTTPAllowAuthHeader{
						Name:  riza.F("name"),
						Value: riza.F("value"),
					}),
					Query: riza.F(riza.CommandExecFuncParamsHTTPAllowAuthQuery{
						Key:   riza.F("key"),
						Value: riza.F("value"),
					}),
				}),
				Host: riza.F("host"),
			}}),
		}),
		Input: riza.F[any](map[string]interface{}{
			"name": "John",
		}),
		Limits: riza.F(riza.CommandExecFuncParamsLimits{
			ExecutionTimeout: riza.F(int64(0)),
			MemorySize:       riza.F(int64(0)),
		}),
		RuntimeRevisionID: riza.F("runtime_revision_id"),
	})
	if err != nil {
		var apierr *riza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
