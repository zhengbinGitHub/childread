<form class="layui-form layui-card-body bgc-f2f C-border-10"  id="search-form">
        <div class="layui-form-item">
            <div class="layui-col-md4 layui-col-xs4 fl-row-center C-marginTop-10">
                <input type="text" name="title" maxlength="20" value="[[.Title]]" placeholder="问题标题搜索" autocomplete="off"
                       class="layui-input col-xs-9" />
            </div>
        </div>
        <div class="fl-row-center">
            <input type="submit" class="layui-btn " value="搜索">
            <button type="reset" class="layui-btn layui-btn-primary" id="layui-form-close">重置</button>
        </div>
    </form>
    <div class="layui-card-body C-marginTop-10">
        <div class="layui-btn-container">
            <span class="layui-btn form_view" data-href='[[urlfor "ArticleController.Edit"]]'>添加文章</span>
        </div>
        <table  lay-filter="table-hide" style="display: none" lay-data="{height:'full-310',height: 600, cellMinWidth: 80,toolbar: '#toolbar', limit: [[.Page_size]] }">
            <thead>
            <tr>
                <th lay-data="{field:'ID'}">ID</th>
                <th lay-data="{field:'Title'}">标题</th>
                <th lay-data="{field:'CreatedAt'}">提问时间</th>
                <th lay-data="{field:'opt', fixed: 'right',width:120}">操作</th>
            </tr>
            </thead>
            <tbody>
            [[range .Lists]]
                <tr>
                    <td>[[.Id]]</td>
                    <td>[[.Title]]</td>
                    <td>[[dateformat .CreatedAt "2006-01-02"]]</td>
                    <td>
                        <a class="G-color-blue C-cursor-pointer C-paddingLr-5" onclick="admin.confirmPopup('确认删除吗？','[[urlfor "admin.ArticleController.Delete" ":id" .Id]]')">删除</a>
                        <a class="G-color-blue C-cursor-pointer C-paddingLr-5 form_view" data-href='[[urlfor "admin.ArticleController.GetOne" ":id" .Id]]'>查看</a>
                    </td>
                </tr>
            [[end]]
            </tbody>
        </table>
        <div id="page"></div>
    </div>