# memobase

このアプリはGo言語で作られたメモ管理アプリです。

コンソール上で動作し以下の機能を持ちます。

* メモ管理フォルダーの作成
* メモ新規作成
* 設定

# メモ管理フォルダーの作成

以下のコマンドを入力することでメモ管理ようのディレクトリを生成できます

``` sh
# 現在のディレクトリにメモ管理ようのディレクトリを作成
$ mb init memobase
create memobase directory
```

上記コマンドを実行すれば以下のディレクトリ構造が出来上がります。

memobase/  
    ├── config.json  
    ├── memo  
    └── template  

config.json：設定ファイルの格納先です  

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

# 作成後、エディタを開く
$ mb new -o
$ mb new --open

# カテゴリーを作成する
$ mb new -c todo
$ mb new --category todo
todo/2016-04-29.txt
```

# configフォルダーにある「config.json」ファイルを編集することにより以下の設定ができます。


```json
{
    "name": "<directoryname>",
    "editor": "<defaultEditor>",
    "newFileDatetime": "<defaultNewFileDatetime>"
}

// example
{
    "name": "memobase",
    "editor": "vim",
    "newFileDatatime": "2016-12-21"
}
```

* デフォルトエディター
デフォルトで使用するエディターを設定できます。上記の例ですとエディターにvimが指定されているため「mb new -o」コマンド入力時はvimが自動的に起動されます
* デフィルトファイルネーム
新規ファイル作成時の日付情報を設定できます。
上記の例ですと新規メモ作成時は年月日（YYYY-MM-dd.txt）で作成されます。

# Todo

今後追加したい機能になります

* メモ編集（簡易アクセス？）
* グローバル設定ファイル（複数のmemobaseフォルダーへの簡易アクセス）
* gitによる差分管理のための.gitignore自動生成
* markdown対応とローカルWEBからのアクセス

