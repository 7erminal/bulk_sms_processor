package utils

import (
	"encoding/json"
	"io"
	"mes_bulk_sms_processor/api"

	"mes_bulk_sms_processor/structs/requests"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func SendSMS(c *beego.Controller, params requests.SendSMSRequest) (data map[string]interface{}) {
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

	// logs.Info("Raw response received is ", read.json())
	// var data map[string]interface{}

	// var data responses.SendBulkSMSResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Raw response received is ")
	logs.Info(data)

	return data
}
