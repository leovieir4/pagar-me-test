package parameter

import "reflect"

type Person struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type Parameter struct {
	Person Person `json:"person"`
}

type RelashionParameter struct {
	ParentId int64 `json:"parent_id"`
	ChildId  int16 `json:"child_id"`
}

type BaconNumberAndKinshipParameter struct {
	FirstPersonId  int64 `json:"first_person_id"`
	SecondPersonId int64 `json:"second_person_id"`
}

func (p Parameter) Validator() bool {

	values := reflect.ValueOf(p)

	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).Kind() == reflect.String {
			continue
		}
		if !validator(values.Field(i)) {
			return false
		}
	}
	return true
}

func validator(values reflect.Value) bool {
	for i := 0; i < values.NumField(); i++ {

		if values.Field(i).Kind() == reflect.Int64 {
			if values.Field(i).Int() == 0 {
				return false
			}
			continue
		}
		if values.Field(i).Len() == 0 {
			return false
		}
	}
	return true
}
