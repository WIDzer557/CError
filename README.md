# CError

**Example custom function:**
```go
Throw(Test()).
	Catch(func(err interface{}) {
		if err != nil {
			fmt.Println(err)
		}
})

func Test() error {
	return fmt.Errorf("404 - Page not found..")
}
```

**Example print function:**
```go
Throw(Test()).
	Catch(catcher.Print)

func Test() error {
	return fmt.Errorf("404 - Page not found..")
}
```
