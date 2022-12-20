package erp

import (
	"encoding/json"
	"errors"
	"log"
	"sync"
)

type erpSessionConfig struct {
	UserName string
	Password string
	Lcid     string
	AcctId   string
}
type erpSession struct {
	config    erpSessionConfig
	sessionId string
	L         *sync.RWMutex
}

const KINGDEE_LOGIN_HEADER = "kdservice-sessionid"

func (session *erpSession) login() (*ErpLoginResponse, error) {
	if session.config.UserName == "" {
		log.Panicln("请先配置账号")
	}
	jsonStr := post("", LOGIN_URL.String(), "", session.config)
	postResponse := new(ErpLoginResponse)
	json.Unmarshal([]byte(jsonStr), postResponse)
	if postResponse.LoginResultType != 1 {
		panic(errors.New(postResponse.Message))
	}
	log.Println("登陆成功:", postResponse.Context.DataCenterName)
	return postResponse, nil
}

type ErpLoginResponse struct {
	SessionId       string `json:"KDSVCSessionId"`
	Context         ErpLoginEntity
	Message         string
	MessageCode     string
	LoginResultType int
}

type ErpLoginEntity struct {
	DataCenterName string
	UserId         int
	UserName       string
	CustomName     string
}

func (session *erpSession) GetSessionId() string {
	session.L.Lock()
	defer session.L.Unlock()
	if session.sessionId != "" {
		return session.sessionId
	}
	response, err := session.login()
	if err != nil {
		log.Println("登陆失败")
		if err == &NO_LOGIN_EXCEPTION {
			log.Println("登陆失效,重制cookied")
			session.sessionId = ""
		}
		panic(err)
	}
	session.sessionId = response.SessionId
	return session.sessionId
}

var defaultSession erpSession

func InitSession(username, password, acctid string) {
	conf := erpSessionConfig{UserName: username, Password: password, AcctId: acctid, Lcid: "2052"}
	defaultSession = erpSession{
		config: conf,
		L:      new(sync.RWMutex),
	}
}
func GetSessionId() string {
	return defaultSession.GetSessionId()
}
