### vue-typescript-beego-admin

    基于vue-cli4、vue、element和beego实现，前后端权限动态配置后台管理基础模板。

### 前端 vue-cli4 生成项目 -> client

#### Project setup
```
yarn install
```

#### Compiles and hot-reloads for development
```
yarn serve
```

#### Compiles and minifies for production
```
yarn build
```

#### Lints and fixes files
```
yarn lint
```

#### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).

基于以下：
    1. https://github.com/PanJiaChen/vue-element-admin
    2. https://github.com/Armour/vue-typescript-admin-template
2个开源项目修改

### 功能
1. 前后端权限管理通过后台控制，前端路由根据登录账号角色动态添加路由
2. element分页抽离二次封装组件
3. 前端路由添加分为公共路由和权限路由
    I. src\router\index.ts 公共路由
    II. src\router\authRouter.ts 权限路由前端只配置component,详细路由信息配置设置：【权限管理-权限因子-打开前端配置】

### 模块
    1. 后台管理员
    2. 权限管理->管理员、角色、权限因子
    3. 角色管理

### beego 后台管理Api接口快速开发 See(https://beego.me/).
    1. 路由

        后台管理路由 routers admin.go

    2. 接口

        后台管理api controllers admin_controllers

### 请求地址

    http://localhost:8000
