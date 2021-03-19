
layui.define(["layer","jquery","message"],function(exports){
    var $ = layui.$
    ,message = layui.message
    ,form = layui.form;
    //图片上传
    $(document).on("click", ".G-upload-con .choice-file", function() {
        let fileEl = $(this),
            faEl = fileEl.parents(".G-upload-con"),
            limitSize = Number(faEl.data("limitsize")),
            limitHight = Number(faEl.data("limithight")),
            limitWidth = Number(faEl.data("limitwidth")),
            inputName = faEl.attr('data-name')
            limitNum = Number(faEl.data("limitnum")),
            limitType = false
        limitHight > 0 ? limitType = true : "";
        //是否限制图片数量
        if (limitNum == faEl.find(".G-img-item").length) {
            message.error(`最多只能上传${limitNum}张图片！`);
            return false
        }
        if (typeof FileReader == "undefind") {
            faEl.InnerHTML = "<p>你的浏览器不支持FileReader接口！</p>";
            fileEl.setAttribute("disabled", "disabled");
        }
        if (fileEl) {
            fileEl[0].onchange = function() {
                //检验是否为图像文件
                var changeResult = fileEl[0].files[0];
                if (changeResult.size >= 1024 * 1024 * limitSize && limitSize > 0) {
                    alert("图片大小不能超过" + limitSize + "M！")
                    return false;
                }
                if (!/image\/\w+/.test(changeResult.type)) {
                    alert("请上传图片")
                    return false;
                }
                var reader = new FileReader();
                reader.readAsDataURL(changeResult);
                reader.onload = function(e) {
                    $('.img-hint').hide();
                    $('.img-num').show();
                    var img = new Image;
                    img.src = this.result;
                    img.onload = function() {
                        var imgEl = `
                                <div class="G-img-item" > 
                                    <img src=${img.src}> 
                                    <div class="G-img-delete cur_pointer" flex="cross:center main:center" style="display:none">
                                        <span class="delete">删除<span/>
                                    </div>
                                    <input type="hidden" name="${inputName}"  value="${img.src}" />
                                </div>
                            `;
                        //判断是否限制大小
                        if (limitType && (img.width != limitWidth || img.height != limitHight)) { //限制大小
                            alert("请上传" + limitWidth + "*" + limitHight + "像素的图片")
                            return false;
                        } else { //不限制大小
                            faEl.find(".G-img-add").before(imgEl);
                            fileEl.replaceWith(fileEl.val("").clone(true))
                            faEl.parents('.photo').addClass('active')
                        }
                    }
                }
            }
        }
    })
    // 显示删除
    $(document).on("mouseover", ".G-upload-con .G-img-item", function() {
        let th = $(this)
        th.find('.G-img-delete').show();
    })
    $(document).on("mouseout", ".G-upload-con .G-img-item", function() {
        let th = $(this)
        th.find('.G-img-delete').hide();
    })
    //删除图片
    $(document).on("click", ".G-upload-con .G-img-delete", function() {
    let th = $(this)
    if (th.parents('.G-upload-con').find('.G-img-item').length == 1) {
        th.parents('.photo').removeClass('active')
    }
    th.parents(".G-img-item").remove();
    })
    exports('uploadAlone', {});    
});  