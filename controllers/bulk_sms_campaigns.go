package controllers

import (
	"encoding/json"
	"errors"
	"mes_bulk_sms_processor/models"
	"mes_bulk_sms_processor/structs/requests"
	"mes_bulk_sms_processor/structs/responses"
	"mes_bulk_sms_processor/utils"
	"reflect"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
)

// BulkSmsCampaignsController operations for BulkSmsCampaigns
type BulkSmsCampaignsController struct {
	beego.Controller
}

// URLMapping ...
func (c *BulkSmsCampaignsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("SendBulkSMS", c.SendBulkSMS)
}

// Post ...
// @Title Post
// @Description create BulkSmsCampaigns
// @Param	body		body 	models.BulkSmsCampaigns	true		"body for BulkSmsCampaigns content"
// @Success 201 {int} models.BulkSmsCampaigns
// @Failure 403 body is empty
// @router / [post]
func (c *BulkSmsCampaignsController) Post() {
	var v models.BulkSmsCampaigns
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddBulkSmsCampaigns(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get BulkSmsCampaigns by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.BulkSmsCampaigns
// @Failure 403 :id is empty
// @router /:id [get]
func (c *BulkSmsCampaignsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetBulkSmsCampaignsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// SendBulkSMS ...
// @Title Send Bulk SMS
// @Description get BulkSmsCampaigns by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.BulkSmsCampaigns
// @Failure 403 :id is empty
// @router /send-bulk-sms/:id [get]
func (c *BulkSmsCampaignsController) SendBulkSMS() {
	idStr := c.Ctx.Input.Param(":id")
	logs.Debug("Request received ", idStr)
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetBulkSmsCampaignsById(id)
	var campaignObj responses.CampaignResp = responses.CampaignResp{Id: v.Id, RecipientEmail: v.RecipientEmail, RecipientNumber: v.RecipientNumber, Title: v.Title, Message: v.Message, ScheduledTime: v.ScheduledTime}
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		logs.Info("Sending ID ", v.Id)
		// rc, err := models.GetBulkSmsBulkRecipientsByCampaign(int64(v.Id))
		rc, err := models.GetBulkSmsBulkRecipientsByCampaign(*v)
		if err == nil {
			// logs.Info("Bulk recipients are ", rc)
			recipients := make([]string, 0)
			for _, recipient := range *rc {
				recipients = append(recipients, recipient.Recipient)
			}

			recipientsStr := strings.Join(recipients, ",")
			if recipientsStr == "" {
				recipientsStr = v.RecipientNumber
			}
			logs.Info("Recipients are ", recipientsStr)

			reqObj := requests.SendSMSRequest{Destination: recipientsStr, Message: v.Message}

			logs.Info("About to send request ")
			resp := utils.SendSMS(&c.Controller, reqObj)
			statCode, exists := resp["code"]

			if exists {

				logs.Info("Status code is ", statCode)
				logs.Info("Checking status code ... ", reflect.TypeOf(statCode))
				var successCode float64 = 0
				if statCode == successCode {
					// var campaign models.BulkSmsCampaigns = models.BulkSmsCampaigns{Id: id, Active: 1}

					v.Active = 1

					logs.Info("Sending::: ", v.Id)
					logs.Info("Sending::: ", v.Title)
					logs.Info("Sending::: ", v.UpdatedById)
					logs.Info("Sending::: ", v.Active)

					if err := models.UpdateBulkSmsBulkrecipientsByCampaignId(v); err == nil {
						logs.Info("Updated recipients")
						logs.Info("About to update with ", v.Id)
						logs.Info("About to update with ", v.Title)
						logs.Info("About to update with ", v.UpdatedById)

						sysUser := "applicationuser"

						if user, err := models.GetUsersByUsername(sysUser); err == nil {

							v.UpdatedById = user

							logs.Info("About to now set ", v.UpdatedById)

							if err := models.UpdateBulkSmsCampaignsById(v); err == nil {
								logs.Info("Campaign updated successfully")

								var resp responses.SendBulkSMSResponse = responses.SendBulkSMSResponse{StatusCode: 200, Result: &campaignObj, StatusDesc: "Successful"}
								c.Data["json"] = resp
							} else {
								logs.Error("Error updating ", err.Error())
								var resp responses.SendBulkSMSResponse = responses.SendBulkSMSResponse{StatusCode: 307, Result: &campaignObj, StatusDesc: "Unable to update campaign"}

								c.Data["json"] = resp
							}
						} else {
							logs.Error("Error updating ", err.Error())
							var resp responses.SendBulkSMSResponse = responses.SendBulkSMSResponse{StatusCode: 308, Result: &campaignObj, StatusDesc: "Unable to find user"}

							c.Data["json"] = resp
						}
					} else {
						var resp responses.SendBulkSMSResponse = responses.SendBulkSMSResponse{StatusCode: 306, Result: &campaignObj, StatusDesc: "Unable to update recipient data"}

						c.Data["json"] = resp
					}
				} else {
					statDesc, exists := resp["message"]
					if exists {
						logs.Info("Status description is ", statDesc)

						var resp responses.SendBulkSMSResponse

						if str, isString := statDesc.(string); isString {
							/* act on str */
							resp = responses.SendBulkSMSResponse{StatusCode: 302, Result: &campaignObj, StatusDesc: str}
						} else {
							/* not string */
							resp = responses.SendBulkSMSResponse{StatusCode: 302, Result: &campaignObj, StatusDesc: "Failed"}
						}

						c.Data["json"] = resp
					} else {
						logs.Info("Status description is ", statDesc)
						var resp responses.SendBulkSMSResponse = responses.SendBulkSMSResponse{StatusCode: 302, Result: &campaignObj, StatusDesc: "Failed"}

						c.Data["json"] = resp
					}
				}
			} else {
				var resp responses.SendBulkSMSResponse = responses.SendBulkSMSResponse{StatusCode: 302, Result: &campaignObj, StatusDesc: "Unknown error"}

				c.Data["json"] = resp
			}

		} else {
			logs.Error("Error returned::: ", err.Error())

			var resp responses.SendBulkSMSResponse = responses.SendBulkSMSResponse{StatusCode: 301, Result: &campaignObj, StatusDesc: err.Error()}

			c.Data["json"] = resp
		}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get BulkSmsCampaigns
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.BulkSmsCampaigns
// @Failure 403
// @router / [get]
func (c *BulkSmsCampaignsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllBulkSmsCampaigns(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the BulkSmsCampaigns
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.BulkSmsCampaigns	true		"body for BulkSmsCampaigns content"
// @Success 200 {object} models.BulkSmsCampaigns
// @Failure 403 :id is not int
// @router /:id [put]
func (c *BulkSmsCampaignsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.BulkSmsCampaigns{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateBulkSmsCampaignsById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the BulkSmsCampaigns
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *BulkSmsCampaignsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteBulkSmsCampaigns(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
