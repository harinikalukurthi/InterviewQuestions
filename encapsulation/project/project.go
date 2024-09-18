package project

type Person struct {
    name string // name is unexported (private)
    Age  int    // Age is exported (public)
}

// Public method to set the name
func (p *Person) SetName(name string) {
    p.name = name
}

// Public method to get the name
func (p *Person) GetName() string {
    return p.name
}