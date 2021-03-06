package httphelper

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//QueryInfo rep
type QueryInfo struct {
	Omitempty     bool
	QueryKey      string //QueryKey is the property of model represented in the URL query param
	DBKey         string //DBKey is the property of model represented in the database
	Kind          reflect.Kind
	Transform     QueryTransform
	SortTransform QueryTransform
}

//QueryTransform query into another value
type QueryTransform func(key string, value interface{}) (string, interface{})

// ReadPagingInfo get paging info from request
func ReadPagingInfo(r *http.Request) (pageNumber int, pageSize int, err error) {
	page := r.FormValue("page")
	size := r.FormValue("size")

	pageNumber, err = strconv.Atoi(page)
	if err != nil {
		return 0, 0, fmt.Errorf("ReadPagingInfo failed parsing pageNumber : " + err.Error())
	}
	pageSize, err = strconv.Atoi(size)
	if err != nil {
		return 0, 0, fmt.Errorf("ReadPagingInfo failed : pageSize " + err.Error())
	}
	return
}

//ReadQuery parameter
func ReadQuery(r *http.Request, p []QueryInfo) (res map[string]interface{}, err error) {
	res = map[string]interface{}{}
	for _, q := range p {
		//if QueryKey is empty, then defaults to DBKey
		jskey := q.QueryKey
		if jskey == "" {
			jskey = q.DBKey
		}

		qv := r.URL.Query().Get(jskey)
		if qv == "" && q.Omitempty {
			continue
		}

		var dbkey = q.DBKey
		var val interface{}
		switch q.Kind {
		case reflect.Bool:
			val, err = strconv.ParseBool(qv)
		case reflect.Float32, reflect.Float64:
			val, err = strconv.ParseFloat(qv, 64)
		case reflect.Int, reflect.Int32, reflect.Int64:
			val, err = strconv.Atoi(qv)
		case reflect.Slice:
			val = strings.Split(qv, ",")
		default:
			val = qv
		}

		if err != nil {
			return nil, err
		}

		if q.Transform != nil {
			dbkey, val = q.Transform(dbkey, val)
		}

		res[dbkey] = val
	}

	return res, nil
}

//ReadSorting parameter for read sorting parameter
func ReadSorting(r *http.Request, p []QueryInfo) map[string]interface{} {
	sortString := r.FormValue("sort")
	if len(sortString) == 0 {
		return nil
	}

	//Transform sort string from `sort=asc(field_a),desc(field_b)` to
	//{field_a: asc, field_b: asc}
	sortMap := parseFunctionString(sortString)
	result := make(map[string]interface{})
	for _, qi := range p {
		urlKey := qi.QueryKey
		if urlKey == "" {
			urlKey = qi.DBKey
		}

		dbkey := qi.DBKey
		val, exists := sortMap[urlKey]
		if !exists {
			continue
		}

		if qi.SortTransform != nil {
			dbkey, val = qi.SortTransform(dbkey, val)
		}

		result[dbkey] = val
	}

	return result
}

//ParseFunctionString parsing function in string
func parseFunctionString(val string) (sorting map[string]interface{}) {
	//creating regexp pattern.
	// pattern desc:
	// Match a single character present in the list below [\((.*?)]
	// \( matches the character ( literally (case sensitive)
	// (.*?) matches a single character in the list (.*?) (case sensitive)
	// Global pattern flags
	// g modifier: global. All matches (don't return after first match)
	// Human Lang : Split string dengan parameter `(` `)`
	sorting = map[string]interface{}{}
	regexpPattern := regexp.MustCompile(`[\((*?)(\),)]`)
	filterPayload := regexpPattern.Split(val, -1)
	for i := 0; i < len(filterPayload); i = i + 3 {
		sorting[filterPayload[i+1]] = filterPayload[i]
	}
	return sorting
}
