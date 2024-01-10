# validate

公用参数校验库


示例：

```
go get -u github.com/tristin2024/validate

import "github.com/tristin2024/validate"

if err := validate.Struct(&req); err != nil {
		logs.Std.Warn(err)
		render.ErrBadRequest(c, "参数校验失败")
		return
	}
```

