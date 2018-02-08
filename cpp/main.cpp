#include <string>
#include <iostream>
#include <iomanip>
#include <boost/format.hpp>
#include "ellesse.hpp"

int main(int argc, char* argv[])
{
    fs::path p{};
    if(argc !=2 ) {
        p = fs::current_path();
    } else {
        p = fs::path{argv[1]};
    }

    try {
        auto dir = Ellesse::Ellesse{p};
        for(const auto& f : dir.items()) {
            std::cout << boost::format("%1$o %2$-20s %3$1s %4$d\n")
                % f.mode
                % f.pathName
                % (f.isDirectory ? "D" : "")
                % f.size;
        }
        return 0;
    } catch(const std::invalid_argument& ia) {
        std::cerr << ia.what() << " is not a valid directory" << std::endl;
        return 1;
    } catch(const fs::filesystem_error& e) {
        std::cerr << e.what() << " is not a valid directory" << std::endl;
        return 1;
    }
}
