#include <string>

#define BLOCKING_READ     false
#define NON_BLOCKING_READ true


#ifdef _MANAGED_CHILD_CPP_

void createPipe(int* readSide, int* writeSide,bool readNonBlock,const std::string& pipeName);
int createChild(int ParentWrite, int ParentRead, int ChildWrite, int ChildRead, char ** argv);

#else

extern void createPipe(int* readSide, int* writeSide,bool readNonBlock,const std::string& pipeName);
extern int createChild(int ParentWrite, int ParentRead, int ChildWrite, int ChildRead, char ** argv);

#endif
