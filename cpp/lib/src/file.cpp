#include <iostream>
#include "file.hpp"

File::File(fs::path p) : path(p)
{
    pathName = path.filename().string();
    fs::file_status s = fs::status(path);

    if(fs::is_directory(path)) {
        size = 0;
        isDirectory = true;
    } else {
        size = fs::file_size(path);
        isDirectory = false;
    }

    mode = s.permissions();
};
