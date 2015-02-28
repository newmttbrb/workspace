#include <unistd.h>
#include <sys/types.h>
#include <errno.h>
#include <stdio.h>
#include <sys/wait.h>
#include <stdlib.h>

int var_glb; /* A global variable*/

int main(void)
{
	pid_t childPID;
	childPID = fork();
	if(childPID >= 0) // fork was successful
	{
		if(childPID == 0) // child process
		{
			char *env[] = {(char *)"FAKE=/var", NULL};
			execle("/bin/ls", "ls", NULL, env);
			return 0;
		}
		else //Parent process
		{
			/*** do nothing here ***/
			printf("fork succeeded\n");
			return 0;
		}
	}
	else // fork failed
	{
		printf("\n Fork failed, quitting!!!!!!\n");
		return 1;
	}
}
