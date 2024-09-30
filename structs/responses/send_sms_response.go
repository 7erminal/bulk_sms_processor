package responses

import (
	"mes_bulk_sms_processor/models"
)

type AccountDTO struct {
	StatusCode int
	Account    *models.Accounts
	StatusDesc string
}
