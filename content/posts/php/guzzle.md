---
title: "Guzzle"
date: 2021-07-23T10:49:13+08:00
draft: false
---

Guzzle is a PHP HTTP client that makes it easy to send HTTP requests and trivial to integrate with web services.

```php
$client = new \GuzzleHttp\Client();
$response = $client->request('GET', 'https://api.github.com/repos/guzzle/guzzle');

echo $response->getStatusCode(); // 200
echo $response->getHeaderLine('content-type'); // 'application/json; charset=utf8'
echo $response->getBody(); // '{"id": 1420053, "name": "guzzle", ...}'

// Send an asynchronous request.
$request = new \GuzzleHttp\Psr7\Request('GET', 'http://httpbin.org');
$promise = $client->sendAsync($request)->then(function ($response) {
    echo 'I completed! ' . $response->getBody();
});

$promise->wait();
```


### 参考资料

- [Guzzle Github](https://github.com/guzzle/guzzle)
- [Guzzle中文文档](https://guzzle-cn.readthedocs.io/zh_CN/latest/#)
