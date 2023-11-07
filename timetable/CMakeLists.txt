cmake_minimum_required(VERSION 3.26)
project(timetable)

set(CMAKE_CXX_STANDARD 17)

add_executable(
        timetable cmd/timetable/application.cpp
)
