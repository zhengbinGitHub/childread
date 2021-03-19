package admin

import (
	m "childread/models"
	. "childread/helperfunc"
	"errors"
	"fmt"
	"strings"
)

// LoginController operations for Login
type LoginController struct {
	CommonController
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Index", c.Index)
	c.Mapping("Password", c.Password)
}

func (c *LoginController) Password() {
	c.TplName = "login/login/password.tpl"
}

// Post ...
// @Title Create
// @Description create Login
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 201 {object} models.Login
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Post() {
	isajax := c.GetString("isajax")
	if isajax == "1" {
		username := c.GetString("mobile")
		password := c.GetString("password")
		user, err := c.checkLogin(username, password)
		if err == nil {
			c.SetSession("userinfo", user)
			accesslist, _ := c.getAccessList(user.Id)
			c.SetSession("accesslist", accesslist)
			c.Rsp(true, "登录成功", nil)
			return
		} else {
			c.Rsp(false, err.Error(), nil)
			return
		}

	}
}

//check login
func (c *LoginController) checkLogin(username string, password string) (user m.User, err error) {
	user = m.GetUserByUsername(username)
	if user.Id == 0 {
		return user, errors.New("用户不存在")
	}
	if user.Password != Pwdhash(password) {
		return user, errors.New("密码错误")
	}
	return user, nil
}

type AccessNode struct {
	Id        int64
	Name      string
	Childrens []*AccessNode
}

func (c *LoginController) getAccessList(uid int64) (map[string]bool, error){
	list, err := m.AccessList(uid)
	if err != nil {
		return nil, err
	}
	alist := make([]*AccessNode, 0)
	for _, l := range list {
		if l["Pid"].(int64) == 0 && l["Level"].(int64) == 1 {
			anode := new(AccessNode)
			anode.Id = l["Id"].(int64)
			anode.Name = l["Name"].(string)
			alist = append(alist, anode)
		}
	}
	for _, l := range list {
		if l["Level"].(int64) == 2 {
			for _, an := range alist {
				if an.Id == l["Pid"].(int64) {
					anode := new(AccessNode)
					anode.Id = l["Id"].(int64)
					anode.Name = l["Name"].(string)
					an.Childrens = append(an.Childrens, anode)
				}
			}
		}
	}
	for _, l := range list {
		if l["Level"].(int64) == 3 {
			for _, an := range alist {
				for _, an1 := range an.Childrens {
					if an1.Id == l["Pid"].(int64) {
						anode := new(AccessNode)
						anode.Id = l["Id"].(int64)
						anode.Name = l["Name"].(string)
						an1.Childrens = append(an1.Childrens, anode)
					}
				}

			}
		}
	}
	accesslist := make(map[string]bool)
	for _, v := range alist {
		for _, v1 := range v.Childrens {
			for _, v2 := range v1.Childrens {
				vname := strings.Split(v.Name, "/")
				v1name := strings.Split(v1.Name, "/")
				v2name := strings.Split(v2.Name, "/")
				str := fmt.Sprintf("%s/%s/%s", strings.ToLower(vname[0]), strings.ToLower(v1name[0]), strings.ToLower(v2name[0]))
				accesslist[str] = true
			}
		}
	}
	return accesslist, nil
}

// GetOne ...
// @Title GetOne
// @Description get Login by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Login
// @Failure 403 :id is empty
// @router /:id [get]
func (c *LoginController) Index() {
	c.TplName = "admin/login/index.tpl"
}

//退出
func (c *LoginController) Logout() {
	c.DelSession("userinfo")
	c.Ctx.Redirect(302, "/admin/login")
}
