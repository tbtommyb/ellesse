cmake_minimum_required(VERSION 2.8.11)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++1z -Wall")

find_package(Boost COMPONENTS system filesystem REQUIRED)
include_directories( ${Boost_INCLUDE_DIR} )

add_library(Ellesse src/ellesse.cpp src/file.cpp)
target_include_directories(Ellesse PUBLIC include)
