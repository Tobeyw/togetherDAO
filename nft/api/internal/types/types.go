// Code generated by goctl. DO NOT EDIT.
package types

type ProjectReq struct {
	ProjectId    string `form:"project_id"`
	ProjectToken string `form:"Project_token"`
}

type ProjectResp struct {
	Ok bool `json:"ok"`
}