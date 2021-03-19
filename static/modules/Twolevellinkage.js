layui.define(["layer","jquery","form"],function(exports){
    var $ = layui.$
    ,admin = layui.admin
    ,table = layui.table
    ,element = layui.element
    ,layer = layui.layer
    ,laydate = layui.laydate
    ,form = layui.form;
    $.ajax({
        type: 'POST',
        url: "",
        data: {areaId:areaId},
        dataType:  'json',
        success:function(e){
            //empty() 方法从被选元素移除所有内容
            $("select[name='data[group_id]']").empty();
            var html = "<option value='0'>选择一级类目</option>";
            $(e.data).each(function (v, k) {
                html += "<option value='" + k.id + "'>" + k.title + "</option>";
            });
            //把遍历的数据放到select表里面
            $("select[name='data[group_id]']").append(html);
            //从新刷新了一下下拉框
            form.render('select');      //重新渲染
        }
    });
    form.on('select(magazine)', function(data){
        var areaId=data.elem.value;
        $.ajax({
            type: 'POST',
            url: "",
            data: {areaId:areaId},
            dataType:  'json',
            success:function(e){
                //empty() 方法从被选元素移除所有内容
                $("select[name='data[type_id]']").empty();
                var html = "<option value='0'>选择二级类目</option>";
                $(e.data).each(function (v, k) {
                    html += "<option value='" + k.id + "'>" + k.title + "</option>";
                });
                //把遍历的数据放到select表里面
                $("select[name='data[type_id]']").append(html);
                //从新刷新了一下下拉框
                form.render('select');      //重新渲染
            }
        });
    });
    exports('region',{});
});