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

// CommandService contains methods and other services that help with interacting
// with the riza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommandService] method instead.
type CommandService struct {
	Options []option.RequestOption
}

// NewCommandService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCommandService(opts ...option.RequestOption) (r *CommandService) {
	r = &CommandService{}
	r.Options = opts
	return
}

// Run a script in a secure, isolated sandbox. Scripts can read from stdin and
// write to stdout or stderr. They can access environment variables and command
// line arguments.
func (r *CommandService) Exec(ctx context.Context, body CommandExecParams, opts ...option.RequestOption) (res *CommandExecResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/execute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CommandExecResponse struct {
	// The exit code returned by the script. Will be `0` on success and non-zero on
	// failure.
	ExitCode int64 `json:"exit_code"`
	// The contents of `stderr` after executing the script.
	Stderr string `json:"stderr"`
	// The contents of `stdout` after executing the script.
	Stdout string                  `json:"stdout"`
	JSON   commandExecResponseJSON `json:"-"`
}

// commandExecResponseJSON contains the JSON metadata for the struct
// [CommandExecResponse]
type commandExecResponseJSON struct {
	ExitCode    apijson.Field
	Stderr      apijson.Field
	Stdout      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommandExecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commandExecResponseJSON) RawJSON() string {
	return r.raw
}

type CommandExecParams struct {
	// The code to execute in the sandbox.
	Code param.Field[string] `json:"code,required"`
	// The interpreter to use when executing code.
	Language param.Field[CommandExecParamsLanguage] `json:"language,required"`
	// List of command line arguments to pass to the script.
	Args param.Field[[]string] `json:"args"`
	// Set of key-value pairs to add to the script's execution environment.
	Env param.Field[map[string]string] `json:"env"`
	// List of allowed hosts for HTTP requests
	Net param.Field[[]string] `json:"net"`
	// Input to pass to the script via `stdin`.
	Stdin param.Field[string] `json:"stdin"`
}

func (r CommandExecParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The interpreter to use when executing code.
type CommandExecParamsLanguage string

const (
	CommandExecParamsLanguagePython     CommandExecParamsLanguage = "PYTHON"
	CommandExecParamsLanguageJavascript CommandExecParamsLanguage = "JAVASCRIPT"
	CommandExecParamsLanguageTypescript CommandExecParamsLanguage = "TYPESCRIPT"
	CommandExecParamsLanguageRuby       CommandExecParamsLanguage = "RUBY"
	CommandExecParamsLanguagePhp        CommandExecParamsLanguage = "PHP"
)

func (r CommandExecParamsLanguage) IsKnown() bool {
	switch r {
	case CommandExecParamsLanguagePython, CommandExecParamsLanguageJavascript, CommandExecParamsLanguageTypescript, CommandExecParamsLanguageRuby, CommandExecParamsLanguagePhp:
		return true
	}
	return false
}
