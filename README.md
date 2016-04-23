# memobase

このアプリはGo言語で作られたメモ管理アプリです。

コンソール上で動作し以下の機能を持ちます。

* メモ管理フォルダーの作成
* メモ新規作成

# メモ管理フォルダーの作成

以下のコマンドを入力することでメモ管理ようのディレクトリを生成できます

``` sh
# 現在のディレクトリにメモ管理ようのディレクトリを作成
$ mb init memobase
create memobase directory
```

上記コマンドを実行すれば以下のディレクトリ構造が出来上がります。

memobase/
    ├── memo
    └── template

memo：作成したメモ帳を保管するフォルダーです  

template：メモ帳のテンプレートを保管するフォルダーです  




# メモの新規作成

``` sh
# 現在時刻でメモを作成
$ mb new
create 2016-04-22-18-52-26.txt

# テンプレートを用いて作成
$ mb new -t mail
create 2016-04-22-18-52-26.txt
$ mb new 0-template mail
create 2016-04-22-18-52-26.txt

# 名前をつける
$ mb new -n tmp
create tmp.txt
$ mb new --name tmp
create tmp.txt
```

