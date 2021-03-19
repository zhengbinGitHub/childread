package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego"
)

type Tag struct {
	Id int64
	Name string `orm:"size(30)" json:"Name"`
	CreatedAt   time.Time `orm:"auto_now_add" json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `orm:"auto_now" json:"updated_at" form:"updated_at"`
}

func (m *Tag) TableName() string {
	return "tags"
}

func init() {
	orm.RegisterModel(new(Tag))
}

// AddTag insert a new Tag into database and returns
// last inserted Id on success.
func AddTag(m *Tag) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTagById retrieves Tag by Id. Returns error if
// Id doesn't exist
func GetTagById(id int64) (v *Tag, err error) {
	o := orm.NewOrm()
	v = &Tag{Id: id}
	if err = o.QueryTable(new(Tag)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}


// GetAllTag retrieves all Tag matches certain condition. Returns empty list if
// no records exist
func GetAllTag(query map[string]interface{}, fields []string, page int64, page_size int64, sortby []string, order []string) (tags []interface{}, count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Tag))
	if len(query) !=0 {
		for k,v := range query{
			k = strings.Replace(k, ".", "__", -1)
			qs = qs.Filter(k, v)
		}
	}
	count, _ = qs.Count()
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
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
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	}
	var l []Tag
	if _, err = qs.OrderBy(sortFields...).Limit(page_size, offset).All(&l, fields...); err == nil {
		if len(fields) > 0 {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				tags = append(tags, m)
			}
		}
		return tags, count, nil
	}
	return nil, count, err
}

// UpdateTag updates Tag by Id and returns error if
// the record to be updated doesn't exist
func UpdateTagById(m *Tag, fields []string) (err error) {
	o := orm.NewOrm()
	v := Tag{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, fields...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTag deletes Tag by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTag(id int64) (err error) {
	o := orm.NewOrm()
	v := Tag{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Tag{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//所有标签
func GetTagList(query map[string]interface{}, fields []string) (tags []*Tag) {
	o := orm.NewOrm()
	group := new(Tag)
	qs := o.QueryTable(group)
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
	qs.All(&tags, fields...)
	return
}

//获取文章标签
func GetArticleTags(articleId int64) (tags []*Tag) {
	driver := beego.AppConfig.String("db.type")
	qb, _ := orm.NewQueryBuilder(driver)
	// 构建查询对象
	qb.Select("tags.name", "tags.id").
		From("tags").
		InnerJoin("article_has_tags").
		On("tags.id = article_has_tags.tag_id").
		Where("article_id = ?").
		OrderBy("tags.id").
		Desc()
	// 导出SQL语句
	sql := qb.String()
	// 执行SQL语句
	o := orm.NewOrm()
	o.Raw(sql, articleId).QueryRows(&tags)
	return
}
