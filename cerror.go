package cerror

type EBuffer []interface{}

func Throw(err ...interface{}) *EBuffer {
	buff := EBuffer(err)
	return &buff
}

func (buff *EBuffer) Catch(catcher func(err interface{})) {
	for _, err := range *buff {
		catcher(err)
	}
}
