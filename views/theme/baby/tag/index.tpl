<!-- 内容开始 -->
  <dt class="f-fl m-mianleft">
    <p class="m-navwz">当前位置：<a href="[[.Domain]]">主页</a> > <a href="[[urlfor "TagController.GetAll" ":id" .TagInfo.Id]]">TAG标签</a> > [[.TagInfo.Name]]</p>
    <ul class="m-listul">
      [[range .Lists]]
      <li class="clearfix" data-cid="[[.Cid]]">
      <a href="/[[.Path]]/info/cms_[[.Id]].html" target='_blank' class='f-fl'><img src="[[.Cover]]"/></a>
        <div class="m-listdiv">
          <p class="m-title"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"><b>[[.Title]]</b></a></p>
          <div class="m-listinfo">[[substr .Memo 0 46]]</div>
          <p class="m-listime">
              <span class="f-fl m-listags">
                [[range .Tags]]
                   <a href="[[urlfor "web.TagController.GetAll" ":id" .Id]]" >[[.Name]]</a>
                [[end]]
              </span>
              <span class="f-fr"><a href="/[[.Path]]/info/cms_[[.Id]].html"><i class="f-ico1"></i>[[.View]]</a></span><span class="f-fr">[[date .CreatedAt "Y-m-d H:i:s"]]</span>
          </p>
        </div>
      </li>
      [[end]]
    </ul>
    [[template "layouts/paginator.tpl" .]]
  </dt>
  <dd class="f-fr m-mainright">
    <div class="m-mdiv">
      <h3 class="m-h3tit ico2">本类<span>排行</span></h3>
      <ul class="m-blphul">
        [[range $index, $elem := .ViewArticles]]
            <li class="f-ix f-t3">
                <a href="/[[$elem.Path]]/info/cms_[[$elem.Id]].html" target="_blank" class="m-nsmall"><em>[[$index | Index]]</em>[[$elem.Title]]</a>
            </li>
        [[end]]
      </ul>
    </div>
    <div class="m-mdiv">
      <h3 class="m-h3tit ico2">热门<span>标签</span></h3>
      <ul class="m-mtags">
        [[range .HotTags]]
         <a href="[[urlfor "TagController.GetAll" ":id" .Id]]">[[.Name]]</a>
         [[end]]
      </ul>
    </div>
  </dd>