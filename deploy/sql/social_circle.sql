CREATE TABLE social_circle (
                               id  VARCHAR(32) NOT NULL ,
                               user_id VARCHAR(32) NOT NULL COMMENT '用户ID（外键关联用户表）',
                               cover_url VARCHAR(512) COMMENT '封面URL（支持CDN）',
                               created_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    -- 外键约束（关键部分）
                               PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户社交圈表';