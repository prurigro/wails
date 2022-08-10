package binding

type As struct {
	B Bs `json:"b"`
}

type Bs struct {
	Name string `json:"name"`
}

func (a As) Get() As {
	return a
}

var NestedFieldTest = BindingTest{
	name: "NestedField",
	structs: []interface{}{
		&As{},
	},
	exemptions:  nil,
	shouldError: false,
	want:        ``,
}
