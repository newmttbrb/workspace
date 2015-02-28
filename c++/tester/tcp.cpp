#include <stdio.h>
#include <stdlib.h>
#include <cstdlib>
#include <cstring>
#include <string.h>
#include <unistd.h>
#include <sys/ioctl.h>
#include <sys/types.h> 
#include <sys/socket.h>
#include <netinet/in.h>
#include <errno.h>
#define _TCP_C_
#include "tcp.h"

void error(const char *msg)
{
    perror(msg);
    exit(1);
}

int socketListen(int port)
{
#define IF_ERROR_THEN_CLEANUP(msg) if (rc < 0) { close(listen_socket); error(msg); }
     int rc;
     int on = 1;
     int listen_socket;
     struct sockaddr_in server_address;

     /* create socket to listen on */
     listen_socket = socket(AF_INET, SOCK_STREAM, 0);
     if (listen_socket < 0) error("ERROR opening socket");

     /* allow to be reusable */
     rc = setsockopt(listen_socket, SOL_SOCKET, SO_REUSEADDR, (char *)&on, sizeof(on));
     IF_ERROR_THEN_CLEANUP("ERROR on reuse"); 

     /* make non blocking */
     rc = ioctl(listen_socket, FIONBIO, (char *)&on);
     IF_ERROR_THEN_CLEANUP("ERROR on non blocking");

     /* bind the socket to the port */
     bzero((char *) &server_address, sizeof(server_address));
     server_address.sin_family = AF_INET;
     server_address.sin_addr.s_addr = INADDR_ANY;
     server_address.sin_port = htons(port);
     rc = bind(listen_socket, (struct sockaddr *) &server_address, sizeof(server_address));
     IF_ERROR_THEN_CLEANUP("ERROR on binding");

     /* set the listen backlog */
     rc = listen(listen_socket,5);
     IF_ERROR_THEN_CLEANUP("ERROR on listen");
     
     return listen_socket;
#undef IF_ERROR_THEN_CLEANUP
}

int socketAccept(int listen_socket)
{
     int client_socket;

     client_socket = accept(listen_socket, NULL, NULL);
     if (client_socket < 0 && errno != EWOULDBLOCK) error("ERROR on accept");
     return client_socket;
}

int socketConnect(int port)
{
#define IF_ERROR_THEN_CLEANUP(msg) if (rc < 0) { close(listen_socket); error(msg); }
     int rc;
     int on = 1;
     int listen_socket;
     struct sockaddr_in server_address;

     /* create socket to listen on */
     listen_socket = socket(AF_INET, SOCK_STREAM, 0);
     if (listen_socket < 0) error("ERROR opening socket");

     /* bind the socket to the port */
     bzero((char *) &server_address, sizeof(server_address));
     server_address.sin_family = AF_INET;
     server_address.sin_addr.s_addr = INADDR_ANY;
     server_address.sin_port = htons(port);
     rc = connect(listen_socket, (struct sockaddr *) &server_address, sizeof(server_address));
     IF_ERROR_THEN_CLEANUP("ERROR on connect");
     return listen_socket;
#undef IF_ERROR_THEN_CLEANUP
}


