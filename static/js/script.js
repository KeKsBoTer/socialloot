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

    $(".vote").on("click", function (e) {
        var $vote = $(this)
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
            voteOnPost(item, dir,function(){
                console.log("success")
                var voteContainer = $post.children(".votes")
                var votes = parseInt(voteContainer.text())
                voteContainer.text(votes+dir)
            })
        else
            console.error("error while voting")
    });
});


function voteOnPost(id, dir, onSuccess) {
    $.ajax({
        type: "POST",
        url: "/api/vote",
        data: {
            id: id,
            dir: parseInt(dir)
        },
        success: function (data) {
            console.log("voted!")
            onSuccess && onSuccess()
        }
    });
}