package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
)

type Topic struct {
	Id int64
	Cover string `orm:"size(100)" json:"cover" form:"cover"  valid:"Required"`
	Url string `orm:"size(100)" json:"url" form:"url"  valid:"Required"`
	Status 		int8     `orm:"default(1)" json:"status" form:"status" valid:"Range(0,1)"`
	CreatedAt   time.Time `orm:"auto_now_add" json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `orm:"auto_now" json:"updated_at" form:"updated_at"`
}

func (n *Topic) TableName() string {
	return "topics"
}

func init() {
	orm.RegisterModel(new(Topic))
}

// AddTopic insert a new Topic into database and returns
// last inserted Id on success.
func AddTopic(m *Topic) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTopicById retrieves Topic by Id. Returns error if
// Id doesn't exist
func GetTopicById(id int64) (v *Topic, err error) {
	o := orm.NewOrm()
	v = &Topic{Id: id}
	if err = o.QueryTable(new(Topic)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTopic retrieves all Topic matches certain condition. Returns empty list if
// no records exist
func GetAllTopic(query map[string]interface{}, fields []string, page int64, page_size int64, sort string) (ml []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Topic))
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
	var l []Topic
	if _, err = qs.Limit(page_size, offset).OrderBy(sort).All(&l, fields...); err == nil {
		if len(fields) > 0 {
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
		count, _ = qs.Count()
		return ml, count, nil
	}
	return nil, 0, err
}

// UpdateTopic updates Topic by Id and returns error if
// the record to be updated doesn't exist
func UpdateTopicById(m *Topic, fields []string) (err error) {
	o := orm.NewOrm()
	v := Topic{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, fields...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTopic deletes Topic by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTopic(id int64) (err error) {
	o := orm.NewOrm()
	v := Topic{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Topic{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//限制专题
func GetLimtTopic(query map[string]interface{}, fields []string, page_size int64, sort string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Topic))
	if len(query) !=0 {
		for k,v := range query{
			k = strings.Replace(k, ".", "__", -1)
			qs = qs.Filter(k, v)
		}
	}
	var l []Topic
	if _, err = qs.Limit(page_size, 0).OrderBy(sort).RelatedSel().All(&l, fields...); err == nil {
		if len(fields) > 0 {
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
