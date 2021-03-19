<div class="g-top-full"></div>
<div class="g-nav-full">
  <div class="g-box-1200"> <a href="[[.Domain]]" class="u-logo f-fl" alt="健康育儿母婴新闻资讯类网站织梦模板(带手机端)"><img src="/static/images/logo.jpg"></a> <span class="tip f-fl"></span>
    <form class="f-fr" name="formsearch" action='[[urlfor "web.SearchController.GetAll"]]'>
      <input name="q" type="text" class="u-search-input f_fl" id="search-keyword" value="[[.Keyword]]" onfocus="if(this.value=='请输入你要搜索的内容'){this.value='';}" onblur="if(this.value==''){this.value='请输入你要搜索的内容';}" />
      <input  type="submit" class="f_fl u-search-btn " value="搜索"/>
    </form>
  </div>
</div>
<div class="g-navlist f-hovertab-box">
  <div class="m-nav1">
    <ul class="g-box-1200 clearfix f-hovertab-btn">
      <li class="m-ind hover m-hover"><a href="/"><i></i>首页</a></li>
      [[range .Navs]]
        <li class="[[.Class_name]]"><a href="/[[.Path]]"><i></i>[[.Name]]</a></li>
      [[end]]
    </ul>
  </div>
  <div class="m-nan2 g-box-1200">
    <div class="m-intxt clearfix f-hovertab-cont"> <i>热点：</i>
        [[range .Hots]]
            <span><a href="/[[.Path]]/info/cms_[[.Id]].html">[[.Title]]</a></span>
        [[end]]
    </div>
    [[range $index, $elems := .Navs]]
    <div class="f-hovertab-cont m-nav2div" style="display:none">
        [[range .Children]]
            <a href="/[[$elems.Path]]/[[.Path]]">[[.Name]]</a>
        [[end]]
    </div>
    [[end]]
  </div>
</div>

<!-- /header -->