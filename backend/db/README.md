# 数据库迁移

- 使用 `db/migrations/001_init.sql` 初始化核心表：users, growth_routes, replies, events, participants, banners。
- 可用 `mysql -uuser -p -h host sf6 < db/migrations/001_init.sql` 快速导入。

