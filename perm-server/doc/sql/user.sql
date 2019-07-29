-- ----------------------------
-- Table structure for cd_sys_user_user
-- ----------------------------
DROP TABLE IF EXISTS `cd_sys_user_user`;
CREATE TABLE `cd_sys_user_user` (
  `id` varchar(50) NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `creator_department_id` varchar(50) DEFAULT NULL,
  `creator_department_name` varchar(255) DEFAULT NULL,
  `creator_id` varchar(50) DEFAULT NULL,
  `creator_name` varchar(255) DEFAULT NULL,
  `is_deleted` int(1) NOT NULL DEFAULT 0,
  `update_time` datetime DEFAULT NULL,
  `updator_department_id` varchar(50) DEFAULT NULL,
  `updator_department_name` varchar(255) DEFAULT NULL,
  `updator_id` varchar(50) DEFAULT NULL,
  `updator_name` varchar(255) DEFAULT NULL,
  `version` int(11) DEFAULT NULL,

  `user_number` varchar(50) DEFAULT NULL COMMENT '编号',
  `user_account` varchar(50) DEFAULT NULL COMMENT '账号',
  `user_password` varchar(50) DEFAULT NULL COMMENT '密码',
  `user_name` varchar(50) DEFAULT NULL COMMENT '姓名',
  `sex` varchar(50) DEFAULT NULL COMMENT '性别',
  `department_id` varchar(50) DEFAULT NULL COMMENT '所属部门',
  `mobile` varchar(50) DEFAULT NULL COMMENT '电话',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `address` varchar(255) DEFAULT NULL COMMENT '地址',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for cd_sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `cd_sys_user_role`;
CREATE TABLE `cd_sys_user_role` (
  `id` varchar(50) NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `creator_department_id` varchar(50) DEFAULT NULL,
  `creator_department_name` varchar(255) DEFAULT NULL,
  `creator_id` varchar(50) DEFAULT NULL,
  `creator_name` varchar(255) DEFAULT NULL,
  `is_deleted` int(1) NOT NULL DEFAULT 0,
  `update_time` datetime DEFAULT NULL,
  `updator_department_id` varchar(50) DEFAULT NULL,
  `updator_department_name` varchar(255) DEFAULT NULL,
  `updator_id` varchar(50) DEFAULT NULL,
  `updator_name` varchar(255) DEFAULT NULL,
  `version` int(11) DEFAULT NULL,

  `p_type` varchar(100) DEFAULT NULL COMMENT '',
  `role_type` varchar(50) DEFAULT NULL COMMENT '角色类型',
  `role_name` varchar(100) DEFAULT NULL COMMENT '角色名称',
  `role_code` varchar(20) DEFAULT NULL COMMENT '角色编号',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for cd_sys_user_resource
-- ----------------------------
DROP TABLE IF EXISTS `cd_sys_user_resource`;
CREATE TABLE `cd_sys_user_resource` (
  `id` varchar(50) NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `creator_department_id` varchar(50) DEFAULT NULL,
  `creator_department_name` varchar(255) DEFAULT NULL,
  `creator_id` varchar(50) DEFAULT NULL,
  `creator_name` varchar(255) DEFAULT NULL,
  `is_deleted` int(1) NOT NULL DEFAULT 0,
  `update_time` datetime DEFAULT NULL,
  `updator_department_id` varchar(50) DEFAULT NULL,
  `updator_department_name` varchar(255) DEFAULT NULL,
  `updator_id` varchar(50) DEFAULT NULL,
  `updator_name` varchar(255) DEFAULT NULL,
  `version` int(11) DEFAULT NULL,

  `resource_type` varchar(50) DEFAULT NULL COMMENT '资源类型',
  `resource_name` varchar(100) DEFAULT NULL COMMENT '名称',
  `resource_level` integer DEFAULT 0 COMMENT '级别',
  `resource_order` integer DEFAULT 0 COMMENT '顺序',
  `resource_sate` integer DEFAULT 0 COMMENT '状态',
  `parent_id` varchar(50) DEFAULT NULL COMMENT '父id',
  `url` varchar(100) DEFAULT NULL COMMENT '地址',
  `icon` varchar(100) DEFAULT NULL COMMENT '图标',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for cd_sys_user_role_resource
-- ----------------------------
DROP TABLE IF EXISTS `cd_sys_user_role_resource`;
CREATE TABLE `cd_sys_user_role_resource` (
  `id` varchar(50) NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `creator_department_id` varchar(50) DEFAULT NULL,
  `creator_department_name` varchar(255) DEFAULT NULL,
  `creator_id` varchar(50) DEFAULT NULL,
  `creator_name` varchar(255) DEFAULT NULL,
  `is_deleted` int(1) NOT NULL DEFAULT 0,
  `update_time` datetime DEFAULT NULL,
  `updator_department_id` varchar(50) DEFAULT NULL,
  `updator_department_name` varchar(255) DEFAULT NULL,
  `updator_id` varchar(50) DEFAULT NULL,
  `updator_name` varchar(255) DEFAULT NULL,
  `version` int(11) DEFAULT NULL,

  `role_id` varchar(50) DEFAULT NULL COMMENT '角色id',
  `resource_id` varchar(100) DEFAULT NULL COMMENT '资源id',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for cd_sys_user_user_role
-- ----------------------------
DROP TABLE IF EXISTS `cd_sys_user_user_role`;
CREATE TABLE `cd_sys_user_user_role` (
  `id` varchar(50) NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `creator_department_id` varchar(50) DEFAULT NULL,
  `creator_department_name` varchar(255) DEFAULT NULL,
  `creator_id` varchar(50) DEFAULT NULL,
  `creator_name` varchar(255) DEFAULT NULL,
  `is_deleted` int(1) NOT NULL DEFAULT 0,
  `update_time` datetime DEFAULT NULL,
  `updator_department_id` varchar(50) DEFAULT NULL,
  `updator_department_name` varchar(255) DEFAULT NULL,
  `updator_id` varchar(50) DEFAULT NULL,
  `updator_name` varchar(255) DEFAULT NULL,
  `version` int(11) DEFAULT NULL,

  `user_id` varchar(50) DEFAULT NULL COMMENT '用户id',
  `role_id` varchar(100) DEFAULT NULL COMMENT '角色id',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for cd_sys_user_department
-- ----------------------------
DROP TABLE IF EXISTS `cd_sys_user_department`;
CREATE TABLE `cd_sys_user_department` (
  `id` varchar(50) NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `creator_department_id` varchar(50) DEFAULT NULL,
  `creator_department_name` varchar(255) DEFAULT NULL,
  `creator_id` varchar(50) DEFAULT NULL,
  `creator_name` varchar(255) DEFAULT NULL,
  `is_deleted` int(1) NOT NULL DEFAULT 0,
  `update_time` datetime DEFAULT NULL,
  `updator_department_id` varchar(50) DEFAULT NULL,
  `updator_department_name` varchar(255) DEFAULT NULL,
  `updator_id` varchar(50) DEFAULT NULL,
  `updator_name` varchar(255) DEFAULT NULL,
  `version` int(11) DEFAULT NULL,

  `dep_name` varchar(100) DEFAULT NULL COMMENT '部门名称',
  `dep_code` varchar(50) DEFAULT NULL COMMENT '部门编号',
  `dep_level` int(11) DEFAULT NULL COMMENT '部门级别',
  `dep_order` int(11) DEFAULT NULL COMMENT '部门顺序',
  `parent_id` varchar(50) DEFAULT NULL COMMENT '上级部门id',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

-- 初始化用户
INSERT INTO cd_sys_user_user (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `user_account`, `user_password`, `user_name`, `sex`, `department_id`, `mobile`, `email`, `address`) VALUES ('5b456ec2c3dc41eca83e8d05bf875ceb', '2019-07-12 16:39:26', '', '', '', '', '0', NULL, '', '', '', '', '0', '', 'x04jpoIrc8/mvNRqAG59Wg==', 'root', '', '', '', '', '');
-- 初始化菜单
INSERT INTO cd_sys_user_resource (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `resource_type`, `resource_name`, `resource_level`, `resource_order`, `resource_sate`, `parent_id`, `url`, `icon`, `description`) VALUES ('4028f1a56a23c8ee016a24108302000a', '2019-07-16 17:16:17', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:16:54', '1', '系统', 'admin', '管理员', '0', '0', '系统首页', '1', '1', '0', '0', 'dashboard', 'el-icon-lx-home', '系统首页');
INSERT INTO cd_sys_user_resource (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `resource_type`, `resource_name`, `resource_level`, `resource_order`, `resource_sate`, `parent_id`, `url`, `icon`, `description`) VALUES ('4028f1a56a23c8ee016a241083020001', '2019-07-16 17:16:17', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:16:54', '1', '系统', 'admin', '管理员', '0', '0', '用户管理', '1', '2', '0', '0', '1', 'el-icon-lx-home', '用户管理');
INSERT INTO cd_sys_user_resource (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `resource_type`, `resource_name`, `resource_level`, `resource_order`, `resource_sate`, `parent_id`, `url`, `icon`, `description`) VALUES ('4028f1a56a23c8ee016a241083020002', '2019-07-16 17:16:17', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:16:54', '1', '系统', 'admin', '管理员', '0', '0', '菜单管理', '2', '1', '0', '4028f1a56a23c8ee016a241083020001', 'permission-list', '', '菜单管理');
INSERT INTO cd_sys_user_resource (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `resource_type`, `resource_name`, `resource_level`, `resource_order`, `resource_sate`, `parent_id`, `url`, `icon`, `description`) VALUES ('4028f1a56a23c8ee016a241083020003', '2019-07-16 17:16:17', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:16:54', '1', '系统', 'admin', '管理员', '0', '0', '角色管理', '2', '2', '0', '4028f1a56a23c8ee016a241083020001', 'role-list', '', '角色管理');
INSERT INTO cd_sys_user_resource (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `resource_type`, `resource_name`, `resource_level`, `resource_order`, `resource_sate`, `parent_id`, `url`, `icon`, `description`) VALUES ('4028f1a56a23c8ee016a241083020004', '2019-07-16 17:16:17', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:16:54', '1', '系统', 'admin', '管理员', '0', '0', '用户管理', '2', '2', '0', '4028f1a56a23c8ee016a241083020001', '2-2', '', '用户管理');
INSERT INTO cd_sys_user_resource (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `resource_type`, `resource_name`, `resource_level`, `resource_order`, `resource_sate`, `parent_id`, `url`, `icon`, `description`) VALUES ('4028f1a56a23c8ee016a241083020005', '2019-07-16 17:16:17', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:16:54', '1', '系统', 'admin', '管理员', '0', '0', '内部用户管理', '3', '1', '0', '4028f1a56a23c8ee016a241083020004', 'new-user', '', '用户管理');
-- 初始化角色
INSERT INTO cd_sys_user_role (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `role_type`, `role_name`, `role_code`, `description`, `p_type`) VALUES ('30be500d2a8a427e94c6130db7403458', '2019-07-16 17:34:07', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:34:38', '1', '系统', 'admin', '管理员', '0', 'p', '管理员', 'role_admin*', '角色管理', '');
-- 初始化角色菜单
INSERT INTO cd_sys_user_role_resource (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `role_id`, `resource_id`) VALUES ('4028f1a56a23c8ee016a24108302999a', '2019-07-16 17:36:34', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:37:10', '1', '系统', 'admin', '管理员', '0', '30be500d2a8a427e94c6130db7403458', '4028f1a56a23c8ee016a24108302000a');
INSERT INTO cd_sys_user_role_resource (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `role_id`, `resource_id`) VALUES ('4028f1a56a23c8ee016a241083029991', '2019-07-16 17:36:34', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:37:10', '1', '系统', 'admin', '管理员', '0', '30be500d2a8a427e94c6130db7403458', '4028f1a56a23c8ee016a241083020001');
INSERT INTO cd_sys_user_role_resource (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `role_id`, `resource_id`) VALUES ('4028f1a56a23c8ee016a241083029992', '2019-07-16 17:36:34', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:37:10', '1', '系统', 'admin', '管理员', '0', '30be500d2a8a427e94c6130db7403458', '4028f1a56a23c8ee016a241083020002');
INSERT INTO cd_sys_user_role_resource (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `role_id`, `resource_id`) VALUES ('4028f1a56a23c8ee016a241083029993', '2019-07-16 17:36:34', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:37:10', '1', '系统', 'admin', '管理员', '0', '30be500d2a8a427e94c6130db7403458', '4028f1a56a23c8ee016a241083020003');
INSERT INTO cd_sys_user_role_resource (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `role_id`, `resource_id`) VALUES ('4028f1a56a23c8ee016a241083029994', '2019-07-16 17:36:34', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:37:10', '1', '系统', 'admin', '管理员', '0', '30be500d2a8a427e94c6130db7403458', '4028f1a56a23c8ee016a241083020004');
INSERT INTO cd_sys_user_role_resource (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `role_id`, `resource_id`) VALUES ('4028f1a56a23c8ee016a241083029995', '2019-07-16 17:36:34', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:37:10', '1', '系统', 'admin', '管理员', '0', '30be500d2a8a427e94c6130db7403458', '4028f1a56a23c8ee016a241083020005');
-- 初始化用户角色
INSERT INTO cd_sys_user_user_role (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `user_id`, `role_id`) VALUES ('4028f1a56a23c8ee016a241083026666', '2019-07-16 17:41:13', '1', '系统', 'admin', '管理员', '0', '2019-07-16 17:41:13', '1', '系统', 'admin', '管理员', '0', '5b456ec2c3dc41eca83e8d05bf875ceb', '30be500d2a8a427e94c6130db7403458');
-- 初始化部门
INSERT INTO cd_sys_user_department (`id`, `create_time`, `creator_department_id`, `creator_department_name`, `creator_id`, `creator_name`, `is_deleted`, `update_time`, `updator_department_id`, `updator_department_name`, `updator_id`, `updator_name`, `version`, `dep_name`, `dep_code`, `dep_order`, `parent_id`) VALUES ('0', '2019-07-19 16:46:04', '1', '系统', 'admin', '管理员', '0', '2019-07-19 16:46:04', '1', '系统', 'admin', '管理员', '0', '系统', '00_00', '0', '-1');

