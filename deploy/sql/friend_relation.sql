CREATE TABLE friend_relation (
                                 id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '关系ID',
                                 user_id VARCHAR(64) NOT NULL COMMENT '用户ID',
                                 friend_id VARCHAR(64) NOT NULL COMMENT '好友ID',
                                 remark VARCHAR(50) DEFAULT '' COMMENT '好友备注',
                                 is_muted TINYINT(1) DEFAULT 0 COMMENT '免打扰 (0:否, 1:是)',
                                 is_topped TINYINT(1) DEFAULT 0 COMMENT '置顶 (0:否, 1:是)',
                                 is_blocked TINYINT(1) DEFAULT 0 COMMENT '拉黑 (0:否, 1:是)',
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                                 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                 UNIQUE KEY uk_user_friend (user_id, friend_id),
                                 INDEX idx_user (user_id),
                                 INDEX idx_friend (friend_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='好友关系及状态';