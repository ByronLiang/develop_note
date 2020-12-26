# ETCD CLI 基本命令

## Keys

获取指定前缀的key的 key与值: `etcdctl get /testdir/ --prefix`

只看key: `etcdctl get /testdir/ --prefix --keys-only`

只看value: `etcdctl get /testdir/ --prefix --print-value-only`

## Watch

key的监听: `etcdctl watch /wt/k`
