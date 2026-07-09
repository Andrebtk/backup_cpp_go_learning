#include "vehicle_purchase.h"

namespace vehicle_purchase {

// needs_license determines whether a license is needed to drive a type of
// vehicle. Only "car" and "truck" require a license.
bool needs_license(std::string kind) {

    if (kind == "car" || kind == "truck"){
        return true;
    }
    return false;
}

// choose_vehicle recommends a vehicle for selection. It always recommends the
// vehicle that comes first in lexicographical order.
std::string choose_vehicle(std::string option1, std::string option2) {
    // TODO: Return the final decision in a sentence.
    std::string res{" is clearly the better choice."};

    if(option1 < option2) {
        return option1 + res;
    }
    return option2 + res;

    
}

// calculate_resell_price calculates how much a vehicle can resell for at a
// certain age.
double calculate_resell_price(double original_price, double age) {
    if(age < 3) {
        return original_price - (original_price * (1.0 - (80.0/100)));
    }else if(age >= 10){
        return original_price - (original_price * (1.0 - (50.0/100)));
    } else {
        return original_price - (original_price * (1.0 - (70.0/100)));
    }
}

}  // namespace vehicle_purchase
