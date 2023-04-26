# ORMTools

ORMTools是一款可以将数据库表结构转化为Golang Struct结构的可视化工具，支持ssh连接，可自定义设置字段转换规则，支持GORM Tag、Json Tag，支持部分GORM模型约定......

[![Go Report Card](https://goreportcard.com/badge/github.com/louismax/ORMTools)](https://goreportcard.com/report/github.com/louismax/ORMTools)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/louismax/ORMTools)
![GitHub package.json version (subfolder of monorepo)](https://img.shields.io/github/package-json/v/louismax/ORMTools?filename=frontend%2Fpackage.json)
![GitHub](https://img.shields.io/github/license/louismax/ORMTools)
![GitHub language count](https://img.shields.io/github/languages/count/louismax/ORMTools?style=flat-square)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/louismax/ORMTools?style=flat-square)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/louismax/ORMTools)
![GitHub commit activity](https://img.shields.io/github/commit-activity/m/louismax/ORMTools)
![GitHub last commit](https://img.shields.io/github/last-commit/louismax/ORMTools)
![GitHub Release Date](https://img.shields.io/github/release-date/louismax/ORMTools)

## webview2运行时依赖
ORMTools需要依赖[Microsoft WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/)运行时，默认情况下，Windows 11 会安装它，但有些机器不会。
在未检测到合适的运行时的时候，应用会主动下载并运行WebView2运行时引导程序
![](https://raw.githubusercontent.com/louismax/img/master/PicGo/20230426113926.png)
![](https://raw.githubusercontent.com/louismax/img/master/PicGo/20230426113935.png)

## 应用功能及配置
- 自定义代码风格
- 自定义忽略显示哪些数据库名（默认忽略mysql系统数据库information_schema、performance_schema 、mysql 、sys）
- 自定义忽略显示数据表名
- 库表结构树表comment的显示控制
- 自定义表名（GORM实现Tabler接口更改默认蛇形复数表名）
- JSON标签显示控制
- Grom Column标签显示控制
- 自定义忽略显示字段名
- 自定义数据库字段类型转换go struct的转换规则(应用默认规则可通过“引入全部规则”功能查看)
- 通过SSH跳板连接数据库
- 导出连接配置
- 导入连接配置

![](https://raw.githubusercontent.com/louismax/img/master/PicGo/cfg%20(1).gif)

![](https://raw.githubusercontent.com/louismax/img/master/PicGo/ipt%20(1).gif)

![](https://raw.githubusercontent.com/louismax/img/master/PicGo/20230426143136.png)

![](https://raw.githubusercontent.com/louismax/img/master/PicGo/dbtree%20(1).gif)
