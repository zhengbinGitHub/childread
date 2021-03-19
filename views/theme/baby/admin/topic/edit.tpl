<div class="layui-fluid">
    <div class="layui-card-body C-marginTop-10">
        <form class="layui-form base_form" action='[[urlfor "TopicController.Post"]]' method="post">
        [[.xsrfdata]]
          <div class="layui-form-item">
              <label class="layui-form-label">封面</label>
              <div class="layui-input-block">
                  <div class="layui-upload" style="display: flex;align-items: flex-end;">
                      <div class="layui-upload-list" style="margin:0">
                          <ul id="layui-upload-box" class="layui-clear uploadPImg" style="padding-left: 0"></ul>
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
              <label class="layui-form-label">URL</label>
              <div class="layui-input-block">
                <input type="text" name="url" required  lay-verify="required" placeholder="请输入连接" autocomplete="off" class="layui-input col-xs-10">
              </div>
            </div>
          <div class="layui-form-item">
            <label class="layui-form-label">显示</label>
            <div class="layui-input-block">
              <input type="checkbox" name="status" lay-skin="switch">
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