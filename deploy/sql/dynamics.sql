CREATE TABLE comments (
                          comment_id VARCHAR(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                          post_id VARCHAR(32) NOT NULL,
                          user_id VARCHAR(32) NOT NULL,
                          content TEXT NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          is_deleted BOOLEAN DEFAULT FALSE,
                          PRIMARY KEY (comment_id),
                          INDEX idx_post (post_id),
                          INDEX idx_user (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE comment_replies (
                                 comment_replie_id VARCHAR(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                 comment_id VARCHAR(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
                                 user_id VARCHAR(32) NOT NULL,
                                 target_user_id VARCHAR(32) NOT NULL,
                                 content TEXT NOT NULL,
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 is_deleted BOOLEAN DEFAULT FALSE,
                                 PRIMARY KEY (comment_replie_id),
                                 FOREIGN KEY (comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE,
                                 INDEX idx_target_user (target_user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE notifications (
                               id VARCHAR(32) NOT NULL,     -- 通知ID（ULID）
                               user_id VARCHAR(32) NOT NULL,   -- 接收用户ID
                               type ENUM('LIKE', 'REPLY') NOT NULL, -- 通知类型
                               trigger_user_id VARCHAR(32) NOT NULL, -- 触发者ID
                               post_id VARCHAR(32) NOT NULL,   -- 关联动态ID
                               comment_id VARCHAR(32),         -- 关联评论ID（仅REPLY类型）
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               is_read BOOLEAN DEFAULT FALSE,  -- 是否已读
                               PRIMARY KEY(id),
                               INDEX idx_user_unread (user_id, is_read) -- 加速未读通知查询
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE post_likes (
                            id VARCHAR(32) NOT NULL,
                            post_id VARCHAR(32) NOT NULL,   -- 动态ID（关联MongoDB的posts._id）
                            user_id VARCHAR(32) NOT NULL,   -- 点赞用户ID
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            is_deleted BOOLEAN DEFAULT FALSE, -- 取消点赞标记（软删除）
                            PRIMARY KEY(id),
                            UNIQUE KEY uniq_post_user (post_id, user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
