<div class="layui-side layui-side-menu">
    <div class="layui-side-scroll">
        <div class="layui-logo">
            <span>店酷云运营平台</span>
        </div>
        <ul class="layui-nav layui-nav-tree" lay-shrink="all" id="LAY-system-side-menu" lay-filter="layadmin-system-side-menu">
           [[range $index, $elem := .Navigations]]
               [[if $elem.children]]]
                   <li class="layui-nav-item [[if eq 0 $index]]layui-nav-itemed [[end]]">
                       <a href="javascript:;">
                           <i class="layui-icon [[$elem.icon]]"></i><cite>[[$elem.name]]</cite>
                       </a>
                       [[range $ck, $cv := $elem.children]]
                           <dl class="layui-nav-child">
                               <dd class="layui-this">
                                   <a lay-href="[[$cv.url]]">[[$cv.name]]</a>
                               </dd>
                           </dl>
                       [[end]]
                   </li>
               [[else]]
                   <li class="layui-nav-item">
                       [[if $elem.icon]] <i class="layui-icon [[$elem.icon]]"></i> [[end]]
                       <a lay-href="[[$elem.url]]">[[$elem.name]]</a>
                   </li>
               [[end]]
           [[end]]
        </ul>

    </div>
</div>