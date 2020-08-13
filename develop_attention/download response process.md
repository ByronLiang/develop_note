# 文件下载Response响应处理

```java
response.setHeader("Content-disposition",
        "attachment;filename=" + percentEncodedFileName +
        ";filename*=utf-8''" + percentEncodedFileName);
```

## 响应头

header的对象`Content-disposition`

### Content-disposition 配置

1. attachment: 配置`disposition-type`; 以附件形式进行下载

2. filename: 配置`disposition-parm`; `filename*`: 百分号编码后的文件名; `filename` 不进行编码的文件名

备注: 例如 `filename*=utf-8''%e2%82%ac%20rates` 第一部分是字符集(utf-8)，中间部分是语言(未填写)，最后的%e2%82%ac%20rates代表了实际值

## 处理/封装

```java
// 对真实文件名进行百分号编码
String percentEncodedFileName = URLEncoder.encode(clientFileName, "utf-8").replaceAll("\\+", "%20");
// 组装contentDisposition的值
StringBuilder contentDispositionValue = new StringBuilder();
contentDispositionValue.append("attachment; filename=")
        .append(percentEncodedFileName)
        .append(";")
        .append("filename*=")
        .append("utf-8''")
        .append(percentEncodedFileName);
response.setHeader("Content-disposition", contentDispositionValue.toString());
```
