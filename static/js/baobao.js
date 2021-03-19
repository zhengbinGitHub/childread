if(typeof(Page) == "undefined" || typeof(Page) == "object"){
    Page = "0"
}

$(function($) {
    // 公用部分JS

    tabase();//选项卡

    //返回顶部结束
    if(Page != "index" ){

        navhover();//当前位置
    }
    // 首页用的
    if(Page == "index" )
    {
        gotop()
        console.log(Page);
        change()//
        Tabflash();//首页幻灯
        inulmun();//首页处理图片不一样的效果
        hovertop();//改变样式


        gundong()//右侧专有，滑动，高亮
        hoverp();//鼠标经过
        indextags()//tags

    }
    //列表页
    if(Page == "list" ){
        hovertop();//改变样式
        gotop();
        listags()


    }
    if(Page == "zhuanti" ){
        hovertop();//改变样式
        // 处理排序
        $(".f-fzhot ul li").each(function(i){
            $(this).find("a em").text(i+1)
        })

    }
    if(Page == "zt" ){//专题页
        gotop()
        // 处理排序
        // $(".f-fzhot ul li").each(function(i){
        //   $(this).find("a em").text(i+1)
        // })
        //增加分享
        $(".m-fxdiv").append('<div class="bdsharebuttonbox"><a href="#" class="bds_qzone" data-cmd="qzone" title="分享到QQ空间"></a><a href="#" class="bds_tsina" data-cmd="tsina" title="分享到新浪微博"></a><a href="#" class="bds_sqq" data-cmd="sqq" title="分享到QQ好友"></a><a href="#" class="bds_weixin" data-cmd="weixin" title="分享到微信"></a><a href="#" class="bds_copy" data-cmd="copy" title="分享到复制网址"></a><a href="#" class="bds_more" data-cmd="more"></a></div>');
        window._bd_share_config={"common":{"bdSnsKey":{},"bdText":"","bdMini":"2","bdMiniList":false,"bdPic":"","bdStyle":"1","bdSize":"32"},"share":{}};with(document)0[(getElementsByTagName('head')[0]||body).appendChild(createElement('script')).src='http://bdimg.share.baidu.com/static/api/js/share.js?v=89860593.js?cdnversion='+~(-new Date()/36e5)];
        // tags随机
        var tags_a = $(".m-tagzt a");
        tags_a.each(function(){
            var x = 4;
            var y = 0;
            var rand = parseInt(Math.random() * (x - y + 1) + y);
            $(this).addClass("tags"+rand);
        });

    }
    //内容页
    if(Page == "main" ){
        hovertop();//改变样式
        html_qnav()//H3作为标题处理
        h3tit();//H3增加样式
        //专有JS
        //右侧固顶
        // var righttop  = $(".m-hot").offset().top;//介绍顶部距离
        // $(window).scroll(function(){//滚动监听
        //   //判断高度
        //   if($(window).scrollTop() > righttop){
        //     $(".m-hot").addClass("f-top");
        //   }else{
        //       $(".m-hot").removeClass("f-top");
        //   }
        // });

        //增加分享
        $(".m-fxdiv").append('<div class="bdsharebuttonbox"><ul><li class="clearfix"><a href="#" class="bds_qzone" data-cmd="qzone" title="分享到QQ空间">QQ空间</a></li><li class="clearfix"><a href="#" class="bds_tsina" data-cmd="tsina" title="分享到新浪微博">新浪微博</a></li><li class="clearfix"><a href="#" class="bds_sqq" data-cmd="sqq" title="分享到QQ好友">QQ好友</a></li><li class="clearfix"><a href="#" class="bds_weixin" data-cmd="weixin" title="分享到微信">微信好友</a></li><li class="clearfix"><a href="#" class="bds_copy" data-cmd="copy" title="分享到复制网址">复制网址</a></li><li class="clearfix"><a href="#" class="bds_more" data-cmd="more">更多分享</a></li></div>');
        window._bd_share_config={"common":{"bdSnsKey":{},"bdText":"","bdMini":"2","bdMiniList":false,"bdPic":"","bdStyle":"1","bdSize":"32"},"share":{}};with(document)0[(getElementsByTagName('head')[0]||body).appendChild(createElement('script')).src='http://bdimg.share.baidu.com/static/api/js/share.js?v=89860593.js?cdnversion='+~(-new Date()/36e5)];
        // 属性下方，没有标题隐藏

        if($(".m-tags a").length<1){
            $(".m-tags p ").hide();
        }
        if($(".m-tuwendiv li").length<1){
            // $(".m-tuwendiv").hide();
            // $(".m-tuwendiv").prev().hide();
            $(".m-tuwendiv ul").css("padding-left","30px");
        }
        // 增加下一页
        var xiayiyeurl = $(".m-navwz a:last").attr("href")
        var xiayiyetxt = '<div id="cms_showpage_text"><b>1</b><a href="'+xiayiyeurl+'">2</a><a href="'+xiayiyeurl+'" class="m-xiayiye">下一页<i class="f-jkico"></i></a></div>';
        $(".m-fenye").append(xiayiyetxt);
        // 增加导读
        var OneP = $(".htmlcontent p").eq(0).text();
        $(".m-mianjianjie div").append(OneP);
        //拉动置顶
        var leftop = $(".m-mianfx").offset().top;
        var lefthtight = $(".m-mianfx").height();
        var height = window.innerHeight
        var pltop = $(".m-xg-ne").offset().top - lefthtight -100;
        $(window).scroll(function(){
            if($(window).scrollTop() > leftop ){
                $(".m-mianfx").addClass("ltop");
            }if($(window).scrollTop() < leftop || $(window).scrollTop() > pltop){
                $(".m-mianfx").removeClass("ltop");
            }
        });

        var navn = $(".m-nav1 ul li.hover").index();
        $(".m-nan2 div").hide().eq(navn).show()
        console.log(navn)
        $(".m-xgwz-item li:lt(2)").addClass("top");
        $(".m-xgwz-item li").eq(0).find("p").append("<span>浏览:8653</span>");
        $(".m-xgwz-item li").eq(1).find("p").append("<span>浏览:1万+</span>");
        // $(window).scroll(function(){//左侧滚动监听
        //   var rightTop = $(".m-mianfx").offset().top;
        //   var leftTop = $(".m-maintxt").offset().top;
        //   var leftHeight =  $(".m-maintxt").height();
        //   var rightHeight = $(".m-mianfx").height();
        //   var winHeight = $(window).height();
        //   var scrollTop = rightHeight+rightTop-winHeight;
        //   var gameTabtop = $(".m-pl").offset().top - winHeight;
        //   var rightMargintop = leftHeight - rightHeight;
        //   var stopFloat = leftTop+rightHeight-winHeight;

        //   console.log($(window).scrollTop()+"========"+scrollTop+"=="+gameTabtop+"==="+stopFloat)
        //   if($(window).scrollTop() >= scrollTop){
        //     $(".m-mianfx").addClass("ltop")
        //   }
        //   if($(window).scrollTop() > gameTabtop){
        //     $(".m-mianfx").removeClass("ltop");
        //     $(".m-mianfx").css({"margin-top":rightMargintop})
        //   }else{
        //     $(".m-mianfx").addClass("ltop")
        //     $(".m-mianfx").css({"margin-top":0})
        //   }
        //   if($(window).scrollTop() < stopFloat){
        //     $(".m-mianfx").removeClass("ltop");
        //     $(".m-mianfx").css({"margin-top":0})
        //   }
        // })

    }
});


function imgup(){//背景滑动JS
    $(".f-thumb li").each(function() {
        var aheight = $(this).find("a").height()
        $(this).find(".thumb-txt").height(aheight).css("top",aheight)

        $(this).hover(function() {

                $(this).find(".thumb-txt").stop().animate({
                    top: 0
                },300);
            },
            function() {
                $(this).find(".thumb-txt").stop().animate({
                    top: aheight
                },300);
            });
    });

}
function hovertop(){//鼠标经过改变样式,第一个单独加样式，前三单独加样式
    $('.f-top li').mouseover(function(){
        $(this).siblings().removeClass('f-ix').end().addClass('f-ix');
    });
    $('.f-top3').each(function() {
        $(this).find('li:first').addClass('f-ix').end()
            .find('li:lt(3)').addClass('f-t3');

    })
}
function hoverp(){//鼠标经增加样式离开去掉用于显示隐藏

    $(".f-tw,.f-hov").mouseover(function(){
        $(this).addClass("f-hover");
    }).mouseout(function(){
        $(this).removeClass("f-hover");
    });
}
function gotop(){//返回顶部
    $("body").append('<div class="g-float"><a href="javascript:;" id="m-top-back"></a></div>')
    // 滚动监听显示隐藏并且返回顶部
    $(window).scroll(function(){
        if($(window).scrollTop()>200){
            $("#m-top-back").animate({height:60},10);
        }else{
            $("#m-top-back").animate({height:0},10);
        }
    })
    $("#m-top-back").click(function(){
        $("body,html").animate({scrollTop:0},300)
    })
}
function navhover(){
    var rootName = $(".m-navwz a").eq(1).text().replace(/\s+/g,'');
    $(".m-nav1 ul li a").each(function(){
        var navName = $(this).text().replace(/\s+/g,'');
        if(navName==rootName){
            $(".m-nav1 ul li ").removeClass("hover f-hover");
            $(this).parent("li").addClass("hover f-hover");
        }
    })
}
function tabase(){//通用选项卡
    if($(".f-hovertab-box").length>0){
        $(".f-hovertab-box").each(function(){
            $(this).find(".f-hovertab-btn").children().first().addClass("f-hover");
            $(this).find(".f-hovertab-cont:gt(0)").hide();
            if($(this).find(".f-hovertab-cont").length<=0){
                $(this).hide();
            }
        })
        $(".f-hovertab-btn").children().hover(function(){
            var tabSpeed = $(this).parents(".f-hovertab-box").attr("data-speed");
            var thisObj = $(this);
            setTimer = setTimeout(function(){
                thisObj.addClass("f-hover").siblings().removeClass("f-hover");
                var n = thisObj.index();
                thisObj.parents(".f-hovertab-box").find(".f-hovertab-cont").eq(n).show().siblings(".f-hovertab-cont").hide();
            },tabSpeed)
        },function(){
            clearTimeout(setTimer);
        })
    }
}
function inulmun() {
    $("m-hotul li").each(function(){
        $(this).find()

    })


}
function Tabflash(){//首页幻灯切换

    var oSj = 5000;
    var i = 0;
    var bar = $(".f-flash .bar");
    var oImg = $(".f-flash .f-flashul");//获取图片盒子
    var oImgfirst= $('.f-flash .f-flashul li:first').clone();//复制第一张图片
    oImg.append(oImgfirst);//将复制的第一张图片放到最后
    var imgNum = $(".f-flash .f-flashul li").size();//获取图片数量

    //根据图片个数添加圆点按钮
    for(var j=1;j<=imgNum-1;j++){
        $('.f-flash .f-lidian').append('<li></li>');
    }

    //给第一个按钮添加选中样式
    $('.f-flash .f-lidian li:first').addClass('index');


    //点击向右轮播
    $(".but-right").click(function(){
        int();
    });

    //点击向左轮播
    $(".but-left").click(function(){
        bar.stop().css('width',0);
        i--;
        if(i == -1){
            $('.f-flash .f-flashul').css('left',-(imgNum-1)*845);//用CSS进行图片位置变换，达到无缝拼接效果
            i = imgNum-2;
        }
        oImg.stop().animate({left:-i*845},500);//动画效果
        clearInterval(oTime);
        oTime = setInterval(function(){
            int();
        },oSj);
        barAniMate();//进度条函数动画效果
        $(".f-flash .f-lidian li").eq(i).addClass('index').siblings().removeClass('index');//给相应的按钮添加样式
    });

    //鼠标移动到圆点后轮播
    $(".f-flash .f-lidian li").hover(function() {
        clearInterval(oTime);//清除定时器
        bar.stop().css('width',0);
        var index = $(this).index();
        i=index;
        oImg.stop().animate({left:-index*845},500);//动画效果
        bar.stop().css('width',0);
        $(this).addClass('index').siblings().removeClass('index');//给相应的按钮添加样式
    },function(){
        barAniMate();//进度条函数动画效果
        oTime = setInterval(function(){
            int();
        },oSj);
    });

    //自动轮播
    var oTime = setInterval(function(){
        int();
    },oSj);

    barAniMate();//进度条函数动画效果

    //进度条函数动画效果
    function barAniMate(){
        bar.animate({width:'100%'},oSj,function(){
            $(this).css('width',0);
        });
    }

    //自动轮播函数
    function int(){
        bar.stop().css('width',0);
        i++;
        if(i == imgNum){
            oImg.css('left',0);//用CSS进行图片位置变换，达到无缝拼接效果
            i = 1;
        }
        oImg.stop().animate({left:-i*845},500);//动画效果
        barAniMate();//进度条函数动画效果
        clearInterval(oTime);
        oTime = setInterval(function(){
            int();
        },oSj);
        if(i == imgNum-1){
            $(".f-flash .f-lidian li").eq(0).addClass('index').siblings().removeClass('index');//给相应的按钮添加样式

        }else{
            $(".f-flash .f-lidian li").eq(i).addClass('index').siblings().removeClass('index');//给相应的按钮添加样式
        }
    }
}

function gundong(){//无插件实现滚动监听高亮
    var topa = $(".g-intop").offset().top;
    var topb = $(".g-inhot").offset().top;
    var topc = $(".g-beiyun").offset().top;
    var topd = $(".g-chanhou").offset().top;
    var tope = $(".g-zaojiao").offset().top;
    var topf = $(".g-youlian").offset().top;
    $(".g-rightnav li").click(function(){

        $(this).removeClass();

    })
    $(".g-rightnav li:eq(0)").click(function(){$("html,body").animate({scrollTop:$('.g-intop').position().top-5}, 500);   });
    $(".g-rightnav li:eq(1)").click(function(){$("html,body").animate({scrollTop:$('.g-inhot').position().top-5}, 500);});
    $(".g-rightnav li:eq(2)").click(function(){$("html,body").animate({scrollTop:$('.g-beiyun').position().top-5}, 500);});
    $(".g-rightnav li:eq(3)").click(function(){$("html,body").animate({scrollTop:$('.g-chanhou').position().top-5}, 500);});
    $(".g-rightnav li:eq(4)").click(function(){$("html,body").animate({scrollTop:$('.g-zaojiao').position().top-5}, 500);});


    $(window).scroll(function(){//滚动监听
        if( $(window).scrollTop() < topa){
            $(".g-rightnav li:eq(0)").addClass("f-hover").siblings().removeClass();
        }
        if(topa < $(window).scrollTop() && $(window).scrollTop() < topb){
            $(".g-rightnav li:eq(1)").addClass("f-hover").siblings().removeClass();
        }
        if(topb < $(window).scrollTop() && $(window).scrollTop() < topc){
            $(".g-rightnav li:eq(2)").addClass("f-hover").siblings().removeClass();
        }
        if(topc < $(window).scrollTop() && $(window).scrollTop() < topd){
            $(".g-rightnav li:eq(3)").addClass("f-hover").siblings().removeClass();
        }
        if(topd < $(window).scrollTop() && $(window).scrollTop() < tope){
            $(".g-rightnav li:eq(4)").addClass("f-hover").siblings().removeClass();
        }
        if(tope < $(window).scrollTop() && $(window).scrollTop() < topf){
            $(".g-rightnav li:eq(5)").addClass("f-hover").siblings().removeClass();
        }
        if(topf < $(window).scrollTop() && $(window).scrollTop() < topg){
            $(".g-rightnav li:eq(6)").addClass("f-hover").siblings().removeClass();
        }
    })
}


function change(){//换一换
    var n = 0;
    $(".m-hyh a").click(function(){
        console.log("asd")
        $(this).find("img").toggleClass("gun");
        if(n == 2){
            n = 0;
            $(this).parents("div").find("ul").eq(0).show().siblings("ul").hide();
        }else{
            $(this).parents("div").find("ul").eq(n+1).show().siblings("ul").hide();
            n=n+1;
        }
        setTimeout(function(){
            $(".m-hyh a").find("img").removeClass("gun");
        },1000);
    });
}

function html_qnav() {//H3右侧处理


    if ($(".htmlcontent h3").length > 0) {
        var Qrnone = "block";
    } else {
        var Qrnone = "none";
    } // 正文内容
    var nstr = '<div id="Qright" style="display:none;">';

    nstr += '<div class="side-catalog" style="display:' + Qrnone + ';"><div class="side-bar"><em class="circle start"></em><em class="circle end"></em></div><div class="catalog-scroller"><dl class="catalog-list"><ul class="culist" Qu-top="0">'

    if ($(".fulllink").length > 0) {
        nstr += '<dt class="catalog-title level1"><em class="pointer"></em><span class="text"><a href="' + $(".fulllink").attr("href") + '" class="title-link">阅读全文</a></span></dt>';
    }

    nstr += '<dt class="catalog-title level1"><em class="pointer"></em><span class="text"><a href="javascript:void(0);" class="title-link" onclick=qnlk(".htmlcontent")>正文内容</a></span></dt>';
    var btlk = "";

    $(".htmlcontent h3,.htmlcontent h3").each(function() {
        var aaa = Qrenav($(this).text());
        var zzz = $(this).index("h3");
        if (aaa.length < 20) {
            if ($(this).attr("class") == "biaoti") {
                nstr += '<dt class="catalog-title level1"><span class="text"><a href="javascript:void(0);" class="title-link" onclick=qnlk("h3:eq(' + zzz + ')")>' + aaa + '</a></span></dt>';
            } else {
                nstr += '<dd class="catalog-title level2"><span class="text"><a href="javascript:void(0);" class="title-link" onclick=qnlk("h3:eq(' + zzz + ')")>' + aaa + '</a></span></dd>';
            }
            btlk += ',h3:eq(' + zzz + ')+' + (zzz * 1 + 1)
        }
    });
    //nstr += '<dt class="catalog-title level1"><em class="pointer"></em><span class="text"><a href="javascript:void(0);" class="title-link" onclick=qnlk(".mComment")>发表评论</a></span></dt>';
    nstr += '<a class="arrow" href="javascript:void(0);" style="top: 38px;"></a></ul></dl></div><div class="right-wrap"><a class="go-up disable"></a><a class="go-down"></a></div></div><div id="gotop"><a class="nas">0</a><a class="got">1</a><div id="aaaa"></div></div></div>';

    $(".m-mainright").append(nstr);
    qncr(".htmlcontent+0" + btlk);

    var Qwcgs = 0;
    if ($(".catalog-list span.text a").length < 10) {
        $(".right-wrap").hide();
        $(".side-bar,.side-catalog").css("height", ($(".catalog-list span.text a").length * 31 + 40) + "px");
        var Qwcgs = ((10 - $(".catalog-list span.text a").length) * 31);
    }

    var winhei = window.innerHeight - 450;
    var Qstop = $(".m-twnews").offset().top +$(".m-twnews").height() -winhei;//获取右侧到最下面的高度

    $(window).scroll(function() {

        if ($(window).scrollTop() > Qstop) {
            $("#Qright").show();
            $("#gotop").css({
                "position": "fixed",
                "bottom": "90px"
            });
            $(".side-catalog").css({
                "position": "fixed",
                "bottom": "200px"
            });
        } else {
            $("#Qright").hide();
        }

    });

    $('.go-up').click(function() {
        $('.catalog-list').animate({
                scrollTop: ($(".catalog-list").scrollTop() - 124) + 'px'
            },
            500);
    });
    $('.go-down').click(function() {
        $('.catalog-list').animate({
                scrollTop: ($(".catalog-list").scrollTop() + 124) + 'px'
            },
            500);
    });

    $(".catalog-list").scroll(function() {
        var zhd = $(".culist").height();
        var dhd = $(".catalog-list").scrollTop();
        if (dhd == 0) {
            $(".go-up").addClass("disable");
        } else if ((zhd - dhd) <= 310) {
            if (zhd > 310) {
                $(".go-up").removeClass("disable");
            }
            $(".go-down").addClass("disable");
        } else {
            $(".go-up").removeClass("disable");
            $(".go-down").removeClass("disable");
        }

    });

    $('#gotop .nas').click(function() {
        $(".side-catalog").toggle();
    });
    $('#gotop .got').click(function() {
        $('body,html').animate({
                scrollTop: '0px'
            },
            500);
    });



}



function qnlk(obj) {
    $('html,body').animate({
            scrollTop: $(obj).offset().top
        },
        100);
}

function qncr(str) {
    var Qrtophd = $(".fulllink").length > 0 ? 37 : 6;
    $(window).scroll(function() {
        var winScroll = $(window).scrollTop();
        var sts = new Array();
        sts = str.split(",");
        for (var i = 0; i < sts.length; i++) {
            if (winScroll > $(sts[i].split("+")[0]).offset().top - 10) {
                $(".arrow").css("top", sts[i].split("+")[1] * 31 + Qrtophd + "px");
                $('.catalog-list').scrollTop(sts[i].split("+")[1] * 31 - 124);
                $(".catalog-title").eq(i).addClass("on").siblings(".catalog-title").removeClass("on");
            }
        }
    });
}

function Qrenav(str) {
    str = str.replace("：", "");
    str = str.replace(":", "");
    return str;
}if(typeof(Page) == "undefined" || typeof(Page) == "object"){
    Page = "0"
}



function h3tit(){

    var h3Size = $(".htmlcontent h3").length;
    for (n = 0; n < h3Size; n++) {
        var h3name = '<i class="before"></i><span></span><i class="after"></i>';
        var h3txt = $(".htmlcontent h3").eq(n).text();
        $(".htmlcontent h3").eq(n).empty();
        $(".htmlcontent h3").eq(n).append(h3name);
        $(".htmlcontent h3").eq(n).find("span").html(h3txt);
        var h4Size = $(".htmlcontent h3").eq(n).nextUntil("h3", "h4").length;
        for (i = 0; i < h4Size; i++) {
            $(".htmlcontent h3").eq(n).nextUntil("h3", "h4").eq(i).addClass("m-h4-" + i);
        };
        if (h4Size != 0) {
            $(".htmlcontent h3").eq(n).nextUntil("h3", "p").addClass("m-left-sod");
            $(".htmlcontent h3").eq(n).nextUntil("h3", "h4").eq(0).css("margin-top", 20);
        };
    };

}

function listags(){

    var cids = [];
    $('.m-listul li').each(function(){
        cids.push($(this).data('cid'))
    })
    ViewMore(cids);
    function ViewMore(cids){
        $.ajax({
            type: "Get",
            url: "/e/api/tags.php",
            data: {cids:cids},
            success: function(data){
                listDate(data)
            }
        });
    };
    function listDate(data){
        var cid, tags, href, tagname,html;
        $('.m-listul li').each(function() {
            cid = $(this).data('cid');
            tags = data[cid];
            html = "";
            if(tags!=undefined){
                tags.forEach(function(e){
                    href = e.href
                    tagname = e.tagname
                    html += "<a href='"+href+"'>"+tagname+"</a>"
                })

                $(this).find("div p span.m-listags").append(html);
            }

        });
    };

}

function indextags(){

    var cids = [];
    $('.m-hotul li').each(function(){
        var items = $(this).data('cid');
        if($.inArray(items,cids)==-1) {
            cids.push(items);
        }
    })
    ViewMore(cids);
    function ViewMore(cids){
        $.ajax({
            type: "Get",
            url: "/e/api/tags.php",
            data: {cids:cids},
            success: function(data){
                listDate(data)
            }
        });
    };
    function listDate(data){
        var cid, tags, href, tagname,html;
        $('.m-hotul li').each(function() {
            cid = $(this).data('cid');
            tags = data[cid];
            html = "";
            if(tags!=undefined){
                tags.forEach(function(e){
                    href = e.href
                    tagname = e.tagname
                    html += "<a href='"+href+"' style='margin-bottom:30px'>"+tagname+"</a>"
                })

                $(this).find(".m-newtags").empty().append(html);
            }

        });
    };

}