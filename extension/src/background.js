function sendToNative(problem) {
    const port = chrome.runtime.connectNative("com.codegrabber.native");

    port.onMessage.addListener((response) => {
        console.log("Native response:", response);
    });

    port.onDisconnect.addListener(() => {
        console.log("Native disconnected:", chrome.runtime.lastError);
    });

    port.postMessage(problem);
}

chrome.action.onClicked.addListener(async () => {
  const [tab] = await chrome.tabs.query({
    active: true,
    currentWindow: true
  });

  if (!tab?.id) return;

  try {
    const response = await chrome.tabs.sendMessage(tab.id, {
      type: "FETCH_PROBLEM"
    });

    if (!response?.success) {
      console.log("Failed to extract problem.");
      return;
    }

    console.log("Received problem from content:", response.payload);

    sendToNative(response.payload);

  } catch (err) {
    console.log("No content script on this page.");
  }
});


chrome.runtime.onMessage.addListener((message) => {
    if (message.type === "PROBLEM_DATA") {
        console.log("Received problem:", message.payload);
        sendToNative(message.payload);
    }
});
