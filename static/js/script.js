$(function () {

    // custom form handling
    $("form").submit(function (e) {
        $form = $(this)
        e.preventDefault();
        $.ajax({
            type: $form.attr('method'),
            url: $form.attr('action'),
            data: $form.serialize(),
            success: function (data) {
                if (data["success"] == true) {
                    if (data["dest"])
                        window.location.href = data["dest"]
                } else {
                    $form.find(".message").text(data["message"])
                }
            }
        });
    });

});

function vote(button, dir) {
    var $vote = $(button)
    // check if allready voted
    if ($vote.hasClass("voted"))
        return
    var $post = $vote.parents(".post")
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
                dir*=2
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
        data: {
            id: id,
            dir: parseInt(dir)
        },
        success: function (data) {
            onSuccess && onSuccess()
        }
    });
}