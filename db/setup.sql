CREATE TABLE `user` (
  `user_id` VARCHAR(128) NOT NULL COMMENT 'ユーザーID',
  `auth_token` VARCHAR(128) NOT NULL COMMENT '認証トークン',
  `name` VARCHAR(64) NOT NULL COMMENT 'ユーザーID',
  PRIMARY KEY (`user_id`),
  INDEX `idx_authToken` (`auth_token` ASC))
ENGINE = InnoDB
COMMENT = 'ユーザー';

INSERT INTO user(`name_id`, `auth_token`, `name`) value ('1', 'test', 'test_user');