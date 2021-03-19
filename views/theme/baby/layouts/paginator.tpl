[[if gt .paginator.PageNums 1]]
<div id="pages">
<ul class="pagination">
    [[if .paginator.HasPrev]]
        <li><a href="[[.paginator.PageLinkFirst]]">上一页</a></li>
        <li><a href="[[.paginator.PageLinkPrev]]">&lt;</a></li>
    [[else]]
        <li class="layui-laypage-prev layui-disabled disabled"><a>上一页</a></li>
        <li class="layui-laypage-prev layui-disabled disabled"><a>&lt;</a></li>
    [[end]]
    [[range $index, $page := .paginator.Pages]]
        [[if $.paginator.IsActive .]]
            <li class="active">
                <span>[[$page]]</span>
            </li>
        [[else]]
            <li>
                <a href="[[$.paginator.PageLink $page]]">[[$page]]</a>
            </li>
        [[end]]
    [[end]]
    [[if .paginator.HasNext]]
        <li><a href="[[.paginator.PageLinkNext]]">&gt;</a></li>
        <li><a href="[[.paginator.PageLinkLast]]" class="layui-laypage-next">下一页</a></li>
    [[else]]
        <li class="disabled"><a>&gt;</a></li>
        <li class="disabled"><a>下一页</a></li>
    [[end]]
</ul>
</div>
[[end]]