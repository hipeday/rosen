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
                       nickname VARCHAR(64) NOT NULL,                          -- 用户昵称
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

-- 创建资源类型表 resource_types
CREATE TABLE resource_types (
                                id SERIAL PRIMARY KEY,                          -- 资源类型ID，自增
                                code VARCHAR(100) NOT NULL UNIQUE,               -- 资源类型唯一标识符，例如：users_management、orders_management等
                                name VARCHAR(100) NOT NULL,                      -- 资源类型名称，例如：用户管理、订单管理等
                                description TEXT,                                -- 资源类型描述
                                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 创建时间
                                modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 更新时间
                                created_by SERIAL NOT NULL,                              -- 创建人
                                modified_by SERIAL NOT NULL,                              -- 更新人
                                CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id),
                                CONSTRAINT fk_modified_by FOREIGN KEY (modified_by) REFERENCES users(id)
);

-- 添加列注释
COMMENT ON COLUMN resource_types.id IS '资源类型ID，自增';
COMMENT ON COLUMN resource_types.code IS '资源类型唯一标识符，例如：users_management、orders_management等';
COMMENT ON COLUMN resource_types.name IS '资源类型名称，用于展示，例如：用户管理、订单管理等';
COMMENT ON COLUMN resource_types.description IS '资源类型描述，简要描述资源的功能';
COMMENT ON COLUMN resource_types.created_at IS '创建时间，记录资源类型创建的时间';
COMMENT ON COLUMN resource_types.modified_at IS '更新时间，记录资源类型最后一次更新时间';
COMMENT ON COLUMN resource_types.created_by IS '创建人ID，记录用户信息被谁创建';
COMMENT ON COLUMN resource_types.modified_by IS '更新时间，记录用户信息最后一次人';

-- 添加表注释
COMMENT ON TABLE resource_types IS '资源类型表，存储所有可能的资源类型，code 是唯一标识，name 用于展示';

-- 创建角色表 roles
CREATE TABLE roles (
                       id SERIAL PRIMARY KEY,                          -- 角色ID，自增
                       code VARCHAR(100) NOT NULL UNIQUE,               -- 角色唯一标识符，例如：admin、user等
                       name VARCHAR(100) NOT NULL,                      -- 角色名称，例如：管理员、普通用户等
                       description TEXT,                                -- 角色描述
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 创建时间
                       modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 更新时间
                       created_by INT NOT NULL,                         -- 创建人ID，关联用户表中的ID
                       modified_by INT NOT NULL,                        -- 更新人ID，关联用户表中的ID
                       CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id),
                       CONSTRAINT fk_modified_by FOREIGN KEY (modified_by) REFERENCES users(id)
);

-- 添加列注释
COMMENT ON COLUMN roles.id IS '角色ID，主键，自增';
COMMENT ON COLUMN roles.code IS '角色唯一标识符，例如：admin、user等';
COMMENT ON COLUMN roles.name IS '角色名称，用于展示，例如：管理员、普通用户等';
COMMENT ON COLUMN roles.description IS '角色描述，简要描述角色的功能';
COMMENT ON COLUMN roles.created_at IS '创建时间，记录角色创建的时间';
COMMENT ON COLUMN roles.modified_at IS '更新时间，记录角色最后一次更新时间';
COMMENT ON COLUMN roles.created_by IS '创建人ID，记录角色信息由哪个用户创建';
COMMENT ON COLUMN roles.modified_by IS '更新时间，记录角色信息由哪个用户更新';

-- 添加表注释
COMMENT ON TABLE roles IS '角色表，存储系统中所有角色的信息';

-- 添加表注释
COMMENT ON TABLE roles IS '角色表，存储系统中所有角色的信息';

-- 创建权限表 permissions
CREATE TABLE permissions (
                             id SERIAL PRIMARY KEY,                                      -- 权限ID，主键，自增
                             resource_type_id INT NOT NULL REFERENCES resource_types(id), -- 外键，引用资源类型表的ID
                             resource_typename VARCHAR(100) NOT NULL,                   -- 资源标识符，例如“查看用户”、“编辑订单”等
                             action_type VARCHAR(50) NOT NULL,                            -- 操作类型，例如“查看”、“编辑”、“删除”等
                             description TEXT,                                            -- 权限描述
                             parent_permissions_id INT,                                   -- 父权限ID，指向权限表中的ID，表示权限继承
                             created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,              -- 创建时间
                             modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,             -- 更新时间
                             created_by INT NOT NULL,                                      -- 创建人ID，关联用户表中的ID
                             modified_by INT NOT NULL,                                     -- 更新人ID，关联用户表中的ID
                             CONSTRAINT fk_parent_permission FOREIGN KEY (parent_permissions_id) REFERENCES permissions(id) ON DELETE SET NULL, -- 外键约束，指向权限表自己
                             CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id),
                             CONSTRAINT fk_modified_by FOREIGN KEY (modified_by) REFERENCES users(id)
);

-- 添加列注释
COMMENT ON COLUMN permissions.id IS '权限ID，主键，自增';
COMMENT ON COLUMN permissions.resource_type_id IS '资源类型ID，引用资源类型表的ID，表示该权限属于哪个资源类型';
COMMENT ON COLUMN permissions.resource_typename IS '资源名称，例如查看用户、编辑订单等';
COMMENT ON COLUMN permissions.action_type IS '操作类型，例如查看、编辑、删除等';
COMMENT ON COLUMN permissions.description IS '权限描述，简要描述该权限的功能';
COMMENT ON COLUMN permissions.parent_permissions_id IS '父权限ID，表示该权限继承自某个父权限';
COMMENT ON COLUMN permissions.created_at IS '创建时间，记录权限创建的时间';
COMMENT ON COLUMN permissions.modified_at IS '更新时间，记录权限最后一次更新的时间';
COMMENT ON COLUMN permissions.created_by IS '创建人ID，记录权限信息由哪个用户创建';
COMMENT ON COLUMN permissions.modified_by IS '更新时间，记录权限信息由哪个用户更新';

-- 添加表注释
COMMENT ON TABLE permissions IS '权限表，存储系统中所有权限的信息，支持细粒度和继承关系的权限控制';