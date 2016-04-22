# memobase

このアプリはGo言語で作られたメモ管理アプリです。

コンソール上で動作し以下の機能を持ちます。

* メモ新規作成

# メモの新規作成

``` sh
# 現在時刻でメモを作成
$ mb new
2016-04-22-18-52-26.txt

# テンプレートを用いて作成
$ mb new -t mail
2016-04-22-18-52-26.txt
$ mb new -template mail
2016-04-22-18-52-26.txt

# 名前をつける
$ mb new -n tmp
tmp.txt
$ mb new -name tmp
tmp.txt
```

