# Websocket learning project

This project's backend is written in GO. However once the Websocket logic is learnt. It can be applied to any other desired languages.
For next step, I am planning to develop a similar project using Java / Spring BOOT as well.

## Introduction

Websocket client applications use **Websocket API** to communicate with Websocket servers over the Websocket protocol.
Establishing a Websocket relies on the **HTTP Upgrade mechanism**, so the request for the protocol upgrading is implicit.

A Websocket server is a TCP application listening on any port of a server that follows a specific protocol simple as that.
Websocket servers are often separate and specialized servers, therefore most of the time we will use a reverse proxy to detect Websocket handshakes, pre process them, and send those clients to a real Websocket server.

### The Websocket handshake
The server must listen for incoming socket connections using a standard TCP socket.

## Websocket API
The Websocket API is an advanced technology which makes it possible to open a two-way interactive communication session between
the user's browser and a server. Thereby, the messages can be sent to a server by client and client receives event-driven responses
without having to poll the server for a reply.

### Interfaces
-   **Websocket**    : The primary interface for connecting to a Websocket server and then sending and receiving data on the connection.
-   **CloseEvent**   : The event sent by the Websocket object when the connection closes.
-   **MessageEvent** : The event sent by the Websocket object when a message is received from the server

## Build & Run
```
$ go build -o server.exe

$ server.exe
```
 
# Resources
-   [Writing Websocket applications](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_client_applications#Receiving_messages_from_the_server)

-   [The Websocket API](https://developer.mozilla.org/en-US/docs/Web/API/Websockets_API)

-   [Websocket protocol](https://tools.ietf.org/html/rfc6455)