#include "reverse_string.h"
#include <string>

namespace reverse_string {

// TODO: add your solution here
    
    std::string reverse_string(std::string msg){

        std::string res{""};
        
        for(int i = msg.length() - 1; i>=0; i--){
            res += msg[i];
        }

        return res;
    }

    
}  // namespace reverse_string
