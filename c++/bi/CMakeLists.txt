#specify the version being used aswell as the language
cmake_minimum_required(VERSION 2.6)
#Name your project here
project(wrapper)

#This tells CMake to fib.c and name it fibonacci
add_executable(wrapper client.cpp ManagedChild.cpp Selecter.cpp)
