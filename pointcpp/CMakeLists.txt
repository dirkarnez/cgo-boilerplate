cmake_minimum_required(VERSION 3.14)

project(pointcpp)
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_POSITION_INDEPENDENT_CODE ON)
SET(CMAKE_EXE_LINKER_FLAGS "${CMAKE_EXE_LINKER_FLAGS} -static")

add_library(${PROJECT_NAME} STATIC 
    "include/point_c_wrapper.h"
    "src/point.hpp"
    "src/point_c_impl.cpp"
    "src/point.cpp"
)

target_include_directories(${PROJECT_NAME}
    PUBLIC
    "include"
)