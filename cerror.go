package cerror

type EBuffer []interface{}

type Catcher interface {
	Catch(func(err interface{}))
}

func Throw(err ...interface{}) Catcher {
	buffer := EBuffer(err)
	return &buffer
}

func (buffer *EBuffer) Catch(catcher func(err interface{})) {
	for _, err := range *buffer {
		catcher(err)
	}
}
