package erp

import (
	"reflect"
)

type ErpUrlEnum string

var (
	BASE_URL  string
	LOGIN_URL ErpUrlEnum = "Kingdee.BOS.WebApi.ServicesStub.AuthService.ValidateUser.common.kdsvc"
	LIST_URL  ErpUrlEnum = "Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.ExecuteBillQuery.common.kdsvc"
)

func (enum ErpUrlEnum) String() string {
	return BASE_URL + reflect.ValueOf(enum).String()
}

func InitBaseUrl(url string) {
	BASE_URL = url
}
