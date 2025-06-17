CREATE TABLE group_apply (
                              apply_id VARCHAR(64)NOT NULL  COMMENT '申请ID',
                              applicant_id VARCHAR(64) NOT NULL COMMENT '申请人ID',
                              target_id VARCHAR(64) NOT NULL COMMENT '目标用户ID',
                              greet_msg VARCHAR(200) DEFAULT '' COMMENT '打招呼消息',
                              status TINYINT(1) DEFAULT 0 COMMENT '状态 (0:待处理, 1:同意, 2:拒绝)',
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间',
                              INDEX idx_applicant (applicant_id),
                              INDEX idx_target (target_id),
                              PRIMARY KEY (`apply_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='群组申请记录';