// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"reflect"

	"github.com/riza-io/riza-api-go/internal/apijson"
	"github.com/riza-io/riza-api-go/internal/requestconfig"
	"github.com/riza-io/riza-api-go/option"
)

type RuntimesPagination[T any] struct {
	Runtimes []T                    `json:"runtimes"`
	JSON     runtimesPaginationJSON `json:"-"`
	cfg      *requestconfig.RequestConfig
	res      *http.Response
}

// runtimesPaginationJSON contains the JSON metadata for the struct
// [RuntimesPagination[T]]
type runtimesPaginationJSON struct {
	Runtimes    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuntimesPagination[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r runtimesPaginationJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *RuntimesPagination[T]) GetNextPage() (res *RuntimesPagination[T], err error) {
	items := r.Runtimes
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	cfg.Apply(option.WithQuery("starting_after", field.Interface().(string)))
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *RuntimesPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &RuntimesPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type RuntimesPaginationAutoPager[T any] struct {
	page *RuntimesPagination[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewRuntimesPaginationAutoPager[T any](page *RuntimesPagination[T], err error) *RuntimesPaginationAutoPager[T] {
	return &RuntimesPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *RuntimesPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Runtimes) == 0 {
		return false
	}
	if r.idx >= len(r.page.Runtimes) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Runtimes) == 0 {
			return false
		}
	}
	r.cur = r.page.Runtimes[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *RuntimesPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *RuntimesPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *RuntimesPaginationAutoPager[T]) Index() int {
	return r.run
}

type ToolsPagination[T any] struct {
	Tools []T                 `json:"tools"`
	JSON  toolsPaginationJSON `json:"-"`
	cfg   *requestconfig.RequestConfig
	res   *http.Response
}

// toolsPaginationJSON contains the JSON metadata for the struct
// [ToolsPagination[T]]
type toolsPaginationJSON struct {
	Tools       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ToolsPagination[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r toolsPaginationJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *ToolsPagination[T]) GetNextPage() (res *ToolsPagination[T], err error) {
	items := r.Tools
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	cfg.Apply(option.WithQuery("starting_after", field.Interface().(string)))
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ToolsPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &ToolsPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type ToolsPaginationAutoPager[T any] struct {
	page *ToolsPagination[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewToolsPaginationAutoPager[T any](page *ToolsPagination[T], err error) *ToolsPaginationAutoPager[T] {
	return &ToolsPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ToolsPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Tools) == 0 {
		return false
	}
	if r.idx >= len(r.page.Tools) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Tools) == 0 {
			return false
		}
	}
	r.cur = r.page.Tools[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ToolsPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *ToolsPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *ToolsPaginationAutoPager[T]) Index() int {
	return r.run
}

type SecretsPagination[T any] struct {
	Secrets []T                   `json:"secrets"`
	JSON    secretsPaginationJSON `json:"-"`
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// secretsPaginationJSON contains the JSON metadata for the struct
// [SecretsPagination[T]]
type secretsPaginationJSON struct {
	Secrets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecretsPagination[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secretsPaginationJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *SecretsPagination[T]) GetNextPage() (res *SecretsPagination[T], err error) {
	items := r.Secrets
	if items == nil || len(items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	value := reflect.ValueOf(items[len(items)-1])
	field := value.FieldByName("ID")
	cfg.Apply(option.WithQuery("starting_after", field.Interface().(string)))
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *SecretsPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &SecretsPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type SecretsPaginationAutoPager[T any] struct {
	page *SecretsPagination[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewSecretsPaginationAutoPager[T any](page *SecretsPagination[T], err error) *SecretsPaginationAutoPager[T] {
	return &SecretsPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *SecretsPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Secrets) == 0 {
		return false
	}
	if r.idx >= len(r.page.Secrets) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Secrets) == 0 {
			return false
		}
	}
	r.cur = r.page.Secrets[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *SecretsPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *SecretsPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *SecretsPaginationAutoPager[T]) Index() int {
	return r.run
}
