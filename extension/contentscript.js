// We just want to notify the extension of our existence
chrome.runtime.sendMessage({}, function(response) {});

volstep = 0.1

function volumeup(v) {
	if (v.volume + volstep > 1)
		return
	v.volume += 0.1
}

function volumedown(v) {
	if (v.volume + volstep < 0)
		return
	v.volume -= 0.1
}

chrome.runtime.onMessage.addListener(function(request, sender, sendResponse) {
	console.log(sender);
	console.log(request);
	var videoElement = document.getElementsByTagName('video')[0]
	if (request.command == 0)
		videoElement.play();
	else if (request.command == 1)
		videoElement.pause();
	else if (request.command == 2)
		volumeup(videoElement)
	else if (request.command == 3)
		volumedown(videoElement)

});
