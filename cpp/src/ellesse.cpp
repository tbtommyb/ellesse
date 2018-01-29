#include "ellesse.hpp"
#include <string>

namespace fs = boost::filesystem;

Ellesse::Ellesse(fs::path p)
{
    fs::file_status s = fs::status(p);
    if(!fs::is_directory(s)) {
        throw std::invalid_argument{p.string()};
    }
    this->query = p;
};

std::vector<fs::path> Ellesse::list()
{
    std::vector<fs::path> results;
    for(auto i = fs::directory_iterator(this->query); i != fs::directory_iterator{}; i++) {
        results.push_back(i->path());
    }
    return results;
};
