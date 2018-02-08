#ifndef __ellesse_file__
#define __ellesse_file__

#include <boost/filesystem.hpp>

namespace fs = boost::filesystem;

struct File {
    File(fs::path p);
    fs::path path;
    std::string pathName;
    unsigned int size;
    enum fs::perms mode;
    bool isDirectory;
};

#endif
