#include <stdio.h>
#include <cstdlib>
#include <iostream>
#include <string>

int main() 
{
  FILE *fp;
  char path[10240];
  std::string input;

  while(1)
  {
    std::getline(std::cin,input); 
    if (input.empty()) continue;
    printf("received '%s'\n",input.c_str());
    fp = popen(input.c_str(),"r");
    if (fp == NULL) perror("popen");
    else {
      while (fgets(path, 10240, fp) != NULL) printf("%s", path);
      if (pclose(fp) == -1) perror("pclose");
    }
  }
}

