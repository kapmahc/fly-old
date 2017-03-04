cmake_minimum_required (VERSION 3.7)
project (fly)
set (fly_VERSION_MAJOR 1)
set (fly_VERSION_MINOR 0)

set (EXTRA_LIBS PocoCrypto PocoData PocoJSON PocoFoundation PocoNet PocoUtil PocoXML)

aux_source_directory (src DIR_SRCS)
add_executable(fly ${DIR_SRCS})
target_link_libraries(fly ${EXTRA_LIBS})
