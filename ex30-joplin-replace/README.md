# 替换\u00a0为真正的空格(' ')

2024-03-15 印象笔记会非常奇怪的使用 \u00a0 做空格。我经常遇到拷贝示例代码执行失败的问题。 不得不处理下。

```
# 查看单条笔记是否包含 '\u00a0'
$ curl 'http://localhost:41184/notes/5460fca4ddd643108c7d3b6008666091?token=dc9dd18b2ea101b5aee131441bf7f6e71a7ddf8886ad0d111ddaed2212f91d28a828d546885a05dc6f17f860133d1d05e6ccb1bfa8b5b17ad73c32c5d022e106&fields=id,title,body,user_updated_time&limit=1' | cat -A

# 检查包含 '\u00a0' 的笔记数量
$ go run . --limit 100
# 替换
$ go run . --limit 100 --replace ' '
```
