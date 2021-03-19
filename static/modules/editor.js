layui.define(["jquery"],function(exports){
    var $ = layui.$;
    var ue = UE.getEditor('editor');

    ue.ready(function() {
        ue.execCommand('serverparam','_token',$('meta[name="csrf_token"]').attr('content'));
    });

    exports('editor', {});
});  