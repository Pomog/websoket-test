let socket = new WebSocket("ws://localhost:8080/ws");

socket.onmessage = (event) => {console.log("received: ", event.data)}

socket.send("test")