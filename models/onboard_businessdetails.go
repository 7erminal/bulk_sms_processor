package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type OnboardBusinessdetails struct {
	Active                     int       `orm:"column(active)"`
	AlternatePhoneNumber       string    `orm:"column(alternatePhoneNumber);size(200);null"`
	Id                         int       `orm:"column(businessDetailId);auto"`
	BusinessRegistrationNumber string    `orm:"column(businessRegistrationNumber);size(200);null"`
	CertOfCorporation          string    `orm:"column(certOfCorporation);size(100);null"`
	CommenceBusinessCert       string    `orm:"column(commenceBusinessCert);size(100);null"`
	CompanyName                string    `orm:"column(companyName);size(200);null"`
	CompanyProfileCert         string    `orm:"column(companyProfileCert);size(100);null"`
	CreatedAt                  time.Time `orm:"column(created_at);type(datetime)"`
	CreatedById                *Users    `orm:"column(created_by_id);rel(fk)"`
	NatureOfBusiness           string    `orm:"column(natureOfBusiness);size(200);null"`
	NumberOfDirectors          int       `orm:"column(numberOfDirectors)"`
	OfficialPhoneNumber        string    `orm:"column(officialPhoneNumber);size(200);null"`
	PostalAddress              string    `orm:"column(postalAddress);size(200);null"`
	StreetAddress              string    `orm:"column(streetAddress);size(200);null"`
	UpdatedAt                  time.Time `orm:"column(updated_at);type(datetime)"`
	UpdatedById                *AuthUser `orm:"column(updated_by_id);rel(fk)"`
	UserIdFile                 string    `orm:"column(userIdFile);size(100);null"`
}

func (t *OnboardBusinessdetails) TableName() string {
	return "onboard_businessdetails"
}

func init() {
	orm.RegisterModel(new(OnboardBusinessdetails))
}

// AddOnboardBusinessdetails insert a new OnboardBusinessdetails into database and returns
// last inserted Id on success.
func AddOnboardBusinessdetails(m *OnboardBusinessdetails) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetOnboardBusinessdetailsById retrieves OnboardBusinessdetails by Id. Returns error if
// Id doesn't exist
func GetOnboardBusinessdetailsById(id int) (v *OnboardBusinessdetails, err error) {
	o := orm.NewOrm()
	v = &OnboardBusinessdetails{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllOnboardBusinessdetails retrieves all OnboardBusinessdetails matches certain condition. Returns empty list if
// no records exist
func GetAllOnboardBusinessdetails(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(OnboardBusinessdetails))
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

	var l []OnboardBusinessdetails
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

// UpdateOnboardBusinessdetails updates OnboardBusinessdetails by Id and returns error if
// the record to be updated doesn't exist
func UpdateOnboardBusinessdetailsById(m *OnboardBusinessdetails) (err error) {
	o := orm.NewOrm()
	v := OnboardBusinessdetails{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteOnboardBusinessdetails deletes OnboardBusinessdetails by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOnboardBusinessdetails(id int) (err error) {
	o := orm.NewOrm()
	v := OnboardBusinessdetails{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&OnboardBusinessdetails{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
