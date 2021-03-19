    <form class="layui-form layui-card-body bgc-f2f C-border-10"  id="search-form">
        <div class="layui-form-item">
            <div class="layui-col-md4 layui-col-xs4 fl-row-center C-marginTop-10">
                <input type="text" name="name" maxlength="20" value="[[.Name]]" placeholder="来源搜索" autocomplete="off"
                       class="layui-input col-xs-9" />
            </div>
        </div>
        <div class="fl-row-center">
            <input type="submit" class="layui-btn " value="搜索">
            <button type="reset" class="layui-btn layui-btn-primary" id="layui-form-close">重置</button>
        </div>
    </form>
    <div id='addSourceAlert' hidden>
        <div class='layui-form layui-card-body C-border-10'>
            <div  class="fl-row-leftNowrap " >
                <span class='layui-col-md4'>来源名称:</span>
                <div class=" layui-input-inline layui-col-md8">
                    <input type="text"  name="title" required  lay-verify="required" placeholder="请输入来源名称" autocomplete="off" class="layui-input alert-ipt">
                </div>
            </div>
        </div>
    </div>
    <div class="layui-card-body C-marginTop-10">
        <div class="layui-btn-container">
            <span class="layui-btn" id='addSource'>添加来源</span>
        </div>
        <table  lay-filter="table-hide" style="display: none" lay-data="{height:'full-310',height: 600, cellMinWidth: 80,toolbar: '#toolbar', limit: [[.Page_size]] }">
            <thead>
            <tr>
                <th lay-data="{field:'ID'}">ID</th>
                <th lay-data="{field:'Name'}">名称</th>
                <th lay-data="{field:'CreatedAt'}">创建时间</th>
                <th lay-data="{field:'opt', fixed: 'right',width:120}">操作</th>
            </tr>
            </thead>
            <tbody>
            [[range .Lists]]
                <tr>
                    <td>[[.Id]]</td>
                    <td>[[.Name]]</td>
                    <td>[[dateformat .CreatedAt "2006-01-02"]]</td>
                    <td>
                        <a class="G-color-blue C-cursor-pointer C-paddingLr-5" onclick="admin.confirmPopup('确认删除吗？','[[urlfor "admin.SourceController.Delete" ":id" .Id]]')">删除</a>
                        <a class="G-color-blue C-cursor-pointer C-paddingLr-5 editAuthor" href='javascript:void(0)' data-id="[[.Id]]" data-name="[[.Name]]">编辑</a>
                    </td>
                </tr>
            [[end]]
            </tbody>
        </table>
        <div id="page"></div>
    </div>