package entity

type Person struct {
	Name string `json:"person_name"`
	Id   int64  `json:"person_id"`
	Age  int64  `json:"person_age"`
}

func (p *Person) Validator() {

}
