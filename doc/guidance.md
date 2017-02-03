# pb2doc 实现详解

## 解析protocol buffer message文件

### 明确数据类型

一个`.proto`文件的基本语法如下:

#### `syntax`

    syntax = "proto3";

#### `package`

    package foo.bar;

#### `import`

    import "foo.proto";

#### `option`

    option java_package = "com.example.foo";

#### `message`

```proto
message SearchRequest {
    string query = 1;
    int32 page_number = 2;
    int32 result_per_page = 3;
}
```

而对于`message`的定义，除了可以是基本的数据类型(`int`,`string`,`bool`),还支持嵌入的`message`, `enum`, `oneof`, `map`, `repeated`.

5. 定义`enum`,如:

```proto
enum Corpus {
    UNIVERSAL = 0;
    WEB = 1;
    IMAGES = 2;
}
```

这些都是在一个`.proto`文件中可以单独存在的结构。而`oneof`, `map<k,v>`, `repeated`这些必须作为message内的一个字段存在。搞清楚`.proto`文件的基本结构，就很容易构造出相应的正则来匹配。

---

### 分析`.proto`文件

以单个`.proto`文件为例，首先读取整个文件内容，这里我是逐行读取，并且去除了空行。

1. 读取`syntax`,判断是不是`proto3`,这里只支持`proto3`的解析

```go
syntax_pattern := "^syntax\\s?=\\s?\"(.+)\"\\s?;"
```

2. 读取`package`

```go
package_pattern := "^package\\s+(.+)\\s*;"
```

3. 读取`import`

```go
pattern := "^import\\s+\"(.+)\"\\s*;"
```

这里注意的是需要为了解析循环`import`的问题，需要标记是否被解析过，如果被解析过，则不再读取。

４．　读取注释

支持

- 单行注释 `regexp.Compile("^\\s*(//.*)|(/\\*.*\\*/)")`
- 多行注释　`regexp.Compile("^\\s*/\\*.*")`

`messgae`的注释要紧跟在定义之前，如:

```proto
/**
 * This is multi comments
 * Please parse it use head
 */
message Person {
  // The name of person
  string name = 1;
  int32 id = 2;  // Unique ID number for this person.
  string email = 3;

  PhoneType type = 4;

  repeated Property properties = 5;
  map<string, Project> projects = 3;  //test map
}

```

对于`message`结构内的每个字段，支持两种格式的注释:

+ 上一行
+ 行尾

**优先显示上一行的注释**


