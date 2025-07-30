package encoder

type Property struct {
	Index   int
	Name    string
	Encoder Encoder
}

type Properties []Property
