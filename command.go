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

// Run a function in a secure, isolated environment. Define a function named
// `execute`. The function will be passed `input` as an object.
func (r *CommandService) ExecFunc(ctx context.Context, body CommandExecFuncParams, opts ...option.RequestOption) (res *CommandExecFuncResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/execute-function"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CommandExecResponse struct {
	// The exit code returned by the script. Will often be '0' on success and non-zero
	// on failure.
	ExitCode int64 `json:"exit_code,required"`
	// The contents of 'stderr' after executing the script.
	Stderr string `json:"stderr,required"`
	// The contents of 'stdout' after executing the script.
	Stdout string                  `json:"stdout,required"`
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

type CommandExecFuncResponse struct {
	Execution CommandExecFuncResponseExecution `json:"execution,required"`
	// The output of the function.
	Output interface{} `json:"output,required"`
	// The status of the output. "valid" means your function executed successfully and
	// returned a valid JSON-serializable object, or void. "json_serialization_error"
	// means your function executed successfully, but returned a nonserializable
	// object. "error" means your function failed to execute.
	OutputStatus CommandExecFuncResponseOutputStatus `json:"output_status,required"`
	JSON         commandExecFuncResponseJSON         `json:"-"`
}

// commandExecFuncResponseJSON contains the JSON metadata for the struct
// [CommandExecFuncResponse]
type commandExecFuncResponseJSON struct {
	Execution    apijson.Field
	Output       apijson.Field
	OutputStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CommandExecFuncResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commandExecFuncResponseJSON) RawJSON() string {
	return r.raw
}

type CommandExecFuncResponseExecution struct {
	// The exit code returned by the script. Will often be '0' on success and non-zero
	// on failure.
	ExitCode int64 `json:"exit_code,required"`
	// The contents of 'stderr' after executing the script.
	Stderr string `json:"stderr,required"`
	// The contents of 'stdout' after executing the script.
	Stdout string                               `json:"stdout,required"`
	JSON   commandExecFuncResponseExecutionJSON `json:"-"`
}

// commandExecFuncResponseExecutionJSON contains the JSON metadata for the struct
// [CommandExecFuncResponseExecution]
type commandExecFuncResponseExecutionJSON struct {
	ExitCode    apijson.Field
	Stderr      apijson.Field
	Stdout      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommandExecFuncResponseExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commandExecFuncResponseExecutionJSON) RawJSON() string {
	return r.raw
}

// The status of the output. "valid" means your function executed successfully and
// returned a valid JSON-serializable object, or void. "json_serialization_error"
// means your function executed successfully, but returned a nonserializable
// object. "error" means your function failed to execute.
type CommandExecFuncResponseOutputStatus string

const (
	CommandExecFuncResponseOutputStatusError                  CommandExecFuncResponseOutputStatus = "error"
	CommandExecFuncResponseOutputStatusJsonSerializationError CommandExecFuncResponseOutputStatus = "json_serialization_error"
	CommandExecFuncResponseOutputStatusValid                  CommandExecFuncResponseOutputStatus = "valid"
)

func (r CommandExecFuncResponseOutputStatus) IsKnown() bool {
	switch r {
	case CommandExecFuncResponseOutputStatusError, CommandExecFuncResponseOutputStatusJsonSerializationError, CommandExecFuncResponseOutputStatusValid:
		return true
	}
	return false
}

type CommandExecParams struct {
	// The code to execute.
	Code param.Field[string] `json:"code,required"`
	// The interpreter to use when executing code.
	Language param.Field[CommandExecParamsLanguage] `json:"language,required"`
	// List of command line arguments to pass to the script.
	Args param.Field[[]string] `json:"args"`
	// Set of key-value pairs to add to the script's execution environment.
	Env param.Field[map[string]string] `json:"env"`
	// List of input files.
	Files param.Field[[]CommandExecParamsFile] `json:"files"`
	// Configuration for HTTP requests and authentication.
	HTTP param.Field[CommandExecParamsHTTP] `json:"http"`
	// Configuration for execution environment limits.
	Limits param.Field[CommandExecParamsLimits] `json:"limits"`
	// The ID of the runtime revision to use when executing code.
	RuntimeRevisionID param.Field[string] `json:"runtime_revision_id"`
	// Input made available to the script via 'stdin'.
	Stdin param.Field[string] `json:"stdin"`
}

func (r CommandExecParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The interpreter to use when executing code.
type CommandExecParamsLanguage string

const (
	CommandExecParamsLanguagePython     CommandExecParamsLanguage = "python"
	CommandExecParamsLanguageJavascript CommandExecParamsLanguage = "javascript"
	CommandExecParamsLanguageTypescript CommandExecParamsLanguage = "typescript"
	CommandExecParamsLanguageRuby       CommandExecParamsLanguage = "ruby"
	CommandExecParamsLanguagePhp        CommandExecParamsLanguage = "php"
)

func (r CommandExecParamsLanguage) IsKnown() bool {
	switch r {
	case CommandExecParamsLanguagePython, CommandExecParamsLanguageJavascript, CommandExecParamsLanguageTypescript, CommandExecParamsLanguageRuby, CommandExecParamsLanguagePhp:
		return true
	}
	return false
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

type CommandExecFuncParams struct {
	// The function to execute. Your code must define a function named "execute" that
	// takes in a single argument and returns a JSON-serializable value.
	Code param.Field[string] `json:"code,required"`
	// The interpreter to use when executing code.
	Language param.Field[CommandExecFuncParamsLanguage] `json:"language,required"`
	// Set of key-value pairs to add to the function's execution environment.
	Env param.Field[map[string]string] `json:"env"`
	// List of input files.
	Files param.Field[[]CommandExecFuncParamsFile] `json:"files"`
	// Configuration for HTTP requests and authentication.
	HTTP param.Field[CommandExecFuncParamsHTTP] `json:"http"`
	// The input to the function. This must be a valid JSON-serializable object. If you
	// do not pass an input, your function will be called with None (Python) or null
	// (JavaScript/TypeScript) as the argument.
	Input param.Field[interface{}] `json:"input"`
	// Configuration for execution environment limits.
	Limits param.Field[CommandExecFuncParamsLimits] `json:"limits"`
	// The ID of the runtime revision to use when executing code.
	RuntimeRevisionID param.Field[string] `json:"runtime_revision_id"`
}

func (r CommandExecFuncParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The interpreter to use when executing code.
type CommandExecFuncParamsLanguage string

const (
	CommandExecFuncParamsLanguagePython     CommandExecFuncParamsLanguage = "python"
	CommandExecFuncParamsLanguageJavascript CommandExecFuncParamsLanguage = "javascript"
	CommandExecFuncParamsLanguageTypescript CommandExecFuncParamsLanguage = "typescript"
)

func (r CommandExecFuncParamsLanguage) IsKnown() bool {
	switch r {
	case CommandExecFuncParamsLanguagePython, CommandExecFuncParamsLanguageJavascript, CommandExecFuncParamsLanguageTypescript:
		return true
	}
	return false
}

type CommandExecFuncParamsFile struct {
	// The contents of the file.
	Contents param.Field[string] `json:"contents"`
	// The relative path of the file.
	Path param.Field[string] `json:"path"`
}

func (r CommandExecFuncParamsFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for HTTP requests and authentication.
type CommandExecFuncParamsHTTP struct {
	// List of allowed HTTP hosts and associated authentication.
	Allow param.Field[[]CommandExecFuncParamsHTTPAllow] `json:"allow"`
}

func (r CommandExecFuncParamsHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// List of allowed HTTP hosts and associated authentication.
type CommandExecFuncParamsHTTPAllow struct {
	// Authentication configuration for outbound requests to this host.
	Auth param.Field[CommandExecFuncParamsHTTPAllowAuth] `json:"auth"`
	// The hostname to allow.
	Host param.Field[string] `json:"host"`
}

func (r CommandExecFuncParamsHTTPAllow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Authentication configuration for outbound requests to this host.
type CommandExecFuncParamsHTTPAllowAuth struct {
	Basic param.Field[CommandExecFuncParamsHTTPAllowAuthBasic] `json:"basic"`
	// Configuration to add an 'Authorization' header using the 'Bearer' scheme.
	Bearer param.Field[CommandExecFuncParamsHTTPAllowAuthBearer] `json:"bearer"`
	Header param.Field[CommandExecFuncParamsHTTPAllowAuthHeader] `json:"header"`
	Query  param.Field[CommandExecFuncParamsHTTPAllowAuthQuery]  `json:"query"`
}

func (r CommandExecFuncParamsHTTPAllowAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CommandExecFuncParamsHTTPAllowAuthBasic struct {
	Password param.Field[string] `json:"password"`
	UserID   param.Field[string] `json:"user_id"`
}

func (r CommandExecFuncParamsHTTPAllowAuthBasic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration to add an 'Authorization' header using the 'Bearer' scheme.
type CommandExecFuncParamsHTTPAllowAuthBearer struct {
	// The token to set, e.g. 'Authorization: Bearer <token>'.
	Token param.Field[string] `json:"token"`
}

func (r CommandExecFuncParamsHTTPAllowAuthBearer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CommandExecFuncParamsHTTPAllowAuthHeader struct {
	Name  param.Field[string] `json:"name"`
	Value param.Field[string] `json:"value"`
}

func (r CommandExecFuncParamsHTTPAllowAuthHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CommandExecFuncParamsHTTPAllowAuthQuery struct {
	Key   param.Field[string] `json:"key"`
	Value param.Field[string] `json:"value"`
}

func (r CommandExecFuncParamsHTTPAllowAuthQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for execution environment limits.
type CommandExecFuncParamsLimits struct {
	// The maximum time allowed for execution (in seconds). Default is 30.
	ExecutionTimeout param.Field[int64] `json:"execution_timeout"`
	// The maximum memory allowed for execution (in MiB). Default is 128.
	MemorySize param.Field[int64] `json:"memory_size"`
}

func (r CommandExecFuncParamsLimits) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
