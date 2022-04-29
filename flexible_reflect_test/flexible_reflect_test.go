package testing

import (
	"errors"
	//"fmt"
	"reflect"
	"testing"
)


type  Employee struct {
	EmployeeID string
	Name  string  // `format:"normal"`
	Age  int
}


func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}


type Customer struct {
	CookieID string
	Name string
	Age int
}


func fillBySettings(st interface{}, settings map[string]interface{} ) error {


	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		// Elem() 获取指针指向的值
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
	}
	if settings == nil {
		return errors.New("settings is nil.")
	}

	var (
		field reflect.StructField
		ok bool
	)

	for k, v := range settings {

		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
	}

	if field.Type == reflect.TypeOf(v) {
		vstr := reflect.ValueOf(st)
		vstr = vstr.Elem()
		vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}



func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{} {"Name": "Mike", "Age": 40}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)

	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}