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

// SandboxService contains methods and other services that help with interacting
// with the riza API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewSandboxService] method instead.
type SandboxService struct {
	Options []option.RequestOption
}

// NewSandboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSandboxService(opts ...option.RequestOption) (r *SandboxService) {
	r = &SandboxService{}
	r.Options = opts
	return
}

// Run a script in a secure, isolated sandbox. Scripts can read from stdin and
// write to stdout or stderr. They can access environment variables and command
// line arguments.
func (r *SandboxService) Execute(ctx context.Context, body SandboxExecuteParams, opts ...option.RequestOption) (res *SandboxExecuteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/execute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SandboxExecuteResponse struct {
	// The exit code returned by the script. Will be `0` on success and non-zero on
	// failure.
	ExitCode int64 `json:"exit_code"`
	// The contents of `stderr` after executing the script.
	Stderr string `json:"stderr"`
	// The contents of `stdout` after executing the script.
	Stdout string                     `json:"stdout"`
	JSON   sandboxExecuteResponseJSON `json:"-"`
}

// sandboxExecuteResponseJSON contains the JSON metadata for the struct
// [SandboxExecuteResponse]
type sandboxExecuteResponseJSON struct {
	ExitCode    apijson.Field
	Stderr      apijson.Field
	Stdout      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SandboxExecuteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sandboxExecuteResponseJSON) RawJSON() string {
	return r.raw
}

type SandboxExecuteParams struct {
	// The code to execute in the sandbox.
	Code param.Field[string] `json:"code,required"`
	// The interpreter to use when executing code.
	Language param.Field[SandboxExecuteParamsLanguage] `json:"language,required"`
	// List of command line arguments to pass to the script.
	Args param.Field[[]string] `json:"args"`
	// Set of key-value pairs to add to the script's execution environment.
	Env param.Field[map[string]string] `json:"env"`
	// Input to pass to the script via `stdin`.
	Stdin param.Field[string] `json:"stdin"`
}

func (r SandboxExecuteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The interpreter to use when executing code.
type SandboxExecuteParamsLanguage string

const (
	SandboxExecuteParamsLanguagePython     SandboxExecuteParamsLanguage = "PYTHON"
	SandboxExecuteParamsLanguageJavascript SandboxExecuteParamsLanguage = "JAVASCRIPT"
	SandboxExecuteParamsLanguageTypescript SandboxExecuteParamsLanguage = "TYPESCRIPT"
	SandboxExecuteParamsLanguageRuby       SandboxExecuteParamsLanguage = "RUBY"
	SandboxExecuteParamsLanguagePhp        SandboxExecuteParamsLanguage = "PHP"
)

func (r SandboxExecuteParamsLanguage) IsKnown() bool {
	switch r {
	case SandboxExecuteParamsLanguagePython, SandboxExecuteParamsLanguageJavascript, SandboxExecuteParamsLanguageTypescript, SandboxExecuteParamsLanguageRuby, SandboxExecuteParamsLanguagePhp:
		return true
	}
	return false
}
