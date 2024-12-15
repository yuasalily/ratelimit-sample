## ratelimit-sample
ratelimitのサンプル実装

### 環境構築
下記を参考にdevcontainerで開発環境を構築
https://code.visualstudio.com/remote/advancedcontainers/connect-multiple-containers

- main
  - READMEやらを管理する。
- mock
  - ratelimitの裏側にあるコンポーネント。ratelimitの動作確認用のモック。
- runner
  - ratelimitを実行する環境。