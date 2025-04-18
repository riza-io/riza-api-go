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

func TestToolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.New(context.TODO(), riza.ToolNewParams{
		Code:              riza.F("code"),
		Language:          riza.F(riza.ToolNewParamsLanguagePython),
		Name:              riza.F("name"),
		Description:       riza.F("description"),
		InputSchema:       riza.F[any](map[string]interface{}{}),
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

func TestToolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.Update(
		context.TODO(),
		"id",
		riza.ToolUpdateParams{
			Code:              riza.F("code"),
			Description:       riza.F("description"),
			InputSchema:       riza.F[any](map[string]interface{}{}),
			Language:          riza.F(riza.ToolUpdateParamsLanguagePython),
			Name:              riza.F("name"),
			RuntimeRevisionID: riza.F("runtime_revision_id"),
		},
	)
	if err != nil {
		var apierr *riza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolListWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.List(context.TODO(), riza.ToolListParams{
		Limit:         riza.F(int64(0)),
		StartingAfter: riza.F("starting_after"),
	})
	if err != nil {
		var apierr *riza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolExecWithOptionalParams(t *testing.T) {
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
	_, err := client.Tools.Exec(
		context.TODO(),
		"id",
		riza.ToolExecParams{
			Env: riza.F([]riza.ToolExecParamsEnv{{
				Name:     riza.F("name"),
				SecretID: riza.F("secret_id"),
				Value:    riza.F("value"),
			}}),
			HTTP: riza.F(riza.ToolExecParamsHTTP{
				Allow: riza.F([]riza.ToolExecParamsHTTPAllow{{
					Auth: riza.F(riza.ToolExecParamsHTTPAllowAuth{
						Basic: riza.F(riza.ToolExecParamsHTTPAllowAuthBasic{
							Password: riza.F("password"),
							SecretID: riza.F("secret_id"),
							UserID:   riza.F("user_id"),
						}),
						Bearer: riza.F(riza.ToolExecParamsHTTPAllowAuthBearer{
							Token:    riza.F("token"),
							SecretID: riza.F("secret_id"),
						}),
						Query: riza.F(riza.ToolExecParamsHTTPAllowAuthQuery{
							Key:      riza.F("key"),
							SecretID: riza.F("secret_id"),
							Value:    riza.F("value"),
						}),
					}),
					Host: riza.F("host"),
				}}),
			}),
			Input:      riza.F[any](map[string]interface{}{}),
			RevisionID: riza.F("revision_id"),
		},
	)
	if err != nil {
		var apierr *riza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolGet(t *testing.T) {
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
	_, err := client.Tools.Get(context.TODO(), "id")
	if err != nil {
		var apierr *riza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
