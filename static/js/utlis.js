
var ajax = $.ajax;
$.extend({
    ajax: function(url, options) {
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
            headers = $.extend(headers, {'X-Xsrftoken':xsrftoken});
            console.log("replaced")
        }
        options.headers = headers;
        return ajax(url, options);
    }
});

jQuery.postJSON = function(url, args, callback) {
    var xsrf, xsrflist;
    xsrf = $.cookie("_xsrf");
    xsrflist = xsrf.split("|");
    args._xsrf = base64_decode(xsrflist[0]);
     $.ajax({url: url, data: $.param(args), dataType: "text", type: "POST",
         success: function(response) {
         callback(eval("(" + response + ")"));
     }});
 };
 
 // add listender for show and hide
 (function ($) {
    $.each(['show', 'hide'], function (i, ev) {
        var el = $.fn[ev];
        $.fn[ev] = function () {
            this.trigger(ev);
            return el.apply(this, arguments);
        };
    });
})(jQuery);
