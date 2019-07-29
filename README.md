> #### 软件架构
> 目前支持单web架构，如果部署成前后端分离，可用nginx中间件代理(已添加跨域访问设置)。
>    * 采用了Casbin做Restful的rbac权限控制；
>    * 采用jwt做用户认证、回话控制；
>    * 采用Mysql+xorm做持久层；
>    * 采用redis做tocken缓存；
>    * Vue前端项目；

***
#### 项目目录结构
```
perm
    +-- perm-server 后端服务
        |   +-- doc 说明文档目录
            |       +-- sql sql文件目录（包含初始化sql）
        |   +-- main 项目启动入口（包含main.go）
        |   +-- sys_common 公共服务目录
            |       +-- caches 缓存工具目录
            |       +-- conf 配置文件目录
            |       +-- db 初始化数据库连接
            |       +-- inits 初始化系统用户，角色目录
            |       +-- middleware 中间件目录
            |       +-- models 公共数据库实体模型
            |       +-- supports 系统常量目录
            |       +-- utils 后端工具目录
            |       +-- vo 页面所需公共数据类型目录
        |   +-- sys_base 业务基础服务（包含：用户，菜单，角色，部门）
                |       +-- models 数据库对应实体模型文件
                |       +-- routes 请求路由处理目录
                |       +-- service 业务逻辑处理目录
                |       +-- vo 业务所需数据类型目录
    +-- perm-vue 前端页面
        |       +-- build 构建静态文件目录
        |       +-- config 前端配置文件目录
        |       +-- src 前端资源文件
                |       +-- components 页面目录
                |       +-- router 前端路由目录
                |       +-- utils 前端工具目录
        |       +-- static 前端静态文件目录

### 服务启动
1. rizla main.go （go get -u github.com/kataras/rizla  热启动插件）


> * 超级管理员登录：
>    > 初始账号：root
>    > 初始密码: 123456


