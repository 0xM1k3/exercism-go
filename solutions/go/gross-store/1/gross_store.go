package gross

//import "fmt"


// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    units := map[string]int {
    	"quarter_of_a_dozen": 3,
        "half_of_a_dozen": 6,
        "dozen": 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    value, exists := units[unit]
    if !exists {
        return false
    }
    
	bill[item] += value
    return true
    
	panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	
    currentQty, itemExists := bill[item]
    if !itemExists {
        return false
    }

    unitValue, unitExists := units[unit]
    if !unitExists {
        return false
    }

    newQty := currentQty - unitValue
	if newQty < 0 {
        return false
    }
    
	if newQty == 0 {
        delete(bill, item)
    } else {
        bill[item] = newQty
    }

    return true
	panic("Please implement the RemoveItem() function")
}


// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    currentQty, itemExists := bill[item]
    return currentQty, itemExists
	panic("Please implement the GetItem() function")
}
