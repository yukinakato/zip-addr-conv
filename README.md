# 郵便番号・住所変換くん

郵便番号や住所を高速に検索・変換するサービス。

https://yubin-jusyo.com/

## 特徴
* Ajax を用いたページ遷移を伴わないストレスレスなオペレーション
* データをメモリ上にのせることによる高速な検索
* マッチ箇所のハイライトやレスポンシブ対応によるユーザーフレンドリーなインターフェース

## 本アプリが解決する課題
事務作業等において、お客様宛の文書を作成する必要が生じる。その際に郵便番号や住所を記載する必要がある。  
本アプリは、この作業を簡易化する効果がある。

**郵便番号 → 住所の変換**
* 郵便番号が既知の場合は、日本語入力システム(IME)の変換機能を用いて容易に住所を知ることができる。
* しかし、IME 内の辞書の更新はラグがあり、また企業によってはオプショナルな更新としてブロックしている場合がある。この場合、最新の地番変更などの情報が反映されておらず、業務に支障をきたす。

**住所 → 郵便番号の変換**
* IME を用いることで変換は可能であるが、完全な住所を入力しなければ目的の郵便番号を得ることができず、効率が悪い。

![conventional](https://user-images.githubusercontent.com/22708232/94155076-814e0400-feb9-11ea-9854-f60191861e9b.gif)

* また、顧客リストに市区町村以下しか記載がなかったり、知らない地名や難読な地名の場合は完全な住所をスムーズに得ることが困難な場合がある。


本アプリを用いることによって、管理者が反映させた最新のデータによる検索結果を提供することができる。  
また、検索キーワードの複数指定や都道府県の絞り込みを行うことで、目的の結果をスピーディーかつ高確率で得ることができる。

## 使用技術
* Go
* jQuery / Chart.js ([パフォーマンス確認画面](https://yubin-jusyo.com/)用)
* testing (Go) / Jest (JavaScript)
* webpack
* nginx / Let's Encrypt
* AWS Lightsail / Route 53
* GitHub Actions
* Docker Compose (開発用環境)
* Excel VBA

## 今後の改修予定
* 脱 VBA (norm パッケージを使用する)
* 検索がワーストケースで 10 ms 前後かかっているため、高速化したい
* ワンクリックでクリップボードにコピーする機能の追加
