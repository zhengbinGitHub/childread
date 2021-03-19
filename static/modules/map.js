
layui.define(["layer","jquery","message"],function(exports){
    var $ = layui.jquery
    ,message = layui.message;
    var map = null;
    var searchService = null;
    var serchType = 2;
    var lat = Number($('input[name="data[lat]"]').val() || 39.916527);
    var lng = Number($('input[name="data[lng]"]').val() || 116.397128);
    var init = function() {
        var myLatlng = new qq.maps.LatLng(lat, lng);
        var myOptions = {
            zoom: 14,
            center: myLatlng,
            disableDoubleClickZoom: true,
            draggable: true,
            scaleControl: true,
            panControl: true,
            scrollwheel: true,
            keyboardShortcuts: true,
            mapTypeControl: true,
            zoomControl: true
        };
        map = new qq.maps.Map(document.getElementById("container"), myOptions);
        var centerPoint = '<div class="centerPoint"><img src="/assets/static/images/map.png" alt="" /><div>'
        $('#container').append(centerPoint)
        //当地图中心属性更改时触发事件
        qq.maps.event.addListener(map, 'center_changed', function() {
            if (serchType == 1) {
                $('#panel_lat').html(map.getCenter().lat);
                $('#panel_lng').html(map.getCenter().lng);

            }
        });
        qq.maps.event.addListener(map, 'click', function(event) {
            marker.setPosition(new qq.maps.LatLng(event.latLng.getLat(), event.latLng.getLng()));
            $('#panel_lat').html(event.latLng.getLat());
            $('#panel_lng').html(event.latLng.getLng());
        });
        //调用Poi检索类
        searchService = new qq.maps.SearchService({
            panel: document.getElementById('infoDiv'),
            //检索成功的回调函数
            complete: function(results) {
                if (!results.detail.pois) {
                    $('#infoDiv').css({'display':'none'});
                    message.error('搜索结果为空，请选择标记新地址');
                    return
                }
                $(document).on('click', '#infoDiv li', function() {
                    var num = $(this).index();
                    $('#panel_lat').html(results.detail.pois[num].latLng.lat);
                    $('#panel_lng').html(results.detail.pois[num].latLng.lng);
                })
            },
            map: map
        });
        //编辑显示定位图标
        $(function() {
            if ($('input[name="data[lng]"]').val() && $('input[name="data[lng]"]').val()) {
                $('.centerPoint').show();
            }
        })
    }
    init();

    function getResult(type) {
        if ($('.province').val() == '' || $('.city').val() == '' || $('.district').val() == '') {
            message.error('请选择省市区');
            return;
        }
        serchType = 2
        var poiText = $("#poiText").val();
        if (poiText.length == 0) {
            message.error('请输入地址信息');
            return;
        }
        $('#panel_lat').html('');
        $('#panel_lng').html('');
        layer.open({
            type: 1,
            title: '选择经纬度',
            shadeClose: false,
            maxmin: true, //开启最大化最小化按钮
            area: ['930px', '565px'],
            content: $(".map_layer").load(),
            cancel: function(index) {
                layer.close(index)
            },
            yes: function(index) {
                layer.close(index)
            }
        });
        $('#infoDiv').show();
        $('.centerPoint').hide();
        var regionText = $(".province option:selected").text().trim();
        searchService.setLocation(regionText);
        searchService.search(poiText);
        searchService.setPageCapacity(4);
        searchService.setError(function() {
            message.error('搜索结果为空，请选择标记新地址');
        })
    }
    //点击设置坐标---出现layer弹窗-地图
    $('body').on('click', '#seach_panel', function() {
        getResult();
    });
    //保存经纬度
    $('body').on('click', '.J_save_lat_lng', function() {
        if(($('#panel_lat').html() == "" || $('#panel_lng').html() == "") && $(".latitude").val() != ""){
            message.closeAll();
            return false;
        }
        changeLat({ "lat": $('#panel_lat').html(), "lng": $('#panel_lng').html() });
        message.closeAll();
    });
    // 关闭
    $('body').on('click','#layui-map-close',function(){
        message.closeAll(); //再执行关闭
    });

    function changeLat(latlng) {
        var lat = latlng.lat;
        var lng = latlng.lng;
        $('.latitude').attr('value', lat);
        $('.longitude').attr('value', lng);
        $('.panel_lat_degree').val(lng + "," + lat);
    }
    $("#changeSearchType").on('click', function() {
        changeType()
    })

    function changeType(events) {
        serchType = 1
        $('#infoDiv').hide();
        $('.centerPoint').show();
        searchService.clear()
    }
    exports('map', {});    
});  