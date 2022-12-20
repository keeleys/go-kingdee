package erp

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/keeleys/go-kingdee/util"
)

type ErpQuery[T any] struct {
	queryModel ErpQueryModel
	columns    []string
	list       []T
}

type ErpQueryModel struct {
	FieldKeys    string
	FilterString string
	FormId       string
	OrderString  string
	StartRow     int
	Limit        int
	TopRowCount  int
}

func (query *ErpQuery[T]) GetList() []T {
	return query.list
}
func (query *ErpQuery[T]) GetModel() ErpQueryModel {
	return query.queryModel
}

func (query *ErpQuery[T]) selectList(sessionId string) []T {
	if query.list != nil {
		return query.list
	}
	var b [][]any
	var err error

	jsonStr := erpPost(sessionId, LIST_URL.String(), "", query.queryModel)
	err = checkQueryError(jsonStr)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal([]byte(jsonStr), &b); err != nil {
		panic(err)
	}
	resultList := make([]T, 0)
	for _, row := range b {
		resultMap := make(map[string]any)
		for j, cell := range row {
			resultMap[query.columns[j]] = cell
		}
		t := new(T)
		util.MapToStruct(&resultMap, t)
		resultList = append(resultList, *t)
	}
	query.list = resultList
	return resultList
}

func BuildQuery[T any](formId string, mode T, handler func(query *ErpQueryModel)) *ErpQuery[T] {
	query := new(ErpQuery[T])
	model := ErpQueryModel{}
	model.Limit = 10
	model.StartRow = 0

	fieldKeys := make([]string, 0)
	modeType := reflect.TypeOf(mode)
	if modeType.Kind() == reflect.Ptr {
		modeType = modeType.Elem()
	}
	for i := 0; i < modeType.NumField(); i++ {
		fieldType := modeType.Field(i)
		key := fieldType.Tag.Get("json")
		if key == "" {
			key = fieldType.Name
		}
		fieldKeys = append(fieldKeys, key)
	}
	model.FieldKeys = strings.Join(fieldKeys, ",")
	model.OrderString = "FID desc"
	model.FormId = formId

	if handler != nil {
		handler(&model)
	}
	query.queryModel = model
	query.columns = fieldKeys
	return query
}
