#specify the version being used aswell as the language
cmake_minimum_required(VERSION 2.6)
#Name your project here
project(mattomation)

add_executable(mattomation main.cpp tcp.cpp Selecter.cpp Reader.cpp ManagedChild.cpp Firewall.cpp)
