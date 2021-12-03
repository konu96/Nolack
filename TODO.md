# TODO
- `CreateNotionPageInteractor` でページが作れないので対応する
  - [Notion API](https://developers.notion.com/reference/post-page) を見ると `database_id` or `page_id` を指定すれば良いと読める
  - ただ、片方しか渡さないともう片方を渡せとエラーになる
  - Public Beta なので API 側の不具合と思ったので、しばらく放置する