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

// CodeService contains methods and other services that help with interacting with
// the riza API. Note, unlike clients, this service does not read variables from
// the environment automatically. You should not instantiate this service directly,
// and instead use the [NewCodeService] method instead.
type CodeService struct {
	Options []option.RequestOption
}

// NewCodeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCodeService(opts ...option.RequestOption) (r *CodeService) {
	r = &CodeService{}
	r.Options = opts
	return
}

func (r *CodeService) Execute(ctx context.Context, body CodeExecuteParams, opts ...option.RequestOption) (res *CodeExecuteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/execute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CodeExecuteResponse struct {
	ExitCode int64                   `json:"exit_code"`
	Stderr   string                  `json:"stderr"`
	Stdout   string                  `json:"stdout"`
	JSON     codeExecuteResponseJSON `json:"-"`
}

// codeExecuteResponseJSON contains the JSON metadata for the struct
// [CodeExecuteResponse]
type codeExecuteResponseJSON struct {
	ExitCode    apijson.Field
	Stderr      apijson.Field
	Stdout      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CodeExecuteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r codeExecuteResponseJSON) RawJSON() string {
	return r.raw
}

type CodeExecuteParams struct {
	Args     param.Field[[]string]                  `json:"args"`
	Code     param.Field[string]                    `json:"code"`
	Env      param.Field[map[string]string]         `json:"env"`
	Language param.Field[CodeExecuteParamsLanguage] `json:"language"`
	Stdin    param.Field[string]                    `json:"stdin"`
}

func (r CodeExecuteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CodeExecuteParamsLanguage string

const (
	CodeExecuteParamsLanguageUnspecified CodeExecuteParamsLanguage = "UNSPECIFIED"
	CodeExecuteParamsLanguagePython      CodeExecuteParamsLanguage = "PYTHON"
	CodeExecuteParamsLanguageJavascript  CodeExecuteParamsLanguage = "JAVASCRIPT"
	CodeExecuteParamsLanguageTypescript  CodeExecuteParamsLanguage = "TYPESCRIPT"
	CodeExecuteParamsLanguageRuby        CodeExecuteParamsLanguage = "RUBY"
	CodeExecuteParamsLanguagePhp         CodeExecuteParamsLanguage = "PHP"
)

func (r CodeExecuteParamsLanguage) IsKnown() bool {
	switch r {
	case CodeExecuteParamsLanguageUnspecified, CodeExecuteParamsLanguagePython, CodeExecuteParamsLanguageJavascript, CodeExecuteParamsLanguageTypescript, CodeExecuteParamsLanguageRuby, CodeExecuteParamsLanguagePhp:
		return true
	}
	return false
}
