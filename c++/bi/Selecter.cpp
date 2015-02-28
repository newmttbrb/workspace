#include <stdio.h>
#include <stdlib.h>
#include <cstring>
#include <string.h>
#include <unistd.h>
#include <sys/ioctl.h>
#include <sys/types.h> 
#include <sys/time.h>
#include <netinet/in.h>
#include <errno.h>
#include <algorithm> 
#include "Selecter.h"

Selecter::Selecter() : max_sd(0), next(0), ready(0) { FD_ZERO(&master); FD_ZERO(&working); }

void Selecter::addSocket(int socket)
{
  FD_SET(socket,&master);
  max_sd =  std::max(max_sd,socket);
}

void Selecter::removeSocket(int socket)
{
  FD_CLR(socket,&master);
  if (socket == max_sd) 
  { 
    while (FD_ISSET(max_sd, &master) == 0) max_sd--; 
  }
}

int Selecter::waitForData()
{
  int rc;

  memcpy(&working, &master, sizeof(master));
  rc = select(max_sd+1, &working, NULL, NULL, NULL);
  if (rc < 0)
  {
    perror("select failed");
    exit(1);
  }
  next = 0;
  ready = rc;
  return rc;
}

int Selecter::getNext()
{
  if (next > max_sd || !ready) return 0;
  while(FD_ISSET(next, &working) == 0) 
  {
    next++;
    if (next > max_sd) return 0;
  }
  ready--;
  return next;
}

void Selecter::closer()
{
 for (int i = 0; i < max_sd; i++)
 {
   if (FD_ISSET(i, &master)) close(i);
 }
}
