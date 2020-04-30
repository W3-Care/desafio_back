const express = require('express')
const app = express();

const http = require('http');
const server = http.Server(app);

const socketIO = require('socket.io');
const io = socketIO(server);

const port = process.env.PORT || 3000;

io.on('connection', (socket) => {

  let previousId;
  socket.join(1);
  console.log('user connected');
   socket.on('new-message', (message) => {
        console.log(message);
        io.to(1).emit('new-message', message);
      });
});

server.listen(port, () => {
    console.log(`started on port: ${port}`);
});