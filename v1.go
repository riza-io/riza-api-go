// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package riza

import (
	"context"
	"net/http"

	"github.com/riza-io/riza-api-go/internal/apijson"
	"github.com/riza-io/riza-api-go/internal/param"
	"github.com/riza-io/riza-api-go/internal/requestconfig"
	"github.com/riza-io/riza-api-go/option"
)

// V1Service contains methods and other services that help with interacting with
// the riza API. Note, unlike clients, this service does not read variables from
// the environment automatically. You should not instantiate this service directly,
// and instead use the [NewV1Service] method instead.
type V1Service struct {
	Options []option.RequestOption
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r *V1Service) {
	r = &V1Service{}
	r.Options = opts
	return
}

func (r *V1Service) Execute(ctx context.Context, body V1ExecuteParams, opts ...option.RequestOption) (res *V1ExecuteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/execute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1ExecuteResponse struct {
	ExitCode string                `json:"exitCode"`
	Stderr   string                `json:"stderr"`
	Stdout   string                `json:"stdout"`
	JSON     v1ExecuteResponseJSON `json:"-"`
}

// v1ExecuteResponseJSON contains the JSON metadata for the struct
// [V1ExecuteResponse]
type v1ExecuteResponseJSON struct {
	ExitCode    apijson.Field
	Stderr      apijson.Field
	Stdout      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V1ExecuteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v1ExecuteResponseJSON) RawJSON() string {
	return r.raw
}

type V1ExecuteParams struct {
	Args     param.Field[[]string]                `json:"args"`
	Code     param.Field[string]                  `json:"code"`
	Env      param.Field[map[string]string]       `json:"env"`
	Language param.Field[V1ExecuteParamsLanguage] `json:"language"`
	Stdin    param.Field[string]                  `json:"stdin"`
}

func (r V1ExecuteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V1ExecuteParamsLanguage string

const (
	V1ExecuteParamsLanguageUnspecified V1ExecuteParamsLanguage = "UNSPECIFIED"
	V1ExecuteParamsLanguagePython      V1ExecuteParamsLanguage = "PYTHON"
	V1ExecuteParamsLanguageJavascript  V1ExecuteParamsLanguage = "JAVASCRIPT"
	V1ExecuteParamsLanguageTypescript  V1ExecuteParamsLanguage = "TYPESCRIPT"
	V1ExecuteParamsLanguageRuby        V1ExecuteParamsLanguage = "RUBY"
	V1ExecuteParamsLanguagePhp         V1ExecuteParamsLanguage = "PHP"
)

func (r V1ExecuteParamsLanguage) IsKnown() bool {
	switch r {
	case V1ExecuteParamsLanguageUnspecified, V1ExecuteParamsLanguagePython, V1ExecuteParamsLanguageJavascript, V1ExecuteParamsLanguageTypescript, V1ExecuteParamsLanguageRuby, V1ExecuteParamsLanguagePhp:
		return true
	}
	return false
}
