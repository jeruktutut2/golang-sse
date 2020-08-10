function onLoaded() {
    console.log("on loaded...");

    var eventSource = new EventSource("/handshake");
    eventSource.onmessage = function(event) {
        console.log("on message... ");
        console.dir(event);
        document.getElementById("counter").innerHTML = event.data;
    }
}