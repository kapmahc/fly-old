cmake_minimum_required (VERSION 3.7)
include(ExternalProject)

# -----------------
project (fly)
set (CMAKE_CXX_STANDARD 14)
set (fly_VERSION_MAJOR 1)
set (fly_VERSION_MINOR 0)
#-------------------
# ExternalProject_Add(
# mstch
#
# GIT_REPOSITORY "https://github.com/no1msd/mstch.git"
# GIT_TAG "master"
# UPDATE_COMMAND ""
# INSTALL_COMMAND ""
# PREFIX external
# )

add_subdirectory(external/mstch)
include_directories(${mstch_INCLUDE_DIR})

#-------------------
set (
EXTRA_LIBS
PocoCrypto PocoData PocoJSON PocoFoundation PocoNet PocoUtil PocoXML PocoZip
mstch
)
#-------------------
aux_source_directory (src DIR_SRCS)
add_executable(fly ${DIR_SRCS})
target_link_libraries(fly ${EXTRA_LIBS})
