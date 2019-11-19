/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 50725
 Source Host           : localhost:3306
 Source Schema         : admin_server

 Target Server Type    : MySQL
 Target Server Version : 50725
 File Encoding         : 65001

 Date: 16/11/2019 19:38:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_auth
-- ----------------------------
DROP TABLE IF EXISTS `admin_auth`;
CREATE TABLE `admin_auth`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pid` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '上级ID，0为顶级',
  `auth_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '权限名称',
  `auth_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT 'URL地址',
  `sort` int(1) UNSIGNED NOT NULL DEFAULT 999 COMMENT '排序，越小越前',
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `is_show` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否显示，0-隐藏，1-显示',
  `user_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '操作者ID',
  `create_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者ID',
  `update_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改者ID',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态，1-正常，0-删除',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '前端页面路由',
  `redirect` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '前端重定向路由',
  `sidebar_hidden` tinyint(1) NOT NULL DEFAULT 0 COMMENT '前端sidebar不显示 默认false显示',
  `breadcrumb` tinyint(1) NOT NULL DEFAULT 1 COMMENT '前端面包屑显示 默认true显示',
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '前端组件名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 17 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '权限因子' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_auth
-- ----------------------------
INSERT INTO `admin_auth` VALUES (1, 0, '所有权限', '/', 1, '', 0, 1, 1, 1, 1, 1505620970, 1505620970, '', '', 0, 1, '');
INSERT INTO `admin_auth` VALUES (2, 1, '权限管理', '/', 0, 'fa-id-card', 1, 1, 0, 1, 1, 0, 1573356785, '/auth', '/auth/admin', 0, 1, '');
INSERT INTO `admin_auth` VALUES (3, 2, '管理员', 'adminuser/list', 1, '', 1, 1, 1, 1, 1, 1505621186, 1573738407, 'admin', '', 0, 1, 'AdminUser');
INSERT INTO `admin_auth` VALUES (4, 2, '角色管理', '/role/list', 2, '', 1, 1, 0, 1, 1, 0, 1573362221, 'role', '', 0, 1, 'AdminRole');
INSERT INTO `admin_auth` VALUES (5, 3, '新增', 'adminuser/add', 1, '', 0, 1, 0, 1, 1, 0, 1573738414, '', '', 0, 1, '');
INSERT INTO `admin_auth` VALUES (6, 3, '修改', 'adminuser/info', 2, '', 0, 1, 0, 1, 1, 0, 1573738417, '', '', 0, 1, '');
INSERT INTO `admin_auth` VALUES (7, 3, '删除', 'adminuser/del', 3, '', 0, 1, 1, 1, 1, 1505621756, 1573738420, '', '', 0, 1, '');
INSERT INTO `admin_auth` VALUES (8, 4, '新增', '/role/add', 1, '', 0, 1, 0, 1, 1, 0, 1505698716, '', '', 0, 1, '');
INSERT INTO `admin_auth` VALUES (9, 4, '修改', '/role/edit', 2, '', 0, 1, 1, 1, 1, 1505621912, 1505621912, '', '', 0, 1, '');
INSERT INTO `admin_auth` VALUES (10, 4, '删除', '/role/ajaxdel', 3, '', 0, 1, 1, 1, 1, 1505621951, 1505621951, '', '', 0, 1, '');
INSERT INTO `admin_auth` VALUES (11, 2, '权限因子', '/auth/list', 3, '', 1, 1, 1, 1, 1, 1505621986, 1572707009, 'auth', '', 0, 1, 'AdminAuth');
INSERT INTO `admin_auth` VALUES (12, 11, '新增', '/auth/add', 1, '', 0, 1, 1, 1, 1, 1505622009, 1505622009, '', '', 0, 1, '');
INSERT INTO `admin_auth` VALUES (13, 11, '修改', '/auth/edit', 2, '', 0, 1, 1, 1, 1, 1505622047, 1505622047, '', '', 0, 1, '');
INSERT INTO `admin_auth` VALUES (14, 11, '删除', '/auth/ajaxdel', 3, '', 0, 1, 1, 1, 1, 1505622111, 1505622111, '', '', 0, 1, '');
INSERT INTO `admin_auth` VALUES (15, 1, '个人中心', 'profile/edit', 1, 'fa-user-circle-o', 1, 1, 0, 1, 1, 0, 1573356785, '/profile', '/profile/edit', 0, 1, '');
INSERT INTO `admin_auth` VALUES (16, 15, '资料修改', '/user/edit', 1, 'fa-edit', 1, 1, 0, 1, 1, 0, 1573830003, 'edit', '', 0, 1, 'AdminUserEdit');

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '角色名称',
  `detail` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '备注',
  `create_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者ID',
  `update_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改这ID',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态1-正常，0-删除',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加时间',
  `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
INSERT INTO `admin_role` VALUES (1, 'API管理员', '拥有API所有权限', 0, 2, 1, 1505874156, 1505874156);
INSERT INTO `admin_role` VALUES (2, '系统管理员', '系统管理员', 0, 0, 1, 1506124114, 1506124114);

-- ----------------------------
-- Table structure for admin_role_auth
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_auth`;
CREATE TABLE `admin_role_auth`  (
  `role_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色ID',
  `auth_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '权限ID',
  PRIMARY KEY (`role_id`, `auth_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '权限和角色关系表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_role_auth
-- ----------------------------
INSERT INTO `admin_role_auth` VALUES (1, 16);
INSERT INTO `admin_role_auth` VALUES (1, 17);
INSERT INTO `admin_role_auth` VALUES (1, 18);
INSERT INTO `admin_role_auth` VALUES (1, 19);
INSERT INTO `admin_role_auth` VALUES (2, 1);
INSERT INTO `admin_role_auth` VALUES (2, 15);
INSERT INTO `admin_role_auth` VALUES (2, 20);
INSERT INTO `admin_role_auth` VALUES (2, 21);
INSERT INTO `admin_role_auth` VALUES (2, 22);
INSERT INTO `admin_role_auth` VALUES (2, 23);
INSERT INTO `admin_role_auth` VALUES (2, 24);
INSERT INTO `admin_role_auth` VALUES (3, 15);
INSERT INTO `admin_role_auth` VALUES (3, 24);

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `login_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `real_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '真实姓名',
  `password` char(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `role_ids` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '角色id字符串，如：2,3,4',
  `phone` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '0' COMMENT '手机号码',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `salt` char(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码盐',
  `last_login` int(11) NOT NULL DEFAULT 0 COMMENT '最后登录时间',
  `last_ip` char(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `status` tinyint(4) NOT NULL DEFAULT 0 COMMENT '状态，1-正常 0禁用',
  `create_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者ID',
  `update_id` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改者ID',
  `create_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改时间',
  `delete` tinyint(1) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_user_name`(`login_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES (1, 'admin', '超级管理员', '5aed881ef3158b043df23320f9d7d952', '0', '13588888888', '535661527@qq.com', 'ojW6', 1573900576, '127.0.0.1', 1, 0, 1, 0, 1573889357, 0);

SET FOREIGN_KEY_CHECKS = 1;
