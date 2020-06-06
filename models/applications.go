package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Applications struct {
	Id          int       `orm:"column(id);pk;auto" json:"id"`
	Aim         *Aims     `orm:"column(aim_id);rel(fk)" json:"aim"`
	UniqId      string    `orm:"column(uniq_id);null" json:"uniq_id"`
	Filename    string    `orm:"column(filename);null" json:"filename"`
	Url         string    `orm:"column(url);null" json:"url"`
	Description string    `orm:"column(description);null" json:"description"`
	Active      bool      `orm:"column(active);null" json:"active"`
	CreatedAt   time.Time `orm:"column(created_at);auto_now_add;type(timestamp with time zone);null" json:"created_at"`
	UpdatedAt   time.Time `orm:"column(updated_at);auto_now;type(timestamp with time zone);null" json:"updated_at"`
	DeletedAt   time.Time `orm:"column(deleted_at);type(timestamp without time zone);null"`
}

func (t *Applications) TableName() string {
	return "applications"
}

func init() {
	orm.RegisterModel(new(Applications))
}

// AddApplications insert a new Applications into database and returns
// last inserted Id on success.
func AddApplications(m *Applications) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetApplicationsById retrieves Applications by Id. Returns error if
// Id doesn't exist
func GetApplicationsById(id int) (v *Applications, err error) {
	o := orm.NewOrm()
	v = &Applications{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllApplications retrieves all Applications matches certain condition. Returns empty list if
// no records exist
func GetAllApplications(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Applications))
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

	var l []Applications
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

// UpdateApplications updates Applications by Id and returns error if
// the record to be updated doesn't exist
func UpdateApplicationsById(m *Applications) (err error) {
	o := orm.NewOrm()
	v := Applications{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteApplications deletes Applications by Id and returns error if
// the record to be deleted doesn't exist
func DeleteApplications(id int) (err error) {
	o := orm.NewOrm()
	v := Applications{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Applications{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
