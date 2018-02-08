cmake_minimum_required(VERSION 3.10)

project("Ellesse")

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++1z -Wall")
add_subdirectory(lib)

find_package(Boost COMPONENTS system filesystem REQUIRED)

add_executable(ellesse main.cpp)

include_directories( ${Boost_INCLUDE_DIR} )

target_include_directories(ellesse PUBLIC lib/include)
target_link_libraries(ellesse ${Boost_LIBRARIES} Ellesse)

