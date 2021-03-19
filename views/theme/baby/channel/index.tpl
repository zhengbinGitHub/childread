<!-- 内容开始 -->
<dl class="g-box-1200 clearfix">
    <dt class="f-fl m-mianleft">
        <p class="m-navwz">当前位置：
            <a href="[[.Domain]]">主页</a>
            [[range $index, $elem := .MenuParents]]
                > <a href="/[[Breadcrumb $.MenuPaths $index]]">[[$elem.Title]]</a>
            [[end]]
          >  当前栏目
        </p>
        <ul class="m-listul">
          [[range .Lists]]
          <li class="clearfix" data-cid="[[.Cid]]"> <a href="/[[$.InfoPath]]/info/cms_[[.Id]].html" target='_blank' class='f-fl'><img src="[[.Cover]]" alt="[[.Title]]"/></a>
            <div class="m-listdiv">
              <p class="m-title"><a href="/[[$.InfoPath]]/info/cms_[[.Id]].html" target="_blank">[[.Title]]</a></p>
              <div class="m-listinfo">[[substr .Memo 0 46]]</div>
              <p class="m-listime">
                  <span class="f-fl m-listags">
                  [[range .Tags]]
                    <a href="[[urlfor "web.TagController.GetAll" ":id" .Id]]" >[[.Name]]</a>
                  [[end]]
                  </span>
                  <span class="f-fr">
                    <a href="/[[$.InfoPath]]/info/cms_[[.Id]].html">
                    <i class="f-ico1"></i>[[.View]]</a>
                  </span>
                  <span class="f-fr">
                    <a href="/[[.Path]]/[[.Group]]">[[.GroupName]]</a> / [[date .CreatedAt "Y-m-d"]]
                  </span>
              </p>
            </div>
          </li>
          [[end]]
        </ul>
        [[template "layouts/paginator.tpl" .]]
  </dt>

  <dd class="f-fr m-mainright">
      <div class="m-mdiv">
        <h3 class="m-h3tit ico2">本类<span>推荐</span></h3>
        <ul class="m-hotgul">
        [[range .CommandArticles]]
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
      <div class="m-mdiv">
        <h3 class="m-h3tit ico2">本类<span>排行</span></h3>
        <ul class="m-blphul">
          [[range $index, $elem := .ViewArticles]]
          <li class="f-ix f-t3"><a href="/[[$.InfoPath]]/info/cms_[[$elem.Id]].html" target="_blank" class="m-nsmall"><em>[[$index | Index]]</em>[[$elem.Title]]</a></li>
          [[end]]
        </ul>
      </div>
      <div class="m-mdiv">
        <h3 class="m-h3tit ico2">热门<span>标签</span></h3>
        <ul class="m-mtags">
           [[range .HotTags]]
           <a href="[[urlfor "web.TagController.GetAll" ":id" .Id]]" target="_blank">[[.Name]]</a>
           [[end]]
        </ul>
      </div>
  </dd>
</dl>