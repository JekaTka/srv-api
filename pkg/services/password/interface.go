package password

//Passwrder interface
type Passwrder interface {
	Generate(raw string) (string, error)
	Compare(p1, p2 string) error
}

type Generate func(raw string) (string, error)
type Compare func(p1, p2 string) error
