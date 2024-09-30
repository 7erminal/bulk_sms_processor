package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type BulkSmsBulkrecipients struct {
	Active       int               `orm:"column(active)"`
	Id           int64             `orm:"column(bulk_recipient_id);auto"`
	CampaignIdId *BulkSmsCampaigns `orm:"column(campaign_id_id);rel(fk)"`
	CreatedAt    time.Time         `orm:"column(created_at);type(datetime)"`
	CreatedBy    int               `orm:"column(created_by)"`
	Processed    int               `orm:"column(processed)"`
	Recipient    string            `orm:"column(recipient);size(200);null"`
	UpdatedAt    time.Time         `orm:"column(updated_at);type(datetime)"`
	UpdatedBy    int               `orm:"column(updated_by)"`
}

func (t *BulkSmsBulkrecipients) TableName() string {
	return "bulk_sms_bulkrecipients"
}

func init() {
	orm.RegisterModel(new(BulkSmsBulkrecipients))
}

// AddBulkSmsBulkrecipients insert a new BulkSmsBulkrecipients into database and returns
// last inserted Id on success.
func AddBulkSmsBulkrecipients(m *BulkSmsBulkrecipients) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBulkSmsBulkrecipientsById retrieves BulkSmsBulkrecipients by Id. Returns error if
// Id doesn't exist
func GetBulkSmsBulkrecipientsById(id int64) (v *BulkSmsBulkrecipients, err error) {
	o := orm.NewOrm()
	v = &BulkSmsBulkrecipients{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetBulkSmsBulkrecipientsById retrieves BulkSmsBulkrecipients by Campaign Id. Returns error if
// Id doesn't exist
func GetBulkSmsBulkRecipientsByCampaign(id BulkSmsCampaigns) (v *[]BulkSmsBulkrecipients, err error) {
	o := orm.NewOrm()
	// v = &BulkSmsBulkrecipients{CampaignIdId: &id}
	var l []BulkSmsBulkrecipients
	if _, err = o.QueryTable(new(BulkSmsBulkrecipients)).Filter("CampaignIdId__Id", id.Id).All(&l); err == nil {
		return &l, nil
	}
	return nil, err
}

// GetAllBulkSmsBulkrecipients retrieves all BulkSmsBulkrecipients matches certain condition. Returns empty list if
// no records exist
func GetAllBulkSmsBulkrecipients(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BulkSmsBulkrecipients))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []BulkSmsBulkrecipients
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateBulkSmsBulkrecipients updates BulkSmsBulkrecipients by Id and returns error if
// the record to be updated doesn't exist
func UpdateBulkSmsBulkrecipientsById(m *BulkSmsBulkrecipients) (err error) {
	o := orm.NewOrm()
	v := BulkSmsBulkrecipients{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBulkSmsBulkrecipients deletes BulkSmsBulkrecipients by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBulkSmsBulkrecipients(id int64) (err error) {
	o := orm.NewOrm()
	v := BulkSmsBulkrecipients{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BulkSmsBulkrecipients{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
