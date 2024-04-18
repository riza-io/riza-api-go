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

// TopLevelService contains methods and other services that help with interacting
// with the riza API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewTopLevelService] method instead.
type TopLevelService struct {
	Options []option.RequestOption
}

// NewTopLevelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTopLevelService(opts ...option.RequestOption) (r *TopLevelService) {
	r = &TopLevelService{}
	r.Options = opts
	return
}

func (r *TopLevelService) Execute(ctx context.Context, body TopLevelExecuteParams, opts ...option.RequestOption) (res *TopLevelExecuteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/execute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TopLevelExecuteResponse struct {
	ExitCode string                      `json:"exitCode"`
	Stderr   string                      `json:"stderr"`
	Stdout   string                      `json:"stdout"`
	JSON     topLevelExecuteResponseJSON `json:"-"`
}

// topLevelExecuteResponseJSON contains the JSON metadata for the struct
// [TopLevelExecuteResponse]
type topLevelExecuteResponseJSON struct {
	ExitCode    apijson.Field
	Stderr      apijson.Field
	Stdout      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TopLevelExecuteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r topLevelExecuteResponseJSON) RawJSON() string {
	return r.raw
}

type TopLevelExecuteParams struct {
	Args     param.Field[[]string]                      `json:"args"`
	Code     param.Field[string]                        `json:"code"`
	Env      param.Field[map[string]string]             `json:"env"`
	Language param.Field[TopLevelExecuteParamsLanguage] `json:"language"`
	Stdin    param.Field[string]                        `json:"stdin"`
}

func (r TopLevelExecuteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TopLevelExecuteParamsLanguage string

const (
	TopLevelExecuteParamsLanguageUnspecified TopLevelExecuteParamsLanguage = "UNSPECIFIED"
	TopLevelExecuteParamsLanguagePython      TopLevelExecuteParamsLanguage = "PYTHON"
	TopLevelExecuteParamsLanguageJavascript  TopLevelExecuteParamsLanguage = "JAVASCRIPT"
	TopLevelExecuteParamsLanguageTypescript  TopLevelExecuteParamsLanguage = "TYPESCRIPT"
	TopLevelExecuteParamsLanguageRuby        TopLevelExecuteParamsLanguage = "RUBY"
	TopLevelExecuteParamsLanguagePhp         TopLevelExecuteParamsLanguage = "PHP"
)

func (r TopLevelExecuteParamsLanguage) IsKnown() bool {
	switch r {
	case TopLevelExecuteParamsLanguageUnspecified, TopLevelExecuteParamsLanguagePython, TopLevelExecuteParamsLanguageJavascript, TopLevelExecuteParamsLanguageTypescript, TopLevelExecuteParamsLanguageRuby, TopLevelExecuteParamsLanguagePhp:
		return true
	}
	return false
}
