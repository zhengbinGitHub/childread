
  <p class="m-navwz m-jieguo">系统搜索到约有<strong></strong>项符合<strong>[[.Keyword]]</strong>的查询结果</p>
  <ul class="m-listul">
    [[range .Lists]]
    <li class="clearfix" data-cid="16"> <a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank" class="f-fl"><img src="[[.Cover]]"></a>
      <div class="m-listdiv">
        <p class="m-title"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank">[[.Title]]</p>
        <div class="m-listinfo">[[substr .Memo 0 46]]</div>
        <p class="m-listime"><span class="f-fl m-listags"></span><span class="f-fr"><a href="/[[.Path]]/info/cms_[[.Id]].html"><i class="f-ico1"></i>[[.View]]</a></span><span class="f-fr">[[date .CreatedAt "Y-m-d"]]</span></p>
      </div>
    </li>
    [[end]]
  </ul>
  [[template "layouts/paginator.tpl" .]]