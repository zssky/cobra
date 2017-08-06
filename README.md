![cobra logo](https://github.com/zssky/cobra/blob/master/resources/cobra.jpg)

# cobra
[![Build Status](https://travis-ci.org/zssky/cobra.svg?branch=master)](https://travis-ci.org/zssky/cobra)


# 什么是cobra
cobra是基于Go语言的数据采集工具，主要用于提供多数据源数据采集到sqlite数据库，基于sqlite数据库再做进一步的数据分析和处理；支持数据以不同的方式导出，提供方便跨界的操作界面。 用户可以自己定义任务内容，采集过程中的处理语句可以直接使用不同的模板来区分。

# 架构
![cobra logo](https://github.com/zssky/cobra/blob/master/resources/cobra_framework.png)

cobra主要分两块:  
## 1.界面部分  
界面部分我们主要使用Nw.js来开发，方便快捷同时支持多平台。

## 2.服务端部分
服务端的开发我们主要基于http协议和grpc协议和前端进行通信，保证用户体验的同时需要保证后端数据采集的效率。

# 界面设计
此处对界面部分的处理做详细说明

# 服务端  
数据采集主要依赖于服务端，服务端的处理逻辑和模块功能说明  
