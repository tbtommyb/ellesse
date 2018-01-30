#include <string>
#include <iostream>
#include <iomanip>
#include <boost/format.hpp>
#include "src/ellesse.hpp"

int main(int argc, char* argv[])
{
    fs::path p{};
    if(argc !=2 ) {
        p = fs::current_path();
    } else {
        p = fs::path{argv[1]};
    }

    try {
        auto results = Ellesse{p};
        for(auto& r : results.list()) {
            std::cout << boost::format("%3% %1% %|60t|%2%\n") % r.path % r.size % r.mode;
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
