package code

import (
	"bytes"
)

type RetResponse struct {
	RetCode int         `json:"code"`
	Err     string      `json:"err"`
	RetData interface{} `json:"data"`
}

type ErrResponse struct {
	Err string `json:"err"`
}

type RetPageData struct {
	PageInfo PageInfo    `json:"info"`
	Data     interface{} `json:"data"`
}

type PageInfo struct {
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
	PageSize  int `json:"page_size"`
	PageNum   int `json:"page_num"`
}

var (
	RetNormal = RetResponse{RetCode: 300000}
	//RetMissParams        = RetResponse{RetCode: 300001, Err: "missing params"}
	RetInvalidParams = RetResponse{RetCode: 300002, Err: "invalid params"}
	RetInvalidUser   = RetResponse{RetCode: 300101, Err: "invalid user"}
	RetOperationErr  = RetResponse{RetCode: 300100, Err: ""}
)

//func GeneralRetInvalidParams(a ...string) RetResponse {
//	ret := RetMissParams
//	buf := bytes.Buffer{}
//	buf.WriteString(ret.Err)
//	buf.WriteString(": ")
//	for _, v := range a {
//		buf.WriteString(v)
//		buf.WriteString(" ")
//	}
//	ret.Err = buf.String()
//	return ret
//}

func GeneralRetInvalidParams(a ...string) RetResponse {
	ret := RetInvalidParams
	buf := bytes.Buffer{}
	buf.WriteString(ret.Err)
	buf.WriteString(": ")
	for _, v := range a {
		buf.WriteString(v)
		buf.WriteString(" ")
	}
	ret.Err = buf.String()
	return ret
}

func GeneralErrRet(err error) ErrResponse {
	return ErrResponse{Err: err.Error()}
}
