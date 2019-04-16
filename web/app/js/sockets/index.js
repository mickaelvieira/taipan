import autobahn from "autobahn";

(function(win) {
  win.addEventListener("DOMContentLoaded", () => {
    console.log("hello socket world");
    var conn = new WebSocket("wss://ws.mickaelvieira.local");
    // var conn = new WebSocket("wss://localhost:8080");

    conn.onopen = function(e) {
      console.log("Connection established!");

      conn.send("Hello Mother Fucker!");
    };

    conn.onmessage = function(e) {
      console.log(e.data);
    };
  });
})(window);
