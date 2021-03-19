<!--  推荐信息 -->
<dl class="g-box-1200 clearfix g-intop">
  <dt class="f-fl m-inleft">
    <div class="m-box f-flash">
      <ul class="m-flashul f-flashul">
        [[range .Banners]]
        <li><a href="[[.Path]]/info/cms_[[.Id]].html" title="[[.Title]]" target="_blank"><img alt="[[.Title]]" src="[[.Cover]]" />
          <p>[[.Title]]</p>
          </a></li>
        [[end]]
      </ul>
      <ul class="m-li f-lidian">
      </ul>
      <div class="but but-left"><</div>
      <div class="but but-right">></div>
      <div class="bar"></div>
    </div>
  </dt>
  <dd class="f-fr m-inright f-hovertab-box">
    <p class="m-headnews-tit f-hovertab-btn f-ico1"><i class="f-hover">热搜头条</i><i>今日更新</i></p>
    <div class="f-hovertab-cont">
      <ul class="m-headnews-list ">
        [[range $index, $elem := .Toutiaos]]
        <li>
          <p class="name"><em>[[$index | Index]]</em><a href="/[[.Path]]/info/cms_[[$elem.Id]].html" title="[[$elem.Title]]" target="_blank">[[$elem.Title]]</a></p>
        </li>
        [[end]]
      </ul>
    </div>
    <div class="f-hovertab-cont" style="display:none">
      <ul class="m-headnews-list ">
        [[range $index, $elem := .Tadays]]
            <li>
              <p class="name"><em>[[$index | Index]]</em><a href="/[[.Path]]/info/cms_[[$elem.Id]].html" title="[[$elem.Title]]" target="_blank">[[$elem.Title]]</a></p>
            </li>
        [[end]]
      </ul>
    </div>
  </dd>
</dl>

<!-- 首页热门 -->
<dl class="g-box-1200 clearfix g-inhot">
  <dt class="f-fl m-inleft clearfix">
    <div class="f-fl m-inph">
      <h5 class="m-intit">
        <div>人气<span>排行</span></div>
        <p>POPULARITY RANKING</p>
      </h5>
      <ul class="m-newhot f-top3">
        [[range $index, $elem := .Viewnews]]
        <li class="clearfix"><span class="f-fl">[[$index | Index]]</span><a href="/[[.Path]]/info/cms_[[$elem.Id]].html" title="[[$elem.Title]]" target="_blank">[[substr $elem.Title 0 26]]</a>
          <p><i class="f-ico1"></i>[[.View]]</p>
        </li>
        [[end]]
      </ul>
    </div>
    <div class="f-fr m-inhot f-hovertab-box">
      <p class="f-hovertab-btn m-hotdianp"><i class="m-hover">热点</i><i>孕期指导</i><i>育儿周刊</i><i>生活</i></p>
      <div class="f-hovertab-cont m-hotdiandiv">
        <ul class="m-hotul">
          [[range .Hotnews]]
          <li class="clearfix m-oneimg" data-cid="[[.Group.Id]]">
            <p class="m-img clearfix"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"><img src="[[.Cover]]" alt="[[.Title]]"></a></p>
            <p class="m-tit"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank">[[.Title]]</a></p>
            <div class="clearfix m-newtags m-ninfo">[[substr .Memo 0 46]]</div>
            <div class="clearfix m-newtime">
              <p class="f-fl"><a href="/[[.Path]]/[[.Group.Path]]">[[.Group.Name]]</a> / [[dateformat .CreatedAt "2006-01-02"]]</p>
              <p class="f-fr m-pl"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"><i class="f-ico1"></i>[[.View]]</a></p>
            </div>
          </li>
          [[end]]
        </ul>
      </div>
      <div class="f-hovertab-cont m-hotdiandiv">
        <ul class="m-hotul">
          [[range .Guidances]]
          <li class="clearfix m-oneimg" data-cid="[[.Group.Id]]">
            <p class="m-img clearfix"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"><img src="[[.Cover]]"></a></p>
            <p class="m-tit"><a href="/[[.Path]]/info/cms_[[.Id]].html" title="[[.Title]]" target="_blank">[[.Title]]</a></p>
            <div class="clearfix m-newtags m-ninfo">[[substr .Memo 0 46]]</div>
            <div class="clearfix m-newtime">
              <p class="f-fl"><a href="/[[.Path]]/[[.Group.Path]]">[[.Group.Name]]</a> / [[dateformat .CreatedAt "2006/01/02"]]</p>
              <p class="f-fr m-pl"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"><i class="f-ico1"></i>[[.View]]</a></p>
            </div>
          </li>
          [[end]]
        </ul>
      </div>
      <div class="f-hovertab-cont m-hotdiandiv">
        <ul class="m-hotul">
        [[range .Weeklies]]
          <li class="clearfix m-oneimg" data-cid="[[.Group.Id]]">
            <p class="m-img clearfix"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"><img src="[[.Cover]]"></a></p>
            <p class="m-tit"><a href="/[[.Path]]/info/cms_[[.Id]].html" title="[[.Title]]" target="_blank">[[.Title]]</a></p>
            <div class="clearfix m-newtags m-ninfo">[[substr .Memo 0 46]]</div>
            <div class="clearfix m-newtime">
              <p class="f-fl"><a href="/[[.Path]]/[[.Group.Path]]">[[.Group.Name]]</a> / [[dateformat .CreatedAt "2006/01/02"]]</p>
              <p class="f-fr m-pl"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"><i class="f-ico1"></i>[[.View]]</a></p>
            </div>
          </li>
          [[end]]
        </ul>
      </div>
      <div class="f-hovertab-cont m-hotdiandiv">
        <ul class="m-hotul">
          [[range .Lifes]]
            <li class="clearfix m-oneimg" data-cid="[[.Group.Id]]">
              <p class="m-img clearfix"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"><img src="[[.Cover]]"></a></p>
              <p class="m-tit"><a href="/[[.Path]]/info/cms_[[.Id]].html" title="[[.Title]]" target="_blank">[[.Title]]</a></p>
              <div class="clearfix m-newtags m-ninfo">[[substr .Memo 0 46]]</div>
              <div class="clearfix m-newtime">
                <p class="f-fl"><a href="/[[.Path]]/[[.Group.Path]]">[[.Group.Name]]</a> / [[dateformat .CreatedAt "2006/01/02"]]</p>
                <p class="f-fr m-pl"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"><i class="f-ico1"></i>[[.View]]</a></p>
              </div>
            </li>
            [[end]]
        </ul>
      </div>
    </div>
  </dt>
  <dd class="f-fr m-inright">
    <h5 class="m-intit">
      <div>独家<span>专题</span></div>
      <p>EXCLUSIVE THEME</p>
    </h5>
    <div class="m-imgnew">
    [[range .Topics]]
        <a href="[[.Url]]" target="_blank"><img src="[[.Cover]]"></a><br>
    [[end]]
      <h5 class="m-intit">
        <div>小编<span>精选</span></div>
        <p>CLASSIC ARTICLES</p>
      </h5>
      <ul class="m-newliul">
        [[range $index, $elem := .Wonderfuls]]
        <li><a href="/[[$elem.Path]]/info/cms_[[$elem.Id]].html" title="[[$elem.Title]]" target="_blank"><em>[[$index | Index]]</em>[[$elem.Title]]</a></li>
        [[end]]
      </ul>
    </div>
  </dd>
</dl>

<!-- 育儿之道 -->
<dl class="g-box-1200 clearfix g-yezd">
  <dt class="f-fl m-yezd f-ico1">育儿之道</dt>
  <dt class="f-fl"><img src="static/images/bao2.jpg"></dt>
  [[range .Yuers]]
    <dd class="f-fl [[.Class_name]]"><a href="/yuer/[[.Path]]" target="_blank"><i class="f-ico1"></i>[[.Title]]</a></dd>
  [[end]]
</dl>

<!-- 备孕 -->
<div class="g-box-1200 g-beiyun">
  <h2 class="clearfix g-tith2">
    <p class="f-fl m-h2left"><i class="f-ico1 m-beiyun"></i><span>备孕</span><i>pregnancy</i></p>
    <p class="f-fr m-h2right"></p>
  </h2>
  <dl class="clearfix">
    <dt class="f-fl m-inleft f-hovertab-box">
      <div class="m-tit3">
        <ul class="clearfix f-lifl f-hovertab-btn">
          [[range .Beiyuns]]
            <li class="f-hover" data-gid="[[.Id]]"><a href="/[[.Path]]" target="_blank">[[.Title]]</a></li>
          [[end]]
        </ul>
      </div>
      [[template "default/beiyun/diet.tpl" .]]
      [[template "default/beiyun/bisxual.tpl" .]]
      [[template "default/beiyun/give-birth.tpl" .]]
      [[template "default/beiyun/pregnant.tpl" .]]
      [[template "default/beiyun/fertility-policy.tpl" .]]
      [[template "default/beiyun/pregnancy.tpl" .]]
    </dt>
    <dd class="f-fr m-inright">
      <h5 class="m-intit">
        <div>热点<span>文章</span></div>
        <p>HOT ARTICLE</p>
      </h5>
      <ul class="m-hotgul">
        [[range .BeiyunHotCovers]]
        <li><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank" class="clearfix">
          <p class="f-fl"><img src="[[.Cover]]" alt="[[.Title]]"></p>
          <strong>[[.Title]]</strong><span>[[dateformat .CreatedAt "06-01-02"]]</span></a></li>
        [[end]]
      </ul>
      <ul class="m-blphul">
      [[range $index, $elem := .BeiyunHotNews]]
        <li class="f-ix f-t3"><a href="/[[$elem.Path]]/info/cms_[[$elem.Id]].html" target="_blank" class="m-nsmall"><em>[[$index | Index]]</em>[[$elem.Title]]</a></li>
        [[end]]
      </ul>
    </dd>
  </dl>
</div>


<!-- 产后 -->
<div class="g-box-1200 clearfix g-chanhou">
  <h2 class="clearfix g-tith2">
    <p class="f-fl m-h2left"><i class="f-ico1 m-chanhou"></i><span>产后</span><i>postpartum</i></p>
    <p class="f-fr m-h2right"></p>
  </h2>
  <dl class="clearfix">
    <dt class="f-fl m-chdt">
    [[range .ChanhouToutiaoLeCovers]]
        <div class="m-imgnew m-saimg">
        <a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"> <img src="[[.Cover]]" alt="[[.Title]]">
        <p>[[.Title]]</p>
        </a>
        </div>
        [[end]]
 </dt>
    <dd class="f-fl m-chdd">
       [[range .ChanhouToutiaoMidCovers]]
      <div class="m-imgnew m-heimg"> <a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"> <img src="[[.Cover]]" alt="[[.Title]]">
        <p>[[.Title]]</p>
        </a>
 </div>
        [[end]]
      <ul class="m-ulimgnew f-top f-top3 ">
      [[range .ChanhouToutiaoMidNews]]
        <li class="f-t3"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank" class="m-nsmall">[[.Title]]</a>
          <div class="m-nbig clearfix"><a href="/[[.Path]]/info/cms_[[.Id]].html" class="m-ntit" target="_blank">[[.Title]]</a>
          <span class="f-fl"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"><img src="[[.Cover]]" alt="[[.Title]]"></a></span>
            <div>[[substr .Memo 0 46]]</div>
          </div>
        </li>
        [[end]]
      </ul>
    </dd>
    <dd class="m-inright f-fr">
      <h5 class="m-intit">
        <div>推荐<span>文章</span></div>
        <p>RECOMMENDED ARTICLES</p>
      </h5>
      <div class="m-imu">
        [[range .ChanhouCommandCovers]]
        <a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank">
        <p><img src="[[.Cover]]" alt="[[.Title]]"></p>
        <strong>[[.Title]]</strong></a>
        [[end]]
      </div>
      <ul class="m-blphul">
        [[range $index, $elem := .ChanhouCommandNews]]
            <li class="f-ix f-t3"><a href="/[[$elem.Path]]/info/cms_[[$elem.Id]].html" title="[[$elem.Title]]" target="_blank" class="m-nsmall"><em>[[$index | Index]]</em>[[$elem.Title]]</a></li>
        [[end]]
      </ul>
    </dd>
  </dl>
</div>

<!-- 早教 -->
<div class="g-box-1200 clearfix g-zaojiao">
  <h2 class="clearfix g-tith2">
    <p class="f-fl m-h2left"><i class="f-ico1 m-zaojiao"></i><span>早教</span><i>Early childhood</i></p>
    <p class="f-fr m-h2right"></p>
  </h2>
  <dl class="clearfix">
    <dt class="f-fl m-chdt">
    [[range .StoryCovers]]
        <div class="m-imgnew m-saimg"><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"> <img src="[[.Cover]]" alt="[[.Title]]"><p>[[.Title]]</p></a></div>
    [[end]]
    </dt>
    <dd class="f-fl m-chdd">
      <h5 class="m-intit">
        <div>故事<span>音乐</span></div>
        <p>CHILDREN STORY</p>
      </h5>
      <ul class="clearfix f-lifl m-inbyul">
        [[range .StoryCommandCovers]]
        <li>
          <div class="f-tw"> <a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"> <img src="[[.Cover]]" alt="[[.Title]]">
            <p><span>[[.Title]]</span></p>
            </a>
            </div>
        </li>
        [[end]]

      </ul>
      <ul class="m-ulimgnew f-top f-top3">
        [[range .StorCommandNews]]
        <li class="f-t3">
        <a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank" class="m-nsmall">[[.Title]]</a>
          <div class="m-nbig clearfix"><a href="/[[.Path]]/info/cms_[[.Id]].html" class="m-ntit" target="_blank">[[.Title]]</a><span class="f-fl">
          <a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"><img src="[[.Cover]]" alt="[[.Title]]"></a></span>
            <div>[[substr .Memo 0 46]]</div>
          </div>
        </li>
        [[end]]
      </ul>
    </dd>
    <dd class="m-inright f-fr">
      <h5 class="m-intit">
        <div>亲子<span>游戏</span></div>
        <p>CHILDREN MUSIC</p>
      </h5>
      [[range .GameCommandCovers]]
      <div class="f-tw f-twbig">
      <a href="/[[.Path]]" target="_blank"> <img src="[[.Cover]]" alt="[[.Title]]">
        <p><span>[[.Title]]</span></p>
        </a>
        </div>
      <div class="m-info">
        <div class="m-div">[[substr .Memo 0 46]]</div>
      </div>
      [[end]]
      <ul class="m-blphul">
        [[range $index, $elem := .GameCommandNews]]
        <li class="f-ix f-t3"><a href="/[[$elem.Path]]/info/cms_[[$elem.Id]].html" target="_blank" class="m-nsmall"><em>[[$index | Index]]</em>[[$elem.Title]]</a></li>
        [[end]]
      </ul>
    </dd>
  </dl>
</div>
<div class="g-box-1200 clearfix g-youlian">
  <h2 class="clearfix g-tith2">
    <p class="f-fl m-h2left"><i class="f-ico1 m-zaojiao"></i><span>友情链接</span><i>Early childhood</i></p>
  </h2>
  <ul class="f-lifl clearfix">
    <li><a href='http://www.eyoucms.com/' target='_blank'>易优CMS</a> </li><li><a href='http://www.sucai58.com/' target='_blank'>建站素材</a> </li><li><a href='http://www.yiyongtong.com/' target='_blank'>微信小程序开发</a> </li><li><a href='http://www.eyoucms.com/' target='_blank'>企业建站系统</a> </li>
  </ul>
</div>

<!-- 监听，返回顶部 -->
<div class="g-rightnav">
  <ul>
    <li><a class="m-rightnav-tj" href="javascript:;">推荐</a></li>
    <li><a class="m-rightnav-nx" href="javascript:;">独家</a></li>
    <li><a class="m-rightnav-ma" href="javascript:;">备孕</a></li>
    <li><a class="m-rightnav-my" href="javascript:;">产后</a></li>
    <li><a class="m-rightnav-yd" href="javascript:;">早教</a></li>
  </ul>
</div>