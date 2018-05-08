$(function () {
    // autofocus work around for some browsers
    $("input[autofocus]").focus();

    // resize textarea according to the length of the user input
    // disables submit button if the textarea is empty or only contains spaces
    $('.comment-form > form > textarea').on('input show', function () {
        this.style.height = 'auto';
        this.style.height = (this.scrollHeight) + 'px';

        var submit = $(this).parents("form").find("input[type=\"submit\"]")
        submit.prop("disabled", this.value.trim().length == 0);
    }).filter(":visible").trigger("show") // trigger show event for all items that are visible by defautl

    // custom form handling
    $("form").submit(function (e) {
        var form = $(this);

        if (form.hasClass("confirm") && !confirm(form.attr("message"))) {
            e.preventDefault();
            return
        }

        if (form.attr('method') === "GET") {
            var input = form.find("input[name=\"query\"]")
            // trim whitespace from input
            var query = input.val().trim()
            input.val(query)
            // don't send form if query is empty
            if (query.length < 1)
                e.preventDefault();
            return
        }
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
                form.find(".invalid").removeClass("invalid")
                form.find(".message, .global-message").text("")
                if (data["success"] == true) {
                    var onSuccess = form.attr("onsuccess")
                    if (onSuccess) {
                        eval(onSuccess);
                    }
                    if (form.hasClass("clear-on-submit"))
                        form.trigger("reset");
                    // redirect to given location
                    if (data["dest"])
                        window.location.href = data["dest"]
                } else {
                    if (data["field"]) {
                        var field = form.find("[name=\"" + data["field"] + "\"]")
                        field.addClass("invalid")
                        var label = field.siblings(".message")
                        if (label.length)
                            label.text(data["message"])
                        else
                            form.find(".global-message").text(data["message"])
                    } else if (data["message"])
                        form.find(".global-message").text(data["message"])
                    else
                        data["message"] = "Somthing went wrong"
                }
            }
        });
    });

    // do not send clicks to children
    // so every click within #user-popup will be recognized as click on #user-popup
    $("#user-popup").click(function (e) {
        e.stopPropagation();
    });

    // toggle user popup list
    $("#user-popup-button").click(function () {
        var $this = $(this)
        var popup = $("#user-popup")
        if (popup.css('display') == 'none') {
            var left = $this.offset().left
            // make sure popup is not going out of screen
            var max = window.innerWidth - popup.outerWidth() - 5;
            if (left > max)
                left = max
            popup.css({
                "top": $this.offset().top + $this.height() + "px",
                "left": left + "px"
            });
            popup.show()
        } else {
            popup.hide();
        }
    });

    // hide user popup if user clicks anywhere but on the pupup
    $(document).click(function (e) {
        if (!$(e.target).parents().addBack().is('#user-popup-button')) {
            $("#user-popup").hide();
        }
    });
});

// function for vote button
// sends vote to server and changes the ui based on the vote action
function vote(button, dir) {
    var $vote = $(button)
    // check if allready voted
    if ($vote.hasClass("voted"))
        return

    var $post = $vote.parents(".vote-container")

    // get post id
    var item = $post.attr("item")
    if (!item) {
        console.error("error while voting")
        return
    }

    // get vote direction (up- or downvote) from class
    var dir = 0
    if ($vote.hasClass("up"))
        dir = 1
    else if ($vote.hasClass("down"))
        dir = -1
    else {
        console.error("invalid vote direction:" + $vote.attr("class"))
        return
    }
    voteOnPost(item, dir, function () {
        $vote.addClass("voted");
        $otherVoteButton = $vote.siblings(".vote-button")
        if ($otherVoteButton.hasClass("voted")) {
            $otherVoteButton.removeClass("voted")
            // if you change your vote on a post you allready voted on
            // the vote count in- or decreases by two
            // since we don't want to reload the page after every vote we just conclude the new vote count
            dir *= 2
        }

        // update vote count
        var voteContainer = $post.find(".vote-count")
        var votes = parseInt(voteContainer.text())
        voteContainer.text(votes + dir)
    })
}


// Sends a ajax request to server to vote for post
// If the user is not authorized, he is redirected to the login page
// Params:
//  id          it the item to vote on
//  dir         is up(1) or down(-1) vote
//  onSuccess   is executed with no arguments if the vote was successfull
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

// expands or collapses the comment in which the given button is located
function toggleComment(btn) {
    var button = $(btn);
    button.closest(".comment").first().toggleClass("collapsed")
}

// shows comment form if reply button is pressed
function showCommentForm(btn) {
    var form = $(btn).parents(".content").children(".comment-form")
    form.show()
    form.find("textarea").trigger("show")
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

// previews selected image in browser
// if the file can be loaded the image is inserted into the source tag of the given preview node 
// also the class "loaded" is added to the preview node to allow custom css for loaded images 
function previewImage(input, previewNode) {
    if (input.files && input.files[0]) {
        var reader = new FileReader();

        reader.onload = function (e) {
            previewNode.attr('src', e.target.result);
            previewNode.addClass("loaded")
        }

        reader.readAsDataURL(input.files[0]);
    }
}