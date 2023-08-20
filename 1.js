const WebSocket = require('ws');

const socket = new WebSocket("ws://127.0.0.1:8000/api/v1/ws");

// 打开连接时触发
socket.onopen = function() {
  console.log("WebSocket 连接已打开");
  
  // 发送消息
  socket.send("Hello, WebSocket!");
};

// 接收到服务器消息时触发
socket.onmessage = function(event) {
  console.log("接收到消息:", event.data);
};

// 关闭连接时触发
socket.onclose = function(event) {
  console.log("WebSocket 连接已关闭");
};