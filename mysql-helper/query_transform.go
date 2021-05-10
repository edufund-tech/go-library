package mysqlhelper

import "fmt"

//LikeQuery like query
func LikeQuery(key string, value interface{}) (string, interface{}) {
	key = key + " LIKE ?"
	value = fmt.Sprintf("%%%s%%", value)

	return key, value
}
