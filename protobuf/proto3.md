# Protobuf3语言指南  
## 定义一个消息类型  
看一个非常简单的例子  
```protobuf
syntax = "proto3";

 message SearchRequest {
   string query = 1;
   int32 page_number = 2;
   int32 result_per_page = 3;
 }
```
- 非注释的第一行指定正在使用```proto3``` ,如果不指定，默认使用```proto2```   
- ```message```关键字为指定一个消息类型
## 消息字段类型
protoType| goType| phpType
- | :-: | :-: | -:
double | float64 | float
float | float32 | float
int32 | int32 | integer
uint32 | uint32 | integer
int64 | int64 | integer
bool | bool | bool
string | string | string
bytes | []byte | string
- 对于strings，默认是一个空string
- 对于bytes，默认是一个空的bytes
- 对于bools，默认是false
- 对于数值类型，默认是0
- 对于枚举，默认是第一个定义的枚举值，必须为0;

## 枚举
当需要定义一个消息类型的时候，可能想为一个字段指定某“预定义值序列”。通过向消息定义中添加一个枚举（enum）并且为每个可能的值定义一个常量就可以了。  
例  
```protobuf
message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
  enum Corpus {  //枚举消息类
    UNIVERSAL = 0;
    WEB = 1;
    IMAGES = 2;
    LOCAL = 3;
    NEWS = 4;
    PRODUCTS = 5;
    VIDEO = 6;
  }
  Corpus corpus = 4;
}
```
- 枚举的第一个常量映射为0，而且每个枚举类型必须将其第一个映射为0
- 枚举常量必须在32位整数值范围内
## 包（package）
```protobuf
package foo.bar;
```
同一目录下的.proto文件要求必须是同一个package
## 定义服务(Service)
 service 包括 request 以及 response 的类型。
 ```protobuf
service SearchService {

  rpc Search (SearchRequest) returns (SearchResponse);

}
 ```
 protocol编译器将产生一个抽象接口SearchService以及一个相应的存根实现。

