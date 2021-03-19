<div class="layui-header">
<ul class="layui-nav layui-layout-right" lay-filter="layadmin-layout-right">
    <li class="layui-nav-item" lay-unselect>
        <a href="javascript:;" layadmin-event="refresh" title="刷新">
            <i class="layui-icon layui-icon-refresh-3" lay-tips="刷新当前标签"></i>
        </a>
    </li>

    <li class="layui-nav-item layui-hide-xs" lay-unselect>
        <a href="javascript:;" layadmin-event="theme">
            <i class="layui-icon layui-icon-theme" lay-tips="切换皮肤"></i>
        </a>
    </li>

    <li class="layui-nav-item layui-hide-xs" lay-unselect>
        <a href="javascript:;" layadmin-event="fullscreen">
            <i class="layui-icon layui-icon-screen-full" lay-tips="全屏"></i>
        </a>
    </li>
    <li class="layui-nav-item" lay-unselect>
        <a href="javascript:;">
            <cite>&nbsp;[[.userinfo.Nickname]]</cite>
        </a>
        <dl class="layui-nav-child">
            <dd><a href="javascript:;" lay-href='[[urlfor "admin.LoginController.Password"]]'>修改密码</a></dd>
            <hr>
            <dd><a href='[[urlfor "admin.LoginController.Logout"]]'>退出系统</a></dd>
        </dl>
    </li>
    <li class="layui-nav-item layui-hide-xs" lay-unselect>
        <a href="javascript:;">
            <i class="layui-icon layui-icon-survey" lay-tips="帮助手册(未完成)"></i>
        </a>
    </li>
</ul>
</div>