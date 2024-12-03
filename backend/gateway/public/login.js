$(document).ready(function () {
    $("#login-button").click(function () {
        const username = $("#user-input").val().trim();

        if (!username) {
            $("#error-message").text("Please enter your username.");
            return;
        }

        $.ajax({
            url: "/user",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify({ user: username }),
            success: function (response) {
                window.location.href = `index.html?user=${encodeURIComponent(username)}`;
            },
            error: function () {
                $("#error-message").text("Failed to authenticate. Please try again.");
            },
        });
    });
});
