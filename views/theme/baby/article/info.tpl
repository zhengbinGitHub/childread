<!-- 内容开始 -->
<dl class="clearfix g-box-1200">
  <dt class="f-fl m-mianleft">
    <p class="m-navwz">当前位置：
        <a href="[[.Domain]]">主页</a>
         [[range $index, $elem := .MenuParents]]
             > <a href="/[[Breadcrumb $.MenuPaths $index]]">[[$elem.Title]]</a>
         [[end]]
           >
    </p>
    <h1>[[.Info.Title]]</h1>
    <p class="m-maininfo">
        <span>时间：[[date .Info.CreatedAt "Y-m-d H:i"]]</span>
        <span>来源：[[.Source.Name]]</span>
        <span>作者：[[.Author.Name]]</span>
        <span>点击：[[.Info.View]]</span>
    </p>
    <div class="m-mianjianjie ico2">
      <div><strong>导读：</strong></div>
    </div>
    <dl class="clearfix">
      <dt class="f-fl m-mianfx">
        <div class="m-fx"><img src="/static/images/saoyisao.gif" alt=""> </div>
        <div class="m-fxdiv"></div>
      </dt>
      <dd class="f-fr m-maintxt htmlcontent">
        <p>[[.Info.Memo]]</p>
      <p>
        [[str2html .Detail.Content]]
      </p>
        <div class="clearfix m-mainbot">
          [[if .Prev]]
          <p class="m-page-up f-fl">上一篇：<a disabled="true" href="/[[$.InfoPath]]/info/cms_[[.Prev.Id]].html">[[.Prev.Title]]</a> </p>
          [[end]]
          <a href="[[.Domain]]" class="u-back-home f-fl"><i></i><strong>网站首页</strong></a>
          <a href="/[[.InfoPath]]/[[.Info.Group.Path]]" class="u-back-list f-fl"><i></i><strong>返回栏目</strong></a>
          [[if .Next]]
          <p class="m-page-down f-fl">下一篇：<a disabled="true" href="/[[$.InfoPath]]/info/cms_[[.Next.Id]].html">[[.Next.Title]]</a> </p>
          [[end]]
        </div>
      </dd>
    </dl>
    <h5 class="m-pl-title"><span>最新<i>文章</i></span></h5>
    <div class="m-xg-ne">
      <ul class="m-xgwz-item">
        [[range $index, $elem := .BestNewArticles]]
            <li>
                <a href="/[[$.InfoPath]]/info/cms_[[.Id]].html"><em>[[$index | Index]]</em><img src="[[.Cover]]" alt="[[.Title]]">
                <p><strong>[[.Title]]</strong></p>
                </a>
            </li>
        [[end]]
      </ul>
    </div>
  </dt>
  <dd class="f-fr m-mainright">
    <div class="m-mdiv">
      <h3 class="m-h3tit ico2">热门<span>标签</span></h3>
      <p class="m-mtags">
        [[range .HotTags]]
           <a href="[[urlfor "web.TagController.GetAll" ":id" .Id]]" target="_blank">[[.Name]]</a>
        [[end]]
      </p>
    </div>
    <div class="m-mdiv m-twnews">
      <h3 class="m-h3tit ico2">热门<span>文章</span></h3>
      <ul class="m-hotgul">
        [[range .HotArticles]]
        <li>
          <a href="/[[$.InfoPath]]/info/cms_[[.Id]].html" target="_blank" class="clearfix">
              <p class="f-fl"><img src="[[.Cover]]" alt="[[.Title]]"></p>
              <strong>[[.Title]]</strong>
              <span>[[date .CreatedAt "Y-m-d"]]</span>
          </a>
        </li>
        [[end]]
      </ul>
    </div>
  </dd>
</dl>