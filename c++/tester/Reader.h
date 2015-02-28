#define BUFFER_SIZE 10240

class Reader
{
public:
  enum READER_RESULTS { FAR_END_CLOSED, SOCKET_ERROR, QUIT_REQUESTED, NO_MORE_DATA, SUCCESS, CHILD_ERROR };

  Reader() { }

  READER_RESULTS readFromSocket(int socket,int writeToChild);
  READER_RESULTS readFromChild(int socket,int readFromChild);

private:
  char buffer[BUFFER_SIZE];
};

