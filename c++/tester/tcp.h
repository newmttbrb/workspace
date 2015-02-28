#ifdef _TCP_C_

int socketListen(int port);
int socketAccept(int server_socket);

void error(const char *msg);

#else

extern int socketListen(int port);
extern int socketAccept(int server_socket);

extern void error(const char *msg);

#endif
