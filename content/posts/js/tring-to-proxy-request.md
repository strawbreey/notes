---
title: "Tring to Proxy Request"
date: 2021-10-21T21:00:46+08:00
draft: false
---


> [HPM] Error occurred while trying to proxy request

有些是代理写错，比如下面，是需要写上 `http://` 

{
  "/web": {
    "target": "qq.cpm.cm.com:8210",
    "secure": true
  }
}

=>

{
  "/web": {
    "target": "http://qq.cpm.cm.com:8210",
    "secure": true
  }
}
