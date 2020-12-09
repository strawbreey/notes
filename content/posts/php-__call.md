---
title: "Php __Call"
date: 2020-12-07T17:22:53+08:00
draft: false
---

__construct()方法到实例化时自动加载function 

__call()方法用来获取没有定义的function  

__get()得到私有变量值  

__toString()  

__set()方法用来设置私有变量值  

__isset()方法判断私有变量值 

__unset()方法删除私有变量值 

__destruct()删除类对象时自动会调用  

```php
class info {  
  
  private $province; //省  
  public  $city;     //城市  
  private $myname;   //姓名  
  
  //__construct()方法到实例化时自动加载function  
  function __construct() {  
    echo "实例化我就起作用<br>";  
    $this->myname("张映");  
  }  
  
  //__call()方法用来获取没有定义的function  
  function __call($name, $param) {  
    echo $name."这个方法没有定义，跑到我这儿来了<br>";  
    print_r($param) ;  
    echo "<br>";  
  }  
  
  //定义一个方法  
  function myname($param) {  
    echo "定义过的就到定义过的这儿来，$param<br>";  
    return $this->myname = $param;  
  }  
  
  //__get()得到私有变量值  
  private function __get($name) {  
    if (isset ($this-> $name)) {  
      return ($this-> $name);  
    } else {  
      return false;  
    }  
  }  
  
  //__toString()  
  private function __toString() {  
    echo '你输入的这个类的名子叫'.__CLASS__."<br>";  
    return __CLASS__;  
  }  
  
  //__set()方法用来设置私有变量值  
  private function __set($name, $value) {  
    echo "设置私有变量时，自动调用了这个__set()方法为私有变量赋值<br>";  
    $this-> $name = $value;  
  }  
  
  //__isset()方法判断私有变量值  
  private function __isset($name) {  
    echo "isset()函数判断私有变量是不是被定义时，自动调用__isset()<br>";  
    return isset ($this-> $name);  
  }  
  
  //__unset()方法删除私有变量值  
  private function __unset($name) {  
    echo "unset()函数删除一个私有变量时，自动调用__unset()<br>";  
    unset ($this-> $name);  
  }  
  
  //__destruct()删除类对象时自动会调用  
  public function __destruct() {  
    echo '删除类对像，就到__destruct（）这儿来<br>';  
  }  
}  

$info = new info();        //会调用__construct()方法和__destruct()方法  
$info->province = "安徽";        //会调用__set()方法  
$info->city = "合肥";                //不会调用__set()方法  
$info->myname("张映");      //会调用myname()方法  
$info->nickname("tank");   //会调用__call()方法  
(string)$info;             //会调用__toString()方法  
var_dump(isset ($info->province)) . "<br>";   //会调用__isset()方法和__get()方法  
echo $info->province . "<br>";     //会调用__get()方法  
unset ($info->province);           //会调用__unset()方法  
echo $info->province;              //已经被删除了， 所以不会显示为安徽了,会调用__isset()方法  
```

PHP5 的对象新增了一个专用方法 __call()，这个方法用来监视一个对象中的其它方法。如果你试着调用一个对象中不存在或被权限控制中的方法，__call 方法将会被自动调用。

__call()函数是php类的默认魔法函数，__call() 在一个对象的上下文中，如果调用的方法不存在的时候，它将被触发。