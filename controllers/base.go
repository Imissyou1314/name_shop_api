package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

type BaseRes struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func (o *BaseController) SetResult(data interface{}, err error) {
	if err != nil {
		o.Data["json"] = BaseRes{
			Message: err.Error(),
			Code:    500,
			Data:    data,
		}
	} else {
		o.Data["json"] = BaseRes{
			Message: "成功",
			Code:    200,
			Data:    data,
		}
	}
	o.ServeJSON()
}
