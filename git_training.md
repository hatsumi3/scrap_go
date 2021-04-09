# git training

## ch1

- バージョン確認
    `git --version`
- 初期化
    `git init`
- 状態確認
    `git status`
- ステージングに追加
    `git add <filename>`
    `git add .`
- コミット
    `git commit -m <commit message>`
- 状態確認
    `git status`
- コミットログ確認
    `git log`
    `git log --graph`

## ch2

- コミットオブジェクト確認
    `git cat-file -p <commit id>`
- ブランチ確認
    `git branch`
- ブランチ作成
    `git branch <branch name>`
- ブランチへ移動
    `git checkout <branch name>`
- ブランチ作成して移動
    `git checkout -b <branch name>`
- ブランチ作成して強制移動
    `git checkout -f <branch name>`

## ch3

- ブランチAの内容を取り込み
    `git merge <branch nameA>`

## ch4

- ブランチAの内容を取り込んだらコンフリクト
    `git merge <branch nameA>`
    `conflict!!! -> fix conflict file`
    `git add <fixed file>`
    `git commit -m <message>`
- mergeでコンフリクト
    `git merge --abort`

## ch5

- コミットを打ち消すコミットを行う
    `git revert <commit id>`

## ch6

特定の地点までファイルを巻き戻す

- コミットだけ
  `git reset --soft HEAD@{1}`
- コミット＋ステージング
  `git reset --mixed HEAD@{1}`
- コミット＋ステージング＋変更内容
  `git reset --hard HEAD@{1}`

## ch7

- 操作履歴を確認する。
  `git reflog`
- 確認して戻したい場所を指定する
  `git reset --hard HEAD@{1}`

## ch8

`git checkout <???>`は意味が変わる。
それぞれの専用コマンドがある

- ファイルの復元
  `git checkout <filename>`
  `git restore`
- ブランチの切り替え
  `git checkout <dirname>`
  `git switch`

なお、新しいコマンドのため、変更可能性あり。
gitをupdateする必要もあり。

- 作業ディレクトリ上の編集内容を取り消す。
  `git restore --worktree <filename>`
  `git checkout <filename>`
  `git reset --hard HEAD`

## ch9

- 変更を確認する
  `git diff <before commitid> <after commitid>`
- 直前のコミットとの差分
  `git diff`

## ch10

- 部分的にコミットを取得する。取得後、コミットされる
  `git cherry-pick <commit id>`
- 部分的にコミットを取得する。取得のみでコミットしない
  `git cherry-pick -n <commit id>`

## ch11

- 編集中のファイル（スタッシュ）を保存する
  `git stash save`
- スタッシュの一覧をみる
  `git stash list`
- スタッシュを復活させる
  `git stash apply stash@{0}`
- スタッシュを削除する
  `git stash drop stash@{0}`
- スタッシュを復活、削除する
  `git stash pop stash@{0}`

## ch12

リポジトリ内の検索

- 検索したい文字列
  `git grep <words>`
- 検索したい文字列 リポジトリ内
  `git grep <words> <commit>`
- 検索したい文字列を含むフォルダ名
  `git grep -l <word>`

## ch13

誰が行を変更したかを確認する

- 変更を確認したいファイル名
  `git blame <filename>`

## ch14

- リモートリポジトリを表示する
  `git remote`
- リモートリポジトリ一覧を表示する
  `git remote -v`
- リモートリポジトリを追加する
  `git remote add <reponame> <repo url>`
- リモートリポジトリを削除する
  `git remote rm <repo name>`

## 参考URL

- エンジニアスタイル
  - [https://www.r-staffing.co.jp/engineer/entry/20190621_1](https://www.r-staffing.co.jp/engineer/entry/20190621_1)