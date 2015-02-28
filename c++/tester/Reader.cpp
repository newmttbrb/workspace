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
#include "Reader.h"

/***
<?xml version="1.0"?>
<commands>
  <enableAutomation autoFormat="true" cpus="0 1"/>
  <pressButton cpuid="0" button="12" duration="200"/>
  <pause duration=\"800\"/>
  <disableAutomation/>
  <returnQueuedMessages/>
</commands>
***/

/*Turret(:outgoing,"10.204.45.168")
Press(:ptu,:release,8)*/


Reader::READER_RESULTS Reader::readFromSocket(int socket,int writer)
{ 
	int rc;

    bzero(buffer,BUFFER_SIZE);
	rc = read(socket,buffer,BUFFER_SIZE);
    if (rc < 0) 
    {
        if (errno != EWOULDBLOCK) { perror("reading from far end"); return SOCKET_ERROR; }
        return NO_MORE_DATA;
    }
    if (!rc) { printf("far end closed?"); return FAR_END_CLOSED; }

    if (!strncmp(buffer,"quit",rc)) return QUIT_REQUESTED;

    printf("--> '%s'\n",buffer);

    if (write(writer,buffer,rc) < 0) { perror("write to child failed"); return CHILD_ERROR; }
                       
    return SUCCESS;
}

Reader::READER_RESULTS Reader::readFromChild(int socket,int readFromChild)
{
	int rc;

    bzero(buffer,BUFFER_SIZE);
	rc = read(readFromChild,buffer,BUFFER_SIZE);
    if (rc < 0)
    {
        if (errno != EWOULDBLOCK) { perror("ERROR reading from child"); return CHILD_ERROR; }
        return NO_MORE_DATA;
    }
    if (!rc) { printf("child closed?"); return FAR_END_CLOSED; }

    printf("<-- %s\n",buffer);

    if (write(socket,buffer,rc) < 0) { perror("write to far end failed"); return SOCKET_ERROR; }

    return SUCCESS;
}


