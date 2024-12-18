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

// Run a script in a secure, isolated environment. Scripts can read from `stdin`
// and write to `stdout` or `stderr`. They can access input files, environment
// variables and command line arguments.
func (r *CommandService) Exec(ctx context.Context, body CommandExecParams, opts ...option.RequestOption) (res *CommandExecResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/execute"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CommandExecResponse struct {
	// The exit code returned by the script. Will often be '0' on success and non-zero
	// on failure.
	ExitCode int64 `json:"exit_code"`
	// The contents of 'stderr' after executing the script.
	Stderr string `json:"stderr"`
	// The contents of 'stdout' after executing the script.
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
	// The code to execute.
	Code param.Field[string] `json:"code,required"`
	// List of allowed hosts for HTTP requests.
	AllowHTTPHosts param.Field[[]string] `json:"allow_http_hosts"`
	// List of command line arguments to pass to the script.
	Args param.Field[[]string] `json:"args"`
	// Set of key-value pairs to add to the script's execution environment.
	Env param.Field[map[string]string] `json:"env"`
	// List of input files.
	Files param.Field[[]CommandExecParamsFile] `json:"files"`
	// Configuration for HTTP requests and authentication.
	HTTP param.Field[CommandExecParamsHTTP] `json:"http"`
	// The interpreter to use when executing code.
	Language param.Field[CommandExecParamsLanguage] `json:"language"`
	// Configuration for execution environment limits.
	Limits   param.Field[CommandExecParamsLimits] `json:"limits"`
	Revision param.Field[string]                  `json:"revision"`
	// The runtime to use when executing code. Deprecated in favor of
	// `runtime_revision_id`.
	Runtime param.Field[string] `json:"runtime"`
	// The ID of the runtime revision to use when executing code.
	RuntimeRevisionID param.Field[string] `json:"runtime_revision_id"`
	// Input made available to the script via 'stdin'.
	Stdin param.Field[string] `json:"stdin"`
}

func (r CommandExecParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CommandExecParamsFile struct {
	// The contents of the file.
	Contents param.Field[string] `json:"contents"`
	// The relative path of the file.
	Path param.Field[string] `json:"path"`
}

func (r CommandExecParamsFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for HTTP requests and authentication.
type CommandExecParamsHTTP struct {
	// List of allowed HTTP hosts and associated authentication.
	Allow param.Field[[]CommandExecParamsHTTPAllow] `json:"allow"`
}

func (r CommandExecParamsHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// List of allowed HTTP hosts and associated authentication.
type CommandExecParamsHTTPAllow struct {
	// Authentication configuration for outbound requests to this host.
	Auth param.Field[CommandExecParamsHTTPAllowAuth] `json:"auth"`
	// The hostname to allow.
	Host param.Field[string] `json:"host"`
}

func (r CommandExecParamsHTTPAllow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Authentication configuration for outbound requests to this host.
type CommandExecParamsHTTPAllowAuth struct {
	Basic param.Field[CommandExecParamsHTTPAllowAuthBasic] `json:"basic"`
	// Configuration to add an 'Authorization' header using the 'Bearer' scheme.
	Bearer param.Field[CommandExecParamsHTTPAllowAuthBearer] `json:"bearer"`
	Header param.Field[CommandExecParamsHTTPAllowAuthHeader] `json:"header"`
	Query  param.Field[CommandExecParamsHTTPAllowAuthQuery]  `json:"query"`
}

func (r CommandExecParamsHTTPAllowAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CommandExecParamsHTTPAllowAuthBasic struct {
	Password param.Field[string] `json:"password"`
	UserID   param.Field[string] `json:"user_id"`
}

func (r CommandExecParamsHTTPAllowAuthBasic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration to add an 'Authorization' header using the 'Bearer' scheme.
type CommandExecParamsHTTPAllowAuthBearer struct {
	// The token to set, e.g. 'Authorization: Bearer <token>'.
	Token param.Field[string] `json:"token"`
}

func (r CommandExecParamsHTTPAllowAuthBearer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CommandExecParamsHTTPAllowAuthHeader struct {
	Name  param.Field[string] `json:"name"`
	Value param.Field[string] `json:"value"`
}

func (r CommandExecParamsHTTPAllowAuthHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CommandExecParamsHTTPAllowAuthQuery struct {
	Key   param.Field[string] `json:"key"`
	Value param.Field[string] `json:"value"`
}

func (r CommandExecParamsHTTPAllowAuthQuery) MarshalJSON() (data []byte, err error) {
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

// Configuration for execution environment limits.
type CommandExecParamsLimits struct {
	// The maximum time allowed for execution (in seconds). Default is 30.
	ExecutionTimeout param.Field[int64] `json:"execution_timeout"`
	// The maximum memory allowed for execution (in MiB). Default is 128.
	MemorySize param.Field[int64] `json:"memory_size"`
}

func (r CommandExecParamsLimits) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
