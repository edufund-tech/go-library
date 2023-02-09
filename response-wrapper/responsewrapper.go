package responsewrapper

import (
	"encoding/json"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type page struct {
	Total   int64 `json:"_total,omitempty"`
	Current int64 `json:"_current,omitempty"`
}

type index struct {
	First int64 `json:"_first,omitempty"`
	Last  int64 `json:"_last,omitempty"`
}

type pagination struct {
	Page  page  `json:"_page,omitempty"`
	Index index `json:"_index,omitempty"`
}

type urlMeta struct {
	Href string `json:"_href,omitempty"`
}

type links struct {
	Self  urlMeta `json:"_self,omitempty"`
	First urlMeta `json:"_first,omitempty"`
	Prev  urlMeta `json:"_prev,omitempty"`
	Next  urlMeta `json:"_next,omitempty"`
	Last  urlMeta `json:"_last,omitempty"`
}
type meta struct {
	TotalData  int64      `json:"_total_data,omitempty"`
	Pagination pagination `json:"_pagination,omitempty"`
	Links      links      `json:"_links,omitempty"`
}

type Wrapper struct {
	Data interface{} `json:"data,omitempty"`

	Meta       meta        `json:"meta,omitempty"`
	Error      interface{} `json:"error,omitempty"`
	Extensions interface{} `json:"extensions,omitempty"`
	Message    string      `json:"message,omitempty"`
	Code       int         `json:"-"`
}

func (w *Wrapper) Respond(wr http.ResponseWriter) {
	b, _ := json.Marshal(w)
	wr.Header().Add("Content-Type", "application/json")
	wr.WriteHeader(w.Code)
	wr.Write(b)
}

// AddPaging
func (w *Wrapper) AddPaging(totaldata, limit, page int64) *Wrapper {

	w.Meta.Pagination.Page.Total = int64(math.Ceil(float64(totaldata) / float64(limit)))
	w.Meta.Pagination.Page.Current = page
	w.Meta.Pagination.Index.Last = limit * page
	w.Meta.Pagination.Index.First = w.Meta.Pagination.Index.Last - limit + 1

	return w
}

// AddLinks FF links
func (w *Wrapper) AddLinks(values url.Values) *Wrapper {

	w.Meta.Links.Self = urlMeta{Href: values.Encode()}
	//set first
	values.Set("page", "1")
	w.Meta.Links.First = urlMeta{Href: values.Encode()}

	//set prev
	if w.Meta.Pagination.Page.Current > 1 {
		values.Set("page", strconv.FormatInt(w.Meta.Pagination.Page.Current-1, 10))
	}
	w.Meta.Links.Prev = urlMeta{Href: values.Encode()}

	//set last
	values.Set("page", strconv.FormatInt(w.Meta.Pagination.Page.Total, 10))
	w.Meta.Links.Last = urlMeta{Href: values.Encode()}

	//set next
	if w.Meta.Pagination.Page.Total > w.Meta.Pagination.Page.Current {
		values.Set("page", strconv.FormatInt(w.Meta.Pagination.Page.Current+1, 10))
		w.Meta.Links.Next = urlMeta{Href: values.Encode()}
	}
	return w
}

// AddMeta added meta list
func (w *Wrapper) AddMeta(r *http.Request, totaldata, limit, page int64) *Wrapper {
	values, _ := url.ParseQuery(r.URL.RawQuery)
	w.AddPaging(totaldata, limit, page)
	w.AddLinks(values)
	w.Meta = meta{
		TotalData:  totaldata,
		Pagination: w.Meta.Pagination,
		Links:      w.Meta.Links,
	}
	return w
}
