package ttest

import "fmt"

func TestSwitch() (err error) {
	switch {
	case 1 == 1:
		fmt.Println("==========1")
	case 2 == 2:
		fmt.Println("==========2")
	default:
		fmt.Println("==========default")
	}
	fmt.Println(">>>>>>>>>>")
	return nil
}
