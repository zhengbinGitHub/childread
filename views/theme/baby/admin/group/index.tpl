<form class="layui-form layui-card-body bgc-f2f C-border-10" id="search-form">
        <div class="layui-form-item">
            <div class="layui-col-md4 fl-row-center">
                <input type="text" name="name" maxlength="20" value="[[.Name]]" placeholder="分类名称搜索" autocomplete="off"
                    class="layui-input col-xs-9" />
            </div>
        </div>

        <div class="fl-row-center">
            <input type="submit" class="layui-btn " value="搜索">
            <button type="reset" class="layui-btn layui-btn-primary" id="layui-form-close">重置</button>
        </div>
    </form>

    <div id='addGroupAlert' hidden>
        <div class='layui-form layui-card-body C-border-10'>
            <div  class="fl-row-leftNowrap " >
                <span class='layui-col-md4'>分类名称:</span>
                <div class=" layui-input-inline layui-col-md8">
                    <input type="text"  name="title" required  lay-verify="required" placeholder="请输入分类名称" autocomplete="off" class="layui-input alert-ipt">
                    <input type="hidden" name="parent_id" value="[[.Parent_id]]" class="alert-pid">
                </div>
            </div>
        </div>
    </div>
    <div class="layui-card-body C-marginTop-10">
        <div class="layui-btn-container">
            <span class="layui-btn" id='addGroup'>添加分类</span>
        </div>

        <table  lay-filter="table-hide" style="display: none" lay-data="{height:'full-310',height: 600, cellMinWidth: 80,toolbar: '#toolbar', defaultToolbar:['filter'], limit: [[.Page_size]] }">
            <thead>
            <tr>
                <th lay-data="{field:'id',sort:'true'}">ID</th>
                <th lay-data="{field:'title'}">分类名称</th>
                <th lay-data="{field:'priority'}">子类</th>
                <th lay-data="{field:'status'}">状态</th>
                <th lay-data="{field:'operate', fixed: 'right',width:160}">操作</th>
            </tr>
            </thead>
            <tbody>
            [[range .Lists]]
                <tr>
                    <td>[[.Id]]</td>
                    <td>[[.Title]]</td>
                    <td><a class="G-color-blue C-cursor-pointer C-paddingLr-5 form_view" data-href="[[urlfor "GroupController.Index" ":level" "1" ":id" .Id]]">下级</a></td>
                    <td><input type="checkbox" lay-filter="tag-group" data-url='[[urlfor "GroupController.Switch" ":id" .Id]]' lay-text="正常|关闭" lay-skin="switch"[[if eq .Status 1]] checked[[end]]></td>
                    <td>
                        <a href='javascript:void(0)' class="G-color-blue C-cursor-pointer C-paddingLr-5 editGroup" data-id="[[.Id]]" data-name="[[.Title]]" >修改</a>
                    </td>
                </tr>
            [[end]]
            </tbody>
        </table>
        <div id="page"></div>
    </div>