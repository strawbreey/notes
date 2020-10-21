---
title: "Php Download"
date: 2020-10-21T15:53:03+08:00
draft: true
---

                header('Access-Control-Allow-Origin: *');
                // header('location:http://www.baidu.com');
                // var_dump('66666');
                header("Location: $file_path");
                header('Content-type: application/octet-stream');//告诉浏览器这是一个文件
                header('Content-Disposition: attachment; filename="'.$file->file_name.'"');//文件描述，页面下载用的文件名，可以实现用不同的文件名下载同一个文件
                header("X-Accel-Buffering: yes");
                // header('X-Accel-Redirect:'.$file_path);
                // header("Location: http://file.waibao.woa.com");


                die; 
16:04