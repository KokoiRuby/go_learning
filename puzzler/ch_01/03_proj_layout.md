### [Layout](https://github.com/golang-standards/project-layout?tab=readme-ov-file)

- `cmd/`：存放应用程序的主要入口点，每个子目录通常对应一个可执行程序。
- `pkg/`：存放可供其他项目导入和使用的包代码。
- `internal/`：存放项目内部使用的包代码，限制了包的可见性，只能在项目内部使用。
- `api/`：存放 API 相关的代码和路由处理逻辑。
- `web/`：存放与 Web 相关的静态文件、模板等。
- `configs/`：存放配置文件，如环境配置、数据库连接等。
- `scripts/`：存放构建、运行、部署等脚本。
- `tests/`：存放测试文件，用于编写单元测试和集成测试。
- `README.md`：项目的说明文档，包含项目的概述、使用方法等信息。

```bash
myproject/
  |- cmd/
  |  |- myapp/
  |  |  |- main.go
  |
  |- pkg/
  |  |- mypackage/
  |  |  |- mypackage.go
  |
  |- internal/
  |  |- myinternalpackage/
  |  |  |- myinternalpackage.go
  |
  |- api/
  |  |- api.go
  |
  |- web/
  |  |- static/
  |  |- templates/
  |
  |- configs/
  |  |- config.yaml
  |
  |- scripts/
  |  |- build.sh
  |  |- run.sh
  |
  |- tests/
  |  |- mypackage_test.go
  |
  |- README.md
```

