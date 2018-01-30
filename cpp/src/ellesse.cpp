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

std::vector<File> Ellesse::list()
{
    std::vector<File> results;
    for(auto i = fs::directory_iterator(this->query); i != fs::directory_iterator{}; i++) {
        fs::path p = i->path();
        unsigned long size;
        if(fs::is_directory(p)) {
            size = 0;
        } else {
            size = fs::file_size(p);
        }
        File f = { p.string(), size, "rwx" };
        results.push_back(f);
    }
    return results;
};
