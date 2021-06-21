package mysqlhelper

import (
	"bytes"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func ApplySqlQuery(db *gorm.DB, filters map[string]interface{}) {
	for key, value := range filters {
		switch {
		case strings.Index(value.(string), "range(") == 0:
			query := fmt.Sprint(key, " BETWEEN ? and ?")
			parseStr := value.(string)[6:]
			splitStr := strings.Split(parseStr[:len(parseStr)-1], ",")
			db.Where(query, splitStr[0], splitStr[1])
		case strings.Index(value.(string), "in(") == 0:
			query := fmt.Sprint(key, " In (?)")
			parseStr := value.(string)[3:]
			splitStr := strings.Split(parseStr[:len(parseStr)-1], ",")
			db.Where(query, splitStr)
		case strings.Index(value.(string), "like(") == 0:
			query := fmt.Sprint(key, " like ?")
			parseStr := value.(string)[5:]
			splitStr := strings.Split(parseStr[:len(parseStr)-1], ",")
			valueStr := fmt.Sprintf("%%%s%%", splitStr[0])
			db.Where(query, valueStr)
		case strings.Index(value.(string), "not(") == 0:
			query := fmt.Sprint(key, " <> ?")
			parseStr := value.(string)[4:]
			splitStr := strings.Split(parseStr[:len(parseStr)-1], ",")
			db.Where(query, splitStr)
		case strings.Index(value.(string), "gt(") == 0:
			query := fmt.Sprint(key, " > ?")
			parseStr := value.(string)[3:]
			splitStr := strings.Split(parseStr[:len(parseStr)-1], ",")
			db.Where(query, splitStr)
		case strings.Index(value.(string), "lt(") == 0:
			query := fmt.Sprint(key, " < ?")
			parseStr := value.(string)[3:]
			splitStr := strings.Split(parseStr[:len(parseStr)-1], ",")
			db.Where(query, splitStr)
		case strings.Index(value.(string), "not_in(") == 0:
			query := fmt.Sprint(key, " not in (?)")
			parseStr := value.(string)[7:]
			splitStr := strings.Split(parseStr[:len(parseStr)-1], ",")
			db.Where(query, splitStr)
		default:
			db.Where(key, value)
		}
	}
}

func ApplySorting(db *gorm.DB, sorts map[string]interface{}) {
	b := new(bytes.Buffer)
	for key, value := range sorts {
		fmt.Fprintf(b, "%s %s ", key, value)
	}
	db.Order(b.String())
}
