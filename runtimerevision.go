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

// RuntimeRevisionService contains methods and other services that help with
// interacting with the riza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuntimeRevisionService] method instead.
type RuntimeRevisionService struct {
	Options []option.RequestOption
}

// NewRuntimeRevisionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRuntimeRevisionService(opts ...option.RequestOption) (r *RuntimeRevisionService) {
	r = &RuntimeRevisionService{}
	r.Options = opts
	return
}

// Creates a new runtime revision.
func (r *RuntimeRevisionService) New(ctx context.Context, id string, body RuntimeRevisionNewParams, opts ...option.RequestOption) (res *Revision, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/runtimes/%s/revisions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all revisions for a runtime.
func (r *RuntimeRevisionService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *RuntimeRevisionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/runtimes/%s/revisions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a runtime revision.
func (r *RuntimeRevisionService) Get(ctx context.Context, runtimeID string, revisionID string, opts ...option.RequestOption) (res *Revision, err error) {
	opts = append(r.Options[:], opts...)
	if runtimeID == "" {
		err = errors.New("missing required runtime_id parameter")
		return
	}
	if revisionID == "" {
		err = errors.New("missing required revision_id parameter")
		return
	}
	path := fmt.Sprintf("v1/runtimes/%s/revisions/%s", runtimeID, revisionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Revision struct {
	ID                      string               `json:"id,required"`
	ManifestFile            RevisionManifestFile `json:"manifest_file,required"`
	RuntimeID               string               `json:"runtime_id,required"`
	Status                  RevisionStatus       `json:"status,required"`
	AdditionalPythonImports string               `json:"additional_python_imports"`
	JSON                    revisionJSON         `json:"-"`
}

// revisionJSON contains the JSON metadata for the struct [Revision]
type revisionJSON struct {
	ID                      apijson.Field
	ManifestFile            apijson.Field
	RuntimeID               apijson.Field
	Status                  apijson.Field
	AdditionalPythonImports apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *Revision) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r revisionJSON) RawJSON() string {
	return r.raw
}

type RevisionManifestFile struct {
	Contents string                   `json:"contents,required"`
	Name     RevisionManifestFileName `json:"name,required"`
	JSON     revisionManifestFileJSON `json:"-"`
}

// revisionManifestFileJSON contains the JSON metadata for the struct
// [RevisionManifestFile]
type revisionManifestFileJSON struct {
	Contents    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RevisionManifestFile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r revisionManifestFileJSON) RawJSON() string {
	return r.raw
}

type RevisionManifestFileName string

const (
	RevisionManifestFileNameRequirementsTxt RevisionManifestFileName = "requirements.txt"
	RevisionManifestFileNamePackageJson     RevisionManifestFileName = "package.json"
)

func (r RevisionManifestFileName) IsKnown() bool {
	switch r {
	case RevisionManifestFileNameRequirementsTxt, RevisionManifestFileNamePackageJson:
		return true
	}
	return false
}

type RevisionStatus string

const (
	RevisionStatusPending   RevisionStatus = "pending"
	RevisionStatusBuilding  RevisionStatus = "building"
	RevisionStatusSucceeded RevisionStatus = "succeeded"
	RevisionStatusFailed    RevisionStatus = "failed"
	RevisionStatusCancelled RevisionStatus = "cancelled"
)

func (r RevisionStatus) IsKnown() bool {
	switch r {
	case RevisionStatusPending, RevisionStatusBuilding, RevisionStatusSucceeded, RevisionStatusFailed, RevisionStatusCancelled:
		return true
	}
	return false
}

type RuntimeRevisionListResponse struct {
	Revisions []Revision                      `json:"revisions,required"`
	JSON      runtimeRevisionListResponseJSON `json:"-"`
}

// runtimeRevisionListResponseJSON contains the JSON metadata for the struct
// [RuntimeRevisionListResponse]
type runtimeRevisionListResponseJSON struct {
	Revisions   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimeRevisionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimeRevisionListResponseJSON) RawJSON() string {
	return r.raw
}

type RuntimeRevisionNewParams struct {
	ManifestFile            param.Field[RuntimeRevisionNewParamsManifestFile] `json:"manifest_file,required"`
	AdditionalPythonImports param.Field[string]                               `json:"additional_python_imports"`
}

func (r RuntimeRevisionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuntimeRevisionNewParamsManifestFile struct {
	Contents param.Field[string]                                   `json:"contents,required"`
	Name     param.Field[RuntimeRevisionNewParamsManifestFileName] `json:"name,required"`
}

func (r RuntimeRevisionNewParamsManifestFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuntimeRevisionNewParamsManifestFileName string

const (
	RuntimeRevisionNewParamsManifestFileNameRequirementsTxt RuntimeRevisionNewParamsManifestFileName = "requirements.txt"
	RuntimeRevisionNewParamsManifestFileNamePackageJson     RuntimeRevisionNewParamsManifestFileName = "package.json"
)

func (r RuntimeRevisionNewParamsManifestFileName) IsKnown() bool {
	switch r {
	case RuntimeRevisionNewParamsManifestFileNameRequirementsTxt, RuntimeRevisionNewParamsManifestFileNamePackageJson:
		return true
	}
	return false
}
