package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
)

type Author struct {
	Id int64
	Name string  `orm:"size(30)" json:"name" form:"name"  valid:"Required"`
	CreatedAt     time.Time `orm:"auto_now_add" json:"created_at" form:"created_at"`
	UpdatedAt     time.Time `orm:"auto_now" json:"updated_at" form:"updated_at"`
}

func (n *Author) TableName() string {
	return "authors"
}

func init() {
	orm.RegisterModel(new(Author))
}

// AddAuthor insert a new Author into database and returns
// last inserted Id on success.
func AddAuthor(m *Author) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAuthorById retrieves Author by Id. Returns error if
// Id doesn't exist
func GetAuthorById(id int64) (v *Author, err error) {
	o := orm.NewOrm()
	v = &Author{Id: id}
	if err = o.QueryTable(new(Author)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAuthor retrieves all Author matches certain condition. Returns empty list if
// no records exist
func GetAllAuthor(query map[string]interface{}, fields []string, page int64, page_size int64, sort string) (authors []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Author))
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
	var l []Author
	if page_size > 0 {
		qs = qs.Limit(page_size, offset)
	}
	if _, err = qs.OrderBy(sort).All(&l, fields...); err == nil {
		if len(fields) > 0 {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				authors = append(authors, m)
			}
		}
		count, _ = qs.Count()
		return authors, count, nil
	}
	return nil, 0, err
}

// UpdateAuthor updates Author by Id and returns error if
// the record to be updated doesn't exist
func UpdateAuthorById(m *Author, fileds []string) (err error) {
	o := orm.NewOrm()
	v := Author{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, fileds...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAuthor deletes Author by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAuthor(id int64) (err error) {
	o := orm.NewOrm()
	v := Author{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Author{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
