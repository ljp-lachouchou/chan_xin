CREATE TABLE comments (
                          id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,  -- 评论ID
                          post_id VARCHAR(32) NOT NULL,                  -- 关联动态ID（与MongoDB的posts._id一致）
                          user_id VARCHAR(32) NOT NULL,                  -- 评论者ID
                          content TEXT NOT NULL,                         -- 评论内容（需敏感词过滤）
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          is_deleted BOOLEAN DEFAULT FALSE,              -- 软删除标记
                          INDEX idx_post (post_id),                      -- 加速动态的评论加载
                          INDEX idx_user (user_id)                       -- 用户评论历史查询
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;