// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package riza

import (
	"context"
	"net/http"

	"github.com/riza-io/riza-api-go/internal/apijson"
	"github.com/riza-io/riza-api-go/internal/requestconfig"
	"github.com/riza-io/riza-api-go/option"
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

func (r *SecretService) List(ctx context.Context, opts ...option.RequestOption) (res *SecretListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
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

type SecretListResponse struct {
	Secrets []Secret               `json:"secrets,required"`
	JSON    secretListResponseJSON `json:"-"`
}

// secretListResponseJSON contains the JSON metadata for the struct
// [SecretListResponse]
type secretListResponseJSON struct {
	Secrets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretListResponseJSON) RawJSON() string {
	return r.raw
}
