package repository

import (
	"MockExamLab/internal/models"
	"MockExamLab/internal/models/consts"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
	"regexp"
	"strings"
)

var (
	columnNameRegexp = regexp.MustCompile(`(?m)column:(\w{1,}).*`)
	paramNameRegexp  = regexp.MustCompile(`(?m)param:(\w{1,}).*`)
)

func orderBy(db *gorm.DB, params models.QueryParams) *gorm.DB {
	return db.Order(clause.OrderByColumn{
		Column: clause.Column{Name: params.OrderBy},
		Desc:   params.OrderDirection == "desc"},
	)
}

func paginate(db *gorm.DB, params models.QueryParams) *gorm.DB {
	if params.All {
		return db
	}

	if params.Page == 0 {
		params.Page = 1
	}

	switch {
	case params.PageSize > 100:
		params.PageSize = 100
	case params.PageSize <= 0:
		params.PageSize = 10
	}

	offset := (params.Page - 1) * params.PageSize
	return db.Offset(offset).Limit(params.PageSize)
}

func getColumnNameForField(field reflect.StructField) string {
	fieldTag := field.Tag.Get("gorm")
	res := columnNameRegexp.FindStringSubmatch(fieldTag)
	if len(res) == 2 {
		return res[1]
	}
	return field.Name
}

func searchField(field reflect.StructField, phrase string) clause.Expression {
	filterTag := field.Tag.Get(consts.TagKey)
	columnName := getColumnNameForField(field)
	if strings.Contains(filterTag, "searchable") {
		return clause.Like{Column: columnName, Value: "%" + phrase + "%"}
	}
	return nil
}

func filterField(field reflect.StructField, phrase string) clause.Expression {
	var paramName string
	if !strings.Contains(field.Tag.Get(consts.TagKey), "filterable") {
		return nil
	}
	columnName := getColumnNameForField(field)
	paramMatch := paramNameRegexp.FindStringSubmatch(field.Tag.Get(consts.TagKey))
	if len(paramMatch) == 2 {
		paramName = paramMatch[1]
	} else {
		paramName = columnName
	}
	re, err := regexp.Compile(fmt.Sprintf(`(?m)%v:(\w{1,}).*`, paramName))
	if err != nil {
		return nil
	}
	filterSubPhraseMatch := re.FindStringSubmatch(phrase)
	if len(filterSubPhraseMatch) == 2 {
		return clause.Eq{Column: columnName, Value: filterSubPhraseMatch[1]}
	}
	return nil
}

func expressionByField(
	db *gorm.DB, phrase string, modelType reflect.Type,
	operator func(reflect.StructField, string) clause.Expression,
	predicate func(...clause.Expression) clause.Expression,
) *gorm.DB {
	numFields := modelType.NumField()
	expressions := make([]clause.Expression, 0, numFields)
	for i := 0; i < numFields; i++ {
		field := modelType.Field(i)
		expression := operator(field, phrase)
		if expression != nil {
			expressions = append(expressions, expression)
		}
	}
	if len(expressions) == 1 {
		db = db.Where(expressions[0])
	} else if len(expressions) > 1 {
		db = db.Where(predicate(expressions...))
	}
	return db
}

func FilterByQuery(params *models.QueryParams, config int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if params == nil {
			return db
		}

		model := db.Statement.Model
		modelType := reflect.TypeOf(model)
		if model != nil && modelType.Kind() == reflect.Ptr && modelType.Elem().Kind() == reflect.Struct {
			if config&consts.SEARCH > 0 && params.Search != "" {
				db = expressionByField(db, params.Search, modelType.Elem(), searchField, clause.Or)
			}
			if config&consts.FILTER > 0 && params.Filter != "" {
				db = expressionByField(db, params.Filter, modelType.Elem(), filterField, clause.And)
			}
		}

		if config&consts.ORDER_BY > 0 {
			db = orderBy(db, *params)
		}
		if config&consts.PAGINATE > 0 {
			db = paginate(db, *params)
		}
		return db
	}
}
