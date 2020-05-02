new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        id: null, // Email address used for grabbing an avatar
        token: null, // Email address used for grabbing an avatar
        username: null, // Our username
        type: null, // Our username
        login: null, // Our username
        chatId: null, // Our username
        joined: false // True if email and username have been filled in
    },

    created: function () {

    },

    methods: {
        send: function () {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        id: this.id,
                        username: this.username,
                        chatId: this.chatId,
                        message: $('<p>').html(this.newMsg).text() // Strip out html
                    }
                    ));
                this.newMsg = ''; // Reset newMsg
            }
        },
        choosePatient: function () {
            var self = this;
            jQuery.ajax({
                type: "POST",
                url: 'http://' + window.location.host + '/api/v1/chats/start',
                dataType: "JSON",
                contentType: 'application/json',
                headers: {
                    "Authorization": "Baerer " + self.token
                },
                success: function (data) {
                    self.chatId = data.id;
                    self.loadWS();
                }
            })

        },
        terminateChat: function () {
            var self = this;
            jQuery.ajax({
                type: "POST",
                url: 'http://' + window.location.host + '/api/v1/chats/terminate/' + self.chatId,
                dataType: "JSON",
                contentType: 'application/json',
                headers: {
                    "Authorization": "Baerer " + self.token
                },
                success: function (data) {
                    self.chatId = null;
                    self.joined = false;
                }
            })

        },
        join: function () {
            var self = this;
            if (!this.login) {
                Materialize.toast('You must enter an login', 2000);
                return
            }
            if (!$('#password').val()) {
                Materialize.toast('You must enter a password', 2000);
                return
            }
            this.login = $('<p>').html(this.login).text();
            jQuery.ajax({
                type: "POST",
                url: 'http://' + window.location.host + '/api/v1/authentication',
                data: JSON.stringify({ "login": this.login, "pass": $('#password').val() }),
                dataType: "JSON",
                contentType: 'application/json',
                success: function (data) {
                    self.sucessLogin(data)
                }
            })
        },
        loadData: function (datain) {
            this.id = datain.UserId;
            this.token = datain.Token;
            this.username = datain.Name;
            this.type = datain.Type;
            this.chatId = datain.ChatId;
        },
        sucessLogin: function (datain) {
            var self = this;
            self.loadData(datain);
            if (this.type === 'Patient' || (this.type === 'Doctor' && self.chatId && self.chatId != ""))
                this.loadWS();
        },
        loadWS: function () {
            var self = this;
            this.ws = new WebSocket('ws://' + window.location.host + '/chatbox?Baerer=' + this.token);
            this.ws.addEventListener('message', function (e) {
                var msg = JSON.parse(e.data);
                if (self.chatId === null) {
                    self.chatId = msg.chatId;
                }
                self.chatContent += '<div class="chip">'
                    + '<img src="' + self.gravatarURL(msg.username) + '">' // Avatar
                    + msg.username
                    + '</div>'
                    + emojione.toImage(msg.message) + '<br/>'; // Parse emojis

                var element = document.getElementById('chat-messages');
                element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
            });
            this.joined = true;
            self.waitForSocketConnection(self.ws, function(){
                self.newMsg = 'Entrou na sala';
                self.send();
            });
        },
        waitForSocketConnection: function (socket, callback) {
            var self = this;
            setTimeout(
                function () {
                    if (socket.readyState === 1) {
                        console.log("Connection is made")
                        if (callback != null){
                            callback();
                        }
                    } else {
                        console.log("wait for connection...")
                        self.waitForSocketConnection(socket, callback);
                    }
        
                }, 5); // wait 5 milisecond for the connection...
        },
        gravatarURL: function (username) {
            return 'http://www.gravatar.com/avatar/' + CryptoJS.MD5(username);
        }
    }
});