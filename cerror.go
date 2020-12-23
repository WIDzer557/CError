package cerror

type EBuffer []interface{}

func Throw(err ...interface{}) *EBuffer {
	buffer := EBuffer(err)
	return &buffer
}

func (buffer *EBuffer) Catch(catcher func(err interface{})) {
	for _, err := range *buffer {
		catcher(err)
	}
}
