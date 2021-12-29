(() => {
  // web/src/config.ts
  function throwIfNull(element) {
    if (element == null) {
      throw "required element was null";
    }
    return element;
  }
  var messageServer = "note";
  var resultDisplay = throwIfNull(document.getElementById("result"));
  var promptDiv = throwIfNull(document.getElementById("prompt"));
  var loadDiv = throwIfNull(document.getElementById("loadingBar"));
  var submitButton = throwIfNull(document.getElementById("sendMessage"));
  var entryText = throwIfNull(document.getElementById("entry"));

  // web/src/main.ts
  function displayData(data) {
    resultDisplay.innerText = data;
    loadDiv.style.display = "none";
    resultDisplay.style.display = "initial";
  }
  function sendMessage() {
    $.ajax({
      method: "POST",
      data: entryText.value,
      url: messageServer,
      success: function(data) {
        displayData(data);
      }
    });
    entryText.value = "";
    promptDiv.style.display = "none";
    loadDiv.style.display = "initial";
  }
  submitButton.onclick = function() {
    sendMessage();
  };
  function main() {
    console.log("test");
  }
  main();
})();
