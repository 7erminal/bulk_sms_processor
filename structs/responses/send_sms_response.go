package responses

import (
	"time"
)

type CampaignResp struct {
	Id              int
	RecipientEmail  string
	RecipientNumber string
	Title           string
	ScheduledTime   time.Time
	Message         string
	CreatedAt       time.Time
}

type SendBulkSMSResponse struct {
	StatusCode int
	Result     *CampaignResp
	StatusDesc string
}
