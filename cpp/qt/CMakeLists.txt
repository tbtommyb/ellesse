cmake_minimum_required(VERSION 2.8.11)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall")

if(APPLE AND EXISTS /usr/local/opt/qt5)
	# Homebrew installs Qt5 (up to at least 5.9.1) in
	# /usr/local/qt5, ensure it can be found by CMake since
	# it is not in the default /usr/local prefix.
	list(APPEND CMAKE_PREFIX_PATH "/usr/local/opt/qt")
endif()

# Find includes in corresponding build directories
set(CMAKE_INCLUDE_CURRENT_DIR ON)

# Instruct CMake to run moc automatically when needed.
set(CMAKE_AUTOMOC ON)

# Find the Qt libraries for Qt Quick/QML
find_package(Qt5Core CONFIG REQUIRED Qml Quick Gui)
find_package(Qt5Quick)
find_package(Qt5Qml)
find_package(Qt5Gui)

# add the qml.qrc file
qt5_add_resources(qml_QRC qml.qrc)

# add source files
set(SOURCE
	${SOURCE}
	main.cpp
)

# Tell CMake to create the helloworld executable
add_executable(ellesse-qt ${SOURCE} ${qml_QRC})

# Use the Qml/Quick modules from Qt 5.
target_link_libraries(ellesse-qt Qt5::Qml Qt5::Quick)
