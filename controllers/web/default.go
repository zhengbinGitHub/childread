package web

import (
	m "childread/models"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//首页
func (c *MainController) Index() {
	c.getBanner()
	c.getToutiao()
	c.getTaday()
	c.getView()
	c.getHot()
	c.getYuer()
	c.getTopic()
	c.getGuidance()
	c.getWeekly()
	c.getWonderful()
	c.getLife()
	c.getPregnancy()
	c.getPostpartum()
	c.getChildhood()
	c.Layout = "layouts/layout.tpl"
	c.TplName = "default/index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "default/style.tpl"
	c.LayoutSections["Scripts"] = "default/scripts.tpl"
}

//育儿
func (c *MainController) getYuer() {
	info,_ := m.GetGroupByFields(map[string]interface{}{"Path": "yuer"}, []string{"Id"})
	groups := m.GroupList(map[string]interface{}{"Parent_id": info.Id})
	c.Data["Yuers"] = groups
}

//图片推广
func (c *MainController) getBanner() {
	banners,_ := m.GetLimtArticle(map[string]interface{}{"IsBanner": 1}, []string{"Id", "Cover", "Title", "Group", "Path"}, 5, "Id")
	c.Data["Banners"] = banners
}

//头条
func (c *MainController) getToutiao()  {
	toutiaos,_ := m.GetLimtArticle(map[string]interface{}{"IsToutiao": 1}, []string{"Id", "Title", "Group", "Path"}, 7, "Id")
	c.Data["Toutiaos"] = toutiaos
}

//今日推荐
func (c *MainController) getTaday() {
	taday,_ := m.GetLimtArticle(map[string]interface{}{"IsToday": 1}, []string{"Id", "Title", "Group", "Path"}, 7 , "Id")
	c.Data["Tadays"] = taday
}

//人气排行
func (c *MainController) getView() {
	filters := make(map[string]interface{}, 0)
	filters["IsHot"] = 0
	filters["IsCommand"] = 0
	filters["IsWonderful"] = 0
	filters["IsBanner"] = 0
	filters["IsToutiao"] = 0
	filters["IsToday"] = 0
	news,_ := m.GetLimtArticle(filters, []string{"Id", "Title", "Group", "View", "Path"}, 9 , "View")
	c.Data["Viewnews"] = news
}

//热点排行
func (c *MainController) getHot() {
	news,_ := m.GetLimtArticle(map[string]interface{}{"IsHot": 1}, []string{"Id", "Title", "Group", "View", "Cover", "Memo", "CreatedAt", "Path"}, 5 , "View")
	c.Data["Hotnews"] = news
}

//小编精选
func (c *MainController) getWonderful() {
	wonderfuls,_ := m.GetLimtArticle(map[string]interface{}{"IsWonderful": 1}, []string{"Id", "Title", "Group", "Path"}, 6 , "-View")
	c.Data["Wonderfuls"] = wonderfuls
}

//独家专题
func (c *MainController) getTopic() {
	topics,_ := m.GetLimtTopic(map[string]interface{}{"Status": 1}, []string{"Id", "Cover", "Url"}, 3 , "-Id")
	c.Data["Topics"] = topics
}

//育期指导
func (c *MainController) getGuidance() {
	filters := make(map[string]interface{}, 0)
	group,_ := m.GetGroupByFields(map[string]interface{}{"Path": "huaiyun"}, []string{"id"})
	groups := m.GroupList(map[string]interface{}{"Parent_id": group.Id})
	groupIds := []int64{}
	for _, item := range groups {
		groupIds = append(groupIds, item.Id)
	}
	filters["Group__in"] = groupIds
	filters["IsCommand"] = 1
	guidances,_ := m.GetLimtArticle(filters, []string{"Id", "Title", "Memo", "Group", "Cover", "View", "CreatedAt", "Path"}, 5, "-Id")
	c.Data["Guidances"] = guidances
}

//育儿周刊
func (c *MainController) getWeekly() {
	filters := make(map[string]interface{}, 0)
	group,_ := m.GetGroupByFields(map[string]interface{}{"Path": "yuer"}, []string{"id"})
	groups := m.GroupList(map[string]interface{}{"Parent_id": group.Id})
	groupIds := []int64{}
	for _, item := range groups {
		groupIds = append(groupIds, item.Id)
	}
	filters["Group__in"] = groupIds
	filters["IsCommand"] = 1
	weeklies,_ := m.GetLimtArticle(filters, []string{"Id", "Title", "Memo", "Group", "Cover", "View", "CreatedAt", "Path"}, 5, "-Id")
	c.Data["Weeklies"] = weeklies
}

//生活
func (c *MainController) getLife() {
	filters := make(map[string]interface{}, 0)
	group,_ := m.GetGroupByFields(map[string]interface{}{"Path": "life"}, []string{"id"})
	groups := m.GroupList(map[string]interface{}{"Parent_id": group.Id})
	groupIds := []int64{}
	for _, item := range groups {
		groupIds = append(groupIds, item.Id)
	}
	filters["Group__in"] = groupIds
	filters["IsCommand"] = 1
	lifes,_ := m.GetLimtArticle(filters, []string{"Id", "Title", "Memo", "Group", "Cover", "View", "CreatedAt", "Path"}, 5, "-Id")
	c.Data["Lifes"] = lifes
}

//备孕
func (c *MainController) getPregnancy() {
	group,_ := m.GetGroupByFields(map[string]interface{}{"Path": "beiyun"}, []string{"id"})
	beiyuns := m.GroupList(map[string]interface{}{"Parent_id": group.Id})
	c.Data["Beiyuns"] = beiyuns
	groupIds := []int64{}
	for _, item := range beiyuns {
		groupIds = append(groupIds, item.Id)
		c.getBeiyunChild(item.Id, item.Path)
	}

	filters := make(map[string]interface{}, 0)
	filters["Group__in"] = groupIds
	filters["IsHot"] = 1
	hotCovers,_ := m.GetLimtArticle(filters, []string{"Id", "Title", "Group", "Cover", "CreatedAt", "Path"}, 2, "-Id")
	c.Data["BeiyunHotCovers"] = hotCovers
	ids := []int64{}
	//reflect.TypeOf(id) 值类型
	for _, item := range hotCovers {
		id := item.(map[string]interface{})["Id"].(int64)
		ids = append(ids, id)
	}
	filters["Id__in_notin"] = ids
	hotNews, _ := m.GetLimtArticle(filters, []string{"Id", "Title", "Group", "Cover", "CreatedAt", "Path"}, 6, "-Id")
	c.Data["BeiyunHotNews"] = hotNews
}

//备孕子类
func (c *MainController) getBeiyunChild(groupId int64, index string) {
	filters := make(map[string]interface{}, 0)
	filters["Group"] = groupId
	filters["IsCommand"] = 1
	commands,_ := m.GetLimtArticle(filters, []string{"Id", "Title", "Group", "Cover", "CreatedAt", "Path"}, 12, "-Id")
	if len(commands) >= 3 {
		c.Data["Beiyun"+index+"Covers"] = commands[:1]
		c.Data["Beiyun"+index+"MidNews"] = commands[1:3]
		c.Data["Beiyun"+index+"News"] = commands[3:]
	}
}

//产后
func (c *MainController) getPostpartum() {
	group,_ := m.GetGroupByFields(map[string]interface{}{"Path": "chanhou"}, []string{"id"})
	chanhous := m.GroupList(map[string]interface{}{"Parent_id": group.Id})
	groupIds := []int64{}
	for _, item := range chanhous {
		groupIds = append(groupIds, item.Id)
	}
	filters := make(map[string]interface{}, 0)
	filters["Group__in"] = groupIds
	filters["IsToutiao"] = 1
	touTiaoNews, _ := m.GetLimtArticle(filters, []string{"Id", "Title", "Group", "Cover", "Memo", "CreatedAt", "Path"}, 9, "-Id")
	c.Data["ChanhouToutiaoLeCovers"] = touTiaoNews[:2]
	c.Data["ChanhouToutiaoMidCovers"] = touTiaoNews[2:3]
	c.Data["ChanhouToutiaoMidNews"] = touTiaoNews[3:]

	delete(filters, "IsToutiao")
	filters["IsCommand"] = 1
	commands,_ := m.GetLimtArticle(filters, []string{"Id", "Title", "Group", "Cover", "CreatedAt", "Path"}, 9, "-Id")
	c.Data["ChanhouCommandCovers"] = commands[:1]
	c.Data["ChanhouCommandNews"] = commands[1:]
}

//早教
func (c *MainController) getChildhood()  {
	filters := make(map[string]interface{}, 0)
	parent,_ := m.GetGroupByFields(map[string]interface{}{"Path": "zaojiao"}, []string{"id"})
	filters["Parent_id"] = parent.Id
	filters["Path__in_notin"] = []string{"game", "music"}
	chanhous := m.GroupList(filters)
	groupIds := []int64{}
	for _, item := range chanhous {
		groupIds = append(groupIds, item.Id)
	}
	delete(filters, "Parent_id")
	delete(filters, "Path__in_notin")
	filters["Group__in"] = groupIds
	covers, _ := m.GetLimtArticle(filters, []string{"Id", "Title", "Group", "Cover", "Memo", "CreatedAt", "Path"}, 2, "Id")
	c.Data["StoryCovers"] = covers

	delete(filters, "Group__in")
	//亲子游戏
	group,_ := m.GetGroupByFields(map[string]interface{}{"Path": "game"}, []string{"id"})
	filters["Group__in"] = group.Id
	filters["IsCommand"] = 1
	news, _ := m.GetLimtArticle(filters, []string{"Id", "Title", "Group", "Cover", "Memo", "CreatedAt", "Path"}, 8, "Id")
	c.Data["GameCommandCovers"] = news[:1]
	c.Data["GameCommandNews"] = news[1:]
	delete(filters, "Group__in")
	//故事
	story,_ := m.GetGroupByFields(map[string]interface{}{"Path": "music"}, []string{"id"})
	filters["Group__in"] = story.Id
	stories, _ := m.GetLimtArticle(filters, []string{"Id", "Title", "Group", "Cover", "Memo", "CreatedAt", "Path"}, 10, "Id")
	c.Data["StoryCommandCovers"] = stories[:2]
	c.Data["StorCommandNews"] = stories[2:]
}
