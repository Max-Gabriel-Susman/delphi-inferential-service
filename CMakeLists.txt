# CMakeLists.txt
cmake_minimum_required(VERSION 3.10)

project(MyProject)

# Google Test
add_subdirectory(googletest)

# Your main executable
add_executable(my_program main.cpp)

# Your test executable
add_executable(tests test.cpp)

# Link your main program against Google Test
target_link_libraries(my_program gtest_main)

# Link your tests against Google Test
target_link_libraries(tests gtest_main)
