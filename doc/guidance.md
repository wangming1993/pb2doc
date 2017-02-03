# pb2doc 实现详解

## 解析protocol buffer message文件

### 明确数据类型

一个`.proto`文件的基本语法如下:

1. 定义`syntax`,如:
    - `syntax = "proto3";`
2. 定义`package`,如:
    - `package foo.bar;`
