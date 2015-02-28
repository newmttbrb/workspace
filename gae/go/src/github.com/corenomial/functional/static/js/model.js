'use strict';

function ModelCtrl($scope) {
    $scope.Data = {
        "target_device" : "",
        "command" : "",
        "arguments" : ""
    };

    $scope.Chats = [];

    $scope.ws = new WebSocket("ws://localhost:1337/ws");

    $scope.ws.onopen = function() { console.log("Connection opened.... " + this.readyState); }

    $scope.ws.onmessage = function(evt) {
        console.log("received the message " + evt.data);
        var m = JSON.parse(evt.data);
        $scope.Chats.push(
            {
                t : m["t"],
                c : m["c"],
                a : m["a"],
                r : m["r"]
            });
        $scope.$apply();
    }

    $scope.ws.onclose = function() { console.log("Connection closed... " + this.readyState); }

    $scope.ws.onerror = function(evt) { console.log("Connection error... " + evt.data); }

    $scope.buttonPressed = function() {
        var sample_data =
        {
            t : $scope.Data.target_device,
            c : $scope.Data.command,
            a : $scope.Data.arguments
        };
        var sample_data_as_string = JSON.stringify(sample_data);
        console.log("sending the message " + sample_data_as_string);
        $scope.ws.send(sample_data_as_string);
    };


}