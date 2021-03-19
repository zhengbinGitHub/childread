[[assets_js "/static/js/turndown.js"]]
[[assets_js "/static/js/turndown-plugin-gfm.js"]]
[[assets_js "/static/editor-md/editormd.js"]]
<script type="text/javascript">
  let uploadUrl = '[[urlfor "ArticleController.Upload" ]]';
  let editor = editormd("content", {
      width:"auto",
      height:"500",
      theme : "",
      toc: false,
      saveHTMLToTextarea: true,
      previewTheme : "",
      editorTheme : "default",
      codeFold : true,
      htmlDecode : true,
      tex : true,
      taskList : true,
      emoji : false,
      flowChart : true,
      sequenceDiagram : true,
      path:"/static/editor-md/lib/",
      autoFocus:false,
      placeholder:"",
      imageUpload : true,
      imageFormats : ["jpg", "jpeg", "gif", "png"],
      imageUploadURL : uploadUrl,
      onload : function() {
          let content = $("#content").find('textarea[name="content"]').val();

          let tables = turndownPluginGfm.tables;
          let strikethrough = turndownPluginGfm.strikethrough;
          let service = new TurndownService();

          service.use([tables, strikethrough]);
          service.keep('video');

          this.setMarkdown(service.turndown(content));
      }
  });
  layui.use(['form', 'upload'], function(){
      var $ = layui.jquery
          ,form = layui.form;
      admin.uploadImgBox('#uploadP', '.uploadPImg', uploadUrl)
      form.render();

      getAjax(0, 1, [[.Pid]])
      [[if gt .Cid 0]]
          getAjax([[.Pid]], 2, [[.Cid]])
      [[end]]
      [[if gt .Sid 0]]
          getAjax([[.Cid]], 3, [[.Sid]])
      [[end]]

      form.on('select(class-A)', function (data) {
          let ocatePath = $(data.elem).find("option:selected").attr("data-path");
          $(document).find('.ocate-path').val(ocatePath)
          return getAjax(data.value, 2)
      });
      form.on('select(class-B)', function (data) {
            return getAjax(data.value, 3)
        });

      function getAjax(id, type, current) {
        if(type != 1 && id == 0){
            layer.msg('分类ID为空')
            return false;
        }
        $.ajax({
            type: "GET",
            dataType: "json",
            url: '[[urlfor "GroupController.GetAjaxGroup"]]?id=' + id,
            success: function (res) {
                str = `<option value="">选择${type == 1 ? '一' : '' || type == 2 ? '二' : '' || type == 3 ? '三' : ''}级品类</option>`;
                if (res) {
                    $.each(res.data.group, function (index, val) {
                        str += `<option data-path="${val.Path}" value="${val.Id}" ${current == val.Id ? 'selected' : ''}>${val.Title}</option>`
                    });
                    switch (type) {
                        case 1:
                            $(".class-A").html(str);
                            break;
                        case 2:
                            $(".class-B").html(str);
                            break;
                        case 3:
                            $(".class-C").html(str);
                            break;
                    }
                    form.render();
                }
            }
        });
    }
  });
</script>