# Django Rest framework & swagger 使用相关总结

## swagger的使用

使用`drf-yasg`包, 可以将Django Rest framework的应用接口文档可视化处理

### 注意点

- 当应用框架的模式属于非测试模式时，需要配置与引入`drf-yasg`的静态文件包来加载接口文档视图文件；当应用属于测试模式时，则无需引入静态文件
