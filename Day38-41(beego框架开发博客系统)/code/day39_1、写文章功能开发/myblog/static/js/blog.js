$(document).ready(function () {
    //注册
    $("register-from").validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 10]
            },
            password: {
                required: true,
                rangelength: [5, 10]
            },
            repassword: {
                required: true,
                rangelength: [5, 10],
                equalTo: "#register-password"
            }
        },
        messages: {
            username: {
                required: "请输入用户名",
                rangelength: "用户名必须是5-10位"
            },
            password: {
                required: "请输入密码",
                rangelength: "密码必须是5-10位"
            },
            repassword: {
                required: "请确认密码",
                rangelength: "密码必须是5-10位",
                equalTo: "两次输入的密码必须相等"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/register";
            // alert("urlStr:"+urlStr)
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/login"
                        }, 1000)
                    }
                },
                err: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            })
        }
    });


    //登录
    $("#login-form").validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 10]
            },
            password: {
                required: true,
                rangelength: [5, 10]
            }
        },
        messages: {
            username: {
                required: "请输入用户名",
                rangelength: "用户名必须是5-10位"
            },
            password: {
                required: "请输入密码",
                rangelength: "密码必须是5-10位"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/login"
            alert("urlStr:" + urlStr)
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message + ":" + status)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/"
                        }, 1000)
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)

                }
            });
        }
    });

    //修改和添加文章的表单
    $("#write-art-form").validate({
        rules: {
            title: "required",
            tags: "required",
            short: {
                required: true,
                minlength: 2
            },
            content: {
                required: true,
                minlength: 2
            }
        },
        messages: {
            title: "请输入标题",
            tags: "请输入标签",
            short: {
                required: "请输入简介",
                minlength: "简介内容最少两个字符"
            },
            content: {
                required: "请输入文章内容",
                minlength: "文章内容最少两个字符"
            }
        },
        submitHandler: function (form) {
            alert("hello")
            var urlStr = "/article/add";
            //判断文章id确定提交的表单的服务器地址
            //若id大于零，说明是修改文章
            var artId = $("#write-article-id").val();
            alert("artId:" + artId);
            if (artId > 0) {
                urlStr = "/article/update"
            }
            alert("urlStr:" + urlStr);
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert(":data:" + data.message);
                    setTimeout(function () {
                        window.location.href = "/"
                    }, 1000)
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            });
        }
    });


    //文件
    $("#album-upload-button").click(function () {
        var filedata = $("#album-upload-file").val();
        if (filedata.length <= 0) {
            alert("请选择文件!");
            return
        }
        //文件上传通过Formdata去储存文件的数据
        var data = new FormData()
        data.append("upload", $("#album-upload-file")[0].files[0]);
        alert(data)
        var urlStr = "/upload"
        $.ajax({
            url: urlStr,
            type: "post",
            dataType: "json",
            contentType: false,
            data: data,
            processData: false,
            success: function (data, status) {
                alert(":data:" + data.message);
                if (data.code == 1) {
                    setTimeout(function () {
                        window.location.href = "/album"
                    }, 1000)
                }
            },
            error: function (data, status) {
                alert("err:" + data.message + ":" + status)
            }
        })
    })
});