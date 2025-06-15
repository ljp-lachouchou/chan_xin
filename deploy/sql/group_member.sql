CREATE TABLE group_member (
                              id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                              group_id VARCHAR(64) NOT NULL COMMENT '群ID',
                              user_id VARCHAR(64) NOT NULL COMMENT '用户ID',
                              group_nickname VARCHAR(50) DEFAULT '' COMMENT '群内昵称',
                              show_nickname TINYINT(1) DEFAULT 1 COMMENT '显示昵称 (0:否, 1:是)',
                              is_admin TINYINT(1) DEFAULT 0 COMMENT '管理员 (0:否, 1:是)',
                              is_muted TINYINT(1) DEFAULT 0 COMMENT '免打扰 (0:否, 1:是)',
                              is_topped TINYINT(1) DEFAULT 0 COMMENT '置顶 (0:否, 1:是)',
                              remark VARCHAR(50) DEFAULT '' COMMENT '群备注',
                              joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
                              UNIQUE KEY uk_group_user (group_id, user_id),
                              INDEX idx_user (user_id),
                              FOREIGN KEY fk_group_id (group_id) REFERENCES group_info(group_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='群成员及个性化设置';