/*
 Navicat Premium Data Transfer

 Source Server         : 本地mysql
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3306
 Source Schema         : dang

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 04/04/2023 17:58:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '姓名',
  `age` tinyint UNSIGNED NULL DEFAULT NULL COMMENT '年龄',
  `email` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `account` bigint NOT NULL COMMENT '账号',
  `password` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
  `sort` bigint NULL DEFAULT NULL COMMENT '显示顺序',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '2022-12-02 17:21:43.088', '2022-12-02 17:21:43.088', NULL, '孟洋', 50, 'k.gnlbyom@qq.com', 368860726176, '123456', NULL);
INSERT INTO `users` VALUES (2, '2022-12-02 17:21:49.011', '2022-12-02 17:21:49.011', NULL, '孟洋', 50, 'k.gnlbyom@qq.com', 368860726176, '32035218251211686X', NULL);
INSERT INTO `users` VALUES (3, '2023-03-27 10:24:46.686', '2023-03-27 10:24:46.686', NULL, 'admin', 40, 'z.kwpvysv@qq.com', 6847204468, '123456', 0);
INSERT INTO `users` VALUES (4, '2023-03-28 13:38:10.986', '2023-03-28 13:38:10.986', NULL, '郑艳', 26, 'm.krcuwh@qq.com', 413186, '612886192910100461', 0);

SET FOREIGN_KEY_CHECKS = 1;
