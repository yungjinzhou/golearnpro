package testing

import (
	"fmt"
	"testing"
)




func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i:=0;i< len(inputs); i++ {
		ret :=  square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, the expected is %d, the actual %d", inputs[i], expected[i], ret)
		}
		}
	}

func TestErrorInCode(t *testing.T){
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
}


func TestFalilInCode(t *testing.T){
	fmt.Println("Start")
	t.Fatal("Fail")
	fmt.Println("End")
}

