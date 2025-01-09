-- 创建枚举类型 user_status，用于用户状态
CREATE TYPE user_status AS ENUM (
    'INACTIVE', -- 未激活
    'ACTIVE',   -- 用户活跃状态
    'FROZEN',   -- 冻结
    'RESIGNED'  -- 已离职
);

-- 创建用户表 users
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,                                  -- 用户ID，主键，自增
                       username VARCHAR(50) NOT NULL UNIQUE,                   -- 用户名，必须唯一
                       password VARCHAR(255) NOT NULL,                         -- 密码哈希
                       totp_secret VARCHAR(64),                                -- TOTP 2FA的密钥
                       email VARCHAR(100) UNIQUE,                              -- 用户邮箱，必须唯一
                       status user_status DEFAULT 'INACTIVE',                  -- 用户状态，默认值为 ACTIVE
                       super_admin BOOLEAN DEFAULT FALSE,                      -- 是否超级管理员，默认否
                       last_login_at TIMESTAMP DEFAULT NULL,                   -- 最近一次登录时间
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,         -- 创建时间
                       modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,        -- 更新时间
                       created_by SERIAL NOT NULL ,                            -- 创建人
                       modified_by SERIAL NOT NULL                             -- 更新人
);

-- 添加列注释
COMMENT ON COLUMN users.id IS '用户ID，主键，自增';
COMMENT ON COLUMN users.username IS '用户名，必须唯一';
COMMENT ON COLUMN users.password IS '密码哈希，使用加密存储用户密码';
COMMENT ON COLUMN users.totp_secret IS 'TOTP 2FA的密钥';
COMMENT ON COLUMN users.email IS '用户邮箱，必须唯一';
COMMENT ON COLUMN users.super_admin IS '是否超级管理员，超级管理员具有所有权限';
COMMENT ON COLUMN users.last_login_at IS '最近一次登录时间，记录用户上次登录的时间';
COMMENT ON COLUMN users.created_at IS '创建时间，记录用户创建的时间';
COMMENT ON COLUMN users.modified_at IS '更新时间，记录用户信息最后一次更新的时间';
COMMENT ON COLUMN users.created_by IS '创建人ID，记录用户信息被谁创建';
COMMENT ON COLUMN users.modified_by IS '更新时间，记录用户信息最后一次人';

-- 添加表注释
COMMENT ON TABLE users IS '用户表，存储系统中所有用户的信息';

-- 创建用户个人资料表 users_profiles
CREATE TABLE users_profiles (
                                id SERIAL PRIMARY KEY,                                   -- 资料ID，主键，自增
                                userid INT NOT NULL UNIQUE REFERENCES users(id)          -- 用户ID，外键，唯一，关联 users 表
                                    ON DELETE CASCADE,                                   -- 当用户被删除时，自动删除对应的资料
                                birthday DATE,                                           -- 出生日期
                                gender CHAR(1) CHECK (gender IN ('M', 'F', 'O')),        -- 性别 ('M': 男, 'F': 女, 'O': 其他)
                                address TEXT,                                            -- 地址
                                totp_verified BOOLEAN NOT NULL DEFAULT FALSE,            -- TOTP是否已经验证 默认否
                                totp_enabled BOOLEAN NOT NULL DEFAULT FALSE,             -- TOTP是否开启
                                mobile VARCHAR(20),                                      -- 移动电话
                                nickname VARCHAR(100),                                   -- 用户昵称
                                avatar VARCHAR(255),                                     -- 头像URL 保存绝对路径 不保存domain
                                bio TEXT,                                                -- 简介/自我描述
                                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,          -- 创建时间
                                modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,         -- 更新时间
                                created_by SERIAL NOT NULL ,                             -- 创建人
                                modified_by SERIAL NOT NULL                              -- 更新人
);

-- 添加列注释
COMMENT ON COLUMN users_profiles.id IS '资料ID，主键，自增';
COMMENT ON COLUMN users_profiles.userid IS '用户ID，外键，关联用户表';
COMMENT ON COLUMN users_profiles.birthday IS '出生日期';
COMMENT ON COLUMN users_profiles.gender IS '性别：M(男)，F(女)，O(其他)';
COMMENT ON COLUMN users_profiles.mobile IS '用户手机号';
COMMENT ON COLUMN users_profiles.nickname IS '用户全名';
COMMENT ON COLUMN users_profiles.avatar IS '头像URL，存储用户头像的链接';
COMMENT ON COLUMN users_profiles.address IS '地址，存储用户的居住地址';
COMMENT ON COLUMN users_profiles.totp_verified IS 'TOTP是否已经验证 默认否';
COMMENT ON COLUMN users_profiles.totp_enabled IS 'TOTP是否开启';
COMMENT ON COLUMN users_profiles.bio IS '个人简介或自我描述';
COMMENT ON COLUMN users_profiles.created_at IS '创建时间，记录资料创建的时间';
COMMENT ON COLUMN users_profiles.modified_at IS '更新时间，记录资料最后一次更新的时间';
COMMENT ON COLUMN users_profiles.created_by IS '创建人ID，记录用户信息被谁创建';
COMMENT ON COLUMN users_profiles.modified_by IS '更新时间，记录用户信息最后一次人';

-- 添加表注释
COMMENT ON TABLE users_profiles IS '用户个人资料表，存储与用户相关的额外信息';

update users_profiles
set avatar = '/avatars/johndoe.jpg',
    address = address
where id = 1;