<script>
    layui.use(['form', 'table','laydate'], function(){
        var laydate = layui.laydate,
            table = layui.table

        table.init("table-hide");
        let SearchHref = location.search;
        if($("input[name=started_at]").val() !="" && $("input[name=ended_at]").val() !=""){
            admin.dateTime('#active','date',$("input[name=started_at]").val()+ ' - ' +$("input[name=ended_at]").val(),'','',true,'yyyy-MM-dd');
        }else{
            if(SearchHref == ""){
                admin.dateTime('#active','date',admin.getLast3Month(3).last + ' - ' +admin.getLast3Month(3).now,'','',true,'yyyy-MM-dd');
                $('input[name="started_at"]').val(admin.getLast3Month(3).last);
                $('input[name="ended_at"]').val(admin.getLast3Month(3).now);
            }else{
                admin.dateTime('#active','date','','','',true,'yyyy-MM-dd');
            }
        }
        admin.paginate("[[.Total]]", "[[.Current_page]]", "[[.paginator.PerPageNums]]");
    });
</script>