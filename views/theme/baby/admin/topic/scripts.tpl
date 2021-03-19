<script>
    layui.use(['form', 'table','laydate'], function(){
        var laydate = layui.laydate,
            table = layui.table

        table.init("table-hide");
        admin.paginate("[[.Total]]", "[[.Current_page]]", "[[.paginator.PerPageNums]]");
    });
</script>