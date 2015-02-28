#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <netdb.h>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <errno.h>
#include <signal.h>
#include <fcntl.h>
#define _MANAGED_CHILD_CPP_
#include "ManagedChild.h"

void createPipe(int* readSide, int* writeSide,bool readNonBlock,const std::string& pipeName)
{
	int pipefd[2];

	if (pipe(pipefd) < 0) { perror(pipeName.c_str()); exit(1); }
	*readSide = pipefd[0];
	*writeSide = pipefd[1];
	if (readNonBlock) fcntl(*readSide, F_SETFL, fcntl(*readSide, F_GETFL) | O_NONBLOCK);
}

int createChild(int ParentWrite, int ParentRead, int ChildWrite, int ChildRead, char ** argv)
{
	int pid;

	switch( pid = fork() )
	{
	    case -1:
	        perror("fork");
	        exit(1);

	    case 0:
	        dup2(ChildWrite,STDOUT_FILENO);
	        dup2(ChildWrite,STDERR_FILENO);
	        dup2(ChildRead ,STDIN_FILENO );
	        close(ParentRead);
	        close(ParentWrite);

	        execvp(argv[0], argv);

	        perror("exec");
	        kill(getppid(), SIGQUIT);
	        exit(1);

	    default:
	        close(ChildRead);
	        close(ChildWrite);
	        return pid;
	}
}

