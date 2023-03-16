package controllers

import (
    AppErrors "apperrors"
    beego "github.com/beego/beego/v2/server/web"
) 
func CreateSuccessfulResponse(c *beego.Controller, code int, data interface{}) {
    c.Data["json"] = map[string]interface{}{
        "success" : true,
        "data": data,
        "message": nil,
        "request_id": "TODO",
    }
    
	c.ServeJSON()
}

func CreateErrorResponse(c *beego.Controller, err AppErrors.BaseErrorInterface) {

    c.Ctx.ResponseWriter.WriteHeader(err.ErrorStatusCode())
    c.Data["json"] = map[string]interface{}{
        "success" : false,
        "data": nil,
        "message": err.ErrorMessage(),
        "request_id": "TODO",
    }
	c.ServeJSON()
}
func SendResponse(c *beego.Controller, successCode int, data interface{}, err AppErrors.BaseErrorInterface) {
    
    if err != nil {
        CreateErrorResponse(c, err)
    } else {
        CreateSuccessfulResponse(c, successCode, data)
    }
}
