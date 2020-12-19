package cerror

type EBuff []interface{}

type Catcher interface {
	Catch(func(interface{}))
}

func Throw(err ...interface{}) Catcher {
	buff := EBuff(err)
	return &buff
}

func (buff *EBuff) Catch(catcher func(err interface{})) {
	for _, err := range *buff {
		catcher(err)
	}
}
