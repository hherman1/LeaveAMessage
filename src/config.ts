

export function throwIfNull<T>(element: T | null): T  {
    if(element == null) {
        throw "required element was null"
    }
    return element
}


export var messageServer = "http://hherman.com/leaveAMessage/server"

export var resultDisplay = <HTMLDivElement>throwIfNull(document.getElementById("result"))

export var promptDiv = <HTMLDivElement>throwIfNull(document.getElementById("prompt"))
export var loadDiv = <HTMLDivElement>throwIfNull(document.getElementById("loadingBar"))


export var submitButton = <HTMLButtonElement>throwIfNull(document.getElementById("sendMessage"))
export var entryText = <HTMLTextAreaElement>throwIfNull(document.getElementById("entry"))
