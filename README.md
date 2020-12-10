# project-gen
the generator of project code for RESTful

### 暂不支持
1. 结构体嵌套
2. 结构体暂未完成处理tag
3. todo: 添加对模版语法的校验
4. todo: 列表查询

db: module是几个table的集合（逻辑唯独），以module为中心，嵌入db *gorm.DB(或其他）,且拥有DB（）用于返回*gorm.DB（或其他）示例

tag: 字段：field-db中字段name / gen-对应FieldModel字段如isAdd,isDelete等