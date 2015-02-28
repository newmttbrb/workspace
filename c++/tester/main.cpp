#include <stdio.h>
#include <stdlib.h>
#include <cstring>
#include <string.h>
#include <unistd.h>
#include <sys/ioctl.h>
#include <sys/types.h> 
#include <sys/socket.h>
#include <sys/time.h>
#include <netinet/in.h>
#include <errno.h>
#include <algorithm> 
#include "Selecter.h"
#include "tcp.h"
#include "Reader.h"
#include "Firewall.h"
#include "ManagedChild.h"

#define PORT_NUMBER 8782


void cleanupSocket(int socket,int* lastSocket)
{
	if (socket == *lastSocket) *lastSocket = 0;
	close(socket);
}

int main(int argc, char **argv)
{
	 if (argc != 2) { printf("error: please supply a program to wrap\n"); exit(1); }

     Selecter s;
     Reader r;
     int socket,lastSocket;
     int rc;

     int ParentWrite,ParentRead;
     int ChildWrite,ChildRead;
     createPipe(&ChildRead ,&ParentWrite,NON_BLOCKING_READ,"ParentToChild pipe");
     createPipe(&ParentRead,&ChildWrite ,    BLOCKING_READ,"ChildToParent pipe");
     createChild(ParentWrite,ParentRead,ChildWrite,ChildRead,argv+1);
     s.addSocket(ParentRead);
     printf("child process initialized\n");

     openFirewallPort(PORT_NUMBER);
     printf("firewall opened\n");

     int listen_socket = socketListen(PORT_NUMBER);
     s.addSocket(listen_socket);
     printf("listening for connect requests on port %d\n",PORT_NUMBER);

     while(1)
     {
       printf("waiting for data\n");
       s.waitForData();
       while(socket = s.getNext())
       {
         if (socket == listen_socket)
         {
           /** ready to accept **/
           rc = socketAccept(listen_socket);
           if (rc == -1) break; 
           s.addSocket(rc);
         }
         else if (socket == ParentRead && lastSocket)
         {
        	 switch(r.readFromChild(lastSocket,ParentRead))
             {
                case Reader::SOCKET_ERROR:   s.removeSocket(socket); cleanupSocket(socket,&lastSocket); goto socket_done;
                case Reader::FAR_END_CLOSED: s.removeSocket(socket); cleanupSocket(socket,&lastSocket); goto socket_done;
                case Reader::CHILD_ERROR:    goto quit;
                default: break;
             }
         }
         else
         {
           lastSocket=socket;
           switch (r.readFromSocket(socket,ParentWrite))
           {
              case Reader::SOCKET_ERROR:   s.removeSocket(socket); cleanupSocket(socket,&lastSocket); goto socket_done;
              case Reader::FAR_END_CLOSED: s.removeSocket(socket); cleanupSocket(socket,&lastSocket); goto socket_done;
              case Reader::CHILD_ERROR:    goto quit;
              case Reader::QUIT_REQUESTED: goto quit;
              default: break;
           }
         }
       }
socket_done: rc = 0;
     } 
quit: 
     s.closer();
     closeFirewallPort(PORT_NUMBER);
     return 0; 
}
