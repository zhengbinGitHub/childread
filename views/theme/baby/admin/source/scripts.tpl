<script>
        layui.use(['form', 'table'], function(){

            var table = layui.table
            var form = layui.form
            table.init("table-hide");

            admin.paginate("[[.Total]]", "[[.Current_page]]", "[[.paginator.PerPageNums]]");
            $(document).on('click',"#addSource",function(){
                    //判断是添加还是修改
                    $('.alert-ipt').val('')
                    layer.open({
                        type: 1,
                        btn: ['确定','取消'],
                        title:'添加一个新的来源',
                        content:$("#addSourceAlert"),
                        data:{'_xsrf':$('meta[name="csrf_token"]').attr("content")},
                        area: ['300px','200px'],
                        yes :function(index,layero){
                            if($('.alert-ipt').val()==''){
                                layer.msg('请输入来源名称', {time: 1000, icon: 5})
                                return false
                            }
                            $.ajax({
                                type: "POST",
                                url:'[[urlfor "SourceController.Post"]]',
                                data:{'_xsrf':$('meta[name="csrf_token"]').attr("content"),
                                'name': $('.alert-ipt').val()
                                },
                                success: function(response) {
                                    if(response.status==false){
                                        layer.msg(response.info, {time: 1000, icon: 5})
                                        return false
                                    }
                                    layer.msg(response.info, {time: 1000, icon: 6})
                                    setTimeout(function () {
                                        window.location.reload()
                                    }, 1000);
                                    layer.close(index)
                                },
                            });
                        }
                    })
                })
                $(document).on('click',".editAuthor",function(){
                    let currentGroup = $(this).attr('data-name');
                    let currentId = $(this).attr('data-id');
                    //判断是添加还是修改
                    $('.alert-ipt').val(currentGroup)
                    layer.open({
                        type: 1,
                        btn: ['确定','取消'],
                        title:'修改当前来源',
                        content:$("#addAuthorAlert"),
                        data:{'_token':$('meta[name="csrf_token"]').attr("content")},
                        area: ['300px','200px'],
                        yes :function(index,layero){
                            if($('.alert-ipt').val()==''){
                                layer.msg('请输入来源', {time: 1000, icon: 5})
                                return false
                            }
                            $.ajax({
                                type: "POST",
                                url:'[[urlfor "SourceController.Put"]]',
                                data:{
                                    '_method': 'PUT',
                                    '_xsrf':$('meta[name="csrf_token"]').attr("content"),
                                    'id': currentId,
                                    'name':$('.alert-ipt').val()
                                },
                                success: function(response) {
                                    if(response.status==0){
                                        layer.msg(response.message, {time: 1000, icon: 5})
                                        return false
                                    }
                                    layer.msg(response.message, {time: 1000, icon: 6})
                                    setTimeout(function () {
                                        window.location.reload()
                                    }, 1000);
                                    layer.close(index)
                                },
                            });
                        }
                    })
                })
        });
    </script>