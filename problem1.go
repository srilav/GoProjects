package problem1

func doComplex() (string, error) {
	return "Success", nil
}

func main() {
	var val string
	num := 32

	switch num {
	case 4:
	// do nothing
	case 16:
	// do nothing
	case 32:
		val, err := doComplex()
		if err != nil {
			panic(err)
		}
		if val == "" {
			// do something else
		}
	case 64:
		// do nothing
	}
	fmt.Println(val)
}
