<div class="f-hovertab-cont clearfix m-bydiv" style="display:none">
    [[range .BeiyunfertilityCovers]]
    <div class="m-imgnew m-big f-fl"> <a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"> <img src="[[.Cover]]" alt="[[.Title]]">
      <p>[[.Title]]</p>
      </a>
    </div>
    [[end]]
    <div class="f-fr m-inbyright">
      <ul class="clearfix f-lifl m-inbyul">
        [[range .BeiyunfertilityMidNews]]
        <li>
          <div class="f-tw"> <a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank"> <img src="[[.Cover]]" alt="[[.Title]]">
            <p><span>[[.Title]]</span></p>
            </a>
            </div>
        </li>
        [[end]]
      </ul>
      <ul class="m-newtopsul">
      [[range .BeiyunfertilityNews]]
        <li><span class="f-fr">[[dateformat .CreatedAt "01-02"]]</span><span class="f-fl"></span><a href="/[[.Path]]/info/cms_[[.Id]].html" target="_blank">[[.Title]]</a></li>
      [[end]]
      </ul>
    </div>
</div>