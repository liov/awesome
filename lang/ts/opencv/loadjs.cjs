
Module = {
    onRuntimeInitialized() {
        // this is our application:
        console.log(cv.getBuildInformation())
    }
}

async function loadRemoteModule(url) {
    try {
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const scriptContent = await response.text();
        eval(scriptContent);
    } catch (error) {
        console.error('Error loading and executing script:', error);
    }
}
loadRemoteModule('https://docs.opencv.org/4.10.0/opencv.js')
