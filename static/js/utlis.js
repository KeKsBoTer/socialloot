// add xsrf token to all ajax requests
// see https://github.com/beego/beedoc/blob/master/en-US/mvc/controller/xsrf.md#usage-in-javascript
var ajax = $.ajax;
$.extend({
    ajax: function (url, options) {
        if (typeof url === 'object') {
            options = url;
            url = undefined;
        }
        options = options || {};
        url = options.url;
        var xsrftoken = $('meta[name=_xsrf]').attr('content');
        var headers = options.headers || {};
        var domain = document.domain.replace(/\./ig, '\\.');
        if (!/^(http:|https:).*/.test(url) || eval('/^(http:|https:)\\/\\/(.+\\.)*' + domain + '.*/').test(url)) {
            headers = $.extend(headers, {
                'X-Xsrftoken': xsrftoken
            });
        }
        options.headers = headers;
        return ajax(url, options);
    }
});

// add xsrf token to http json requests
// see https://github.com/beego/beedoc/blob/master/en-US/mvc/controller/xsrf.md#usage-in-javascript
jQuery.postJSON = function (url, args, callback) {
    var xsrf, xsrflist;
    xsrf = $.cookie("_xsrf");
    xsrflist = xsrf.split("|");
    args._xsrf = base64_decode(xsrflist[0]);
    $.ajax({
        url: url,
        data: $.param(args),
        dataType: "text",
        type: "POST",
        success: function (response) {
            callback(eval("(" + response + ")"));
        }
    });
};