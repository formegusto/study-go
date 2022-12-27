package condition

func CanIDrink(age int) bool {
	if age < 18 {
		return false
	}
	return true 
}

func VECanIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true 
}

func SwitchCanIDrink(age int) bool {
	// switch age {
	// case 10:
	// 	return false
	// case 18:
	// 	return true
	// }

	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}

	switch koreanAge:=age+2; koreanAge{
	case 10:
		return false
	case 18:
		return true
	}
	return false 
}