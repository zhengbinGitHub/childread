package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
)

type GroupHasTag struct {
	Id int64
	GroupId int64 `json:"group_id"`
	Tag *Tag `orm:"rel(fk)" json:"tag_id"`
	CreatedAt   time.Time `orm:"auto_now_add" json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `orm:"auto_now" json:"updated_at" form:"updated_at"`
}

func (m *GroupHasTag) TableName() string {
	return "group_has_tags"
}

func init() {
	orm.RegisterModel(new(GroupHasTag))
}

// AddGroupHasTag insert a new GroupHasTag into database and returns
// last inserted Id on success.
func AddGroupHasTag(m *GroupHasTag) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGroupHasTagById retrieves GroupHasTag by Id. Returns error if
// Id doesn't exist
func GetGroupHasTagById(id int64) (v *GroupHasTag, err error) {
	o := orm.NewOrm()
	v = &GroupHasTag{Id: id}
	if err = o.QueryTable(new(GroupHasTag)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGroupHasTag retrieves all GroupHasTag matches certain condition. Returns empty list if
// no records exist
func GetAllGroupHasTag(query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GroupHasTag))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []GroupHasTag
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if limit > 0 && offset > 0 {
		qs = qs.Limit(limit, offset)
	}
	if _, err = qs.All(&l, fields...); err == nil {
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

// UpdateGroupHasTag updates GroupHasTag by Id and returns error if
// the record to be updated doesn't exist
func UpdateGroupHasTagById(m *GroupHasTag) (err error) {
	o := orm.NewOrm()
	v := GroupHasTag{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGroupHasTag deletes GroupHasTag by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGroupHasTag(id int64) (err error) {
	o := orm.NewOrm()
	v := GroupHasTag{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GroupHasTag{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
