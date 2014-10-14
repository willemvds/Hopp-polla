var activeTab = null;

chrome.runtime.onMessage.addListener(function(request, sender, sendResponse) {
    console.log(sender);
    console.log(request);
    if (sender.tab) {
        activeTab = sender.tab;
    }
});

var ws = null;
var addr = "ws://localhost:42050/ws"

function init() {
    try {
        ws = new WebSocket(addr);
    } catch(err) {
        console.log(err);
        setTimeout(function() { this.connect(addr); }, 5000);
        return
    }
    setup(ws);
};

function setup(ws) {
    ws.onmessage = function (evt) {
        console.log(evt);
        var data = JSON.parse(evt.data);

        if (activeTab === null) {
            return
        }

        chrome.tabs.sendMessage(activeTab.id, {command: data.command}, function(response) {
            console.log(response);
        });
    };

    ws.onclose = function(err) {
        console.log(err);
        setTimeout(function() { init(); }, 5000);
    }
}

init();