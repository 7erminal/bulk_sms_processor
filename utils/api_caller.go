package utils

import (
	"encoding/json"
	"io"
	"mes_bulk_sms_processor/api"

	"mes_bulk_sms_processor/structs/requests"
	"mes_bulk_sms_processor/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func SendSMS(c *beego.Controller, params requests.SendSMSRequest) {
	host, _ := beego.AppConfig.String("deywuro")

	logs.Info("Sending to destination ", params.Destination)

	request := api.NewRequest(
		host,
		"/api/sms",
		api.POST)
	request.Params["username"], _ = beego.AppConfig.String("deywuroUsername")
	request.Params["password"], _ = beego.AppConfig.String("deywuroPassword")
	request.Params["source"], _ = beego.AppConfig.String("deywuroSource")
	request.Params["destination"] = params.Destination
	request.Params["message"] = params.Message
	request.Params["ol"] = "false"
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.AccountDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data["access_token"])
	// logs.Info("Expires in ", data["expires_in"])
	// logs.Info("Scope is ", data["scope"])
	// logs.Info("Token Type is ", data["token_type"])
	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data.Access_token)
	// logs.Info("Expires in ", data.Expires_in)
	// logs.Info("Scope is ", data.Scope)
	// logs.Info("Token Type is ", data.Token_type)
}
