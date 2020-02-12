# Laravel Eloquent 操作笔记

### 关联关系保存

sample:

```php
$user = User::where('name', 'nick')->first();
$user->age = 18;
$user->wechat->nickName = "Remember";
$user->save();
$user->wechat->save();
// 可被直接替换成
$user->push();
```
源自`Illuminate\Database\Eloquent\Model`类
push()方法: `Save the model and all of its relationships`