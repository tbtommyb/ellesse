#include <boost/filesystem.hpp>
#include <string>

namespace fs = boost::filesystem;

class Ellesse {
public:
    Ellesse() = default;
    ~Ellesse() = default;
    static std::vector<fs::path> list(std::string path);
};
