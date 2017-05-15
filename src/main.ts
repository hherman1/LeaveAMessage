// import * as Config from "./config"
import * as Config from "./config"


function displayData(data:any) {
    Config.resultDisplay.innerText = data
    Config.loadDiv.style.display = "none"
    Config.resultDisplay.style.display = "initial"
}

function sendMessage() {
    $.ajax({
        method:"POST",
        data:Config.entryText.value,
        url:Config.messageServer,
        success: function(data) {
            displayData(data)
        }
    })
    Config.entryText.value = ""
    Config.promptDiv.style.display = "none"
    Config.loadDiv.style.display = "initial"
}

Config.submitButton.onclick = function() {
    sendMessage()
}


function main() {
    console.log("test");
    
}

main();