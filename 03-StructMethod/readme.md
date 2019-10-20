# Methods

Note - One thing you should note when working with methods is that, like functions, they create **copies of the arguments passed into it**. To avoid this, we can use pointer receivers when defining our methods.

```go
func (pointer *Pointer) myMethod()
```

* Struct Method Intermediate