layui.define(["layer","jquery","form"],function(exports){
    var $ = layui.$
    ,admin = layui.admin
    ,table = layui.table
    ,element = layui.element
    ,layer = layui.layer
    ,laydate = layui.laydate
    ,form = layui.form
    ,square_id = $(".square_main").data('id');
    var region = {};
    region.getRegion = function getRegion(parent_id, ele) {
        var region_id = arguments[2] ? arguments[2] : '';
        $.get({
            url: "/region/server",
            data: { parent_id: parent_id },
            success: function(data) {
                var html = '';
                for(let key  in data.regions){
                    html += '<option value="' + data.regions[key].id + '"';
                    if (region_id == data.regions[key].id) html += ' selected="selected"';
                    html += '>' + data.regions[key].name + '</option>';
                }
                $("." + ele).append(html);
                form.render('select');
                if(region_id == ""){
                    if(ele != "province"){
                        $("." + ele).siblings('.layui-form-select').addClass('layui-form-selected')
                    }
                }
            },
            error: function(xOptions, textStatus) {
                console.log(xOptions);
            }
        });
    }
    var province = $(".province").data("id");
    var city = $(".city").data("id");
    var district = $(".district").data("id")
    if (province == null || province == '') {
        region.getRegion(1, 'province', province);
    } else if (city == null || city == '') {
        region.getRegion(1, 'province', province);
        region.getRegion(province, 'city', city);
    } else {
        region.getRegion(1, 'province', province);
        region.getRegion(province, 'city', city);
        region.getRegion(city, 'district', district);
        aJaxSquare(city,square_id)
    }
    form.on('select(province)',function(data){
        var value = data.value;
        $(".city").html("<option value=''>选择市</option>");
        $(".district").html("<option value=''>选择区</option>");
        $(".square_main").html("<option value=''>选择商圈</option>");
        $('.latitude,.longitude').attr('value', '');
        $('.panel_lat_degree').val('');
        if(value != '' && value != null) {
            region.getRegion(value, 'city');
        } else {
            form.render('select');
        }
    });
    form.on('select(city)',function(data){
        var value = data.value;
        $(".district").html("<option value=''>选择区</option>");
        $(".square_main").html("<option value=''>选择商圈</option>");
        $('.latitude,.longitude').attr('value', '');
        $('.panel_lat_degree').val('');
        if(value != '' && value != null) {
            region.getRegion(value, 'district');
            aJaxSquare(value,'')
        } else {
            form.render('select');
        }
    });
    form.on('select(district)',function(data){
        $('.latitude,.longitude').attr('value', '');
        $('.panel_lat_degree').val('');
        var value = data.value;
    });
    function aJaxSquare(value,id){
        $.get({
            url: "/square/server",
            data: { city_id: value },
            success: function(data) {
                $(".square_main").empty();
                var html = '<option value="">选择商圈</option>';
                for(let key  in data.squares){
                    html += '<option value="' + data.squares[key].id + '"';
                    if (id == data.squares[key].id) html += ' selected="selected"';
                    html += '>' + data.squares[key].name + '</option>';
                }
                $(".square_main").append(html);
                form.render('select');
            },
            error: function(xOptions, textStatus) {
                console.log(xOptions);
            }
        });
    }
    exports('region',{});
});
