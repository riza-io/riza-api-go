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

func (r *ToolService) New(ctx context.Context, body ToolNewParams, opts ...option.RequestOption) (res *Tool, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

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

func (r *ToolService) List(ctx context.Context, opts ...option.RequestOption) (res *ToolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

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
	ID          string      `json:"id,required"`
	Code        string      `json:"code,required"`
	Description string      `json:"description,required"`
	InputSchema interface{} `json:"input_schema,required"`
	Name        string      `json:"name,required"`
	RevisionID  string      `json:"revision_id,required"`
	JSON        toolJSON    `json:"-"`
}

// toolJSON contains the JSON metadata for the struct [Tool]
type toolJSON struct {
	ID          apijson.Field
	Code        apijson.Field
	Description apijson.Field
	InputSchema apijson.Field
	Name        apijson.Field
	RevisionID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Tool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolJSON) RawJSON() string {
	return r.raw
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
	Execution ToolExecResponseExecution `json:"execution,required"`
	Output    string                    `json:"output"`
	JSON      toolExecResponseJSON      `json:"-"`
}

// toolExecResponseJSON contains the JSON metadata for the struct
// [ToolExecResponse]
type toolExecResponseJSON struct {
	Execution   apijson.Field
	Output      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolExecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolExecResponseJSON) RawJSON() string {
	return r.raw
}

type ToolExecResponseExecution struct {
	ExitCode int64                         `json:"exit_code,required"`
	Stderr   string                        `json:"stderr,required"`
	Stdout   string                        `json:"stdout,required"`
	JSON     toolExecResponseExecutionJSON `json:"-"`
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

type ToolNewParams struct {
	Code        param.Field[string]                `json:"code,required"`
	Name        param.Field[string]                `json:"name,required"`
	Description param.Field[string]                `json:"description"`
	InputSchema param.Field[interface{}]           `json:"input_schema"`
	Language    param.Field[ToolNewParamsLanguage] `json:"language"`
}

func (r ToolNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolNewParamsLanguage string

const (
	ToolNewParamsLanguagePython     ToolNewParamsLanguage = "PYTHON"
	ToolNewParamsLanguageJavascript ToolNewParamsLanguage = "JAVASCRIPT"
	ToolNewParamsLanguageTypescript ToolNewParamsLanguage = "TYPESCRIPT"
	ToolNewParamsLanguageRuby       ToolNewParamsLanguage = "RUBY"
	ToolNewParamsLanguagePhp        ToolNewParamsLanguage = "PHP"
)

func (r ToolNewParamsLanguage) IsKnown() bool {
	switch r {
	case ToolNewParamsLanguagePython, ToolNewParamsLanguageJavascript, ToolNewParamsLanguageTypescript, ToolNewParamsLanguageRuby, ToolNewParamsLanguagePhp:
		return true
	}
	return false
}

type ToolUpdateParams struct {
	Code        param.Field[string]                   `json:"code"`
	Description param.Field[string]                   `json:"description"`
	InputSchema param.Field[interface{}]              `json:"input_schema"`
	Language    param.Field[ToolUpdateParamsLanguage] `json:"language"`
	Name        param.Field[string]                   `json:"name"`
}

func (r ToolUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolUpdateParamsLanguage string

const (
	ToolUpdateParamsLanguagePython     ToolUpdateParamsLanguage = "PYTHON"
	ToolUpdateParamsLanguageJavascript ToolUpdateParamsLanguage = "JAVASCRIPT"
	ToolUpdateParamsLanguageTypescript ToolUpdateParamsLanguage = "TYPESCRIPT"
	ToolUpdateParamsLanguageRuby       ToolUpdateParamsLanguage = "RUBY"
	ToolUpdateParamsLanguagePhp        ToolUpdateParamsLanguage = "PHP"
)

func (r ToolUpdateParamsLanguage) IsKnown() bool {
	switch r {
	case ToolUpdateParamsLanguagePython, ToolUpdateParamsLanguageJavascript, ToolUpdateParamsLanguageTypescript, ToolUpdateParamsLanguageRuby, ToolUpdateParamsLanguagePhp:
		return true
	}
	return false
}

type ToolExecParams struct {
	Env        param.Field[[]ToolExecParamsEnv] `json:"env"`
	HTTP       param.Field[ToolExecParamsHTTP]  `json:"http"`
	Input      param.Field[interface{}]         `json:"input"`
	RevisionID param.Field[string]              `json:"revision_id"`
}

func (r ToolExecParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ToolExecParamsEnv struct {
	Name     param.Field[string] `json:"name,required"`
	SecretID param.Field[string] `json:"secret_id"`
	Value    param.Field[string] `json:"value"`
}

func (r ToolExecParamsEnv) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
