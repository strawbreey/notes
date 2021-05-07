<?php
/*
 * 此数组将在遍历的过程中改变其中某些单元的值
 */
$people = Array(
  Array('name' => 'Kalle', 'salt' => 856412), 
  Array('name' => 'Pierre', 'salt' => 215863)
);

for($i = 0; $i < count($people); ++$i) {
  $people[$i]['salt'] = rand(000000, 999999);
}

// 以上代码可能执行很慢，因为每次循环时都要计算一遍数组的长度。由于数组的长度始终不变，可以用一个中间变量来储存数组长度以优化而不是不停调用 count()：


$people = Array(
  Array('name' => 'Kalle', 'salt' => 856412), 
  Array('name' => 'Pierre', 'salt' => 215863)
);

for($i = 0, $size = count($people); $i < $size; ++$i)
{
    $people[$i]['salt'] = rand(000000, 999999);
}

?>