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

## ch5000

## 参考URL

- エンジニアスタイル
  - [https://www.r-staffing.co.jp/engineer/entry/20190621_1](https://www.r-staffing.co.jp/engineer/entry/20190621_1)