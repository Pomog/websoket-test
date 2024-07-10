let socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = () => {
    console.log("Connected to WebSocket server");
    socket.send("Hello Server, no ");
};

socket.onmessage = (event) => {
    console.log("Received: ", event.data);
};

socket.onerror = (error) => {
    console.error("WebSocket error: ", error);
};

socket.onclose = () => {
    console.log("WebSocket connection closed");
};