CREATE TABLE post_likes (
                            id BIGINT AUTO_INCREMENT PRIMARY KEY,
                            post_id VARCHAR(32) NOT NULL,   -- 动态ID（关联MongoDB的posts._id）
                            user_id VARCHAR(32) NOT NULL,   -- 点赞用户ID
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            is_deleted BOOLEAN DEFAULT FALSE, -- 取消点赞标记（软删除）
                            UNIQUE KEY uniq_post_user (post_id, user_id)
);