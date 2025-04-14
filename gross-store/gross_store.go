package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{}
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exist := units[unit]
	if !exist {
		return false
	}
	quantity, exist := bill[item]

	if !exist {
		bill[item] = units[unit]
	} else {
		bill[item] = quantity + units[unit]
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitQuant, exist := units[unit]
	if !exist {
		return false
	}
	quantity, exist := bill[item]

	if !exist || quantity < 0 {
		return false
	}
	if quantity-unitQuant == 0 {
		delete(bill, item)
	}
	if quantity-unitQuant < 0 {
		return false
	}
	if quantity-unitQuant > 0 {
		bill[item] = quantity - unitQuant
	}
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	currItem, exist := bill[item]
	if !exist {
		return 0, false
	}
	return currItem, true
}
