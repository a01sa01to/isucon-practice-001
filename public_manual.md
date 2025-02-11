# 社内ISUCON 当日レギュレーション

## 禁止事項

以下の事項は特別に禁止する。

  * 他のチームへの妨害と主催者がみなす全ての行為

## サーバ事項

参加者は主催者が用意したAmazon Web Services（以下AWS）のEC2インスタンスを1台利用する。

## ソフトウェア事項

コンテストにあたり、参加者は与えられたソフトウェア、もしくは自分で競技時間内に実装したソフトウェアを用いる。

高速化対象のソフトウェアとして主催者からRuby, PHPによるWebアプリケーションが与えられる。ただし各々の性能が一致することを主催者は保証しない。どれをベースに用いてもよいし、独自で実装したものを用いてもよい。

競技における高速化対象のアプリケーションとして与えられたアプリケーションから、以下の機能は変更しないこと。

  * アクセス先のURI（ポート、およびHTTPリクエストパス）
  * レスポンス(HTML)のDOM構造
  * JavaScript/CSSファイルの内容
  * 画像および動画等のメディアファイルの内容

各サーバにおけるソフトウェアの入れ替え、設定の変更、アプリケーションコードの変更および入れ替えなどは一切禁止しない。起動したインスタンス以外の外部リソースを利用する行為（他のインスタンスに処理を委譲するなど）は禁止する。

許可される事項には、例として以下のような作業が含まれる。

  * DBスキーマの変更やインデックスの作成・削除
  * キャッシュ機構の追加、jobqueue機構の追加による遅延書き込み
  * 他の言語による再実装

ただし以下の事項に留意すること。

  * コンテスト進行用のメンテナンスコマンドが正常に動作するよう互換性を保つこと
  * 各サーバの設定およびデータ構造は任意のタイミングでのサーバ再起動に耐えること
  * サーバ再起動後にすべてのアプリケーションコードが正常動作する状態を維持すること
  * ベンチマーク実行時にアプリケーションに書き込まれたデータは再起動後にも取得できること

# 採点

採点は採点条件（後述）をクリアした参加者の間で、性能値（後述）の高さを競うものとする。

採点条件として、以下の各チェックの検査を通過するものとする。

  * 負荷走行中、POSTしたデータが、POSTへのHTTPレスポンスを返してから即座に関連するURI GETのレスポンスデータに反映されていること
  * レスポンスHTMLのDOM構造が変化していないこと
  * ブラウザから対象アプリケーションにアクセスした結果、ページ上の表示および各種動作が正常であること

性能値として、以下の指標を用いる。

  * 計測ツールの実行時間は1分間とする
    * 細かい閾値ならびに配点についての詳細は当日のマニュアルに記載する
  * 計測時間内のHTTPリクエスト成功数をベースとする
    * リクエストの種類毎に配点を変更する
    * エラーの数により減点する
