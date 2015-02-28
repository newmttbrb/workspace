#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <netdb.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>
#include "ManagedChild.h"
#include "Selecter.h"

#define BUFFER_SIZE 2056

int getUserInput(char* buffer, size_t size)
{
	  bzero(buffer,size);
	  int rc = read(0,buffer,size);
	  switch(rc)
	  {
		case -1: if (errno == EWOULDBLOCK) return 0;
			     perror("read from keyboard"); 
			     exit(1);
		case  0: /*printf("0 returned\n");*/ 
		         return 0;
		default: /*printf("read '%s' from user (%d)\n",buffer,rc);*/ 
		         break;
	  }
	  return rc;
}

bool forwardToChild(int writefd, char* buffer, size_t size)
{
	if (size <= 1) { printf("nothing to send to child\n"); return false; }
    //printf("\tsending '%s' to child (%d)\n",buffer,size);
	int rc = write(writefd,buffer,size);
	switch(rc)
	{
	   case -1: perror("write to child"); return false;
	   case  0: 
	   default: /*printf("send to child returned %d\n",rc);*/ return true;
	}
}

void readFromChild(int readfd,char* buffer, size_t size)
{
	  bzero(buffer,size);
	  switch(read(readfd,buffer,size))
	  {
		    case  0: printf("'' read from child"); break;
		    case -1: if (errno == EWOULDBLOCK) return;
			         perror("read from child"); exit(1);
		    default: printf("[%s]\n",buffer); break;
	  }
}

int main(int argc, char ** argv)
{
  char buf[BUFFER_SIZE];
  //Selecter s;

  int ParentWrite,ParentRead;
  int ChildWrite,ChildRead;
  int socket;

  if (argc != 2) { printf("error: please supply a program to wrap\n"); exit(1); }

  createPipe(&ChildRead ,&ParentWrite,BLOCKING_READ,"ParentToChild pipe");
  createPipe(&ParentRead,&ChildWrite ,BLOCKING_READ,"ChildToParent pipe");
  createChild(ParentWrite,ParentRead,ChildWrite,ChildRead,argv+1);

  //s.addSocket(ParentRead);
  //s.addSocket(0);

  while(1)
  {
    if (forwardToChild(ParentWrite,buf,getUserInput(buf,BUFFER_SIZE)))
      readFromChild(ParentRead,buf,BUFFER_SIZE);
  }
}
