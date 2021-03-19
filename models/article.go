package models

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id 			int64
	Title 		string  `orm:"size(200)" json:"title" form:"title"  valid:"Required"`
	Level 		int 	`orm:"default(0)" json:"level"`
	Memo  		string  `orm:"size(200)" json:"memo" form:"memo"  valid:"Required"`
	Path  		string  `orm:"size(20)" json:"path" form:"path"`
	Cover 		string  `orm:"size(200)" json:"cover" form:"cover"  valid:"Required"`
	Status 		int8     `orm:"default(1)" json:"status" form:"status" valid:"Range(0,1)"`
	IsHot 		int8     `orm:"default(0)" json:"is_hot" form:"is_hot" valid:"Range(0,1)"`
	IsCommand 	int8     `orm:"default(0)" json:"is_command" form:"Is_command" valid:"Range(0,1)"`
	IsWonderful	int8     `orm:"default(0)" form:"is_wonderful" valid:"Range(0,1)"`
	IsBanner	int8     `orm:"default(0)" form:"is_banner" valid:"Range(0,1)"`
	IsToutiao	int8     `orm:"default(0)" form:"is_toutiao" valid:"Range(0,1)"`
	IsToday		int8     `orm:"default(0)" form:"is_today" valid:"Range(0,1)"`
	Sort		int      `orm:"default(1)" json:"view" form:"sort" valid:"Numeric"`
	View     	int       `orm:"default(0)" json:"view"`
	Agree       int       `orm:"default(0)" json:"agree"`
	Comment     int       `orm:"default(0)" json:"comment"`
	AuthorId    int64     `orm:"default(0)" json:"author_id" form:"author_id" valid:"Required"`
	SourceId    int64     `orm:"default(0)" json:"source_id" form:"source_id" valid:"Required"`
	CreatedAt   time.Time `orm:"auto_now_add" json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `orm:"auto_now" json:"updated_at" form:"updated_at"`
	Group  		*Group  `orm:"rel(fk)"`
}

func (m *Article) TableName() string {
	return "articles"
}

func init() {
	orm.RegisterModel(new(Article))
}

// AddArticle insert a new Article into database and returns
// last inserted Id on success.
func AddArticle(m *Article) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetArticleById retrieves Article by Id. Returns error if
// Id doesn't exist
func GetArticleById(id int64) (v *Article, err error) {
	o := orm.NewOrm()
	v = &Article{Id: id}
	if err = o.QueryTable(new(Article)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllArticle retrieves all Article matches certain condition. Returns empty list if
// no records exist
func GetAllArticle(query map[string]interface{}, fields []string, page int64, page_size int64, sort string) (
	articles []Article, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))
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
	if _, err = qs.Limit(page_size, offset).OrderBy(sort).RelatedSel().All(&articles, fields...); err == nil {
		count, _ = qs.Count()
		return articles, count, nil
	}
	return nil, 0, err
}

// UpdateArticle updates Article by Id and returns error if
// the record to be updated doesn't exist
func UpdateArticleById(m *Article, fields []string) (err error) {
	o := orm.NewOrm()
	v := Article{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, fields...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteArticle deletes Article by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArticle(id int64) (err error) {
	o := orm.NewOrm()
	v := Article{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Article{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//聚合数据
func CountArticle(query map[string]interface{}) (total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))
	if len(query) !=0 {
		for k,v := range query{
			k = strings.Replace(k, ".", "__", -1)
			qs = qs.Filter(k, v)
		}
	}
	if total, err = qs.Count(); err == nil{
		fmt.Println("Number of records deleted in database:", total)
	}
	return
}

//限制条数文章
func GetLimtArticle(query map[string]interface{}, fields []string, page_size int64, sort string) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Article))
	if len(query) !=0 {
		for k,v := range query{
			k = strings.Replace(k, ".", "__", -1)
			if strings.Contains(k, "notin") {
				qs = qs.Exclude(strings.Replace(k, "_notin", "", -1), v)
			} else {
				qs = qs.Filter(k, v)
			}
		}
	}
	var l []Article
	if _, err = qs.Limit(page_size, 0).OrderBy(sort).RelatedSel().All(&l); err == nil {
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

// GetArticle retrieves Article by field. Returns error if
// Id doesn't exist
func GetArticleByField(query map[string]interface{}) (v *Article, err error) {
	o := orm.NewOrm()
	// query k=v
	qs := o.QueryTable(new(Article))
	for k, item := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, item)
	}
	v = &Article{}
	if err = qs.RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}
