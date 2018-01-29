#ifndef __ellesse__
#define __ellesse__

#include <stdexcept>
#include <boost/filesystem.hpp>
#include <string>

namespace fs = boost::filesystem;

class Ellesse {
public:
    Ellesse(fs::path p);
    ~Ellesse() = default;
    std::vector<fs::path> list();
private:
    fs::path query;
};

#endif
