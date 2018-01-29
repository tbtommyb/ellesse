#include <string>
#include <iostream>
#include "src/ellesse.hpp"

int main(int argc, char* argv[])
{
    std::string p{};
    if(argc !=2 ) {
        p = ".";
    } else {
        p = argv[1];
    }
    auto results{Ellesse::list(p)};
    for (auto& r : results) {
        std::cout << r.filename().string() << std::endl;
    }
}
