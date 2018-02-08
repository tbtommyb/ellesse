#ifndef __ellesse__
#define __ellesse__

#include <boost/filesystem.hpp>
#include "file.hpp"

namespace fs = boost::filesystem;

class Ellesse {
public:
    Ellesse(fs::path p);
    ~Ellesse() = default;
    std::vector<File> items();
private:
    fs::path query;
};

#endif
