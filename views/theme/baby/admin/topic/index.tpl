    <div class="layui-card-body C-marginTop-10">
        <div class="layui-btn-container">
            <span class="layui-btn form_view" data-href='[[urlfor "TopicController.Edit"]]'>添加专题</span>
        </div>
        <table  lay-filter="table-hide" style="display: none" lay-data="{height:'full-310',height: 600, cellMinWidth: 80,toolbar: '#toolbar', limit: [[.Page_size]] }">
            <thead>
            <tr>
                <th lay-data="{field:'ID'}">ID</th>
                <th lay-data="{field:'Cover'}">图片</th>
                <th lay-data="{field:'Url'}">连接</th>
                <th lay-data="{field:'CreatedAt'}">创建时间</th>
                <th lay-data="{field:'opt', fixed: 'right',width:120}">操作</th>
            </tr>
            </thead>
            <tbody>
            [[range .Lists]]
                <tr>
                    <td>[[.Id]]</td>
                    <td><img src="[[.Cover]]"></td>
                    <td>[[.Url]]</td>
                    <td>[[dateformat .CreatedAt "2006-01-02"]]</td>
                    <td>
                        <a class="G-color-blue C-cursor-pointer C-paddingLr-5" onclick="admin.confirmPopup('确认删除吗？','[[urlfor "admin.TopicController.Delete" ":id" .Id]]')">删除</a>
                        <a class="G-color-blue C-cursor-pointer C-paddingLr-5 form_view" href='javascript:void(0)' data-href='[[urlfor "admin.TopicController.GetOne" ":id" .Id]]'>编辑</a>
                    </td>
                </tr>
            [[end]]
            </tbody>
        </table>
        <div id="page"></div>
    </div>