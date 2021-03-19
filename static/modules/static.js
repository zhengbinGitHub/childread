;!function (win) {

    var Admin = function () {
    };
    Admin.prototype.paginate = function (count, curr, limit, limits) {
      layui.laypage.render({
        elem: 'page',
        count: count,
        curr: curr,
        limit: limit,
        limits: limits ? limits : [20, 30, 40, 50],
        layout: ['count', 'prev', 'page', 'next', 'limit', 'refresh', 'skip'],
        jump: function(obj, first){
          if(!first){
            location.href= window.location.pathname + '?' + $("#search-form").serialize() + '&page='+obj.curr+'&limit='+obj.limit;
          }
        }
      });
    };
    Admin.prototype.tableDataDelete = function (url, th, isRefresh,tips="你确认删除吗？",btnText='删除') {
      layui.layer.confirm(tips, {
        btn: [btnText, '取消']
      }, function () {
        $.ajax({
          type: "DELETE",
          url: url,
          data:{'_token':$('meta[name="csrf_token"]').attr("content")},
          success: function(response) {
            if(response.error) {
              layui.layer.close();
              layui.layer.msg(response.message, {time: 2000, icon: 5})
            } else {
              layui.layer.close();
              layui.layer.msg(response.message, {time: 2000, icon: 6})
              setTimeout(function() {
                if (isRefresh) {
                  window.location = window.location.href
                  return false;
                }
                $(th).parent().parent().parent().remove();
              }, 1000);
            }
          },
        });
      }, function () {
        layui.layer.close();
      });
    };
    Admin.prototype.openLayerForm = function (url, title, method, width, height, noRefresh, formId) {
      var formId = formId ? formId : "#layer-form";
      $.get(url, function(view) {
        layui.layer.open({
          type: 1,
          title: title,
          anim: 2,
          shadeClose: true,
          content: view,
          success: function() {
            layui.form.render();
          },
          area:[
            width ? width : '50%',
            height ? height : '500px'
          ],
          btn: ['确认', '重置'],
          yes: function (index, layero) {
            var formObj = $(formId);
            $.ajax({
              type: method ? method : 'POST',
              url: formObj.attr("action"),
              dataType: "json",
              data: formObj.serialize(),
              success: function(response) {
                if (response.status === 'success') {
                  layui.layer.close(index);
                  layui.layer.msg(response.message, {time: 2000, icon: 6})
                  if (!noRefresh) {
                    window.location = window.location.href
                  }
                } else {
                  layui.layer.msg(response.message, {time: 3000, icon: 5})
                }
              },
            });
          },
          btn2: function (index, layero) {
            $(formId)[0].reset();
            return false;
          }
        });
      });
    };
    /** 
     * 时间参数
     * elem: DOM元素
     * type: 年月日类型
     * value: 初始默认值
     * min：最小值
     * max：最大值
     * is：单选/多选
     * format：自定义格式
    */
   Admin.prototype.dateTime = function(elem,type,value,min,max,is = false,format = ''){
      layui.laydate.render({ 
        elem: elem
        ,type: type
        ,value: value
        ,min: min ? min : '1900-1-1 00:00:00'
        ,max: max ? max : '2099-12-31 23:59:59'
        ,range: is
        ,format: format ? format : 'yyyy-MM-dd'
        ,done: function(value, date, endDate) {
          if(date.year == undefined){
            $('input[name="started_at"]').val('');
            $('input[name="ended_at"]').val('');
            return
          }
            date.month >=10 ? date.month : date.month = '0'+ date.month;
            date.date >= 10 ? date.date : date.date = '0'+ date.date;
            date.hours >= 10 ? date.hours : date.hours = '0'+ date.hours;
            date.minutes >= 10 ? date.minutes : date.minutes = '0'+ date.minutes;
            date.seconds >= 10 ? date.seconds : date.seconds = '0'+ date.seconds;
            endDate.month >= 10 ? endDate.month : endDate.month = '0'+ endDate.month;
            endDate.date >=10 ? endDate.date : endDate.date = '0'+ endDate.date;
            endDate.hours >= 10 ? endDate.hours : endDate.hours = '0'+ endDate.hours;
            endDate.minutes >= 10 ? endDate.minutes : endDate.minutes = '0'+ endDate.minutes;
            endDate.seconds >= 10 ? endDate.seconds : endDate.seconds = '0'+ endDate.seconds;
            if(type == 'date' && is){
              $('input[name="started_at"]').val(date.year+'-'+date.month+'-'+date.date);
              $('input[name="ended_at"]').val(endDate.year+'-'+endDate.month+'-'+endDate.date);
            }else if(type == 'datetime' && is){
              $('input[name="started_at"]').val(date.year+'-'+date.month+'-'+date.date+' '+date.hours+':'+date.minutes+':'+date.seconds);
              $('input[name="ended_at"]').val(endDate.year+'-'+endDate.month+'-'+endDate.date+' '+endDate.hours+':'+endDate.minutes+':'+endDate.seconds);
            }
        }
      })
    }
    Admin.prototype.popupContent = function(type,title,content,width,height,
      successCallback = function(index,layera) {},
      yesCallback = function(index,layera) {},
      btnCallback = function(index) {}){
      layer.open({
        type: type,
        title:title
        ,content:content
        ,area:[width + 'px', height + 'px']
        ,btn: ['确认', '取消']
        ,success: function(index,layero){
            successCallback(index,layero)
        },yes :function(index,layero){
            yesCallback(index,layero)
        },btn2:function(index){
            btnCallback(index)
          }
      });
    }
    Admin.prototype.confirmPopup = function(content,url,width,height,yes = function(){},cancel = function(){}){
       layer.confirm(content, {
        btn: ['确认','取消'],
        data:{'_xsrf':$('meta[name="csrf_token"]').attr("content")},
        area: [width + 'px',height + 'px']
      }, function(){
        if(url == '' || url == undefined){
          yes();
        }else{
          $.ajax({
              type: 'post',
              dataType: "json",
              data:{'_xsrf':$('meta[name="csrf_token"]').attr("content"), '_method': 'DELETE'},
              url: url,
              success: function (data) {
                if (data.status) {
                    message.success(data.message);
                    var index = parent.layer.getFrameIndex(window.name); //获取窗口索引
                    setTimeout(function() {
                        if (data.url) {
                            layer.close(index);//关闭弹出的子页面窗口
                            location = data.url;
                        } else
                          layui.layer.close();
                          location.reload();
                    }, 1000);
                } else {
                    message.error(data.message);
                }
              }
          });
        }
      }, function(){
        cancel()
      });
    }
    Admin.prototype.xmSelect = function(el,filter,type,data = [],name,value = [],tips = '请选择',callback=function(data){}){
      xmSelect.render({
          el: el,  
	        height: '150px',
          filterable: filter,
        	radio: type,
          data: data,
          name: name,
          initValue: value,
          clickClose: type,
          autoRow:true,
          tips: tips,
          on: function(getValue){
            callback(getValue.arr)
          }
      });
    }
    Admin.prototype.options = function (){
        return options = {
          beforeSerialize: function() {
              // $(':submit').attr('disabled', true); 
          },
          success: function(data) {
              var index = parent.layer.getFrameIndex(window.name); //获取窗口索引
              $(':submit').attr('disabled', false);
              if (data.error == undefined) {
                  message.success(data.message);
                  setTimeout(function() {
                      if (data.url) {
                          parent.layer.close(index); //关闭弹出的子页面窗口
                          if (index == undefined) {
                              window.location = data.url;
                              return false;
                          } else
                              parent.location = data.url;
                      } else
                        parent.layer.close(index);
                        parent.location.reload();
                  }, 1000);
              } else {
                $(':submit').attr('disabled', false);
                message.error(data.message);
              }
          }
        }
    };
      // 限制输入框小数点后几位 
    /** 
     * admin.numLength()* 
     * obj 当前文本框DOM元素
     * type 0 只能输入整数 1 只能输入一位小数 2 两位小数
     */

    Admin.prototype.numLength = function(obj,type){
      switch(type){
        case 'just': obj.value = obj.value == 0 ? 1 : Math.ceil(Math.abs(obj.value)); break;
        case 0: obj.value = obj.value.replace(/[^-|\d*]/,''); break;
        case 1: obj.value = obj.value.replace(/^(\-)*(\d+)\.(\d).*$/,'$1$2.$3'); break;
        case 2: obj.value = obj.value == 0 ? 1 : Math.abs(obj.value.replace(/^(\-)*(\d+)\.(\d\d).*$/,'$1$2.$3'));
      }
    }
    /**
     * 上传图片公共
     * elem 按钮id
     * box 最后填充地址
     */
    Admin.prototype.uploadImgBox = function (elem, box, url){
      layui.upload.render({
          elem: elem
          , url: url
          , multiple: true
          , field: 'files'
          , accept: 'images'
          , size: "5120"
          , auto: true
          , data: {'_xsrf':$('meta[name=csrf_token]').attr('content')}
          , done: function (res) {
              let name = this.name;
              let num;
              this.num == undefined ? num = 5 : num = this.num;
              let length = $(box).find('li').length;
              //如果上传失败
              if (res.status == true){
                  if(length < num){
                      $(box).append('<li><img src="' + res.data.url + '" /><input type="hidden" name="'+name+'" value="'+ res.data.url + '"><span hidden="" class="img-delete" style="display: none;"><i class="icon-shanchu iconfont"></i></span></li>')
                      return layer.msg(res.message);
                  }
                  return layer.msg('最多可上传'+num+'张图片')
              }
              return layer.msg(res.message);
          }
      });
    };
    Admin.prototype.errorList = function (data){
      var str = "";
        data.data.error.forEach((item,index) =>{
          if(item.name2 != undefined){
            str += `<p class="C-marginLeft-10">${index + 1}.<span>${item.name2}</span><span class="G-color-red">${item.name}</span>${item.message}</p>`
          }else{
            str += `<p class="C-marginLeft-10">${index + 1}.<span class="G-color-red">${item.name}</span>${item.message}</p>`
          }
        })
        let content = `
          <div  class="layui-form-item C-padding-15">
            <p class="C-marginTop-20">我们接收到您一共添加了 <span class="G-color-red"> ${data.total} </span> 条数据，成功 <span class="G-color-red"> ${data.success_num} </span>条，失败<span class="G-color-red">${data.error_num}</span>条。</p>
            <div ${data.data.error.length == 0 ? 'hidden': ""}>
              <h1 class="C-marginTop-20 C-marginBottom-15">失败信息</h1>
              ${ str}
            </div>
          </div>
        `
      layer.open({
        type: 1,
        title:'提示信息'
        ,content:content
        ,closeBtn :0
        ,area:['600px', '400px']
        ,btn: ['关闭']
        ,yes :function(index,layero){
          if(data.url){
            parent.layer.close(index);
            parent.location.reload();
            // 这里是品类管理品牌 关闭后 要新建一个tabs页面展示型号列表
            if($(".class-A").data('id')!=""){
              $(".create-model").click();
            }
          }else{
            parent.layer.close(index);
            parent.location.reload();
          }
        }
      });
    }
        /**  
     * 下拉分页加载
     * elem : 大盒子
     * scrollElem：可以滚动的DOM
     * url: 请求的地址
     * colspan  当前表格有多少列，因为下面的tips要占一整行，所以要合并
     * tips:自定义下拉到底部文字
    */
   
   Admin.prototype.flowLoad = function(elem,scrollElem,url,colspan,tips="木有啦"){
    layui.flow.load({
        elem: elem //指定列表容器
        , scrollElem: scrollElem
        , mb: 100          //距离底端多少像素触发auto加载
        , end: '<td colspan='+colspan+' style="color:red">'+tips+'</td>'    //加载所有后显示文本，默认'没有更多了'
        ,done: function(page, next){ //到达临界点（默认滚动触发），触发下一页
          if(page == 1){next();return;
          }
          let href = location.href;
          let ajaxUrl = ''
          if(href.split("?")[1] != undefined){
            ajaxUrl = href+"&page="+ page + "&alert=1"
          }else{
            ajaxUrl = href+"?page="+ page + "&alert=1"
          }
            $.get( ajaxUrl, function(res){
                next(res, $(elem).find('tr[class!="layui-flow-more"]').length == (page -1) * 20);                
                let index = parent.layer.getFrameIndex(window.name)
                admin.scrollCheck(index,'child');
            });
        }
    });
  }
        /** 
     *  请求的时候 分页下拉是单独的页面，而按钮在创建页面，是两个页面，
     *  所以这里就要判断是哪个页面要请求进分页页面选中，body 是 父页面请求子弹窗的DOM
     *  parent.window.$("***") 是子弹窗获取父页面的DOM元素。
    */
   Admin.prototype.scrollCheck = function(index,type = "parent"){
    if(type =="parent"){
      var body = layer.getChildFrame('body', index);
      for(var i = 0; i < body.find('tbody .js_table_id').length; i++){
          let iframeid = body.find('tbody .js_table_id').eq(i).find('input').val();
          for(var j = 0; j < $("#table-admin").find("tr").length;j++){
              let bodyid = $("#table-admin").find("tr").eq(j).find('input[type="hidden"]').val();
              if(bodyid == iframeid){
                  body.find("tbody input[type='checkbox']").eq(i).prop({'checked':true,'disabled':'disabled'});
              }
          }
      }
    }else{
      parent.window.$("#table-admin").find('tr').length
      for(var i = 0; i < $('#demo .js_table_id').length; i++){
          let iframeid = $('#demo .js_table_id').eq(i).find('input').val();
          for(var j = 0; j < parent.window.$("#table-admin").find("tr").length;j++){
              let bodyid = parent.window.$("#table-admin").find("tr").eq(j).find('input[type="hidden"]').val();
              if(bodyid == iframeid){
                  $("#demo input[type='checkbox']").eq(i).prop({'checked':true,'disabled':'disabled'});
              }
          }
      }
    }
  }
  Admin.prototype.mergeTale = function (res,columsName,columsIndex){
    var data = res.data;
    var mergeIndex = 0;
    var mark = 1;
    var columsName = columsName;
    var columsIndex = columsIndex;
    var trArr;
    let arr = [];//记录合并行的行ID
    for (var k = 0; k < columsName.length; k++){
      trArr = $(".layui-table-body>.layui-table").find("tr");
      for (var i = 1; i < res.data.length; i++) {
        var tdCurArr = trArr.eq(i).find("td").eq(columsIndex[k]);
        var tdPreArr = trArr.eq(mergeIndex).find("td").eq(columsIndex[k]);
        arr.push(mergeIndex);
        trArr.eq(i-1).attr("data-key",mergeIndex)
        if (data[i - 1][columsName[k]] === data[i][columsName[k]]) {
          trArr.eq(i).attr("data-key",mergeIndex)
          mark += 1;
          tdPreArr.each(function () {
            $(this).attr("rowspan", mark);
          });
          tdCurArr.each(function () {
            $(this).css('display','none')
          });
        }else {
          mergeIndex = i;
          mark = 1;
          trArr.eq(i).attr("data-key",mergeIndex)
        }
        arr.push(mergeIndex);
      }

      //以下是给表格添加隔行背景色逻辑
      let newArr = [...new Set(arr)];
      let odd=[],even=[];//记录偶数奇数行，便于做背景色赋值
      newArr.map(function (item,index) {
        if(index%2){
          odd.push(item)
        }else{
          even.push(item)
        }
      });
      trArr.each((index,el)=>{
        if(!!~odd.indexOf(Number($(el).attr('data-key')))){
          $(el).css("background-color", "#f2f2f2")
        }
      })
    } 
  };
  Admin.prototype.deleteShow = function (){
    if($("#table-admin tr").length){
        $(".deleteShow").show();
    }else{
        $(".deleteShow").hide();
    }
  };
  Admin.prototype.getLast3Month = function (normal = 1){
    var now = new Date();
    var year = now.getFullYear();
    var month = now.getMonth() + 1;//0-11表示1-12月
    var day = now.getDate();
    var dateObj = {};
    if (parseInt(month) < 10) {
        month = "0" + month;
    }
    if (parseInt(day) < 10) {
        day = "0" + day;
    }
    dateObj.now = year + '-' + month + '-' + day;
    if (parseInt(month) == 1) {//如果是1月份，则取上一年的10月份
        dateObj.last = (parseInt(year) - 1) + '-10-' + day;
        return dateObj;
    }
    if (parseInt(month) == 2) {//如果是2月份，则取上一年的11月份
        dateObj.last = (parseInt(year) - 1) + '-11-' + day;
        return dateObj;
    }
    if (parseInt(month) == 3) {//如果是3月份，则取上一年的12月份
        dateObj.last = (parseInt(year) - 1) + '-12-' + day;
        return dateObj;
    }
    var preSize = new Date(year, parseInt(month) - normal, 0).getDate();//开始时间所在月的总天数
    if (preSize < parseInt(day)) {
　　　// 开始时间所在月的总天数<本月总天数，比如当前是5月30日，在2月中没有30，则取下个月的第一天(3月1日)为开始时间
        let resultMonth = parseInt(month) - 2 < 10 ? ('0' + parseInt(month) - 2) : (parseInt(month) - 2);  
        dateObj.last = year + '-' + resultMonth + '-01';  
        return dateObj;
    }
    if (parseInt(month) <= 10) {
        dateObj.last = year + '-0' + (parseInt(month) - normal) + '-' + day;
        return dateObj;
    } else {
        dateObj.last = year + '-' + (parseInt(month) - normal) + '-' + day;
        return dateObj;
    }
  }  
    window.admin = new Admin();
  }(window);
