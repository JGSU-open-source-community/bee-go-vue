# bee-go-vue
Build a simple todo list that use Go and Vue.js to implement.

##Stared
you need have installed go env

```
create table 
CREATE TABLE `task` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `done` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;


git clone https://github.com/JingDa-open-source-community/bee-go-vue.git
cd ~/src/bee-go-vue
go build 
./bee-go-vue
```
open you browser input localhost:8080, following you will see
![](http://p1.bqimg.com/567571/a18fff866cf2afe9.jpg)
