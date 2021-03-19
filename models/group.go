package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"errors"
	"strings"
)

type Group struct {
	Id     int64
	Name   string  `orm:"size(100)" form:"Name"  valid:"Required"`
	Title  string  `orm:"size(100)" form:"Title"  valid:"Required"`
	Status int     `orm:"default(1)" form:"Status" valid:"Range(0,1)"`
	Sort   int     `orm:"default(1)" form:"Sort" valid:"Numeric"`
	Level   int     `orm:"default(1)" form:"Level" valid:"Numeric"`
	Parent_id   int64     `orm:"default(0)"`
	Nodes  []*Node `orm:"reverse(many)"`
	Class_name string `orm:"size(10)" from:"Class_name"`
	Path string `orm:"size(20)"`
	Article []*Article `orm:"reverse(many)"`
}

func (n *Group) TableName() string {
	return "groups"
}

func init() {
	orm.RegisterModel(new(Group))
}

func checkGroup(g *Group) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&g)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

//get group list
func GetGrouplist(query map[string]interface{}, page int64, page_size int64, sort string) (groups []orm.Params, count int64) {
	o := orm.NewOrm()
	group := new(Group)
	qs := o.QueryTable(group)
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
	qs.Limit(page_size, offset).OrderBy(sort).Values(&groups)
	count, _ = qs.Count()
	return groups, count
}

func AddGroup(g *Group) (int64, error) {
	if err := checkGroup(g); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	group := new(Group)
	group.Name = g.Name
	group.Title = g.Title
	group.Sort = g.Sort
	group.Level = g.Level
	group.Parent_id = g.Parent_id
	group.Status = g.Status
	id, err := o.Insert(group)
	return id, err
}

func UpdateGroup(g *Group) (int64, error) {
	if err := checkGroup(g); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	group := make(orm.Params)
	if len(g.Name) > 0 {
		group["Name"] = g.Name
	}
	if len(g.Title) > 0 {
		group["Title"] = g.Title
	}
	if g.Status >= 0 {
		group["Status"] = g.Status
	}
	if g.Sort != 0 {
		group["Sort"] = g.Sort
	}
	if len(group) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Group
	num, err := o.QueryTable(table).Filter("Id", g.Id).Update(group)
	return num, err
}

func DelGroupById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Group{Id: Id})
	return status, err
}

func GroupList(query map[string]interface{}) (groups []*Group) {
	o := orm.NewOrm()
	group := new(Group)
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
	qs.All(&groups, "Id", "Title", "Parent_id", "Level", "Class_name", "Path")
	return groups
}

type TreeList struct {
	Id int64			`json:"id"`
	Name string		`json:"name"`
	Class_name string `json:"class_name"`
	Pid int64			`json:"pid"`
	Path string `json:"path"`
	Sort int 		`json:"sort"`
	Children []*TreeList	`json:"children"`
}

/**
菜单列表
*/
func MenuList() []*TreeList {
	return getMenu(0)
}

/**
递归获取树形菜单
*/
func getMenu(pid int64) []*TreeList {
	o := orm.NewOrm()
	var menu []Group
	_,_ = o.QueryTable(new(Group)).Filter("parent_id", pid).OrderBy("sort").All(&menu)
	treeList := []*TreeList{}
	for _, v := range menu{
		child := getMenu(v.Id)
		node := &TreeList{
			Id:v.Id,
			Name:v.Title,
			Class_name:v.Class_name,
			Path:v.Path,
			Sort:v.Sort,
			Pid:v.Parent_id,
		}
		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList
}

//单条数据
func GetGroupByFields(query map[string]interface{}, fields []string) (group *Group, err error)  {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Group))
	if len(query) !=0 {
		for k,v := range query{
			k = strings.Replace(k, ".", "__", -1)
			if strings.Contains(k, "notin") {
				qs = qs.Exclude(strings.Replace(k, "_notin", "", 1), v)
			} else {
				qs = qs.Filter(k, v)
			}
		}
	}
	//filedStr := strings.Join(fields,",")
	group = &Group{}
	if err = qs.One(group, fields...); err == nil {
		return group, nil
	}
	return nil, err
}

