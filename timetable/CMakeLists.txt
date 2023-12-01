cmake_minimum_required(VERSION 3.26)
project(timetable)

set(CMAKE_CXX_STANDARD 17)
include_directories(/opt/homebrew/include)
include_directories(/opt/homebrew/include/mongocxx/v_noabi)
include_directories(/opt/homebrew/include/bsoncxx/v_noabi)
include_directories(/opt/homebrew/include/bsoncxx/v_noabi/bsoncxx)
add_executable(
        timetable cmd/timetable/application.cpp
        cmd/timetable/controllers/system_controller.h
        cmd/timetable/controllers/timetable_controller.h

)
