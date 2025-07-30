package decoder

const propertiesSize = 26

type Property struct {
	Index   int
	Name    string
	Decoder Decoder
}

type Properties struct {
	primary [propertiesSize]Property
	other   []Property
}

func (p *Properties) Add(prop Property) {
	key := prop.Name[0] % propertiesSize
	if len(p.primary[key].Name) == 0 {
		p.primary[key] = prop
		return
	}
	p.other = append(p.other, prop)
}

func (p *Properties) Find(name string) Property {
	key := name[0] % propertiesSize
	prop := p.primary[key]
	if prop.Name == name {
		return prop
	}
	for _, prop := range p.other {
		if prop.Name == name {
			return prop
		}
	}
	return Property{}
}
