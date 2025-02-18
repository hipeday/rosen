-- 初始化一条系统账号 密码为 0ktRa09LFqcC
insert into users (
                   id, username, password, totp_secret, email, status, super_admin, last_login_at, created_at,
                   modified_at, created_by, modified_by)
values (-1, '系统账号', '9d2111b4dd91c01b5029339b2eb98a2d', null, 'official@hipeday.org', 'ACTIVE', true, null, now(), now(), -1, -1);