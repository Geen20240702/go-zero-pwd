// Code generated by goctl. DO NOT EDIT.
package types

type CreatePwdReq struct {
	Name    string `form:"name"`
	Url     string `form:"url"`
	Account string `form:"account"`
	Pwd     string `form:pwd`
	Remark  string `form:"remark"`
}

type CreatePwdResp struct {
	Code string `from:"code"`
	Msg  string `form:"msg"`
}

type DelPwdReq struct {
	ID []int64 `form:"id"`
}

type DelPwdResp struct {
	Code string `from:"code"`
	Msg  string `form:"msg"`
}

type GetPwdListReq struct {
	PageNum  int64 `form:"pageNum"`
	PageSize int64 `form:"pageNum"`
}

type GetPwdListResp struct {
	Code string `from:"code"`
	Msg  string `form:"msg"`
	Data []PWD  `form:data`
}

type PWD struct {
	ID      int64  `form:"id"`
	Name    string `form:"name"`
	Url     string `form:"url"`
	Account string `form:"account"`
	Pwd     string `form:pwd`
	Remark  string `form:"remark"`
}

type UpdatePwdReq struct {
	Name    string `form:"name"`
	Url     string `form:"url"`
	Account string `form:"account"`
	Pwd     string `form:pwd`
	Remark  string `form:"remark"`
}

type UpdatePwdResp struct {
	Code string `from:"code"`
	Msg  string `form:"msg"`
}
