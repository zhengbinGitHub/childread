package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
)

type Source struct {
	Id int64
	Name string  `orm:"size(30)" json:"name" form:"name"  valid:"Required"`
	CreatedAt     time.Time `orm:"auto_now" json:"created_at" form:"created_at"`
	UpdatedAt     time.Time `orm:"auto_now_add" json:"updated_at" form:"updated_at"`
}

func (n *Source) TableName() string {
	return "sources"
}

func init() {
	orm.RegisterModel(new(Source))
}

// AddSource insert a new Source into database and returns
// last inserted Id on success.
func AddSource(m *Source) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSourceById retrieves Source by Id. Returns error if
// Id doesn't exist
func GetSourceById(id int64) (v *Source, err error) {
	o := orm.NewOrm()
	v = &Source{Id: id}
	if err = o.QueryTable(new(Source)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSource retrieves all Source matches certain condition. Returns empty list if
// no records exist
func GetAllSource(query map[string]interface{}, fields []string, page int64, page_size int64, sort string) (sources []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Source))
	if len(query) !=0 {
		for k,v := range query{
			k = strings.Replace(k, ".", "__", -1)
			qs = qs.Filter(k, v)
		}
	}
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	if page_size > 0 {
		qs = qs.Limit(page_size, offset)
	}
	var l []Source
	if _, err = qs.OrderBy(sort).All(&l, fields...); err == nil {
		if len(fields) > 0 {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				sources = append(sources, m)
			}
		}
		count, _ = qs.Count()
		return sources, count, nil
	}
	return nil, 0, err
}

// UpdateSource updates Source by Id and returns error if
// the record to be updated doesn't exist
func UpdateSourceById(m *Source, fields []string) (err error) {
	o := orm.NewOrm()
	v := Source{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, fields...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSource deletes Source by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSource(id int64) (err error) {
	o := orm.NewOrm()
	v := Source{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Source{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
