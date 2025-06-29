// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package riza

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/riza-io/riza-api-go/internal/apijson"
	"github.com/riza-io/riza-api-go/internal/apiquery"
	"github.com/riza-io/riza-api-go/internal/param"
	"github.com/riza-io/riza-api-go/internal/requestconfig"
	"github.com/riza-io/riza-api-go/option"
	"github.com/riza-io/riza-api-go/packages/pagination"
	"github.com/tidwall/gjson"
)

// ExecutionService contains methods and other services that help with interacting
// with the riza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExecutionService] method instead.
type ExecutionService struct {
	Options []option.RequestOption
}

// NewExecutionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExecutionService(opts ...option.RequestOption) (r *ExecutionService) {
	r = &ExecutionService{}
	r.Options = opts
	return
}

// Returns a list of executions in your project.
func (r *ExecutionService) List(ctx context.Context, query ExecutionListParams, opts ...option.RequestOption) (res *pagination.DefaultPagination[Execution], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/executions"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a list of executions in your project.
func (r *ExecutionService) ListAutoPaging(ctx context.Context, query ExecutionListParams, opts ...option.RequestOption) *pagination.DefaultPaginationAutoPager[Execution] {
	return pagination.NewDefaultPaginationAutoPager(r.List(ctx, query, opts...))
}

// Retrieves an execution.
func (r *ExecutionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Execution, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/executions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Execution struct {
	ID        string            `json:"id,required"`
	Duration  int64             `json:"duration,required"`
	ExitCode  int64             `json:"exit_code,required"`
	Language  ExecutionLanguage `json:"language,required"`
	StartedAt time.Time         `json:"started_at,required" format:"date-time"`
	Details   ExecutionDetails  `json:"details"`
	JSON      executionJSON     `json:"-"`
}

// executionJSON contains the JSON metadata for the struct [Execution]
type executionJSON struct {
	ID          apijson.Field
	Duration    apijson.Field
	ExitCode    apijson.Field
	Language    apijson.Field
	StartedAt   apijson.Field
	Details     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Execution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionJSON) RawJSON() string {
	return r.raw
}

type ExecutionLanguage string

const (
	ExecutionLanguagePython     ExecutionLanguage = "python"
	ExecutionLanguageJavascript ExecutionLanguage = "javascript"
	ExecutionLanguageTypescript ExecutionLanguage = "typescript"
	ExecutionLanguageRuby       ExecutionLanguage = "ruby"
	ExecutionLanguagePhp        ExecutionLanguage = "php"
)

func (r ExecutionLanguage) IsKnown() bool {
	switch r {
	case ExecutionLanguagePython, ExecutionLanguageJavascript, ExecutionLanguageTypescript, ExecutionLanguageRuby, ExecutionLanguagePhp:
		return true
	}
	return false
}

type ExecutionDetails struct {
	// This field can have the runtime type of
	// [ExecutionDetailsToolExecutionDetailsRequest],
	// [ExecutionDetailsFunctionExecutionDetailsRequest],
	// [ExecutionDetailsScriptExecutionDetailsRequest].
	Request interface{} `json:"request,required"`
	// This field can have the runtime type of
	// [ExecutionDetailsToolExecutionDetailsResponse],
	// [ExecutionDetailsFunctionExecutionDetailsResponse],
	// [ExecutionDetailsScriptExecutionDetailsResponse].
	Response interface{}          `json:"response,required"`
	Type     ExecutionDetailsType `json:"type,required"`
	ToolID   string               `json:"tool_id"`
	JSON     executionDetailsJSON `json:"-"`
	union    ExecutionDetailsUnion
}

// executionDetailsJSON contains the JSON metadata for the struct
// [ExecutionDetails]
type executionDetailsJSON struct {
	Request     apijson.Field
	Response    apijson.Field
	Type        apijson.Field
	ToolID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r executionDetailsJSON) RawJSON() string {
	return r.raw
}

func (r *ExecutionDetails) UnmarshalJSON(data []byte) (err error) {
	*r = ExecutionDetails{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ExecutionDetailsUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [ExecutionDetailsToolExecutionDetails],
// [ExecutionDetailsFunctionExecutionDetails],
// [ExecutionDetailsScriptExecutionDetails].
func (r ExecutionDetails) AsUnion() ExecutionDetailsUnion {
	return r.union
}

// Union satisfied by [ExecutionDetailsToolExecutionDetails],
// [ExecutionDetailsFunctionExecutionDetails] or
// [ExecutionDetailsScriptExecutionDetails].
type ExecutionDetailsUnion interface {
	implementsExecutionDetails()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ExecutionDetailsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecutionDetailsToolExecutionDetails{}),
			DiscriminatorValue: "tool",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecutionDetailsFunctionExecutionDetails{}),
			DiscriminatorValue: "function",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecutionDetailsScriptExecutionDetails{}),
			DiscriminatorValue: "script",
		},
	)
}

type ExecutionDetailsToolExecutionDetails struct {
	Request  ExecutionDetailsToolExecutionDetailsRequest  `json:"request,required"`
	Response ExecutionDetailsToolExecutionDetailsResponse `json:"response,required"`
	ToolID   string                                       `json:"tool_id,required"`
	Type     ExecutionDetailsToolExecutionDetailsType     `json:"type,required"`
	JSON     executionDetailsToolExecutionDetailsJSON     `json:"-"`
}

// executionDetailsToolExecutionDetailsJSON contains the JSON metadata for the
// struct [ExecutionDetailsToolExecutionDetails]
type executionDetailsToolExecutionDetailsJSON struct {
	Request     apijson.Field
	Response    apijson.Field
	ToolID      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsToolExecutionDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsToolExecutionDetailsJSON) RawJSON() string {
	return r.raw
}

func (r ExecutionDetailsToolExecutionDetails) implementsExecutionDetails() {}

type ExecutionDetailsToolExecutionDetailsRequest struct {
	// Set of key-value pairs to add to the tool's execution environment.
	Env []ExecutionDetailsToolExecutionDetailsRequestEnv `json:"env"`
	// Configuration for HTTP requests and authentication.
	HTTP ExecutionDetailsToolExecutionDetailsRequestHTTP `json:"http"`
	// The input to the tool. This must be a valid JSON-serializable object. It will be
	// validated against the tool's input schema.
	Input interface{} `json:"input"`
	// The Tool revision ID to execute. This optional parmeter is used to pin
	// executions to specific versions of the Tool. If not provided, the latest
	// (current) version of the Tool will be executed.
	RevisionID string                                          `json:"revision_id"`
	JSON       executionDetailsToolExecutionDetailsRequestJSON `json:"-"`
}

// executionDetailsToolExecutionDetailsRequestJSON contains the JSON metadata for
// the struct [ExecutionDetailsToolExecutionDetailsRequest]
type executionDetailsToolExecutionDetailsRequestJSON struct {
	Env         apijson.Field
	HTTP        apijson.Field
	Input       apijson.Field
	RevisionID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsToolExecutionDetailsRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsToolExecutionDetailsRequestJSON) RawJSON() string {
	return r.raw
}

// Set of key-value pairs to add to the tool's execution environment.
type ExecutionDetailsToolExecutionDetailsRequestEnv struct {
	Name     string                                             `json:"name,required"`
	SecretID string                                             `json:"secret_id"`
	Value    string                                             `json:"value"`
	JSON     executionDetailsToolExecutionDetailsRequestEnvJSON `json:"-"`
}

// executionDetailsToolExecutionDetailsRequestEnvJSON contains the JSON metadata
// for the struct [ExecutionDetailsToolExecutionDetailsRequestEnv]
type executionDetailsToolExecutionDetailsRequestEnvJSON struct {
	Name        apijson.Field
	SecretID    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsToolExecutionDetailsRequestEnv) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsToolExecutionDetailsRequestEnvJSON) RawJSON() string {
	return r.raw
}

// Configuration for HTTP requests and authentication.
type ExecutionDetailsToolExecutionDetailsRequestHTTP struct {
	// List of allowed HTTP hosts and associated authentication.
	Allow []ExecutionDetailsToolExecutionDetailsRequestHTTPAllow `json:"allow"`
	JSON  executionDetailsToolExecutionDetailsRequestHTTPJSON    `json:"-"`
}

// executionDetailsToolExecutionDetailsRequestHTTPJSON contains the JSON metadata
// for the struct [ExecutionDetailsToolExecutionDetailsRequestHTTP]
type executionDetailsToolExecutionDetailsRequestHTTPJSON struct {
	Allow       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsToolExecutionDetailsRequestHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsToolExecutionDetailsRequestHTTPJSON) RawJSON() string {
	return r.raw
}

// List of allowed HTTP hosts and associated authentication.
type ExecutionDetailsToolExecutionDetailsRequestHTTPAllow struct {
	// Authentication configuration for outbound requests to this host.
	Auth ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuth `json:"auth"`
	// The hostname to allow.
	Host string                                                   `json:"host"`
	JSON executionDetailsToolExecutionDetailsRequestHTTPAllowJSON `json:"-"`
}

// executionDetailsToolExecutionDetailsRequestHTTPAllowJSON contains the JSON
// metadata for the struct [ExecutionDetailsToolExecutionDetailsRequestHTTPAllow]
type executionDetailsToolExecutionDetailsRequestHTTPAllowJSON struct {
	Auth        apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsToolExecutionDetailsRequestHTTPAllow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsToolExecutionDetailsRequestHTTPAllowJSON) RawJSON() string {
	return r.raw
}

// Authentication configuration for outbound requests to this host.
type ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuth struct {
	Basic ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuthBasic `json:"basic"`
	// Configuration to add an 'Authorization' header using the 'Bearer' scheme.
	Bearer ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuthBearer `json:"bearer"`
	Query  ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuthQuery  `json:"query"`
	JSON   executionDetailsToolExecutionDetailsRequestHTTPAllowAuthJSON   `json:"-"`
}

// executionDetailsToolExecutionDetailsRequestHTTPAllowAuthJSON contains the JSON
// metadata for the struct
// [ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuth]
type executionDetailsToolExecutionDetailsRequestHTTPAllowAuthJSON struct {
	Basic       apijson.Field
	Bearer      apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsToolExecutionDetailsRequestHTTPAllowAuthJSON) RawJSON() string {
	return r.raw
}

type ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuthBasic struct {
	Password string                                                            `json:"password"`
	SecretID string                                                            `json:"secret_id"`
	UserID   string                                                            `json:"user_id"`
	JSON     executionDetailsToolExecutionDetailsRequestHTTPAllowAuthBasicJSON `json:"-"`
}

// executionDetailsToolExecutionDetailsRequestHTTPAllowAuthBasicJSON contains the
// JSON metadata for the struct
// [ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuthBasic]
type executionDetailsToolExecutionDetailsRequestHTTPAllowAuthBasicJSON struct {
	Password    apijson.Field
	SecretID    apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuthBasic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsToolExecutionDetailsRequestHTTPAllowAuthBasicJSON) RawJSON() string {
	return r.raw
}

// Configuration to add an 'Authorization' header using the 'Bearer' scheme.
type ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuthBearer struct {
	// The token to set, e.g. 'Authorization: Bearer <token>'.
	Token    string                                                             `json:"token"`
	SecretID string                                                             `json:"secret_id"`
	JSON     executionDetailsToolExecutionDetailsRequestHTTPAllowAuthBearerJSON `json:"-"`
}

// executionDetailsToolExecutionDetailsRequestHTTPAllowAuthBearerJSON contains the
// JSON metadata for the struct
// [ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuthBearer]
type executionDetailsToolExecutionDetailsRequestHTTPAllowAuthBearerJSON struct {
	Token       apijson.Field
	SecretID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuthBearer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsToolExecutionDetailsRequestHTTPAllowAuthBearerJSON) RawJSON() string {
	return r.raw
}

type ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuthQuery struct {
	Key      string                                                            `json:"key"`
	SecretID string                                                            `json:"secret_id"`
	Value    string                                                            `json:"value"`
	JSON     executionDetailsToolExecutionDetailsRequestHTTPAllowAuthQueryJSON `json:"-"`
}

// executionDetailsToolExecutionDetailsRequestHTTPAllowAuthQueryJSON contains the
// JSON metadata for the struct
// [ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuthQuery]
type executionDetailsToolExecutionDetailsRequestHTTPAllowAuthQueryJSON struct {
	Key         apijson.Field
	SecretID    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsToolExecutionDetailsRequestHTTPAllowAuthQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsToolExecutionDetailsRequestHTTPAllowAuthQueryJSON) RawJSON() string {
	return r.raw
}

type ExecutionDetailsToolExecutionDetailsResponse struct {
	// The execution details of the function.
	Execution ExecutionDetailsToolExecutionDetailsResponseExecution `json:"execution,required"`
	// The returned value of the Tool's execute function.
	Output interface{} `json:"output,required"`
	// The status of the output. "valid" means your Tool executed successfully and
	// returned a valid JSON-serializable object, or void. "json_serialization_error"
	// means your Tool executed successfully, but returned a nonserializable object.
	// "error" means your Tool failed to execute.
	OutputStatus ExecutionDetailsToolExecutionDetailsResponseOutputStatus `json:"output_status,required"`
	JSON         executionDetailsToolExecutionDetailsResponseJSON         `json:"-"`
}

// executionDetailsToolExecutionDetailsResponseJSON contains the JSON metadata for
// the struct [ExecutionDetailsToolExecutionDetailsResponse]
type executionDetailsToolExecutionDetailsResponseJSON struct {
	Execution    apijson.Field
	Output       apijson.Field
	OutputStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ExecutionDetailsToolExecutionDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsToolExecutionDetailsResponseJSON) RawJSON() string {
	return r.raw
}

// The execution details of the function.
type ExecutionDetailsToolExecutionDetailsResponseExecution struct {
	// The ID of the execution.
	ID string `json:"id,required"`
	// The execution time of the function in milliseconds.
	Duration int64 `json:"duration,required"`
	// The exit code returned by the function. Will often be '0' on success and
	// non-zero on failure.
	ExitCode int64 `json:"exit_code,required"`
	// The contents of 'stderr' after executing the function.
	Stderr string `json:"stderr,required"`
	// The contents of 'stdout' after executing the function.
	Stdout string                                                    `json:"stdout,required"`
	JSON   executionDetailsToolExecutionDetailsResponseExecutionJSON `json:"-"`
}

// executionDetailsToolExecutionDetailsResponseExecutionJSON contains the JSON
// metadata for the struct [ExecutionDetailsToolExecutionDetailsResponseExecution]
type executionDetailsToolExecutionDetailsResponseExecutionJSON struct {
	ID          apijson.Field
	Duration    apijson.Field
	ExitCode    apijson.Field
	Stderr      apijson.Field
	Stdout      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsToolExecutionDetailsResponseExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsToolExecutionDetailsResponseExecutionJSON) RawJSON() string {
	return r.raw
}

// The status of the output. "valid" means your Tool executed successfully and
// returned a valid JSON-serializable object, or void. "json_serialization_error"
// means your Tool executed successfully, but returned a nonserializable object.
// "error" means your Tool failed to execute.
type ExecutionDetailsToolExecutionDetailsResponseOutputStatus string

const (
	ExecutionDetailsToolExecutionDetailsResponseOutputStatusError                  ExecutionDetailsToolExecutionDetailsResponseOutputStatus = "error"
	ExecutionDetailsToolExecutionDetailsResponseOutputStatusJsonSerializationError ExecutionDetailsToolExecutionDetailsResponseOutputStatus = "json_serialization_error"
	ExecutionDetailsToolExecutionDetailsResponseOutputStatusValid                  ExecutionDetailsToolExecutionDetailsResponseOutputStatus = "valid"
)

func (r ExecutionDetailsToolExecutionDetailsResponseOutputStatus) IsKnown() bool {
	switch r {
	case ExecutionDetailsToolExecutionDetailsResponseOutputStatusError, ExecutionDetailsToolExecutionDetailsResponseOutputStatusJsonSerializationError, ExecutionDetailsToolExecutionDetailsResponseOutputStatusValid:
		return true
	}
	return false
}

type ExecutionDetailsToolExecutionDetailsType string

const (
	ExecutionDetailsToolExecutionDetailsTypeTool ExecutionDetailsToolExecutionDetailsType = "tool"
)

func (r ExecutionDetailsToolExecutionDetailsType) IsKnown() bool {
	switch r {
	case ExecutionDetailsToolExecutionDetailsTypeTool:
		return true
	}
	return false
}

type ExecutionDetailsFunctionExecutionDetails struct {
	Request  ExecutionDetailsFunctionExecutionDetailsRequest  `json:"request,required"`
	Response ExecutionDetailsFunctionExecutionDetailsResponse `json:"response,required"`
	Type     ExecutionDetailsFunctionExecutionDetailsType     `json:"type,required"`
	JSON     executionDetailsFunctionExecutionDetailsJSON     `json:"-"`
}

// executionDetailsFunctionExecutionDetailsJSON contains the JSON metadata for the
// struct [ExecutionDetailsFunctionExecutionDetails]
type executionDetailsFunctionExecutionDetailsJSON struct {
	Request     apijson.Field
	Response    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsFunctionExecutionDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsFunctionExecutionDetailsJSON) RawJSON() string {
	return r.raw
}

func (r ExecutionDetailsFunctionExecutionDetails) implementsExecutionDetails() {}

type ExecutionDetailsFunctionExecutionDetailsRequest struct {
	// The function to execute. Your code must define a function named "execute" that
	// takes in a single argument and returns a JSON-serializable value.
	Code string `json:"code,required"`
	// The interpreter to use when executing code.
	Language ExecutionDetailsFunctionExecutionDetailsRequestLanguage `json:"language,required"`
	// Set of key-value pairs to add to the function's execution environment.
	Env map[string]string `json:"env"`
	// List of input files.
	Files []ExecutionDetailsFunctionExecutionDetailsRequestFile `json:"files"`
	// Configuration for HTTP requests and authentication.
	HTTP ExecutionDetailsFunctionExecutionDetailsRequestHTTP `json:"http"`
	// The input to the function. This must be a valid JSON-serializable object. If you
	// do not pass an input, your function will be called with None (Python) or null
	// (JavaScript/TypeScript) as the argument.
	Input interface{} `json:"input"`
	// Configuration for execution environment limits.
	Limits ExecutionDetailsFunctionExecutionDetailsRequestLimits `json:"limits"`
	// The ID of the runtime revision to use when executing code.
	RuntimeRevisionID string                                              `json:"runtime_revision_id"`
	JSON              executionDetailsFunctionExecutionDetailsRequestJSON `json:"-"`
}

// executionDetailsFunctionExecutionDetailsRequestJSON contains the JSON metadata
// for the struct [ExecutionDetailsFunctionExecutionDetailsRequest]
type executionDetailsFunctionExecutionDetailsRequestJSON struct {
	Code              apijson.Field
	Language          apijson.Field
	Env               apijson.Field
	Files             apijson.Field
	HTTP              apijson.Field
	Input             apijson.Field
	Limits            apijson.Field
	RuntimeRevisionID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ExecutionDetailsFunctionExecutionDetailsRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsFunctionExecutionDetailsRequestJSON) RawJSON() string {
	return r.raw
}

// The interpreter to use when executing code.
type ExecutionDetailsFunctionExecutionDetailsRequestLanguage string

const (
	ExecutionDetailsFunctionExecutionDetailsRequestLanguagePython     ExecutionDetailsFunctionExecutionDetailsRequestLanguage = "python"
	ExecutionDetailsFunctionExecutionDetailsRequestLanguageJavascript ExecutionDetailsFunctionExecutionDetailsRequestLanguage = "javascript"
	ExecutionDetailsFunctionExecutionDetailsRequestLanguageTypescript ExecutionDetailsFunctionExecutionDetailsRequestLanguage = "typescript"
)

func (r ExecutionDetailsFunctionExecutionDetailsRequestLanguage) IsKnown() bool {
	switch r {
	case ExecutionDetailsFunctionExecutionDetailsRequestLanguagePython, ExecutionDetailsFunctionExecutionDetailsRequestLanguageJavascript, ExecutionDetailsFunctionExecutionDetailsRequestLanguageTypescript:
		return true
	}
	return false
}

type ExecutionDetailsFunctionExecutionDetailsRequestFile struct {
	// The contents of the file.
	Contents string `json:"contents"`
	// The relative path of the file.
	Path string                                                  `json:"path"`
	JSON executionDetailsFunctionExecutionDetailsRequestFileJSON `json:"-"`
}

// executionDetailsFunctionExecutionDetailsRequestFileJSON contains the JSON
// metadata for the struct [ExecutionDetailsFunctionExecutionDetailsRequestFile]
type executionDetailsFunctionExecutionDetailsRequestFileJSON struct {
	Contents    apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsFunctionExecutionDetailsRequestFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsFunctionExecutionDetailsRequestFileJSON) RawJSON() string {
	return r.raw
}

// Configuration for HTTP requests and authentication.
type ExecutionDetailsFunctionExecutionDetailsRequestHTTP struct {
	// List of allowed HTTP hosts and associated authentication.
	Allow []ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllow `json:"allow"`
	JSON  executionDetailsFunctionExecutionDetailsRequestHTTPJSON    `json:"-"`
}

// executionDetailsFunctionExecutionDetailsRequestHTTPJSON contains the JSON
// metadata for the struct [ExecutionDetailsFunctionExecutionDetailsRequestHTTP]
type executionDetailsFunctionExecutionDetailsRequestHTTPJSON struct {
	Allow       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsFunctionExecutionDetailsRequestHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsFunctionExecutionDetailsRequestHTTPJSON) RawJSON() string {
	return r.raw
}

// List of allowed HTTP hosts and associated authentication.
type ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllow struct {
	// Authentication configuration for outbound requests to this host.
	Auth ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuth `json:"auth"`
	// The hostname to allow.
	Host string                                                       `json:"host"`
	JSON executionDetailsFunctionExecutionDetailsRequestHTTPAllowJSON `json:"-"`
}

// executionDetailsFunctionExecutionDetailsRequestHTTPAllowJSON contains the JSON
// metadata for the struct
// [ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllow]
type executionDetailsFunctionExecutionDetailsRequestHTTPAllowJSON struct {
	Auth        apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsFunctionExecutionDetailsRequestHTTPAllowJSON) RawJSON() string {
	return r.raw
}

// Authentication configuration for outbound requests to this host.
type ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuth struct {
	Basic ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBasic `json:"basic"`
	// Configuration to add an 'Authorization' header using the 'Bearer' scheme.
	Bearer ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBearer `json:"bearer"`
	Header ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthHeader `json:"header"`
	Query  ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthQuery  `json:"query"`
	JSON   executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthJSON   `json:"-"`
}

// executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthJSON contains the
// JSON metadata for the struct
// [ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuth]
type executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthJSON struct {
	Basic       apijson.Field
	Bearer      apijson.Field
	Header      apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthJSON) RawJSON() string {
	return r.raw
}

type ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBasic struct {
	Password string                                                                `json:"password"`
	UserID   string                                                                `json:"user_id"`
	JSON     executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBasicJSON `json:"-"`
}

// executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBasicJSON contains
// the JSON metadata for the struct
// [ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBasic]
type executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBasicJSON struct {
	Password    apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBasic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBasicJSON) RawJSON() string {
	return r.raw
}

// Configuration to add an 'Authorization' header using the 'Bearer' scheme.
type ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBearer struct {
	// The token to set, e.g. 'Authorization: Bearer <token>'.
	Token string                                                                 `json:"token"`
	JSON  executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBearerJSON `json:"-"`
}

// executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBearerJSON contains
// the JSON metadata for the struct
// [ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBearer]
type executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBearerJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBearer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthBearerJSON) RawJSON() string {
	return r.raw
}

type ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthHeader struct {
	Name  string                                                                 `json:"name"`
	Value string                                                                 `json:"value"`
	JSON  executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthHeaderJSON `json:"-"`
}

// executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthHeaderJSON contains
// the JSON metadata for the struct
// [ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthHeader]
type executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthHeaderJSON) RawJSON() string {
	return r.raw
}

type ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthQuery struct {
	Key   string                                                                `json:"key"`
	Value string                                                                `json:"value"`
	JSON  executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthQueryJSON `json:"-"`
}

// executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthQueryJSON contains
// the JSON metadata for the struct
// [ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthQuery]
type executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthQueryJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsFunctionExecutionDetailsRequestHTTPAllowAuthQueryJSON) RawJSON() string {
	return r.raw
}

// Configuration for execution environment limits.
type ExecutionDetailsFunctionExecutionDetailsRequestLimits struct {
	// The maximum time allowed for execution (in seconds). Default is 30.
	ExecutionTimeout int64 `json:"execution_timeout"`
	// The maximum memory allowed for execution (in MiB). Default is 128.
	MemorySize int64                                                     `json:"memory_size"`
	JSON       executionDetailsFunctionExecutionDetailsRequestLimitsJSON `json:"-"`
}

// executionDetailsFunctionExecutionDetailsRequestLimitsJSON contains the JSON
// metadata for the struct [ExecutionDetailsFunctionExecutionDetailsRequestLimits]
type executionDetailsFunctionExecutionDetailsRequestLimitsJSON struct {
	ExecutionTimeout apijson.Field
	MemorySize       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ExecutionDetailsFunctionExecutionDetailsRequestLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsFunctionExecutionDetailsRequestLimitsJSON) RawJSON() string {
	return r.raw
}

type ExecutionDetailsFunctionExecutionDetailsResponse struct {
	// The execution details of the function.
	Execution ExecutionDetailsFunctionExecutionDetailsResponseExecution `json:"execution,required"`
	// The output of the function.
	Output interface{} `json:"output,required"`
	// The status of the output. "valid" means your function executed successfully and
	// returned a valid JSON-serializable object, or void. "json_serialization_error"
	// means your function executed successfully, but returned a nonserializable
	// object. "error" means your function failed to execute.
	OutputStatus ExecutionDetailsFunctionExecutionDetailsResponseOutputStatus `json:"output_status,required"`
	JSON         executionDetailsFunctionExecutionDetailsResponseJSON         `json:"-"`
}

// executionDetailsFunctionExecutionDetailsResponseJSON contains the JSON metadata
// for the struct [ExecutionDetailsFunctionExecutionDetailsResponse]
type executionDetailsFunctionExecutionDetailsResponseJSON struct {
	Execution    apijson.Field
	Output       apijson.Field
	OutputStatus apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ExecutionDetailsFunctionExecutionDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsFunctionExecutionDetailsResponseJSON) RawJSON() string {
	return r.raw
}

// The execution details of the function.
type ExecutionDetailsFunctionExecutionDetailsResponseExecution struct {
	// The ID of the execution.
	ID string `json:"id,required"`
	// The execution time of the function in milliseconds.
	Duration int64 `json:"duration,required"`
	// The exit code returned by the function. Will often be '0' on success and
	// non-zero on failure.
	ExitCode int64 `json:"exit_code,required"`
	// The contents of 'stderr' after executing the function.
	Stderr string `json:"stderr,required"`
	// The contents of 'stdout' after executing the function.
	Stdout string                                                        `json:"stdout,required"`
	JSON   executionDetailsFunctionExecutionDetailsResponseExecutionJSON `json:"-"`
}

// executionDetailsFunctionExecutionDetailsResponseExecutionJSON contains the JSON
// metadata for the struct
// [ExecutionDetailsFunctionExecutionDetailsResponseExecution]
type executionDetailsFunctionExecutionDetailsResponseExecutionJSON struct {
	ID          apijson.Field
	Duration    apijson.Field
	ExitCode    apijson.Field
	Stderr      apijson.Field
	Stdout      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsFunctionExecutionDetailsResponseExecution) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsFunctionExecutionDetailsResponseExecutionJSON) RawJSON() string {
	return r.raw
}

// The status of the output. "valid" means your function executed successfully and
// returned a valid JSON-serializable object, or void. "json_serialization_error"
// means your function executed successfully, but returned a nonserializable
// object. "error" means your function failed to execute.
type ExecutionDetailsFunctionExecutionDetailsResponseOutputStatus string

const (
	ExecutionDetailsFunctionExecutionDetailsResponseOutputStatusError                  ExecutionDetailsFunctionExecutionDetailsResponseOutputStatus = "error"
	ExecutionDetailsFunctionExecutionDetailsResponseOutputStatusJsonSerializationError ExecutionDetailsFunctionExecutionDetailsResponseOutputStatus = "json_serialization_error"
	ExecutionDetailsFunctionExecutionDetailsResponseOutputStatusValid                  ExecutionDetailsFunctionExecutionDetailsResponseOutputStatus = "valid"
)

func (r ExecutionDetailsFunctionExecutionDetailsResponseOutputStatus) IsKnown() bool {
	switch r {
	case ExecutionDetailsFunctionExecutionDetailsResponseOutputStatusError, ExecutionDetailsFunctionExecutionDetailsResponseOutputStatusJsonSerializationError, ExecutionDetailsFunctionExecutionDetailsResponseOutputStatusValid:
		return true
	}
	return false
}

type ExecutionDetailsFunctionExecutionDetailsType string

const (
	ExecutionDetailsFunctionExecutionDetailsTypeFunction ExecutionDetailsFunctionExecutionDetailsType = "function"
)

func (r ExecutionDetailsFunctionExecutionDetailsType) IsKnown() bool {
	switch r {
	case ExecutionDetailsFunctionExecutionDetailsTypeFunction:
		return true
	}
	return false
}

type ExecutionDetailsScriptExecutionDetails struct {
	Request  ExecutionDetailsScriptExecutionDetailsRequest  `json:"request,required"`
	Response ExecutionDetailsScriptExecutionDetailsResponse `json:"response,required"`
	Type     ExecutionDetailsScriptExecutionDetailsType     `json:"type,required"`
	JSON     executionDetailsScriptExecutionDetailsJSON     `json:"-"`
}

// executionDetailsScriptExecutionDetailsJSON contains the JSON metadata for the
// struct [ExecutionDetailsScriptExecutionDetails]
type executionDetailsScriptExecutionDetailsJSON struct {
	Request     apijson.Field
	Response    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsScriptExecutionDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsScriptExecutionDetailsJSON) RawJSON() string {
	return r.raw
}

func (r ExecutionDetailsScriptExecutionDetails) implementsExecutionDetails() {}

type ExecutionDetailsScriptExecutionDetailsRequest struct {
	// The code to execute.
	Code string `json:"code,required"`
	// The interpreter to use when executing code.
	Language ExecutionDetailsScriptExecutionDetailsRequestLanguage `json:"language,required"`
	// List of command line arguments to pass to the script.
	Args []string `json:"args"`
	// Set of key-value pairs to add to the script's execution environment.
	Env map[string]string `json:"env"`
	// List of input files.
	Files []ExecutionDetailsScriptExecutionDetailsRequestFile `json:"files"`
	// Configuration for HTTP requests and authentication.
	HTTP ExecutionDetailsScriptExecutionDetailsRequestHTTP `json:"http"`
	// Configuration for execution environment limits.
	Limits ExecutionDetailsScriptExecutionDetailsRequestLimits `json:"limits"`
	// The ID of the runtime revision to use when executing code.
	RuntimeRevisionID string `json:"runtime_revision_id"`
	// Input made available to the script via 'stdin'.
	Stdin string                                            `json:"stdin"`
	JSON  executionDetailsScriptExecutionDetailsRequestJSON `json:"-"`
}

// executionDetailsScriptExecutionDetailsRequestJSON contains the JSON metadata for
// the struct [ExecutionDetailsScriptExecutionDetailsRequest]
type executionDetailsScriptExecutionDetailsRequestJSON struct {
	Code              apijson.Field
	Language          apijson.Field
	Args              apijson.Field
	Env               apijson.Field
	Files             apijson.Field
	HTTP              apijson.Field
	Limits            apijson.Field
	RuntimeRevisionID apijson.Field
	Stdin             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ExecutionDetailsScriptExecutionDetailsRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsScriptExecutionDetailsRequestJSON) RawJSON() string {
	return r.raw
}

// The interpreter to use when executing code.
type ExecutionDetailsScriptExecutionDetailsRequestLanguage string

const (
	ExecutionDetailsScriptExecutionDetailsRequestLanguagePython     ExecutionDetailsScriptExecutionDetailsRequestLanguage = "python"
	ExecutionDetailsScriptExecutionDetailsRequestLanguageJavascript ExecutionDetailsScriptExecutionDetailsRequestLanguage = "javascript"
	ExecutionDetailsScriptExecutionDetailsRequestLanguageTypescript ExecutionDetailsScriptExecutionDetailsRequestLanguage = "typescript"
	ExecutionDetailsScriptExecutionDetailsRequestLanguageRuby       ExecutionDetailsScriptExecutionDetailsRequestLanguage = "ruby"
	ExecutionDetailsScriptExecutionDetailsRequestLanguagePhp        ExecutionDetailsScriptExecutionDetailsRequestLanguage = "php"
)

func (r ExecutionDetailsScriptExecutionDetailsRequestLanguage) IsKnown() bool {
	switch r {
	case ExecutionDetailsScriptExecutionDetailsRequestLanguagePython, ExecutionDetailsScriptExecutionDetailsRequestLanguageJavascript, ExecutionDetailsScriptExecutionDetailsRequestLanguageTypescript, ExecutionDetailsScriptExecutionDetailsRequestLanguageRuby, ExecutionDetailsScriptExecutionDetailsRequestLanguagePhp:
		return true
	}
	return false
}

type ExecutionDetailsScriptExecutionDetailsRequestFile struct {
	// The contents of the file.
	Contents string `json:"contents"`
	// The relative path of the file.
	Path string                                                `json:"path"`
	JSON executionDetailsScriptExecutionDetailsRequestFileJSON `json:"-"`
}

// executionDetailsScriptExecutionDetailsRequestFileJSON contains the JSON metadata
// for the struct [ExecutionDetailsScriptExecutionDetailsRequestFile]
type executionDetailsScriptExecutionDetailsRequestFileJSON struct {
	Contents    apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsScriptExecutionDetailsRequestFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsScriptExecutionDetailsRequestFileJSON) RawJSON() string {
	return r.raw
}

// Configuration for HTTP requests and authentication.
type ExecutionDetailsScriptExecutionDetailsRequestHTTP struct {
	// List of allowed HTTP hosts and associated authentication.
	Allow []ExecutionDetailsScriptExecutionDetailsRequestHTTPAllow `json:"allow"`
	JSON  executionDetailsScriptExecutionDetailsRequestHTTPJSON    `json:"-"`
}

// executionDetailsScriptExecutionDetailsRequestHTTPJSON contains the JSON metadata
// for the struct [ExecutionDetailsScriptExecutionDetailsRequestHTTP]
type executionDetailsScriptExecutionDetailsRequestHTTPJSON struct {
	Allow       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsScriptExecutionDetailsRequestHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsScriptExecutionDetailsRequestHTTPJSON) RawJSON() string {
	return r.raw
}

// List of allowed HTTP hosts and associated authentication.
type ExecutionDetailsScriptExecutionDetailsRequestHTTPAllow struct {
	// Authentication configuration for outbound requests to this host.
	Auth ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuth `json:"auth"`
	// The hostname to allow.
	Host string                                                     `json:"host"`
	JSON executionDetailsScriptExecutionDetailsRequestHTTPAllowJSON `json:"-"`
}

// executionDetailsScriptExecutionDetailsRequestHTTPAllowJSON contains the JSON
// metadata for the struct [ExecutionDetailsScriptExecutionDetailsRequestHTTPAllow]
type executionDetailsScriptExecutionDetailsRequestHTTPAllowJSON struct {
	Auth        apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsScriptExecutionDetailsRequestHTTPAllow) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsScriptExecutionDetailsRequestHTTPAllowJSON) RawJSON() string {
	return r.raw
}

// Authentication configuration for outbound requests to this host.
type ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuth struct {
	Basic ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBasic `json:"basic"`
	// Configuration to add an 'Authorization' header using the 'Bearer' scheme.
	Bearer ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBearer `json:"bearer"`
	Header ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthHeader `json:"header"`
	Query  ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthQuery  `json:"query"`
	JSON   executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthJSON   `json:"-"`
}

// executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthJSON contains the JSON
// metadata for the struct
// [ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuth]
type executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthJSON struct {
	Basic       apijson.Field
	Bearer      apijson.Field
	Header      apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthJSON) RawJSON() string {
	return r.raw
}

type ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBasic struct {
	Password string                                                              `json:"password"`
	UserID   string                                                              `json:"user_id"`
	JSON     executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBasicJSON `json:"-"`
}

// executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBasicJSON contains the
// JSON metadata for the struct
// [ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBasic]
type executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBasicJSON struct {
	Password    apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBasic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBasicJSON) RawJSON() string {
	return r.raw
}

// Configuration to add an 'Authorization' header using the 'Bearer' scheme.
type ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBearer struct {
	// The token to set, e.g. 'Authorization: Bearer <token>'.
	Token string                                                               `json:"token"`
	JSON  executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBearerJSON `json:"-"`
}

// executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBearerJSON contains
// the JSON metadata for the struct
// [ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBearer]
type executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBearerJSON struct {
	Token       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBearer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthBearerJSON) RawJSON() string {
	return r.raw
}

type ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthHeader struct {
	Name  string                                                               `json:"name"`
	Value string                                                               `json:"value"`
	JSON  executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthHeaderJSON `json:"-"`
}

// executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthHeaderJSON contains
// the JSON metadata for the struct
// [ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthHeader]
type executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthHeaderJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthHeaderJSON) RawJSON() string {
	return r.raw
}

type ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthQuery struct {
	Key   string                                                              `json:"key"`
	Value string                                                              `json:"value"`
	JSON  executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthQueryJSON `json:"-"`
}

// executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthQueryJSON contains the
// JSON metadata for the struct
// [ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthQuery]
type executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthQueryJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsScriptExecutionDetailsRequestHTTPAllowAuthQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsScriptExecutionDetailsRequestHTTPAllowAuthQueryJSON) RawJSON() string {
	return r.raw
}

// Configuration for execution environment limits.
type ExecutionDetailsScriptExecutionDetailsRequestLimits struct {
	// The maximum time allowed for execution (in seconds). Default is 30.
	ExecutionTimeout int64 `json:"execution_timeout"`
	// The maximum memory allowed for execution (in MiB). Default is 128.
	MemorySize int64                                                   `json:"memory_size"`
	JSON       executionDetailsScriptExecutionDetailsRequestLimitsJSON `json:"-"`
}

// executionDetailsScriptExecutionDetailsRequestLimitsJSON contains the JSON
// metadata for the struct [ExecutionDetailsScriptExecutionDetailsRequestLimits]
type executionDetailsScriptExecutionDetailsRequestLimitsJSON struct {
	ExecutionTimeout apijson.Field
	MemorySize       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ExecutionDetailsScriptExecutionDetailsRequestLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsScriptExecutionDetailsRequestLimitsJSON) RawJSON() string {
	return r.raw
}

type ExecutionDetailsScriptExecutionDetailsResponse struct {
	// The ID of the execution.
	ID string `json:"id,required"`
	// The execution time of the script in milliseconds.
	Duration int64 `json:"duration,required"`
	// The exit code returned by the script. Will often be '0' on success and non-zero
	// on failure.
	ExitCode int64 `json:"exit_code,required"`
	// The contents of 'stderr' after executing the script.
	Stderr string `json:"stderr,required"`
	// The contents of 'stdout' after executing the script.
	Stdout string                                             `json:"stdout,required"`
	JSON   executionDetailsScriptExecutionDetailsResponseJSON `json:"-"`
}

// executionDetailsScriptExecutionDetailsResponseJSON contains the JSON metadata
// for the struct [ExecutionDetailsScriptExecutionDetailsResponse]
type executionDetailsScriptExecutionDetailsResponseJSON struct {
	ID          apijson.Field
	Duration    apijson.Field
	ExitCode    apijson.Field
	Stderr      apijson.Field
	Stdout      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecutionDetailsScriptExecutionDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executionDetailsScriptExecutionDetailsResponseJSON) RawJSON() string {
	return r.raw
}

type ExecutionDetailsScriptExecutionDetailsType string

const (
	ExecutionDetailsScriptExecutionDetailsTypeScript ExecutionDetailsScriptExecutionDetailsType = "script"
)

func (r ExecutionDetailsScriptExecutionDetailsType) IsKnown() bool {
	switch r {
	case ExecutionDetailsScriptExecutionDetailsTypeScript:
		return true
	}
	return false
}

type ExecutionDetailsType string

const (
	ExecutionDetailsTypeTool     ExecutionDetailsType = "tool"
	ExecutionDetailsTypeFunction ExecutionDetailsType = "function"
	ExecutionDetailsTypeScript   ExecutionDetailsType = "script"
)

func (r ExecutionDetailsType) IsKnown() bool {
	switch r {
	case ExecutionDetailsTypeTool, ExecutionDetailsTypeFunction, ExecutionDetailsTypeScript:
		return true
	}
	return false
}

type ExecutionListParams struct {
	// The number of items to return. Defaults to 100. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// If true, only show executions where the exit code is not 0, indicating an
	// execution error. Defaults to false.
	OnlyNonZeroExitCodes param.Field[bool] `query:"only_non_zero_exit_codes"`
	// The ID of the item to start after. To get the next page of results, set this to
	// the ID of the last item in the current page.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [ExecutionListParams]'s query parameters as `url.Values`.
func (r ExecutionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
