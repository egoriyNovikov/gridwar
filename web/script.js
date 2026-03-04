const socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = function () {
  socket.send("ping");
};

socket.onmessage = function (event) {
  const data = event.data;
  const messages = document.getElementById("messages");
  messages.innerHTML += data + "<br>";
  if (data === "pong") {
    console.log("Получен pong:", data);
  } else {
    try {
      console.log(JSON.parse(data));
    } catch {
      console.log(data);
    }
  }
};