# CError

**Example custom function:**
```go
Throw(fmt.Errorf("404 - Page not found..")).
	Catch(func(err interface{}) {
		if err != nil {
			fmt.Printf(err)
		}
})
```
