CREATE TABLE comment_replies (
                                 id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                                 comment_id BIGINT UNSIGNED NOT NULL,           -- 关联的一级评论ID
                                 user_id VARCHAR(32) NOT NULL,                  -- 回复者ID
                                 target_user_id VARCHAR(32) NOT NULL,           -- 被回复用户ID（通知推送关键字段）
                                 content TEXT NOT NULL,
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 is_deleted BOOLEAN DEFAULT FALSE,
                                 FOREIGN KEY (comment_id) REFERENCES comments(id) ON DELETE CASCADE,
                                 INDEX idx_target_user (target_user_id)         -- 加速用户被回复的通知查询
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;