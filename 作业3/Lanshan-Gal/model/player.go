package model
type people interface{
	Add(name string, description string)
}
type Player struct {
	Name string
	ShortDescription string
}

func (m *Player) Add(name string, description string){
	m.Name = name
	m.ShortDescription = description
}