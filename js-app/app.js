(function() {
    'use strict';

    // Request type
    var MSG_ERROR = 1 + 0,
        MSG_GET_USER = 1 + 1,
        MSG_CREATE_CHARACTER = 1 + 2,
        MSG_REMOVE_CHARACTER = 1 + 3,
        MSG_CHECK_CHARACTER = 1 + 4,
        MSG_USER_ACTION = 1 + 5;

    // Response type
    var MSG_USER_DATA = 1 + 0,
        MSG_CHARACTER_DATA = 1 + 1;

    var WsWrapper = function(addr) {
        this.ws = new WebSocket(addr);
        this.eBinds = {
            message: [],
            error: [],
            open: [],
            close: []
        };

        window.onbeforeunload = function(event) {
            this.ws.close();
        }.bind(this);

        this.ws.onmessage = function(e){
            var msg = JSON.parse(e.data);
            console.log('Got message with type', msg.type, msg);
            this.eBinds.message.forEach(function(cb){cb(msg.data);}.bind(this));
            if (this.eBinds['message.'+msg.type]) {
                this.eBinds['message.'+msg.type].forEach(function(cb){cb(msg.data);}.bind(this));
            }
        }.bind(this);
        this.ws.onerror = function(e){
            this.eBinds.error.forEach(function(cb){cb(e);}.bind(this));
        }.bind(this);
        this.ws.onclose = function(e){
            this.eBinds.close.forEach(function(cb){cb(e);}.bind(this));
        }.bind(this);
        this.ws.onopen = function(e){
            console.log("connection ready;");
            this.eBinds.open.forEach(function(cb){cb(e);}.bind(this));
        }.bind(this);

        return {
            send: function(data) {
                this.ws.send(JSON.stringify(data));
            }.bind(this),

            on: function(eventName, type, cb) {
                if (arguments.length === 2) {
                    cb = type;
                    this.eBinds[eventName] = this.eBinds[eventName] || [];
                    this.eBinds[eventName].push(cb);
                } else {
                    eventName += '.' + type;
                    this.eBinds[eventName] = this.eBinds[eventName] || [];
                    this.eBinds[eventName].push(cb);
                }
            }.bind(this),

            trigger: function(eventName, type) {
                if (type) {
                    eventName = eventName + '.' + type;
                }
                this.eBinds[eventName] = this.eBinds[eventName] || [];
                this.eBinds[eventName].forEach(function(cb) {cb();}.bind(this));
            }.bind(this)
        };
    };

    var Game = function(connection) {
        this.c = connection;
        this.user = {};
        this.characters = [];
        this.classes = ['warrior', 'priest', 'mage'];
        this.uid = "";

        this.c.on('message', function(data) {
            console.log(data);
            this.uid = data.uid;
            window.localStorage.setItem('uid', this.uid);
        }.bind(this));

        this.c.on('message', MSG_USER_DATA, function(user) {
            this.uid = user.uid;
            this.user = user;
            console.log('Got user data:', this.user);
            this.c.trigger('user:ready');
        }.bind(this));

        this.c.on('message', MSG_CHARACTER_DATA, function(characters) {
            if (characters) {
                this.characters = characters;
                console.log('Got characer data:', this.characters);
                this.c.trigger('character:ready');
            }
        }.bind(this));

        return {
            newCharacter: function(){},
            hasCharacter: function(){},
            checkLocalUser: function() {
                var uid = localStorage.getItem('uid');

                console.log('Checking existance of local user');
                if (uid) {
                    console.log('Local user exists, requesting details');
                    var msg = {type: MSG_GET_USER, uid: uid};

                    this.c.send(msg);
                    this.c.on('message', function(msg) {
                        if (msg.type !== MSG_GET_USER) {
                            return;
                        }

                        this.user = msg.data;
                        this.characters = this.user.characters || [];

                        console.log('Got user data:', this.user);
                        console.log('There are characters, which already exists:', game.getCharacters());
                    }.bind(this));


                    return true;
                }

                console.log('No local user found');
                return false;
            }.bind(this),

            newUser: function() {
                console.log("Creating new user");
                var msg = {type: MSG_GET_USER, uid: null};

                this.c.send(msg);
                this.c.on('message', MSG_GET_USER, function(msg) {
                    this.user = msg.data || {};
                    this.characters = this.user.characters || [];

                    console.log('Got user data:', this.user);
                    console.log('There are characters, which already exists:', game.getCharacters());
                }.bind(this));
            }.bind(this),

            test: function(run) {
                if (run) {
                    var user = {
                        name: 'Joil ' + Math.random(),
                        email: 'joil+' + Math.random() + '@testermail.com'
                    };

                    this.c.send({type: MSG_CREATE_USER, data: user});
                }
            }.bind(this),

            createNewCharacter: function() {
                var uid = this.uid || localStorage.getItem('uid');
                console.log('Creating new character');
                this.c.send({type: MSG_CREATE_CHARACTER, uid: uid});

                this.c.on('message', MSG_CREATE_CHARACTER, function(msg) {
                    this.characters = this.characters.concat(msg.data.characters);
                    console.log('Got character data:', this.characters[0]);
                }.bind(this));
            }.bind(this),

            getCharacters: function() {
                return this.characters;
            }.bind(this)
        };
    };

    var ws = new WsWrapper('ws://127.0.0.1:9090/ws');

    var game = new Game(ws);

    ws.on('open', function() {
        if (!game.checkLocalUser()) {
            game.newUser();
        }
        for (var i = 0; i < 3000; i++) {
            setTimeout(game.test.bind(this, false), Math.random() * 1000 + i);
        }

        ws.on('user:ready', function() {
            game.createNewCharacter();
            // game.createNewCharacter();
        });
    });

})();