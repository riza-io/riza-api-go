// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package riza

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/riza-io/riza-api-go/internal/apijson"
	"github.com/riza-io/riza-api-go/internal/param"
	"github.com/riza-io/riza-api-go/internal/requestconfig"
	"github.com/riza-io/riza-api-go/option"
)

// ToolService contains methods and other services that help with interacting with
// the riza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolService] method instead.
type ToolService struct {
	Options []option.RequestOption
}

// NewToolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewToolService(opts ...option.RequestOption) (r *ToolService) {
	r = &ToolService{}
	r.Options = opts
	return
}

// Create a tool in your project.
func (r *ToolService) New(ctx context.Context, body ToolNewParams, opts ...option.RequestOption) (res *Tool, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update the source code and input schema of a tool.
func (r *ToolService) Update(ctx context.Context, id string, body ToolUpdateParams, opts ...option.RequestOption) (res *Tool, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tools/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of tools in your project.
func (r *ToolService) List(ctx context.Context, opts ...option.RequestOption) (res *ToolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Execute a tool with a given input. The input is validated against the tool's
// input schema.
func (r *ToolService) Exec(ctx context.Context, id string, body ToolExecParams, opts ...option.RequestOption) (res *ToolExecResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tools/%s/execute", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a tool.
func (r *ToolService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Tool, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/tools/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Tool struct {
	// The ID of the tool.
	ID string `json:"id,required"`
	// The code of the tool. You must define a function named "execute" that takes in a
	// single argument and returns a JSON-serializable value. The argument will be the
	// "input" passed when executing the tool, and will match the input schema.
	Code string `json:"code,required"`
	// A description of the tool.
	Description string `json:"description,required"`
	// The input schema of the tool. This must be a valid JSON Schema object.
	InputSchema interface{} `json:"input_schema,required"`
	// The language of the tool's code.
	Language ToolLanguage `json:"language,required"`
	// The name of the tool.
	Name string `json:"name,required"`
	// The ID of the tool's current revision. This is used to pin executions to a
	// specific version of the tool, even if the tool is updated later.
	RevisionID string `json:"revision_id,required"`
	// The ID of the custom runtime revision that the tool uses for executions. This
	// pins executions to specific version of a custom runtime runtime, even if the
	// runtime is updated later.
	RuntimeRevisionID string   `json:"runtime_revision_id"`
	JSON              toolJSON `json:"-"`
}

// toolJSON contains the JSON metadata for the struct [Tool]
type toolJSON struct {
	ID                apijson.Field
	Code              apijson.Field
	Description       apijson.Field
	InputSchema       apijson.Field
	Language          apijson.Field
	Name              apijson.Field
	RevisionID        apijson.Field
	RuntimeRevisionID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Tool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolJSON) RawJSON() string {
	return r.raw
}

// The language of the tool's code.
type ToolLanguage string

const (
	ToolLanguagePython     ToolLanguage = "python"
	ToolLanguageJavascript ToolLanguage = "javascript"
	ToolLanguageTypescript ToolLanguage = "typescript"
)

func (r ToolLanguage) IsKnown() bool {
	switch r {
	case ToolLanguagePython, ToolLanguageJavascript, ToolLanguageTypescript:
		return true
	}
	return false
}

type ToolListResponse struct {
	Tools []Tool               `json:"tools,required"`
	JSON  toolListResponseJSON `json:"-"`
}

// toolListResponseJSON contains the JSON metadata for the struct
// [ToolListResponse]
type toolListResponseJSON struct {
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolListResponseJSON) RawJSON() string {
	return r.raw
}

type ToolExecResponse struct {
	// The execution details of the function.
	Execution ToolExecResponseExecution `json:"execution,required"`
	// The returned value of the Tool's execute function.
	Output interface{} `json:"output,required"`
	// The status of the output. "valid" means your Tool executed successfully and
	// returned a valid JSON-serializable object, or void. "json_serialization_error"
	// means your Tool executed successfully, but returned a nonserializable object.
	// "error" means your Tool failed to execute.
	OutputStatus ToolExecResponseOutputStatus `json:"output_status,required"`
	JSON         toolExecResponseJSON         `json:"-"`
}

// toolExecResponseJSON contains the JSON metadata for the struct
// [ToolExecResponse]
type toolExecResponseJSON struct {
	Execution    apijson.Field
	Output       apijson.Field
	OutputStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ToolExecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolExecResponseJSON) RawJSON() string {
	return r.raw
}

// The execution details of the function.
type ToolExecResponseExecution struct {
	// The exit code returned by the function. Will often be '0' on success and
	// non-zero on failure.
	ExitCode int64 `json:"exit_code,required"`
	// The contents of 'stderr' after executing the function.
	Stderr string `json:"stderr,required"`
	// The contents of 'stdout' after executing the function.
	Stdout string                        `json:"stdout,required"`
	JSON   toolExecResponseExecutionJSON `json:"-"`
}

// toolExecResponseExecutionJSON contains the JSON metadata for the struct
// [ToolExecResponseExecution]
type toolExecResponseExecutionJSON struct {
	ExitCode    apijson.Field
	Stderr      apijson.Field
	Stdout      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolExecResponseExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolExecResponseExecutionJSON) RawJSON() string {
	return r.raw
}

// The status of the output. "valid" means your Tool executed successfully and
// returned a valid JSON-serializable object, or void. "json_serialization_error"
// means your Tool executed successfully, but returned a nonserializable object.
// "error" means your Tool failed to execute.
type ToolExecResponseOutputStatus string

const (
	ToolExecResponseOutputStatusError                  ToolExecResponseOutputStatus = "error"
	ToolExecResponseOutputStatusJsonSerializationError ToolExecResponseOutputStatus = "json_serialization_error"
	ToolExecResponseOutputStatusValid                  ToolExecResponseOutputStatus = "valid"
)

func (r ToolExecResponseOutputStatus) IsKnown() bool {
	switch r {
	case ToolExecResponseOutputStatusError, ToolExecResponseOutputStatusJsonSerializationError, ToolExecResponseOutputStatusValid:
		return true
	}
	return false
}

type ToolNewParams struct {
	// The code of the tool. You must define a function named "execute" that takes in a
	// single argument and returns a JSON-serializable value. The argument will be the
	// "input" passed when executing the tool, and will match the input schema.
	Code param.Field[string] `json:"code,required"`
	// The language of the tool's code.
	Language param.Field[ToolNewParamsLanguage] `json:"language,required"`
	// The name of the tool.
	Name param.Field[string] `json:"name,required"`
	// A description of the tool.
	Description param.Field[string] `json:"description"`
	// The input schema of the tool. This must be a valid JSON Schema object.
	InputSchema param.Field[interface{}] `json:"input_schema"`
	// The ID of the runtime revision to use when executing the tool.
	RuntimeRevisionID param.Field[string] `json:"runtime_revision_id"`
}

func (r ToolNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The language of the tool's code.
type ToolNewParamsLanguage string

const (
	ToolNewParamsLanguagePython     ToolNewParamsLanguage = "python"
	ToolNewParamsLanguageJavascript ToolNewParamsLanguage = "javascript"
	ToolNewParamsLanguageTypescript ToolNewParamsLanguage = "typescript"
)

func (r ToolNewParamsLanguage) IsKnown() bool {
	switch r {
	case ToolNewParamsLanguagePython, ToolNewParamsLanguageJavascript, ToolNewParamsLanguageTypescript:
		return true
	}
	return false
}

type ToolUpdateParams struct {
	// The code of the tool. You must define a function named "execute" that takes in a
	// single argument and returns a JSON-serializable value. The argument will be the
	// "input" passed when executing the tool, and will match the input schema.
	Code param.Field[string] `json:"code"`
	// A description of the tool.
	Description param.Field[string] `json:"description"`
	// The input schema of the tool. This must be a valid JSON Schema object.
	InputSchema param.Field[interface{}] `json:"input_schema"`
	// The language of the tool's code.
	Language param.Field[ToolUpdateParamsLanguage] `json:"language"`
	// The name of the tool.
	Name param.Field[string] `json:"name"`
	// The ID of the custom runtime revision that the tool uses for executions. This is
	// used to pin executions to a specific version of a custom runtime, even if the
	// runtime is updated later.
	RuntimeRevisionID param.Field[string] `json:"runtime_revision_id"`
}

func (r ToolUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The language of the tool's code.
type ToolUpdateParamsLanguage string

const (
	ToolUpdateParamsLanguagePython     ToolUpdateParamsLanguage = "python"
	ToolUpdateParamsLanguageJavascript ToolUpdateParamsLanguage = "javascript"
	ToolUpdateParamsLanguageTypescript ToolUpdateParamsLanguage = "typescript"
)

func (r ToolUpdateParamsLanguage) IsKnown() bool {
	switch r {
	case ToolUpdateParamsLanguagePython, ToolUpdateParamsLanguageJavascript, ToolUpdateParamsLanguageTypescript:
		return true
	}
	return false
}

type ToolExecParams struct {
	// Set of key-value pairs to add to the tool's execution environment.
	Env param.Field[[]ToolExecParamsEnv] `json:"env"`
	// Configuration for HTTP requests and authentication.
	HTTP param.Field[ToolExecParamsHTTP] `json:"http"`
	// The input to the tool. This must be a valid JSON-serializable object. It will be
	// validated against the tool's input schema.
	Input param.Field[interface{}] `json:"input"`
	// The Tool revision ID to execute. This optional parmeter is used to pin
	// executions to specific versions of the Tool. If not provided, the latest
	// (current) version of the Tool will be executed.
	RevisionID param.Field[string] `json:"revision_id"`
}

func (r ToolExecParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set of key-value pairs to add to the tool's execution environment.
type ToolExecParamsEnv struct {
	Name     param.Field[string] `json:"name,required"`
	SecretID param.Field[string] `json:"secret_id"`
	Value    param.Field[string] `json:"value"`
}

func (r ToolExecParamsEnv) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for HTTP requests and authentication.
type ToolExecParamsHTTP struct {
	// List of allowed HTTP hosts and associated authentication.
	Allow param.Field[[]ToolExecParamsHTTPAllow] `json:"allow"`
}

func (r ToolExecParamsHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// List of allowed HTTP hosts and associated authentication.
type ToolExecParamsHTTPAllow struct {
	// Authentication configuration for outbound requests to this host.
	Auth param.Field[ToolExecParamsHTTPAllowAuth] `json:"auth"`
	// The hostname to allow.
	Host param.Field[string] `json:"host"`
}

func (r ToolExecParamsHTTPAllow) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Authentication configuration for outbound requests to this host.
type ToolExecParamsHTTPAllowAuth struct {
	Basic param.Field[ToolExecParamsHTTPAllowAuthBasic] `json:"basic"`
	// Configuration to add an 'Authorization' header using the 'Bearer' scheme.
	Bearer param.Field[ToolExecParamsHTTPAllowAuthBearer] `json:"bearer"`
	Query  param.Field[ToolExecParamsHTTPAllowAuthQuery]  `json:"query"`
}

func (r ToolExecParamsHTTPAllowAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolExecParamsHTTPAllowAuthBasic struct {
	Password param.Field[string] `json:"password"`
	SecretID param.Field[string] `json:"secret_id"`
	UserID   param.Field[string] `json:"user_id"`
}

func (r ToolExecParamsHTTPAllowAuthBasic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration to add an 'Authorization' header using the 'Bearer' scheme.
type ToolExecParamsHTTPAllowAuthBearer struct {
	// The token to set, e.g. 'Authorization: Bearer <token>'.
	Token    param.Field[string] `json:"token"`
	SecretID param.Field[string] `json:"secret_id"`
}

func (r ToolExecParamsHTTPAllowAuthBearer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolExecParamsHTTPAllowAuthQuery struct {
	Key      param.Field[string] `json:"key"`
	SecretID param.Field[string] `json:"secret_id"`
	Value    param.Field[string] `json:"value"`
}

func (r ToolExecParamsHTTPAllowAuthQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
