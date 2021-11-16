---
title: "Get Content Disposition"
date: 2021-11-04T10:46:07+08:00
draft: false
---

```js
// 在 angular 实现文件下载功能， 默认只能在前端代码中手动添加文件类型及文件名。
export class DownloadService {
  constructor(private http: HttpClient) { }

  /**
   * Blob请求
   */
  public requestBlob(url: string, data?: any): Observable<any> {
      return this.http.request('post', url, {
          body: data,
          observe: 'response',
          responseType: 'blob',
      });
  }

  /**
   * Blob文件转换下载
   */
  public downFile(result: any, fileName: string, fileType?: string) {
      const data = result.body;
      const blob = new Blob([data], {
              type: fileType || data.type,
          });
      const objectUrl = URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.setAttribute('style', 'display:none');
      a.setAttribute('href', objectUrl);
      a.setAttribute('download', fileName);
      a.click();
      URL.revokeObjectURL(objectUrl);
  }

  public export(url: string, data: any, fileName: string, fileType?: any) {
      this.requestBlob(url, data).subscribe(result => {
          const headers = result.headers as HttpHeaders;
          this.downFile(result, fileName,
              fileType || headers.get('Content-Type'));
      });
  }

}

private downloadService: DownloadService
this.downloadService.export('http://localhost/export', {}, '流水记录.xlsx');
```

### 参考地址
- https://zodream.cn/blog/id/201.html

