<!doctype html>
<html lang="zh">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>运营平台登录</title>
    [[assets_css "/static/css/common.css"]]
    [[assets_css "/static/layui/css/layui.css"]]
    [[assets_css "/static/css/login.css"]]
</head>
<body  class="layui-layout-body" layadmin-themealias="default">
<div class="layui-layout layui-layout-admin" id="LAY_app" >
    <div class="login-con  layui-form">
        <div class='login-hd'>
            <img class='hd-logo' src="/static/images/logo.png" >
            <p class='hd-title'>店酷云运营平台</p>
            <p class='hd-prop'>让装修简单有品味</p>
        </div>
        <form method="post" class="layui-form">
            [[.xsrfdata]]
            <div class="layui-form-item login-item">
                <label class="login-icon layui-icon layui-icon-username" for="login-username"></label>
                <input type="text" name="mobile" required lay-verify="required" placeholder="用户名" class="layui-input">
            </div>
            <div class="layui-form-item login-item">
                <label class="login-icon layui-icon layui-icon-password" for="login-password"></label>
                <input type="password" name="password" required lay-verify="required" placeholder="密码" class="layui-input">
            </div>
            <div class="layui-form-item">
                <input type="hidden" name="isajax" value="1">
                <button class="layui-btn layui-btn-fluid" lay-submit="" lay-filter="login">登  录</button>
            </div>
        </form>
    </div>
</div>
[[assets_js "/static/js/jqbs.min.js"]]
[[assets_js "/static/layui/layui.js"]]
<script>
        layui.use(['form'], function(){
            var form = layui.form
                ,$ = layui.jquery,
                layer = layui.layer,
                url = '[[urlfor "LoginController.Post"]]';

            form.on('submit(login)', function(data){
                $.post(url, data.field, function(response){
                    if(response.status == false){
                        layer.msg(response.info)
                        return false;
                    } else {
                        location.href = '[[urlfor "HomeController.Index"]]'
                    }
                })
              });
            form.render();
        });
    </script>
</body>
</html>