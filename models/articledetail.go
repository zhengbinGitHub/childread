package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
)

type ArticleDetail struct {
	Id int64
	ArticleId  int64  `orm:"default(0)" form:"article_id" valid:"Required"`
	Content    string    `orm:"type(text)" json:"content" form:"content"`
	CreatedAt     time.Time `orm:"auto_now_add" json:"created_at" form:"created_at"`
	UpdatedAt     time.Time `orm:"auto_now" json:"updated_at" form:"updated_at"`
}

func (n *ArticleDetail) TableName() string {
	return "article_details"
}

// 多字段索引
func (u *ArticleDetail) TableIndex() [][]string {
	return [][]string{
		[]string{"ArticleId"},
	}
}

func init() {
	orm.RegisterModel(new(ArticleDetail))
}

// AddArticleDetail insert a new ArticleDetail into database and returns
// last inserted Id on success.
func AddArticleDetail(m *ArticleDetail) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetArticleDetailById retrieves ArticleDetail by Id. Returns error if
// Id doesn't exist
func GetArticleDetail(query map[string]interface{}) (v *ArticleDetail, err error) {
	o := orm.NewOrm()
	// query k=v
	qs := o.QueryTable(new(ArticleDetail))
	for k, item := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, item)
	}
	v = &ArticleDetail{}
	if err = qs.RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllArticleDetail retrieves all ArticleDetail matches certain condition. Returns empty list if
// no records exist
func GetAllArticleDetail(query map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ArticleDetail))
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

	var l []ArticleDetail
	qs = qs.OrderBy(sortFields...).RelatedSel()
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

// UpdateArticleDetail updates ArticleDetail by Id and returns error if
// the record to be updated doesn't exist
func UpdateArticleDetailByArticleId(m *ArticleDetail, fields []string) (err error) {
	o := orm.NewOrm()
	v := ArticleDetail{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, fields...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteArticleDetail deletes ArticleDetail by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArticleDetail(id int64) (err error) {
	o := orm.NewOrm()
	v := ArticleDetail{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ArticleDetail{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
