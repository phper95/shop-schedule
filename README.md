# shop-schedule

#### 介绍
shop-schedule是海量数据高并发场景，构建Go+ES8企业级搜索微服务课程商城系统的调度服务，

[课程地址 **点此 打开**](https://coding.imooc.com/class/579.html?mc_marking=bb86c9071ed9b7cf12612a2a85203372)

项目基于golang开发，主要用于商城的一些脚本和调度服务

#### 软件技术栈
Go + mongo+Elasticsearch


#### 安装教程

1、安装go>=1.15,这个可以https://studygolang.com/dl下载

2、开启mod： go env -w GO111MODULE=on

3、配置代理：go env -w GOPROXY=https://goproxy.cn,direct 这个让下载依赖速度更快

5、配置私有仓库：go env -w  GOPRIVATE=*gitee.com

6、下载项目：git clone https://gitee.com/phper95/shop-main.git

7、go mod tidy 安装所需依赖

8、导入sql/shop.sql,修改cconfig,yml 里数据库与redis配置

9、本地运行go run main.go


#### 使用说明

1.  config/config.yml下配置mongo等组件的连接地址
2.  编译运行项目


