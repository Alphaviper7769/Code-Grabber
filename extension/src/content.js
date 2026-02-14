import { sites } from "./sites/sites.js";


function waitForDescription(timeout = 10000) {
  return new Promise((resolve, reject) => {
    const start = Date.now();

    const interval = setInterval(() => {
      const el = document.querySelector(
        '[data-track-load="description_content"]'
      );

      if (el) {
        clearInterval(interval);
        resolve();
      }

      if (Date.now() - start > timeout) {
        clearInterval(interval);
        reject("Timeout waiting for description");
      }
    }, 100);
  });
}

async function detectAndExtract() {
  const activeSite = sites.find(site => site.matches());

  if (!activeSite) {
    console.log("[CodeGrabber] Unsupported site.");
    return null;
  }

  console.log(`[CodeGrabber] Active site: ${activeSite.name}`);

  try {
    await waitForDescription();
  } catch (err) {
    console.warn("[CodeGrabber] Description not found:", err);
    return null;
  }

  const problem = await activeSite.extractProblem();

  if (!problem) {
    console.log("[CodeGrabber] No problem extracted.");
    return null;
  }

  console.log("[CodeGrabber] Extracted:", problem);

  return problem;
}



chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
  if (message.type === "FETCH_PROBLEM") {

    (async () => {
      const problem = await detectAndExtract();

      if (!problem) {
        sendResponse({ success: false });
        return;
      }

      sendResponse({
        success: true,
        payload: problem
      });
    })();


    return true;
  }
});


console.log("CodeGrabber content injected on:", window.location.href);
console.log("Content script ready. Waiting for extension click...");
