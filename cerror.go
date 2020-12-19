package cerror

type Catcher interface {
	Catch(func(interface{}))
}

type EBuff []interface{}

func Throw(err ...interface{}) Catcher {
	buff := EBuff(err)
	return &buff
}

func (buff *EBuff) Catch(catcher func(err interface{})) {
	for _, err := range *buff {
		catcher(err)
	}
}
