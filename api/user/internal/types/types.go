// Code generated by goctl. DO NOT EDIT.
package types

type UserReq struct {
	Mobile string `json:"mobile"`
	Passwd string `json:"passwd"`
}

type UserReply struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
