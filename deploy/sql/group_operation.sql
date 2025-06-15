CREATE TABLE group_operation (
                                 id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                                 group_id VARCHAR(64) NOT NULL COMMENT '群ID',
                                 operator_id VARCHAR(64) NOT NULL COMMENT '操作者ID',
                                 target_id VARCHAR(64) COMMENT '被操作成员ID',
                                 action_type ENUM('KICK', 'SET_ADMIN', 'CHANGE_NAME') NOT NULL COMMENT '操作类型',
                                 extra_info JSON COMMENT '扩展信息 (如新群名)',
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
                                 INDEX idx_group (group_id),
                                 FOREIGN KEY fk_group (group_id) REFERENCES group_info(group_id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='群敏感操作审计';