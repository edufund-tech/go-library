package responsewrapper

import (
	"encoding/json"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

type page struct {
	Total   int64 `json:"_total"`
	Current int64 `json:"_current"`
}

type index struct {
	First int64 `json:"_first"`
	Last  int64 `json:"_last"`
}

type pagination struct {
	Page  page  `json:"_page"`
	Index index `json:"_index"`
}

type urlMeta struct {
	Href string `json:"_href"`
}

type links struct {
	Self  urlMeta `json:"_self"`
	First urlMeta `json:"_first"`
	Prev  urlMeta `json:"_prev"`
	Next  urlMeta `json:"_next"`
	Last  urlMeta `json:"_last"`
}
type meta struct {
	TotalData  int64      `json:"_total_data"`
	Pagination pagination `json:"_pagination"`
	Links      links      `json:"_links"`
}

type wrapper struct {
	Data interface{} `json:"data,omitempty"`

	Meta    *meta       `json:"meta"`
	Error   interface{} `json:"error"`
	Message string      `json:"message"`
	Code    int         `json:"-"`
}

func (w wrapper) respond(wr http.ResponseWriter) {
	b, _ := json.Marshal(w)
	wr.Header().Add("Content-Type", "application/json")
	wr.WriteHeader(w.Code)
	wr.Write(b)
}

//AddPaging
func (m meta) AddPaging(totaldata, limit, page int64) meta {

	m.Pagination.Page.Total = int64(math.Ceil(float64(totaldata) / float64(limit)))
	m.Pagination.Page.Current = page
	m.Pagination.Index.Last = limit * page
	m.Pagination.Index.First = m.Pagination.Index.Last - limit + 1

	return m
}

//AddLinks FF links
func (m meta) AddLinks(values url.Values) meta {

	m.Links.Self = urlMeta{Href: values.Encode()}
	//set first
	values.Set("page", "1")
	m.Links.First = urlMeta{Href: values.Encode()}

	//set prev
	if m.Pagination.Page.Current > 1 {
		values.Set("page", strconv.FormatInt(m.Pagination.Page.Current, 10))
	}
	m.Links.Prev = urlMeta{Href: values.Encode()}

	//set last
	values.Set("page", strconv.FormatInt(m.Pagination.Page.Total, 10))
	m.Links.Last = urlMeta{Href: values.Encode()}

	//set next
	if m.Pagination.Page.Total > m.Pagination.Page.Current {
		values.Set("page", strconv.FormatInt(m.Pagination.Page.Current+1, 10))
		m.Links.Next = urlMeta{Href: values.Encode()}
	}
	return m
}

//AddMeta added meta list
func (w wrapper) AddMeta(r *http.Request, totaldata, limit, page int64) wrapper {
	values, _ := url.ParseQuery(r.URL.RawQuery)
	w.Meta.AddPaging(totaldata, limit, page)
	w.Meta.AddLinks(values)
	m := meta{
		TotalData:  w.Meta.TotalData,
		Pagination: w.Meta.Pagination,
		Links:      w.Meta.Links,
	}

	w.Meta = &m
	return w
}
