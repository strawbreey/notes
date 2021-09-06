---
title: "Js Cos Upload"
date: 2021-08-27T15:24:57+08:00
draft: false
---

### 使用 cos sdk 实现原生上传

```js
    
    var Bucket = 'wbp-1258344700';
    var Region = 'ap-guangzhou';     /* 存储桶所在地域，必须字段 */
    var protocol = location.protocol === 'https:' ? 'https:' : 'https:';
    var prefix = protocol + '//' + Bucket + '.cos.' + Region + '.myqcloud.com/';

    var info = {
        TmpSecretId: null,
        TmpSecretKey: null,
        SecurityToken: null,
    }

    // 对更多字符编码的 url encode 格式
    var camSafeUrlEncode = function (str) {
        return encodeURIComponent(str)
            .replace(/!/g, '%21')
            .replace(/'/g, '%27')
            .replace(/\(/g, '%28')
            .replace(/\)/g, '%29')
            .replace(/\*/g, '%2A');
    };

    // 随机数
    var randomKey = function () {
        var $chars = 'ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678';
        var maxPos = $chars.length;
        var pwd = '';
        for (i = 0; i < 32; i++) {
                pwd += $chars.charAt(Math.floor(Math.random() * maxPos));
        }
        return 'test/cpm/' + pwd
    }

    // 初始化实例
    var cos = new COS({
        getAuthorization: function (options, callback) {
            // 异步获取临时密钥
            $.get('/cos/sts', function (data) {
                data = JSON.parse(data).data
                var credentials = data && data.credentials;
                if (!data || !credentials) return console.error('credentials invalid');
                callback({
                    TmpSecretId: credentials.tmpSecretId,
                    TmpSecretKey: credentials.tmpSecretKey,
                    SecurityToken: credentials.sessionToken,
                    // 建议返回服务器时间作为签名的开始时间，避免用户浏览器本地时间偏差过大导致签名错误
                    StartTime: data.startTime,     // 时间戳，单位秒，如：1580000000
                    ExpiredTime: data.expiredTime, // 时间戳，单位秒，如：1580000900
                });
            });
        }
    });

    // 监听表单提交
    document.getElementById('upload').addEventListener("change", function (e) {
        var file = document.getElementById('upload').files[0];
        if (!file) {
            document.getElementById('msg').innerText = '未选择上传文件';
            return;
        }

        cos.putObject({
            Bucket: Bucket, 
            Region: Region,  
            Key: randomKey() +  '/' + file.name,
            StorageClass: 'STANDARD',
            Body: file, // 上传文件对象
            onProgress: function(progressData) {
                console.log(JSON.stringify(progressData));
            }
        }, function(err, data) {
            console.log(err || data);
        });
    });

```