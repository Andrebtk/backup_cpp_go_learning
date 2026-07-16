package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	hashMap := map[string]int{}
    hashMap["quarter_of_a_dozen"] = 3
    hashMap["half_of_a_dozen"] = 6
    hashMap["dozen"] = 12
    hashMap["small_gross"] = 120
    hashMap["gross"] = 144
    hashMap["great_gross"] = 1728
    return hashMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	val, exists := units[unit]
    if !exists {
        return false
    }

	_, ex := bill[item]
    if !ex {
        bill[item] = val
    } else {
        bill[item] += val
    }
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	val, exists := bill[item]
    if !exists {
        return false
    }

    val2, exists2 := units[unit]
    if !exists2 {
        return false
    }

	if val - val2 < 0 {
        return false
    }

	if val - val2 == 0 {
        delete(bill, item)
        return true
    }

    bill[item] -= val2
    return true

    
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	val, exists := bill[item]
    if !exists {
        return 0, false
    }

    return val, true
}
