function sendToNative(problem) {
    const port = chrome.runtime.connectNative("com.codegrabber.native");

    port.postMessage(problem);

    port.onMessage.addListener((response) => {
        console.log("Native response:", response);
    });

    port.onDisconnect.addListener(() => {
        console.log("Native disconnected");
    });
}



chrome.action.onClicked.addListener(async () => {
    const [tab] = await chrome.tabs.query({
        active: true,
        currentWindow: true
    });

    if (!tab?.id) return;

    try {
        await chrome.tabs.sendMessage(tab.id, { type: "FETCH_PROBLEM" });
        console.log("Requested problem from content script");
    } catch (err) {
        console.log("No content script on this page.");
    }
});



chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
    if (message.type === "PROBLEM_DATA") {
        console.log("Received problem:", message.payload);

        sendToNative(message.payload);
    }
});
