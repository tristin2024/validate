# validate

公用参数校验库


示例：

```
go get -u git.beeps.cn/jiazujiang/validate

import "git.beeps.cn/jiazujiang/validate"

if err := validate.Struct(&req); err != nil {
		logs.Std.Warn(err)
		render.ErrBadRequest(c, "参数校验失败")
		return
	}
```

