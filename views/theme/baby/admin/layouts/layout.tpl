<!doctype html>
<html lang="zh">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="csrf_token" content="[[.token]]"/>
    <title>店酷云</title>
    <link rel="stylesheet" href="//at.alicdn.com/t/font_1471540_2lo3mrhr461.css?t=aaaa">
    [[assets_css "/static/layui/css/layui.css"]]
    [[assets_css "/static/css/common.css"]]
    [[assets_css "/static/css/app.css"]]
    <link rel="stylesheet" href="http://at.alicdn.com/t/font_1382198_ti4d3ypz0y.css">
    [[.HtmlHead]]
</head>
<body layadmin-themealias="default" style="height:100%">
<div class="layui-fluid">
[[.LayoutContent]]
</div>
[[assets_js "/static/vender/xmSelect/xm-select.js"]]
[[assets_js "/static/modules/static.js"]]
[[assets_js "/static/vender/jquery/1.12.3/jquery-1.12.3.js"]]
[[assets_js "/static/vender/jqueryform/jquery-form.js"]]
[[assets_js "/static/layui/layui.js"]]
[[assets_js "/static/js/index.js"]]
<script src="https://map.qq.com/api/js?v=2.exp&key=ECUBZ-FRJW3-HD43M-YPU3P-LOIW5-SPFUT"></script>
<script>
    layui.config({
        base: '/static/' //静态资源所在路径
    }).extend({
        index: 'lib/index', //主入口模块
    }).use('index');
</script>
[[.Scripts]]
</body>
</html>