package domain

type Localidade struct {
	name string
}

func NewLocalidade(name string) *Localidade {
	return &Localidade{
		name: name,
	}
}

func (l *Localidade) Name() string {
	return l.name
}
