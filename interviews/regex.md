# 正则相关笔记

## 提取网页里图片资源

如 `<img width=120px src="xxx.yy">` 或 `<img src="xxx.yy" class="ss" width="100px">`

通过`preg_match_all('/<img[^>]*src="([^"]*)"[^>]*>/i',$html, $matchs);`可提取图片资源的链接
