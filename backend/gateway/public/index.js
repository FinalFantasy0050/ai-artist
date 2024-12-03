$(document).ready(function () {
    const urlParams = new URLSearchParams(window.location.search);
    const user = urlParams.get("user")

    console.log(`Logged in as: ${user}`);

    let activeTab = "image";

    $(".nav-button").click(function () {
        $(".nav-button").removeClass("active");
        $(this).addClass("active");

        activeTab = $(this).attr("id").replace("btn-", "");
        $("#text-input").val("");
        $("#loading").text("");
        $("#response-display").html("");
        console.log(`Active tab set to: ${activeTab}`);
    });

    $("#submit").click(function () {
        const prompt = $("#text-input").val();

        if (!prompt.trim()) {
            alert("Please enter a prompt!");
            return;
        }

        $("#loading").text("Generating...");
        console.log(`Sending request for ${activeTab} with prompt: ${prompt}`);

        let url = "";
        switch (activeTab) {
            case "image":
                url = "/infer/image";
                break;
            case "writer":
                url = "/infer/writer";
                break;
            case "character":
                url = "/infer/character";
                break;
        }

        $.ajax({
            url: url,
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify({ prompt: prompt, user: user }),
            success: function (response) {
                console.log("Response received:", response);
                $("#loading").text("");

                if (response.image) {
                    const imageBase64 = response.image;
                    const imageUrl = `data:image/png;base64,${imageBase64}`;
                    $("#response-display").append(`
                        <div class="response-image">
                            <a href="${imageUrl}" download="generated-image.png" class="download-icon" title="Download Image">⬇</a>
                            <img src="${imageUrl}" alt="Generated Image" />
                        </div>
                    `);
                }

                if (response.text) {
                    const textContent = response.text.replace(/\n/g, '<br>');
                    $("#response-display").append(`
                        <div class="response-text">
                            <p>${textContent}</p>
                        </div>
                    `);
                }
            },
            error: function (xhr, status, error) {
                console.error("Request failed:", error);
                $("#loading").text("");
                $("#response-display").text("Failed to generate response.");
            },
        });
    });
});
