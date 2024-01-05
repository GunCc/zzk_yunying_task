# zzk_yunying_task
云影天光后端开发实习生笔试题


# 目录介绍

| 目录名 | 作用 | 
| -- | -- |
| api | 存放路由所调用的API |
| config | 存放与配置相关的模型 | 
| core | 存放核心代码 | 
| global | 全局变量 | 
| initialize | 初始化函数 |
| middleware | 存放中间件 |
| model | 对象模型 | 
| router | 存放相关路由 |
| service | 存放控制器（操作数据库等）|
| utils | 存放工具类 |
| test | 测试文件 |


# 特殊文件介绍

`config.yaml` 设置配置相关数据的文件
`main.go` 程序主入口


# 运行

亲爱的面试官你好，我来介绍一下项目执行顺序

## 1. 拉取项目
git clone https://github.com/GunCc/zzk_yunying_task.git
cd zzk_yunying_task

## 2. 下载依赖
go mod tidy


## 3. 初始化文档
swag init

## 4. 启动项目
go run main.go


## 5. 访问页面
http://127.0.0.1:9000/swagger/index.html

