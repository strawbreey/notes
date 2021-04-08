---
title: "Css Vertical Center"
date: 2021-04-08T16:35:20+08:00
draft: false
---

## 垂直居中

1. line-height

```html

  <div id="father">
    <span id="child">
      子元素
    </span>
  </div>

  <style type="text/css">
    #father{
      height:100px;
      line-height:100px;
    }
  </style>
```

2. flex

```html
  <div id="father">
    <div id="child"></div>
  </div>

 <style type="text/css">
    #father{
      display: flex;
      align-items: center;
    }

  </style>
```

3. absolute

```html
  <div id="father">
    <div id="child"></div>
  </div>

  <style type="text/css">
    #father{
      position:relative;
    }
    #child{
      top:50%;
      transform: translateY(-50%);
      position:absolute;
    }
  </style>
```


4. 父元素设置了高度

```html
  <div id="father">
    <div id="child"></div>
  </div>

  <style type="text/css">
    #father{
      height: 100px
    }
    #child{
      position: relative;
      top: 50%;
      transform: translate(-50%);

    }
  </style>
```

5. 表格居中


```html
<div id="box"><div id="content"></div></div>

<style type="text/css">
#box { 
  display: table;
  height: 400px;
  background: #c00;
}

#content {
  color: #fff;
  text-align: center;
  display: table-cell;
  vertical-align: middle;
}
</style>



作者：李天昭
链接：https://www.zhihu.com/question/20543196/answer/15432218
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
```