cmake_minimum_required(VERSION 3.10)

add_library(Ellesse src/ellesse.cpp src/file.cpp)
target_include_directories(Ellesse PUBLIC include)
