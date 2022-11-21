# whosbug webservice
## main.go

项目入口

## docs

项目所有相关文档

## api

api入口，相当于controllers

调用底层代码(对数据库进行处理的代码），对业务进行处理

### v1

重构后的版本命名为v1

#### users

对账号进行处理

#### commits

对C插件解析出的git commits历史信息进行处理

#### objects

对传入的objects(函数进行处理，计算出objects在语法定义链上的nodes

#### owners

计算出最后的责任人

## models

与数据库进行交互的函数(提供给api调用）、项目中使用的数据结构 middlewears

中间件(JWT登录验证）

## routes

路由组

### router

路由器，路由入口，根据URL调用api接口

## utils

工具包，公共功能(函数）

### setting.go

从config.ini配置文件读取参数

### errorMsg.go

约定业务处理中的错误码

## configs

配置文件目录

### config.ini

ini配置文件(数据库，网站的配置参数等）

## tests

测试， 目前有Apifox作为代替

## third_party

第三方库，swagger等

## build

docker配置，打包时运行的脚本等

## README.md

- 整个项目的运行，部署方法

- 采用的技术栈介绍，使用的组件的版本

- 项目作者
