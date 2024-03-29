# tbf

[![Build Status](https://travis-ci.org/jiro4989/tbf.svg?branch=master)](https://travis-ci.org/jiro4989/tbf)

技術書典のサークル情報のリストを取得するコマンドです。

## 使い方

```bash
tbf tbf07 | head -n 1
```

> EZ-NET	イージーネット	https://techbookfest.org/event/tbf07/circle/5639817928900608	プログラミング、とりわけ勉強会の楽しさを伝えたい、そんな気持ちを発端に始まった技術同人誌のシリーズ。今回は Swift 入門書として『Swift らしい表現を目指そう』と『Swift イニシャライザー大全』を頒布します。どちらとも表側から Swift を眺める本になっているので "雑学" というより言語を扱うときの "素養" として活用できる本、楽しくコードを書くために欠かせない "言語" としての入門書として仕立てています。そして今回は新刊として Swift の "変数" についての解説書と、海外のローカルカンファレンス iOSCon に行ってみた体験記的な書籍を頒布します。

```bash
tbf tbf06 -o tbf06.tsv
```

## インストール

以下のコマンドでインストールするかReleasesからダウンロードしてください。

```bash
go get github.com/jiro4989/tbf
```

## LICENSE

MIT
