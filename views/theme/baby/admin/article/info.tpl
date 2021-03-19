<div class="layui-fluid">
    <div class="layui-card-body C-marginTop-10">
        <form class="layui-form base_form" action='[[urlfor "admin.ArticleController.Put" ":id" .Info.Id]]' method="post">
        <input type="hidden" name="_method" value="PUT">
        [[.xsrfdata]]
          <div class="layui-form-item">
            <label class="layui-form-label">标题</label>
            <div class="layui-input-block">
              <input type="text" name="title" value="[[.Info.Title]]" required  lay-verify="required" placeholder="请输入标题" autocomplete="off" class="layui-input col-xs-10">
            </div>
          </div>
          <div class="layui-form-item">
              <label class="layui-form-label">封面</label>
              <div class="layui-input-block">
                  <div class="layui-upload" style="display: flex;align-items: flex-end;">
                      <div class="layui-upload-list" style="margin:0">
                          <ul id="layui-upload-box" class="layui-clear uploadPImg" style="padding-left: 0">
                            [[if .Info.Cover]]
                            <li>
                                <img src="[[.Info.Cover]]">
                                <input type="hidden" name="img" value="[[.Info.Cover]]">
                                <span hidden="" class="img-delete" style="display: none;">
                                <i class="icon-shanchu iconfont"></i></span>
                            </li>
                            [[end]]
                          </ul>
                      </div>
                      <button type="button" class="layui-btn" id="uploadP" lay-data="{name:'img',num:1}">
                          <i class="layui-icon">&#xe67c;</i>上传图片
                      </button>
                      <div class="C-lineHeight-1 C-marginLeft-10">
                          <span>建议上传图片尺寸:800×800</span>
                      </div>
                  </div>
              </div>
            </div>
          <div class="layui-form-item">
            <label class="layui-form-label">描述</label>
            <div class="layui-input-block">
              <textarea name="memo" placeholder="请输入内容" class="layui-textarea col-xs-10">[[.Info.Memo]]</textarea>
            </div>
          </div>
          <div class="layui-form-item">
            <label class="layui-form-label">标签</label>
            <div class="layui-input-block">
              [[range .Tags]]
              <input type="checkbox" [[if InArray .Id $.TagIds]]checked[[end]] value="[[.Id]]" name="tag[]" lay-skin="primary" title="[[.Name]]">
              [[end]]
            </div>
          </div>
          <div class="layui-form-item">
            <label class="layui-form-label">分类</label>
            <div class="layui-input-block">
              <div class="layui-inline">
                  <select class="class-A" name="cate" lay-filter="class-A" lay-verify="required" class="province"
                          data-id="0">
                      <option value="">选择分类</option>
                  </select>
                  <input value="[[.Info.Path]]" name="ocate_path" class="ocate-path" type="hidden">
              </div>
              <div class="layui-inline">
                  <select class="class-B" name="cate_child" class="city" lay-filter="class-B" lay-verify="required" data-id="0">
                      <option value="">选择分类</option>
                  </select>
              </div>
              <div class="layui-inline">
                  <select class="class-C" name="cate_sun" class="district" lay-verify="required" data-id="">
                      <option value="">选择分类</option>
                  </select>
              </div>
            </div>
          </div>
          <div class="layui-form-item">
            <label class="layui-form-label">&nbsp;</label>
            <div class="layui-input-block">
              <input type="checkbox" name="is_hot" title="热门" [[if .Info.IsHot]]checked[[end]]>
              <input type="checkbox" name="is_banner" title="Banner" [[if .Info.IsBanner]]checked[[end]]>
              <input type="checkbox" name="is_today" title="今日推荐" [[if .Info.IsToday]]checked[[end]]>
              <input type="checkbox" name="is_toutiao" title="头条" [[if .Info.IsToutiao]]checked[[end]]>
              <input type="checkbox" name="is_command" title="推荐" [[if .Info.IsCommand]]checked[[end]]>
              <input type="checkbox" name="is_wonderful" title="精选" [[if .Info.IsWonderful]]checked[[end]]>
            </div>
          </div>
          <div class="layui-form-item">
            <label class="layui-form-label">显示</label>
            <div class="layui-input-block">
              <input type="checkbox" [[if .Info.Status]]checked[[end]] value="[[.Info.Status]]" name="status" lay-skin="switch">
            </div>
          </div>
          <div class="layui-form-item">
              <label class="layui-form-label">作者</label>
              <div class="layui-input-inline">
                <select name="author" lay-verify="required">
                  <option value=""></option>
                  [[range .Authors]]
                  <option value="[[.Id]]" [[if eq .Id $.Info.AuthorId]]selected[[end]]>[[.Name]]</option>
                  [[end]]
                </select>
              </div>
            </div>
            <div class="layui-form-item">
                <label class="layui-form-label">来源</label>
                <div class="layui-input-inline">
                  <!--range 相当于一个闭包需要使用$. 来引用上下文-->
                  <select name="source" lay-verify="required">
                    <option value=""></option>
                    [[range .Sources]]
                    <option value="[[.Id]]" [[if eq .Id $.Info.SourceId]]selected[[end]]>[[.Name]]</option>
                    [[end]]
                  </select>
                </div>
          </div>
          <div class="layui-form-item">
            <label class="layui-form-label">排序</label>
            <div class="layui-input-inline">
              <input type="text" value="[[.Info.Sort]]" name="sort" required  lay-verify="required" placeholder="请输入排序" autocomplete="off" class="layui-input col-xs-3">
            </div>
          </div>
          <div class="layui-form-item layui-form-text">
            <label class="layui-form-label">文本域</label>
            <div class="layui-input-block" style="text-align:left">
                <div id="content">
                <textarea style="display:none" placeholder="请输入内容" class="layui-textarea col-xs-10" name="content">[[.Detail.Content]]</textarea>
                </div>
             </div>
          </div>
          <div id="secondBox"></div>
          <div class="layui-form-item  layui-layout-admin">
              <div class="layui-input-block">
                  <div class="layui-footer" style="left: 0;">
                      <button type="submit" class="layui-btn layui-btn-normal"> 确认 </button>
                      <a class="layui-btn layui-btn-primary" id="layui-form-close"> 返回 </a>
                  </div>
              </div>
          </div>
        </form>
    </div>
</div>