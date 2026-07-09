#include <string>

namespace log_line {
    std::string message(std::string line) {
        // return the message
        int index = line.find(": ");
        return line.substr(index + 2);
    }

    std::string log_level(std::string line) {
        // return the log level

        int endIndex = line.find("]");
        return line.substr(1, (endIndex - 1));
    }

    std::string reformat(std::string line) {
        // return the reformatted message
        std::string mess = message(line);
        std::string level = log_level(line);

        return mess + " (" + level + ")";
    }
}  // namespace log_line
