const express = require('express')
const app = express();

const http = require('http');
const server = http.Server(app);

const socketIO = require('socket.io');
const io = socketIO(server);

const port = process.env.PORT || 3000;

io.on('connection', (socket) => {

  console.log('user connected');
  socket.on('new-message', (message) => {
      console.log(message);
      io.to(message.roomId).emit('new-message', message);     
  });

    socket.on('new-room', (room) => {
      console.log('New Room ', room);
      io.emit('new-room', room);
    });
    socket.on('join-room', (room) => {
      socket.join(room);
      console.log('Joined Room ', room);
    });
});

server.listen(port, () => {
    console.log(`started on port: ${port}`);
});