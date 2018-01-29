#include "ellesse.hpp"
#include <string>

namespace fs = boost::filesystem;

std::vector<fs::path> Ellesse::list(std::string path)
{
    std::vector<fs::path> results;
    for(auto i = fs::directory_iterator(path); i != fs::directory_iterator{}; i++) {
        if(!fs::is_directory(i->path())) {
            results.push_back(i->path());
        }
    }
    return results;
}
