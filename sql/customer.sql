CREATE TABLE IF NOT EXISTS `cloud-mall`.`customer_address`  (
  `id` int(11) NOT NULL COMMENT '自增ID',
  `customer_id` int(16) NOT NULL COMMENT '用户ID',
  `zip` int(11) NULL COMMENT '邮编',
  `provice_id` int(11) NULL COMMENT '省/自治市 ID',
  `city_id` int(0) NULL COMMENT '地区表中城市的ID',
  `district_id` int(0) NULL COMMENT '区ID',
  `address` varchar(255) NULL COMMENT '具体门牌号地址',
  `id_default` tinyint(2) NULL COMMENT '是否是默认地址',
  `created_time` datetime(0) NULL,
  `updated_time` datetime(0) NULL,
  `deleted_time` datetime(0) NULL,
  PRIMARY KEY (`id`, `customer_id`)
);

CREATE TABLE IF NOT EXISTS `cloud-mall`.`customer_balance_log`  (
  `id` int(11) NOT NULL COMMENT '余额日志ID',
  `customer_id` int(16) NOT NULL COMMENT '用户ID',
  `source` int(11) NULL COMMENT '记录来源：1订单，2退货单',
  `source_sn` int(11) NULL COMMENT '相关单据ID',
  `amount` decimal(11, 2) NULL COMMENT '变动金额',
  `created_time` datetime(0) NULL,
  `updated_time` datetime(0) NULL,
  `deleted_time` datetime(0) NULL,
  PRIMARY KEY (`id`, `customer_id`)
);

CREATE TABLE IF NOT EXISTS `cloud-mall`.`customer_info`  (
  `id` int(11) NOT NULL COMMENT '自增ID',
  `customer_id` int(16) NOT NULL COMMENT '用户ID',
  `customer_name` varchar(255) NULL COMMENT '用户真实姓名',
  `identity_card_type` tinyint(2) NULL COMMENT '证件类型：1 身份证，2 军官证，3 护照',
  `identity_card_no` varchar(20) NULL COMMENT '证件号码',
  `mobile_phone` int(11) NULL COMMENT '手机号码',
  `email` varchar(255) NULL COMMENT '电子邮件',
  `gender` tinyint(2) NULL COMMENT '性别：1 男生 2 女生 3 其他',
  `user_point` int(11) NULL COMMENT '用户积分',
  `register_time` datetime(0) NULL COMMENT '用户注册时间',
  `birthday` datetime(0) NULL COMMENT '用户生日',
  `customer_level` tinyint(2) NULL COMMENT '会员级别',
  `user_money` decimal(11, 2) NULL COMMENT '会员余额',
  `created_time` datetime(0) NULL,
  `updated_time` datetime(0) NULL,
  `deleted_time` datetime(0) NULL,
  PRIMARY KEY (`id`, `customer_id`)
);

CREATE TABLE IF NOT EXISTS `cloud-mall`.`customer_level`  (
  `id` int(11) NOT NULL COMMENT '自增ID',
  `customer_level` tinyint(2) NOT NULL COMMENT '会员级别',
  `customer_level_name` varchar(255) NULL COMMENT '会员级别名',
  `min_point` int(11) NULL COMMENT '该级别最低分',
  `max_point` int(11) NULL COMMENT '该级别最高分',
  `created_time` datetime(0) NULL,
  `updated_time` datetime(0) NULL,
  `deleted_time` datetime(0) NULL,
  PRIMARY KEY (`id`, `customer_level`)
);

CREATE TABLE IF NOT EXISTS `cloud-mall`.`customer_login`  (
  `id` int(11) NOT NULL COMMENT '自增ID',
  `customer_id` int(16) NOT NULL COMMENT '用户ID',
  `username` varchar(255) NULL COMMENT '用户名',
  `password` varchar(255) NULL COMMENT '加密后的密码',
  `state` tinyint(2) NULL COMMENT '用户状态',
  `created_time` datetime(0) NULL,
  `updated_time` datetime(0) NULL,
  `deleted_time` datetime(0) NULL,
  PRIMARY KEY (`id`, `customer_id`)
);

CREATE TABLE IF NOT EXISTS `cloud-mall`.`customer_login_log`  (
  `id` int(11) NOT NULL COMMENT '登陆日志ID',
  `customer_id` int(16) NOT NULL COMMENT '用户ID',
  `login_time` datetime(0) NULL COMMENT '登录时间',
  `login_ip` varchar(128) NULL COMMENT '登录IP',
  `login_type` tinyint(2) NULL COMMENT '登陆类型：0未成功，1成功',
  `created_time` datetime(0) NULL,
  `updated_time` datetime(0) NULL,
  `deleted_time` datetime(0) NULL,
  PRIMARY KEY (`id`, `customer_id`)
);

CREATE TABLE IF NOT EXISTS `cloud-mall`.`customer_point_log`  (
  `id` int(11) NOT NULL COMMENT '积分日志ID',
  `customer_id` int(16) NOT NULL COMMENT '用户ID',
  `point_source` varchar(255) NULL COMMENT '积分来源：0订单，1登陆，2活动',
  `refer_number` int(11) NULL COMMENT '积分来源相关编号',
  `change_point` int(11) NULL COMMENT '变更积分数',
  `created_time` datetime(0) NULL,
  `updated_time` datetime(0) NULL,
  `deleted_time` datetime(0) NULL,
  PRIMARY KEY (`id`, `customer_id`)
);

