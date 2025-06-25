CREATE TABLE notifications (
  id VARCHAR(32) PRIMARY KEY,     -- 通知ID（ULID）
  user_id VARCHAR(32) NOT NULL,   -- 接收用户ID
  type ENUM('LIKE', 'REPLY') NOT NULL, -- 通知类型
  trigger_user_id VARCHAR(32) NOT NULL, -- 触发者ID
  post_id VARCHAR(32) NOT NULL,   -- 关联动态ID
  comment_id VARCHAR(32),         -- 关联评论ID（仅REPLY类型）
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  is_read BOOLEAN DEFAULT FALSE,  -- 是否已读
  INDEX idx_user_unread (user_id, is_read) -- 加速未读通知查询
);