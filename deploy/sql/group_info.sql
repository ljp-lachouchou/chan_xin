CREATE TABLE group_info (
                            group_id VARCHAR(24) PRIMARY KEY COMMENT '群ID',
                            name VARCHAR(100) NOT NULL COMMENT '群名称',
                            owner_id VARCHAR(24) NOT NULL COMMENT '群主ID',
                            max_members INT UNSIGNED DEFAULT 200 COMMENT '群人数上限',
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            INDEX idx_owner (owner_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='群基础信息';