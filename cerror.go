package cerror

type EBuffer []interface{}

type Catcher interface {
	Catch(func(interface{}))
}

func Throw(err ...interface{}) Catcher {
	buff := EBuffer(err)
	return &buff
}

func (buff *EBuffer) Catch(fn func(err interface{})) {
	for _, err := range *buff {
		fn(err)
	}
}
