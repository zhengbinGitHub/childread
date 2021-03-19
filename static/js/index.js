let timestamp = new Date().getTime();
layui.config({
    base: '/static/modules/',
    version:timestamp
}).use('message');
//
layui.use(['table','form','element','jquery','upload','message'],function(e){
    var $ = layui.jquery
    upload = layui.upload
        ,form = layui.form
        ,layer = layui.layer
        ,element = layui.element
        ,table = layui.table
        ,message = layui.message;
    // 表单提交 form标签上 + base_form 类名
    var options = {
        beforeSerialize: function() {
            $(':submit').attr('disabled', true);
        },
        success: function(data) {
            var index = parent.layer.getFrameIndex(window.name); //获取窗口索引
            if (data.status) {
                message.success(data.message);
                setTimeout(function() {
                    if (data.url) {
                        parent.layer.close(index);//关闭弹出的子页面窗口
                        if(index == undefined){
                            window.location = data.url;
                            element.render();
                            return false;
                        }else
                            parent.location = data.url;
                        element.render();
                        return false;
                    } else
                        parent.layer.close(index);
                    parent.location.reload();
                }, 1000);
            } else {
                $(':submit').attr('disabled', false);
                message.error(data.message);
            }
        }
    };
    $('.base_form').ajaxForm(options);
    $(document).on('mouseover','.uploadPImg li,.uploadBImg li,.posterImg li',function(){
        let th = $(this);
        th.find('.img-delete').css('display','block');
    })
    $(document).on('mouseout','.uploadPImg li,.uploadBImg li,.posterImg li',function(){
        let th = $(this);
        th.find('.img-delete').css('display','none');
    })
    $(document).on("click", ".img-delete", function() {
        let th = $(this);
        th.parents("li").remove();
    })
    //关闭页面
    $('#layui-form-close').on('click', function(){
        var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
        console.log(index);
        // if(index == undefined){
        //     parent.layui.admin.events.closeThisTabs();
        // }
        parent.layer.close(index); //再执行关闭
    });
    $(document).on('click','.form_view',function(){
        var url = $(this).data('href'),
            title = $(this).data('title');
        layer.close(layer.index);
        var index = layer.open({
            type: 2
            ,title: title
            ,content: url
            ,maxmin: false
            ,offset: ['0']
            ,area : ['100%','100%']
            ,resize:true//禁止拖拉框的大小
            ,fix : true // 固定
        });
    });
    form.on('switch', function(data){
        let value = 0,
            url = this.getAttribute('data-url');
        !this.checked ? data.elem.value = 0 :value = 1;data.elem.value = 1
        if(url == null || url == ''){
            return false;
        }else{
            $.ajax({
                type: "POST",
                async: false,
                data: {value: value,'_xsrf':$('meta[name="csrf_token"]').attr("content")},
                headers:{'_xsrf':$('meta[name="csrf_token"]').attr("content")},
                dataType: "json",
                url:url,
                success: function(res) {
                    if (res.status) {
                        message.success(res.message);
                        table.render();
                    }else{
                        message.error(res.message);
                        data.elem.checked = false;
                        form.render();
                    }
                }
            });
        }
    });
    // 重置
    $(document).on("click","button[type='reset']",function(e){
        $(this).parents('form').find('input[type!="hidden"]').val("");
        $(this).parents('form').find('select').val("");
        $(this).parents('form').find('input[name="started_at"],input[name="ended_at"]').val("");
        $(this).parents('form').find('input[name="started_time"],input[name="ended_time"]').val("");
        $(this).parents('form').find('.xm-select-default').val("");
        e.preventDefault();
        $(this).siblings("input[type='submit']").click();
    });
    // 弹窗内全选或全不选
    $(document).on('click','#allbox',function(){
        $(".js_table_id input[type='checkbox']").prop("checked", this.checked);
    });
})