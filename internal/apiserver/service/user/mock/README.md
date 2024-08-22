# 通过mockgen生成

- 命令可用flag
  - `-source` 指定需要mock的接口文件所在位置
  - `-destination` 指定mock生成出来的文件存放位置
  - `-package` 指定mock文件使用的包名
  - `-import` 设置mock文件依赖的包
  - `-build_flags` 传递给build的参数

具体使用参考：https://github.com/golang/mock

```shell
mockgen -destination internal/apiserver/service/user/mock/mock_user.go -package user github.com/rppkg/godfrey/internal/apiserver/service/user IService
```