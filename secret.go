// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package riza

import (
	"context"
	"net/http"
	"net/url"

	"github.com/riza-io/riza-api-go/internal/apijson"
	"github.com/riza-io/riza-api-go/internal/apiquery"
	"github.com/riza-io/riza-api-go/internal/param"
	"github.com/riza-io/riza-api-go/internal/requestconfig"
	"github.com/riza-io/riza-api-go/option"
	"github.com/riza-io/riza-api-go/packages/pagination"
)

// SecretService contains methods and other services that help with interacting
// with the riza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecretService] method instead.
type SecretService struct {
	Options []option.RequestOption
}

// NewSecretService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSecretService(opts ...option.RequestOption) (r *SecretService) {
	r = &SecretService{}
	r.Options = opts
	return
}

// Create a secret in your project.
func (r *SecretService) New(ctx context.Context, body SecretNewParams, opts ...option.RequestOption) (res *Secret, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of secrets in your project.
func (r *SecretService) List(ctx context.Context, query SecretListParams, opts ...option.RequestOption) (res *pagination.SecretsPagination[Secret], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v1/secrets"
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

// Returns a list of secrets in your project.
func (r *SecretService) ListAutoPaging(ctx context.Context, query SecretListParams, opts ...option.RequestOption) *pagination.SecretsPaginationAutoPager[Secret] {
	return pagination.NewSecretsPaginationAutoPager(r.List(ctx, query, opts...))
}

type Secret struct {
	ID   string     `json:"id,required"`
	Name string     `json:"name,required"`
	JSON secretJSON `json:"-"`
}

// secretJSON contains the JSON metadata for the struct [Secret]
type secretJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Secret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretJSON) RawJSON() string {
	return r.raw
}

type SecretNewParams struct {
	Name  param.Field[string] `json:"name,required"`
	Value param.Field[string] `json:"value,required"`
}

func (r SecretNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecretListParams struct {
	// The number of items to return. Defaults to 100. Maximum is 100.
	Limit param.Field[int64] `query:"limit"`
	// The ID of the item to start after. To get the next page of results, set this to
	// the ID of the last item in the current page.
	StartingAfter param.Field[string] `query:"starting_after"`
}

// URLQuery serializes [SecretListParams]'s query parameters as `url.Values`.
func (r SecretListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
