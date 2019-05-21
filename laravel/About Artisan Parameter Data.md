### 关于Console的命令行参数接收原理笔记

#### Artisan命令行的主要例子：`email:send {user}`

##### Method 1

- 当调用命令时，则可使用`$schedule->command('email:send John');`如果在命令里面传参数，直接把参数放在命令后面就可以了;

##### Method 2
- 基于Laravel的源码
`public function command($command, array $parameters = [])`

- 可以将请求的参数以数组形式，且数组里排列的数据需对应请求的参数

处理请求参数的函数`compileParameters(array $parameters)`

    return collect($parameters)->map(function ($value, $key) {
        if (is_array($value)) {
            $value = collect($value)->map(function ($value) {
                return ProcessUtils::escapeArgument($value);
            })->implode(' ');
        } elseif (! is_numeric($value) && ! preg_match('/^(-.$|--.*)/i', $value)) {
            $value = ProcessUtils::escapeArgument($value);
        }

        return is_numeric($key) ? $value : "{$key}={$value}";
    })->implode(' ');

