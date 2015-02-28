#include <sys/select.h>

class Selecter
{
public:
  Selecter();

  void addSocket(int socket);
  void removeSocket(int socket);
  int waitForData();
  int getNext();
  void closer();

private:
  fd_set master, working;
  int max_sd;
  int next;
  int ready;
};