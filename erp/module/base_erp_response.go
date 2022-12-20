package module

type BaseResultResponse struct {
	Result BaseResponseStatusReponse
}

type BaseResponseStatusReponse struct {
	ResponseStatus        BaseResponseStatus
	ConvertResponseStatus BaseResponseStatus
	Id                    int
	Number                string
	NeedReturnData        []any
}

type BaseResponseStatus struct {
	ErrorCode       int
	IsSuccess       bool
	Errors          []ErrorEntity
	SuccessEntitys  []SuccessEntity
	SuccessMessages []string
	MsgCode         int
}

type ErrorEntity struct {
	FieldName string
	Message   string
	DIndex    int
}

type SuccessEntity struct {
	Id     int
	Number string
	DIndex int
}
