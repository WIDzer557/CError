package cerror

type Catcher interface {
	Catch(func(interface{}))
}

type EBuffer []interface{}

func Throw(err ...interface{}) Catcher {
	buff := EBuffer(err)
	return &buff
}

func (buff *EBuffer) Catch(catcher func(err interface{})) {
	for _, err := range *buff {
		catcher(err)
	}
}
