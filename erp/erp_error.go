package erp

import (
	"encoding/json"
	"strings"

	"github.com/keeleys/go-kingdee/erp/module"
)

const checkStr = "ErrorCode"

type ErpError struct {
	ErrorMsg  string
	ErrorCode int
}

var (
	NO_LOGIN_EXCEPTION       = ErpError{ErrorCode: 10403, ErrorMsg: "会话信息已丢失，请重新登录"}
	AppROVE_STATUS_EXCEPTION = ErpError{ErrorCode: 10404, ErrorMsg: "只能删除创建，暂存，重新审核状态的数据"}
	errorList                = [...]ErpError{
		NO_LOGIN_EXCEPTION,
		AppROVE_STATUS_EXCEPTION,
	}
)

func (err *ErpError) Error() string {
	return err.ErrorMsg
}

func NewErpError(msg string) *ErpError {
	return &ErpError{ErrorMsg: msg, ErrorCode: -1}
}

func checkError(response module.BaseResultResponse) error {
	if !response.Result.ResponseStatus.IsSuccess {
		errList := make([]string, 0)
		for _, v := range response.Result.ResponseStatus.Errors {
			errList = append(errList, v.Message)
		}
		errMsg := strings.Join(errList, ",")

		for _, v := range errorList {
			if strings.Contains(errMsg, v.ErrorMsg) {
				return &v
			}
		}
		return NewErpError(errMsg)

	}
	return nil
}

func checkQueryError(jsonStr string) error {
	if jsonStr != "" && strings.Contains(jsonStr, checkStr) {
		b := make([][]module.BaseResultResponse, 0)
		err := json.Unmarshal([]byte(jsonStr), &b)
		if err != nil {
			return err
		}
		return checkError(b[0][0])
	}
	return nil
}
