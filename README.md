go-wcwidth
==========

### Install

```
go get github.com/shinichy/go-wcwidth
```

### Example
```
import (
	"github.com/shinichy/go-wcwidth"
)

wcwidth.Wcwidth('あ') // returns 2
wcwidth.Wcswidth("aあ★")) // returns 5
```

### License
BSD-like (2-clause, see wcwidth.go)
