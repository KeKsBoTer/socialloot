(function ($) {
    $.each(['show', 'hide'], function (i, ev) {
        var el = $.fn[ev];
        $.fn[ev] = function () {
            this.trigger(ev);
            return el.apply(this, arguments);
        };
    });
})(jQuery);

$(function () {
    // autofocus work around
    $("input[autofocus]").focus();
    $('textarea').on("show", function () {
        var me = $(this)
        me.attr('style', 'height:' + (this.scrollHeight) + 'px;overflow-y:hidden;');
        if (this.value.length > 0)
            me.addClass("filled")
        else
            me.removeClass("filled")
    }).on('input', function () {
        this.style.height = 'auto';
        this.style.height = (this.scrollHeight) + 'px';
        if (this.value.length > 0)
            $(this).addClass("filled")
        else
            $(this).removeClass("filled")
    }).filter(":visible").trigger("show")

    // custom form handling
    $("form").submit(function (e) {
        var form = $(this);
        if (form.attr('method') === "GET")
            return
        var formdata = false;
        if (window.FormData) {
            formdata = new FormData(form[0]);
        }

        e.preventDefault();
        $.ajax({
            type: form.attr('method'),
            url: form.attr('action'),
            data: formdata ? formdata : form.serialize(),
            cache: false,
            contentType: false,
            processData: false,
            success: function (data) {
                if (data["success"] == true) {
                    var onSuccess = form.attr("onsuccess")
                    if (onSuccess) {
                        eval(onSuccess);
                    }
                    if (form.hasClass("clear-on-submit"))
                        form.trigger("reset");
                    if (data["dest"])
                        window.location.href = data["dest"]
                } else {
                    form.find(".message").text(data["message"])
                }
            }
        });
    });

    $("#user-popup-button").click(function () {
        var $this = $(this)
        var popup = $("#user-popup")
        if (popup.css('display') == 'none') {
            popup.css({
                "top": $this.offset().top + $this.height() + "px",
                "left": $this.offset().left + "px"
            });
            popup.show()
        } else {
            popup.hide();
        }
    });
    $("#user-popup").click(function (e) {
        e.stopPropagation();
    });

    $(document).click(function (e) {
        if (!$(e.target).parents().addBack().is('#user-popup-button')) {
            $("#user-popup").hide();
        }
    });
});

const vote_container = ".vote-container";

function vote(button, dir) {
    var $vote = $(button)
    // check if allready voted
    if ($vote.hasClass("voted"))
        return
    var $post = $vote.parents(vote_container)
    var item = $post.attr("item")
    var dir = 0
    if ($vote.hasClass("up"))
        dir = 1
    else if ($vote.hasClass("down"))
        dir = -1
    else {
        console.error("invalid vote direction:" + $vote.attr("class"))
        return
    }
    if (item)
        voteOnPost(item, dir, function () {
            $vote.addClass("voted");
            $otherVoteButton = $vote.siblings(".vote-button")
            if ($otherVoteButton.hasClass("voted")) {
                $otherVoteButton.removeClass("voted")
                dir *= 2
            }

            // update vote count
            var voteContainer = $post.find(".vote-count")
            var votes = parseInt(voteContainer.text())
            voteContainer.text(votes + dir)
        })
    else
        console.error("error while voting")
}


function voteOnPost(id, dir, onSuccess) {
    $.ajax({
        type: "POST",
        url: "/api/vote",
        dataType: "json",
        data: {
            id: id,
            dir: parseInt(dir)
        },
        success: function (data) {
            if (data["success"])
                onSuccess && onSuccess()
            else
                console.error(data["message"])
        },
        statusCode: {
            401: function () {
                window.location = "/login?dest=" + encodeURIComponent(window.location.pathname);
            }
        },
    });
}

function showCommentForm(elem) {
    $(elem).parents(".content").children(".comment-form").show()
}

function toggleComment(elem) {
    var button = $(elem);
    button.closest(".comment").first().toggleClass("collapsed")
}

function scrollList(button, dir) {
    var container = $(button).siblings().filter(".list-container")
    var list = container.find(".list");
    if (list.width() > container.width()) {
        var transformMatrix = list.css("-webkit-transform") ||
            list.css("-moz-transform") ||
            list.css("-ms-transform") ||
            list.css("-o-transform") ||
            list.css("transform");
        var matrix = transformMatrix.replace(/[^0-9\-.,]/g, '').split(',');
        var x = parseInt(matrix[12] || matrix[4] || 0);
        var amount = 100;
        switch (dir) {
            case "left":
                x -= amount
                break;
            case "right":
                x += amount
                break;
        }
        var max = list.width() - container.width()
        if (x > 0)
            x = 0;
        else if (x < -max)
            x = -max
        list.css("transform", "translateX(" + x + "px)")
    }
}

function previewURL(input, previewNode) {
    if (input.files && input.files[0]) {
        var reader = new FileReader();

        reader.onload = function (e) {
            previewNode.attr('src', e.target.result);
            previewNode.addClass("loaded")
        }

        reader.readAsDataURL(input.files[0]);
    }
}