package erp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/keeleys/go-kingdee/erp/module"
)

func erpPost(sessionId, url string, formId string, model any) string {
	wrap := module.BaseErpPostWrap{Data: model}
	return post(sessionId, url, formId, wrap)
}
func post(sessionId, url string, formId string, model any) string {
	body, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	log.Printf("请求url:%s,sessionId:%s,请求内容:%s\n", url, sessionId, string(body))
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Content-Type", "application/json")
	if sessionId != "" {
		r.Header.Add(KINGDEE_LOGIN_HEADER, sessionId)
	}

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	jsonByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	resultStr := string(jsonByte)
	log.Printf("返回结果%s\n", resultStr)
	return resultStr
}

func PostSave(url string, formId string, model any) *module.BaseResultResponse {
	result := &module.BaseResultResponse{}
	sessionId := GetSessionId()
	jonsStr := erpPost(sessionId, url, formId, model)
	json.Unmarshal([]byte(jonsStr), result)
	checkError(*result)
	return result
}

func PostSelect[T any](formId string, model T) []T {
	return PostSelectHandler(formId, model, func(query *ErpQueryModel) {})
}

func PostSelectHandler[T any](formId string, model T, handler func(query *ErpQueryModel)) []T {
	query := BuildQuery(formId, model, handler)
	return query.selectList(GetSessionId())
}
