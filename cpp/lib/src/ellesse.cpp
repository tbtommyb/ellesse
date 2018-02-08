#include <string>
#include "ellesse.hpp"

namespace fs = boost::filesystem;

Ellesse::Ellesse(fs::path p)
{
    fs::file_status s = fs::status(p);
    if(!fs::is_directory(s)) {
        throw std::invalid_argument{p.string()};
    }
    this->query = p;
};

std::vector<File> Ellesse::items()
{
    std::vector<File> results;
    for(auto i = fs::directory_iterator(this->query); i != fs::directory_iterator{}; i++) {
        fs::path p = i->path();

        const File f{p};
        results.push_back(f);
    }
    return results;
};
