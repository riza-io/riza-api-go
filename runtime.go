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

// RuntimeService contains methods and other services that help with interacting
// with the riza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuntimeService] method instead.
type RuntimeService struct {
	Options   []option.RequestOption
	Revisions *RuntimeRevisionService
}

// NewRuntimeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuntimeService(opts ...option.RequestOption) (r *RuntimeService) {
	r = &RuntimeService{}
	r.Options = opts
	r.Revisions = NewRuntimeRevisionService(opts...)
	return
}

// Creates a runtime.
func (r *RuntimeService) New(ctx context.Context, body RuntimeNewParams, opts ...option.RequestOption) (res *Runtime, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/runtimes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of runtimes in your project.
func (r *RuntimeService) List(ctx context.Context, opts ...option.RequestOption) (res *RuntimeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/runtimes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a runtime.
func (r *RuntimeService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Runtime, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/runtimes/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Runtime struct {
	ID                      string              `json:"id,required"`
	Engine                  RuntimeEngine       `json:"engine,required"`
	Language                RuntimeLanguage     `json:"language,required"`
	Name                    string              `json:"name,required"`
	RevisionID              string              `json:"revision_id,required"`
	Status                  RuntimeStatus       `json:"status,required"`
	AdditionalPythonImports string              `json:"additional_python_imports"`
	ManifestFile            RuntimeManifestFile `json:"manifest_file"`
	JSON                    runtimeJSON         `json:"-"`
}

// runtimeJSON contains the JSON metadata for the struct [Runtime]
type runtimeJSON struct {
	ID                      apijson.Field
	Engine                  apijson.Field
	Language                apijson.Field
	Name                    apijson.Field
	RevisionID              apijson.Field
	Status                  apijson.Field
	AdditionalPythonImports apijson.Field
	ManifestFile            apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *Runtime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeJSON) RawJSON() string {
	return r.raw
}

type RuntimeEngine string

const (
	RuntimeEngineWasi    RuntimeEngine = "wasi"
	RuntimeEngineMicrovm RuntimeEngine = "microvm"
	RuntimeEngineV8      RuntimeEngine = "v8"
)

func (r RuntimeEngine) IsKnown() bool {
	switch r {
	case RuntimeEngineWasi, RuntimeEngineMicrovm, RuntimeEngineV8:
		return true
	}
	return false
}

type RuntimeLanguage string

const (
	RuntimeLanguagePython     RuntimeLanguage = "python"
	RuntimeLanguageJavascript RuntimeLanguage = "javascript"
)

func (r RuntimeLanguage) IsKnown() bool {
	switch r {
	case RuntimeLanguagePython, RuntimeLanguageJavascript:
		return true
	}
	return false
}

type RuntimeStatus string

const (
	RuntimeStatusPending   RuntimeStatus = "pending"
	RuntimeStatusBuilding  RuntimeStatus = "building"
	RuntimeStatusSucceeded RuntimeStatus = "succeeded"
	RuntimeStatusFailed    RuntimeStatus = "failed"
	RuntimeStatusCancelled RuntimeStatus = "cancelled"
)

func (r RuntimeStatus) IsKnown() bool {
	switch r {
	case RuntimeStatusPending, RuntimeStatusBuilding, RuntimeStatusSucceeded, RuntimeStatusFailed, RuntimeStatusCancelled:
		return true
	}
	return false
}

type RuntimeManifestFile struct {
	Contents string                  `json:"contents,required"`
	Name     RuntimeManifestFileName `json:"name,required"`
	JSON     runtimeManifestFileJSON `json:"-"`
}

// runtimeManifestFileJSON contains the JSON metadata for the struct
// [RuntimeManifestFile]
type runtimeManifestFileJSON struct {
	Contents    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeManifestFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeManifestFileJSON) RawJSON() string {
	return r.raw
}

type RuntimeManifestFileName string

const (
	RuntimeManifestFileNameRequirementsTxt RuntimeManifestFileName = "requirements.txt"
	RuntimeManifestFileNamePackageJson     RuntimeManifestFileName = "package.json"
)

func (r RuntimeManifestFileName) IsKnown() bool {
	switch r {
	case RuntimeManifestFileNameRequirementsTxt, RuntimeManifestFileNamePackageJson:
		return true
	}
	return false
}

type RuntimeListResponse struct {
	Runtimes []Runtime               `json:"runtimes,required"`
	JSON     runtimeListResponseJSON `json:"-"`
}

// runtimeListResponseJSON contains the JSON metadata for the struct
// [RuntimeListResponse]
type runtimeListResponseJSON struct {
	Runtimes    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeListResponseJSON) RawJSON() string {
	return r.raw
}

type RuntimeNewParams struct {
	Language                param.Field[RuntimeNewParamsLanguage]     `json:"language,required"`
	ManifestFile            param.Field[RuntimeNewParamsManifestFile] `json:"manifest_file,required"`
	Name                    param.Field[string]                       `json:"name,required"`
	AdditionalPythonImports param.Field[string]                       `json:"additional_python_imports"`
	Engine                  param.Field[RuntimeNewParamsEngine]       `json:"engine"`
}

func (r RuntimeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuntimeNewParamsLanguage string

const (
	RuntimeNewParamsLanguagePython     RuntimeNewParamsLanguage = "python"
	RuntimeNewParamsLanguageJavascript RuntimeNewParamsLanguage = "javascript"
)

func (r RuntimeNewParamsLanguage) IsKnown() bool {
	switch r {
	case RuntimeNewParamsLanguagePython, RuntimeNewParamsLanguageJavascript:
		return true
	}
	return false
}

type RuntimeNewParamsManifestFile struct {
	Contents param.Field[string]                           `json:"contents,required"`
	Name     param.Field[RuntimeNewParamsManifestFileName] `json:"name,required"`
}

func (r RuntimeNewParamsManifestFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuntimeNewParamsManifestFileName string

const (
	RuntimeNewParamsManifestFileNameRequirementsTxt RuntimeNewParamsManifestFileName = "requirements.txt"
	RuntimeNewParamsManifestFileNamePackageJson     RuntimeNewParamsManifestFileName = "package.json"
)

func (r RuntimeNewParamsManifestFileName) IsKnown() bool {
	switch r {
	case RuntimeNewParamsManifestFileNameRequirementsTxt, RuntimeNewParamsManifestFileNamePackageJson:
		return true
	}
	return false
}

type RuntimeNewParamsEngine string

const (
	RuntimeNewParamsEngineWasi    RuntimeNewParamsEngine = "wasi"
	RuntimeNewParamsEngineMicrovm RuntimeNewParamsEngine = "microvm"
	RuntimeNewParamsEngineV8      RuntimeNewParamsEngine = "v8"
)

func (r RuntimeNewParamsEngine) IsKnown() bool {
	switch r {
	case RuntimeNewParamsEngineWasi, RuntimeNewParamsEngineMicrovm, RuntimeNewParamsEngineV8:
		return true
	}
	return false
}
