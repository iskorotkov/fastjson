package encoder

type Property struct {
	Index           int
	Name            string
	PrecomputedName []byte
	Encoder         Encoder
	OmitEmpty       bool
}

type Properties []Property

func precomputeFieldName(name string) []byte {
	buf := make([]byte, 0, len(name)+3)
	buf = append(buf, '"')
	buf = append(buf, name...)
	buf = append(buf, '"', ':')
	return buf
}
