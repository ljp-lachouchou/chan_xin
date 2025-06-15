CREATE TABLE friend_relation (
                                 id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL COMMENT '关系ID',
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
                                 INDEX idx_friend (friend_id),
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='好友关系及状态';
CREATE TABLE friend_apply (
                              apply_id VARCHAR(64)NOT NULL  COMMENT '申请ID',
                              applicant_id VARCHAR(64) NOT NULL COMMENT '申请人ID',
                              target_id VARCHAR(64) NOT NULL COMMENT '目标用户ID',
                              greet_msg VARCHAR(200) DEFAULT '' COMMENT '打招呼消息',
                              status TINYINT(1) DEFAULT 0 COMMENT '状态 (0:待处理, 1:同意, 2:拒绝)',
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '申请时间',
                              INDEX idx_applicant (applicant_id),
                              INDEX idx_target (target_id),
                              PRIMARY KEY (`apply_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='好友申请记录';
CREATE TABLE group_info (
                            group_id VARCHAR(64) NOT NULL COMMENT '群ID',
                            name VARCHAR(100) NOT NULL COMMENT '群名称',
                            owner_id VARCHAR(24) NOT NULL COMMENT '群主ID',
                            max_members INT UNSIGNED DEFAULT 200 COMMENT '群人数上限',
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            INDEX idx_owner (owner_id),
                            PRIMARY KEY (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='群基础信息';
CREATE TABLE group_member (
                              id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL,
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
                              PRIMARY KEY (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='群成员及个性化设置';
CREATE TABLE group_operation (
                                 id BIGINT UNSIGNED AUTO_INCREMENT NOT NULL,
                                 group_id VARCHAR(64) NOT NULL COMMENT '群ID',
                                 operator_id VARCHAR(64) NOT NULL COMMENT '操作者ID',
                                 target_id VARCHAR(64) COMMENT '被操作成员ID',
                                 action_type ENUM('KICK', 'SET_ADMIN', 'CHANGE_NAME') NOT NULL COMMENT '操作类型',
                                 extra_info JSON COMMENT '扩展信息 (如新群名)',
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
                                 INDEX idx_group (group_id),
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='群敏感操作审计';