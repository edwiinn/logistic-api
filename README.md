# Logistic API
Framework or Library Used :
+ Echo : https://github.com/labstack/echo
+ Gorm : https://github.com/jinzhu/gorm


Database : MySQL

### Query Parameter
|Parameter|Argument|Description|
|---|---|---|
|sort_by<br>(optional)|[origin][destination]<br>[shipment][distance]|Sort all jobs according<br>to argument|
|order_by<br>(optional, require sort_by)|[ascending][asc]<br>[descending][dsc]| Additional parameter for <br>*sort_by*, order sorted data|