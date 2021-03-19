<script type="text/javascript">
  let uploadUrl = '[[urlfor "ArticleController.Upload" ]]';
  layui.use(['form', 'upload'], function(){
      var $ = layui.jquery
          ,form = layui.form;
      admin.uploadImgBox('#uploadP', '.uploadPImg', uploadUrl)
      form.render();
    });
</script>