### 1. 使用gin框架 并结合传统MVC架构练习业务的增删改查
### 2. 抛异常不用原生的err，用自己定义的异常类
### 3. 使用 router -> controller -> service -> dao 
### 4. 入参、出参都需要封装到结构体当中来解析
### 5. 为了解耦，dao层的返回的数据库对象(db.***)，都要format为 seriliaze包下的对象(有选择性的接收需要的属性)
### 6. dao中dao.go文件定义了全局的数据库操作对象DB，方便每一个dbmodel表去使用
### 7. dbmodel中抽象的模型不必配置，默认会映射到数据库中的表，表明变小写，驼峰规则变_
### 8.待更新......
