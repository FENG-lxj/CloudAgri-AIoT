/*
 Navicat Premium Data Transfer

 Source Server         : MySQL (替换为你的数据库信息)
 Source Server Type    : MySQL
 Source Server Version : 50729
 Source Host           : your-host:3306
 Source Schema         : iota

 Target Server Type    : MySQL
 Target Server Version : 50729
 File Encoding         : 65001

 Date: 21/06/2025 13:22:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ai_models
-- ----------------------------
DROP TABLE IF EXISTS `ai_models`;
CREATE TABLE `ai_models`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '设备模型主键ID',
  `types` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '模型类型',
  `devid` int(11) NOT NULL COMMENT '关联设备的ID',
  `cropper` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '种植的作物名称',
  `longitude` double NOT NULL COMMENT '设备所在位置的经度',
  `dimension` double NOT NULL COMMENT '设备所在位置的纬度',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建的时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '记录最后更新的时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '记录软删除的时间戳（如果是未删除状态则为NULL）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'AI模型数据表，存储与AI相关的设备及其位置和相关作物信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of ai_models
-- ----------------------------

-- ----------------------------
-- Table structure for apikey
-- ----------------------------
DROP TABLE IF EXISTS `apikey`;
CREATE TABLE `apikey`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'API密钥模型主键ID',
  `apikey` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'API密钥',
  `status` int(11) NOT NULL COMMENT '状态（例如：启用、禁用）',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建的时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '记录最后更新的时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '记录软删除的时间戳（如果是未删除状态则为NULL）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '存储API密钥及其状态的数据表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of apikey
-- ----------------------------
INSERT INTO `apikey` VALUES (1, 'your-jwt-token-example', 1, '2025-04-23 16:50:02', '2025-04-23 16:50:06', NULL);

-- ----------------------------
-- Table structure for cropper_image
-- ----------------------------
DROP TABLE IF EXISTS `cropper_image`;
CREATE TABLE `cropper_image`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '作物模型主键ID',
  `ICID` int(11) NULL DEFAULT NULL COMMENT '内部关联ID，不对外展示',
  `Cropper` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '种植的作物名称，不对外展示',
  `Pest` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '害虫信息',
  `IdentifyID` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '识别ID',
  `Datetime` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '日期时间',
  `ImagePath` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图片路径，不对外展示',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建的时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '记录最后更新的时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '记录软删除的时间戳（如果是未删除状态则为NULL）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '存储与作物相关的检测信息，包括害虫和识别详情' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of cropper_image
-- ----------------------------
INSERT INTO `cropper_image` VALUES (2, 1, '玉米', '草地贪夜蛾', 'YMCDT120250424161736K', '2025-04-24 00:05:54', '/xm/iota/web/public/storage/uploaded/Cropper/YMCDT120250424161736K.png', NULL, NULL, NULL);
INSERT INTO `cropper_image` VALUES (3, 1, '水稻', '稻瘟病', 'BLAST120250604180308E', '2025-06-04 18:03:08', '/xm/iota/web/public/storage/uploaded/Cropper/BLAST120250604180308E.jpg', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for dev_manage
-- ----------------------------
DROP TABLE IF EXISTS `dev_manage`;
CREATE TABLE `dev_manage`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '设备管理模型主键ID',
  `types` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '设备类型，不对外展示',
  `devid` int(11) NOT NULL COMMENT '设备ID',
  `cropper` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '种植的作物名称',
  `longitude` double NOT NULL COMMENT '设备所在位置的经度',
  `dimension` double NOT NULL COMMENT '设备所在位置的纬度',
  `status` int(11) NULL DEFAULT NULL COMMENT '设备当前状态',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建的时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '记录最后更新的时间',
  `LatestNewsTime` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '最新信息时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '存储设备管理相关信息，包括设备类型、位置及状态' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of dev_manage
-- ----------------------------
INSERT INTO `dev_manage` VALUES (4, 'AE', 1, '玉米', 111.63773156370571, 40.88971939380419, 1, '2025-04-23 22:45:26', '2025-04-23 22:45:34', '2025-06-05 11:34:47');
INSERT INTO `dev_manage` VALUES (5, 'AS', 1, '大米', 111.0377315637057, 40.09971939380419, 1, '2025-04-23 22:45:29', '2025-04-23 22:45:35', '2025-06-09 23:30:44');
INSERT INTO `dev_manage` VALUES (6, 'IC', 1, '水稻', 110.43773156370571, 42.87971939380419, 1, '2025-04-23 22:45:30', '2025-04-23 22:45:37', '2025-06-05 11:34:47');

-- ----------------------------
-- Table structure for fire_image
-- ----------------------------
DROP TABLE IF EXISTS `fire_image`;
CREATE TABLE `fire_image`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '火灾模型主键ID',
  `ICID` int(11) NULL DEFAULT NULL COMMENT '内部关联ID，不对外展示',
  `IdentifyID` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '识别ID',
  `Datetime` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '日期时间',
  `ImagePath` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图片路径，不对外展示',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建的时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '记录最后更新的时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '记录软删除的时间戳（如果是未删除状态则为NULL）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '存储与火灾检测相关的信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of fire_image
-- ----------------------------
INSERT INTO `fire_image` VALUES (6, 1, 'FIRE120250604174642C', '2025-06-04 17:46:42', '/xm/iota/web/public/storage/uploaded/Fire/FIRE120250604174642C.jpg', NULL, NULL, NULL);
INSERT INTO `fire_image` VALUES (7, 1, 'FIRE120250604175045Q', '2025-06-04 17:50:45', '/xm/iota/web/public/storage/uploaded/Fire/FIRE120250604175045Q.jpg', NULL, NULL, NULL);
INSERT INTO `fire_image` VALUES (8, 1, 'FIRE120250604175058U', '2025-06-04 17:50:58', '/xm/iota/web/public/storage/uploaded/Fire/FIRE120250604175058U.jpg', NULL, NULL, NULL);
INSERT INTO `fire_image` VALUES (9, 1, 'FIRE120250604180048L', '2025-06-04 18:00:48', '/xm/iota/web/public/storage/uploaded/Fire/FIRE120250604180048L.jpg', NULL, NULL, NULL);
INSERT INTO `fire_image` VALUES (10, 1, 'FIRE120250604180058U', '2025-06-04 18:00:58', '/xm/iota/web/public/storage/uploaded/Fire/FIRE120250604180058U.jpg', NULL, NULL, NULL);
INSERT INTO `fire_image` VALUES (11, 1, 'FIRE120250604180111G', '2025-06-04 18:01:11', '/xm/iota/web/public/storage/uploaded/Fire/FIRE120250604180111G.jpg', NULL, NULL, NULL);
INSERT INTO `fire_image` VALUES (12, 1, 'FIRE120250604180121H', '2025-06-04 18:01:21', '/xm/iota/web/public/storage/uploaded/Fire/FIRE120250604180121H.jpg', NULL, NULL, NULL);
INSERT INTO `fire_image` VALUES (13, 1, 'FIRE120250604180131U', '2025-06-04 18:01:31', '/xm/iota/web/public/storage/uploaded/Fire/FIRE120250604180131U.jpg', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for invade_image
-- ----------------------------
DROP TABLE IF EXISTS `invade_image`;
CREATE TABLE `invade_image`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '入侵模型主键ID',
  `ICID` int(11) NULL DEFAULT NULL COMMENT '内部关联ID，不对外展示',
  `InvadeClass` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '入侵类别',
  `IdentifyID` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '识别ID',
  `Datetime` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '日期时间',
  `ImagePath` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '图片路径，不对外展示',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建的时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '记录最后更新的时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '记录软删除的时间戳（如果是未删除状态则为NULL）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 183 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '存储与入侵检测相关的信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of invade_image
-- ----------------------------
INSERT INTO `invade_image` VALUES (132, 1, '人物入侵', 'PERSO120250604174732Z', '2025-06-04 17:47:32', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604174732Z.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (133, 1, '人物入侵', 'PERSO120250604174754K', '2025-06-04 17:47:54', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604174754K.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (134, 1, '人物入侵', 'PERSO120250604174804X', '2025-06-04 17:48:04', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604174804X.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (135, 1, '人物入侵', 'PERSO120250604174816V', '2025-06-04 17:48:16', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604174816V.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (136, 1, '人物入侵', 'PERSO120250604175051S', '2025-06-04 17:50:51', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175051S.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (137, 1, '人物入侵', 'PERSO120250604175102Q', '2025-06-04 17:51:02', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175102Q.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (138, 1, '人物入侵', 'PERSO120250604175112Y', '2025-06-04 17:51:12', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175112Y.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (139, 1, '人物入侵', 'PERSO120250604175122D', '2025-06-04 17:51:22', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175122D.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (140, 1, '人物入侵', 'PERSO120250604175132D', '2025-06-04 17:51:32', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175132D.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (141, 1, '人物入侵', 'PERSO120250604175143Q', '2025-06-04 17:51:43', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175143Q.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (142, 1, '人物入侵', 'PERSO120250604175237W', '2025-06-04 17:52:37', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175237W.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (143, 1, '人物入侵', 'PERSO120250604175437C', '2025-06-04 17:54:38', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175437C.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (144, 1, '人物入侵', 'PERSO120250604175448L', '2025-06-04 17:54:48', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175448L.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (145, 1, '人物入侵', 'PERSO120250604175458K', '2025-06-04 17:54:58', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175458K.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (146, 1, '人物入侵', 'PERSO120250604175508V', '2025-06-04 17:55:08', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175508V.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (147, 1, '人物入侵', 'PERSO120250604175523A', '2025-06-04 17:55:23', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175523A.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (148, 1, '人物入侵', 'PERSO120250604175534C', '2025-06-04 17:55:34', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175534C.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (149, 1, '人物入侵', 'PERSO120250604175544S', '2025-06-04 17:55:44', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175544S.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (150, 1, '人物入侵', 'PERSO120250604175554C', '2025-06-04 17:55:54', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175554C.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (151, 1, '人物入侵', 'PERSO120250604175604V', '2025-06-04 17:56:04', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175604V.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (152, 1, '人物入侵', 'PERSO120250604175614Y', '2025-06-04 17:56:14', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175614Y.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (153, 1, '人物入侵', 'PERSO120250604175624G', '2025-06-04 17:56:24', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175624G.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (154, 1, '人物入侵', 'PERSO120250604175634C', '2025-06-04 17:56:34', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175634C.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (155, 1, '人物入侵', 'PERSO120250604175645N', '2025-06-04 17:56:45', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175645N.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (156, 1, '人物入侵', 'PERSO120250604175656F', '2025-06-04 17:56:56', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175656F.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (157, 1, '人物入侵', 'PERSO120250604175707N', '2025-06-04 17:57:07', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175707N.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (158, 1, '人物入侵', 'PERSO120250604175717Q', '2025-06-04 17:57:17', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175717Q.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (159, 1, '人物入侵', 'PERSO120250604175730Q', '2025-06-04 17:57:30', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175730Q.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (160, 1, '人物入侵', 'PERSO120250604175740Q', '2025-06-04 17:57:40', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175740Q.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (161, 1, '人物入侵', 'PERSO120250604175751F', '2025-06-04 17:57:51', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175751F.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (162, 1, '人物入侵', 'PERSO120250604175803C', '2025-06-04 17:58:03', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175803C.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (163, 1, '人物入侵', 'PERSO120250604175814R', '2025-06-04 17:58:14', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175814R.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (164, 1, '人物入侵', 'PERSO120250604175824O', '2025-06-04 17:58:24', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175824O.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (165, 1, '人物入侵', 'PERSO120250604175835O', '2025-06-04 17:58:35', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175835O.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (166, 1, '人物入侵', 'PERSO120250604175845Y', '2025-06-04 17:58:45', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175845Y.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (167, 1, '人物入侵', 'PERSO120250604175855I', '2025-06-04 17:58:55', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175855I.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (168, 1, '人物入侵', 'PERSO120250604175905J', '2025-06-04 17:59:05', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175905J.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (169, 1, '人物入侵', 'PERSO120250604175916W', '2025-06-04 17:59:16', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175916W.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (170, 1, '人物入侵', 'PERSO120250604175928X', '2025-06-04 17:59:28', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175928X.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (171, 1, '人物入侵', 'PERSO120250604175938P', '2025-06-04 17:59:38', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175938P.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (172, 1, '人物入侵', 'PERSO120250604175953Y', '2025-06-04 17:59:53', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604175953Y.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (173, 1, '人物入侵', 'PERSO120250604180004T', '2025-06-04 18:00:04', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604180004T.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (174, 1, '人物入侵', 'PERSO120250604180014D', '2025-06-04 18:00:14', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604180014D.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (175, 1, '人物入侵', 'PERSO120250604180024K', '2025-06-04 18:00:24', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604180024K.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (176, 1, '人物入侵', 'PERSO120250604180048D', '2025-06-04 18:00:48', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604180048D.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (177, 1, '人物入侵', 'PERSO120250604180109L', '2025-06-04 18:01:09', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604180109L.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (178, 1, '人物入侵', 'PERSO120250604180133Y', '2025-06-04 18:01:33', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604180133Y.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (179, 1, '人物入侵', 'PERSO120250604180209B', '2025-06-04 18:02:09', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604180209B.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (180, 1, '人物入侵', 'PERSO120250604180319H', '2025-06-04 18:03:19', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604180319H.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (181, 1, '人物入侵', 'PERSO120250604180346Y', '2025-06-04 18:03:46', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604180346Y.jpg', NULL, NULL, NULL);
INSERT INTO `invade_image` VALUES (182, 1, '人物入侵', 'PERSO120250604180356O', '2025-06-04 18:03:56', '/xm/iota/web/public/storage/uploaded/Invade/PERSO120250604180356O.jpg', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for monitor_logs
-- ----------------------------
DROP TABLE IF EXISTS `monitor_logs`;
CREATE TABLE `monitor_logs`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志模型主键ID',
  `icid` int(11) NULL DEFAULT NULL COMMENT '内部关联ID',
  `datetime` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '日期时间',
  `Level1Class` int(11) NULL DEFAULT NULL COMMENT '一级分类',
  `Log` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '日志内容',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建的时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '记录最后更新的时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '记录软删除的时间戳（如果是未删除状态则为NULL）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1227 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '存储监控日志信息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of monitor_logs
-- ----------------------------
INSERT INTO `monitor_logs` VALUES (801, 1, '2025-04-30 14:00:05', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (802, 1, '2025-04-30 14:00:17', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (803, 1, '2025-04-30 14:00:21', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (804, 1, '2025-04-30 14:00:57', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (805, 1, '2025-04-30 14:01:33', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (806, 1, '2025-04-30 14:01:50', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (807, 1, '2025-04-30 14:01:58', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (808, 1, '2025-04-30 14:02:18', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (809, 1, '2025-04-30 14:02:54', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (810, 1, '2025-04-30 14:02:58', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (811, 1, '2025-04-30 14:03:18', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (812, 1, '2025-04-30 14:03:27', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (813, 1, '2025-04-30 14:04:27', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (814, 1, '2025-04-30 14:04:31', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (815, 1, '2025-04-30 14:04:55', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (816, 1, '2025-04-30 14:04:59', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (817, 1, '2025-04-30 14:05:16', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (818, 1, '2025-04-30 14:05:20', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (819, 1, '2025-04-30 14:05:24', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (820, 1, '2025-04-30 14:06:32', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (821, 1, '2025-04-30 14:06:44', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (822, 1, '2025-04-30 14:06:53', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (823, 1, '2025-04-30 14:07:21', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (824, 1, '2025-04-30 14:07:29', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (825, 1, '2025-04-30 14:07:53', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (826, 1, '2025-06-04 17:31:24', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (827, 1, '2025-06-04 17:31:36', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (828, 1, '2025-06-04 17:31:40', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (829, 1, '2025-06-04 17:32:17', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (830, 1, '2025-06-04 17:32:53', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (831, 1, '2025-06-04 17:33:09', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (832, 1, '2025-06-04 17:33:17', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (833, 1, '2025-06-04 17:33:37', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (834, 1, '2025-06-04 17:34:25', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (835, 1, '2025-06-04 17:34:37', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (836, 1, '2025-06-04 17:34:41', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (837, 1, '2025-06-04 17:35:17', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (838, 1, '2025-06-04 17:35:53', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (839, 1, '2025-06-04 17:36:10', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (840, 1, '2025-06-04 17:36:18', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (841, 1, '2025-06-04 17:36:38', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (842, 1, '2025-06-04 17:37:14', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (843, 1, '2025-06-04 17:37:19', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (844, 1, '2025-06-04 17:37:39', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (845, 1, '2025-06-04 17:37:47', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (846, 1, '2025-06-04 17:38:48', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (847, 1, '2025-06-04 17:38:52', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (848, 1, '2025-06-04 17:39:16', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (849, 1, '2025-06-04 17:39:20', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (850, 1, '2025-06-04 17:39:36', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (851, 1, '2025-06-04 17:39:40', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (852, 1, '2025-06-04 17:39:44', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (853, 1, '2025-06-04 17:39:49', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (854, 1, '2025-06-04 17:40:53', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (855, 1, '2025-06-04 17:41:05', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (856, 1, '2025-06-04 17:41:13', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (857, 1, '2025-06-04 17:41:41', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (858, 1, '2025-06-04 17:41:49', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (859, 1, '2025-06-04 17:42:14', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (860, 1, '2025-06-04 17:42:26', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (861, 1, '2025-06-04 17:42:30', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (862, 1, '2025-06-04 17:42:39', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (863, 1, '2025-06-04 17:42:47', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (864, 1, '2025-06-04 17:43:08', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (865, 1, '2025-06-04 17:43:32', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (866, 1, '2025-06-04 17:44:40', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (867, 1, '2025-06-04 17:44:45', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (868, 1, '2025-06-04 17:46:42', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (869, 1, '2025-06-04 17:47:32', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (870, 1, '2025-06-04 17:47:54', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (871, 1, '2025-06-04 17:48:04', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (872, 1, '2025-06-04 17:48:16', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (873, 1, '2025-06-04 17:50:45', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (874, 1, '2025-06-04 17:50:51', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (875, 1, '2025-06-04 17:50:58', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (876, 1, '2025-06-04 17:51:02', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (877, 1, '2025-06-04 17:51:12', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (878, 1, '2025-06-04 17:51:22', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (879, 1, '2025-06-04 17:51:32', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (880, 1, '2025-06-04 17:51:43', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (881, 1, '2025-06-04 17:52:37', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (882, 1, '2025-06-04 17:54:37', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (883, 1, '2025-06-04 17:54:48', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (884, 1, '2025-06-04 17:54:58', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (885, 1, '2025-06-04 17:55:08', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (886, 1, '2025-06-04 17:55:23', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (887, 1, '2025-06-04 17:55:34', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (888, 1, '2025-06-04 17:55:44', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (889, 1, '2025-06-04 17:55:54', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (890, 1, '2025-06-04 17:56:04', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (891, 1, '2025-06-04 17:56:14', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (892, 1, '2025-06-04 17:56:24', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (893, 1, '2025-06-04 17:56:34', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (894, 1, '2025-06-04 17:56:45', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (895, 1, '2025-06-04 17:56:56', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (896, 1, '2025-06-04 17:57:07', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (897, 1, '2025-06-04 17:57:17', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (898, 1, '2025-06-04 17:57:30', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (899, 1, '2025-06-04 17:57:40', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (900, 1, '2025-06-04 17:57:51', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (901, 1, '2025-06-04 17:58:03', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (902, 1, '2025-06-04 17:58:14', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (903, 1, '2025-06-04 17:58:24', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (904, 1, '2025-06-04 17:58:35', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (905, 1, '2025-06-04 17:58:45', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (906, 1, '2025-06-04 17:58:55', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (907, 1, '2025-06-04 17:59:05', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (908, 1, '2025-06-04 17:59:16', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (909, 1, '2025-06-04 17:59:28', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (910, 1, '2025-06-04 17:59:38', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (911, 1, '2025-06-04 17:59:53', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (912, 1, '2025-06-04 18:00:04', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (913, 1, '2025-06-04 18:00:14', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (914, 1, '2025-06-04 18:00:24', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (915, 1, '2025-06-04 18:00:48', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (916, 1, '2025-06-04 18:00:48', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (917, 1, '2025-06-04 18:00:58', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (918, 1, '2025-06-04 18:01:09', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (919, 1, '2025-06-04 18:01:11', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (920, 1, '2025-06-04 18:01:21', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (921, 1, '2025-06-04 18:01:31', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (922, 1, '2025-06-04 18:01:33', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (923, 1, '2025-06-04 18:02:09', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (924, 1, '2025-06-04 18:03:08', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (925, 1, '2025-06-04 18:03:19', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (926, 1, '2025-06-04 18:03:46', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (927, 1, '2025-06-04 18:03:56', 2, '发现人物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (928, 1, '2025-06-04 21:51:23', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (929, 1, '2025-06-04 21:51:35', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (930, 1, '2025-06-04 21:51:39', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (931, 1, '2025-06-04 21:52:16', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (932, 1, '2025-06-04 21:52:52', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (933, 1, '2025-06-04 21:53:08', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (934, 1, '2025-06-04 21:53:17', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (935, 1, '2025-06-04 21:53:37', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (936, 1, '2025-06-04 21:54:13', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (937, 1, '2025-06-04 21:54:17', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (938, 1, '2025-06-04 21:54:38', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (939, 1, '2025-06-04 21:54:46', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (940, 1, '2025-06-04 21:55:46', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (941, 1, '2025-06-04 21:55:50', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (942, 1, '2025-06-04 21:56:15', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (943, 1, '2025-06-04 21:56:19', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (944, 1, '2025-06-04 21:56:35', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (945, 1, '2025-06-04 21:56:39', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (946, 1, '2025-06-04 21:56:43', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (947, 1, '2025-06-04 21:57:52', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (948, 1, '2025-06-04 21:58:04', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (949, 1, '2025-06-04 21:58:14', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (950, 1, '2025-06-04 21:58:42', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (951, 1, '2025-06-04 21:58:50', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (952, 1, '2025-06-04 21:59:15', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (953, 1, '2025-06-04 21:59:27', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (954, 1, '2025-06-04 21:59:31', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (955, 1, '2025-06-04 21:59:39', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (956, 1, '2025-06-04 21:59:47', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (957, 1, '2025-06-04 22:00:07', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (958, 1, '2025-06-04 22:00:32', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (959, 1, '2025-06-04 22:01:41', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (960, 1, '2025-06-04 22:01:45', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (961, 1, '2025-06-04 22:01:49', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (962, 1, '2025-06-04 22:02:05', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (963, 1, '2025-06-04 22:02:10', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (964, 1, '2025-06-04 22:02:18', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (965, 1, '2025-06-04 22:02:26', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (966, 1, '2025-06-04 22:02:43', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (967, 1, '2025-06-04 22:03:03', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (968, 1, '2025-06-04 22:03:35', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (969, 1, '2025-06-04 22:03:39', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (970, 1, '2025-06-04 22:03:44', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (971, 1, '2025-06-04 22:03:48', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (972, 1, '2025-06-04 22:04:12', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (973, 1, '2025-06-04 22:04:24', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (974, 1, '2025-06-04 22:04:40', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (975, 1, '2025-06-04 22:04:44', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (976, 1, '2025-06-04 22:04:53', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (977, 1, '2025-06-04 22:05:13', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (978, 1, '2025-06-04 22:05:17', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (979, 1, '2025-06-04 22:05:42', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (980, 1, '2025-06-04 22:05:46', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (981, 1, '2025-06-04 22:05:50', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (982, 1, '2025-06-04 22:06:02', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (983, 1, '2025-06-04 22:06:31', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (984, 1, '2025-06-04 22:06:35', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (985, 1, '2025-06-04 22:06:50', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (986, 1, '2025-06-04 22:07:06', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (987, 1, '2025-06-04 22:07:10', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (988, 1, '2025-06-04 22:08:23', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (989, 1, '2025-06-04 22:08:43', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (990, 1, '2025-06-04 22:08:47', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (991, 1, '2025-06-04 22:08:55', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (992, 1, '2025-06-04 22:08:59', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (993, 1, '2025-06-04 22:09:03', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (994, 1, '2025-06-04 22:09:07', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (995, 1, '2025-06-04 22:09:11', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (996, 1, '2025-06-04 22:09:40', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (997, 1, '2025-06-04 22:10:08', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (998, 1, '2025-06-04 22:10:36', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (999, 1, '2025-06-04 22:10:49', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1000, 1, '2025-06-04 22:10:57', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1001, 1, '2025-06-04 22:11:09', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1002, 1, '2025-06-04 22:11:26', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1003, 1, '2025-06-04 22:11:30', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1004, 1, '2025-06-04 22:11:38', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1005, 1, '2025-06-04 22:11:42', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1006, 1, '2025-06-04 22:12:02', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1007, 1, '2025-06-04 22:12:14', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1008, 1, '2025-06-04 22:12:18', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1009, 1, '2025-06-04 22:12:23', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1010, 1, '2025-06-04 22:12:39', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1011, 1, '2025-06-04 22:12:43', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1012, 1, '2025-06-04 22:12:47', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1013, 1, '2025-06-04 22:12:55', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1014, 1, '2025-06-04 22:13:15', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1015, 1, '2025-06-04 22:13:23', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1016, 1, '2025-06-04 22:13:27', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1017, 1, '2025-06-04 22:13:31', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1018, 1, '2025-06-04 22:13:40', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1019, 1, '2025-06-04 22:13:44', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1020, 1, '2025-06-04 22:13:52', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1021, 1, '2025-06-04 22:13:56', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1022, 1, '2025-06-04 22:14:04', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1023, 1, '2025-06-04 22:14:12', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1024, 1, '2025-06-04 22:14:57', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1025, 1, '2025-06-04 22:15:54', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1026, 1, '2025-06-04 22:16:23', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1027, 1, '2025-06-04 22:16:31', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1028, 1, '2025-06-04 22:16:39', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1029, 1, '2025-06-04 22:16:43', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1030, 1, '2025-06-04 22:16:47', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1031, 1, '2025-06-04 22:17:07', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1032, 1, '2025-06-04 22:17:16', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1033, 1, '2025-06-04 22:17:29', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1034, 1, '2025-06-05 10:49:31', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1035, 1, '2025-06-05 10:49:43', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1036, 1, '2025-06-05 10:49:47', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1037, 1, '2025-06-05 10:50:24', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1038, 1, '2025-06-05 10:51:00', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1039, 1, '2025-06-05 10:51:16', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1040, 1, '2025-06-05 10:51:24', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1041, 1, '2025-06-05 10:51:45', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1042, 1, '2025-06-05 10:52:21', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1043, 1, '2025-06-05 10:52:25', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1044, 1, '2025-06-05 10:52:45', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1045, 1, '2025-06-05 10:52:53', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1046, 1, '2025-06-05 10:53:54', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1047, 1, '2025-06-05 10:53:58', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1048, 1, '2025-06-05 10:54:22', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1049, 1, '2025-06-05 10:54:26', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1050, 1, '2025-06-05 10:54:43', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1051, 1, '2025-06-05 10:54:47', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1052, 1, '2025-06-05 10:54:51', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1053, 1, '2025-06-05 10:55:59', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1054, 1, '2025-06-05 10:56:12', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1055, 1, '2025-06-05 10:56:20', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1056, 1, '2025-06-05 10:56:48', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1057, 1, '2025-06-05 10:56:56', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1058, 1, '2025-06-05 10:57:20', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1059, 1, '2025-06-05 10:57:33', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1060, 1, '2025-06-05 10:57:37', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1061, 1, '2025-06-05 10:57:45', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1062, 1, '2025-06-05 10:57:53', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1063, 1, '2025-06-05 10:58:13', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1064, 1, '2025-06-05 10:58:37', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1065, 1, '2025-06-05 10:59:46', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1066, 1, '2025-06-05 10:59:50', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1067, 1, '2025-06-05 10:59:54', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1068, 1, '2025-06-05 11:00:10', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1069, 1, '2025-06-05 11:00:14', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1070, 1, '2025-06-05 11:00:23', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1071, 1, '2025-06-05 11:00:31', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1072, 1, '2025-06-05 11:00:47', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1073, 1, '2025-06-05 11:01:07', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1074, 1, '2025-06-05 11:01:39', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1075, 1, '2025-06-05 11:01:43', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1076, 1, '2025-06-05 11:01:48', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1077, 1, '2025-06-05 11:01:52', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1078, 1, '2025-06-05 11:02:16', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1079, 1, '2025-06-05 11:02:28', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1080, 1, '2025-06-05 11:02:44', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1081, 1, '2025-06-05 11:02:48', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1082, 1, '2025-06-05 11:02:56', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1083, 1, '2025-06-05 11:03:17', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1084, 1, '2025-06-05 11:03:21', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1085, 1, '2025-06-05 11:03:45', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1086, 1, '2025-06-05 11:03:49', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1087, 1, '2025-06-05 11:03:53', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1088, 1, '2025-06-05 11:04:05', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1089, 1, '2025-06-05 11:04:34', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1090, 1, '2025-06-05 11:04:38', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1091, 1, '2025-06-05 11:04:50', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1092, 1, '2025-06-05 11:05:06', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1093, 1, '2025-06-05 11:05:10', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1094, 1, '2025-06-05 11:06:23', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1095, 1, '2025-06-05 11:06:43', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1096, 1, '2025-06-05 11:06:47', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1097, 1, '2025-06-05 11:06:55', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1098, 1, '2025-06-05 11:06:59', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1099, 1, '2025-06-05 11:07:03', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1100, 1, '2025-06-05 11:07:07', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1101, 1, '2025-06-05 11:07:11', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1102, 1, '2025-06-05 11:07:40', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1103, 1, '2025-06-05 11:08:08', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1104, 1, '2025-06-05 11:08:36', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1105, 1, '2025-06-05 11:08:49', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1106, 1, '2025-06-05 11:08:57', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1107, 1, '2025-06-05 11:09:09', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1108, 1, '2025-06-05 11:09:25', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1109, 1, '2025-06-05 11:09:29', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1110, 1, '2025-06-05 11:09:37', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1111, 1, '2025-06-05 11:09:41', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1112, 1, '2025-06-05 11:10:01', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1113, 1, '2025-06-05 11:10:14', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1114, 1, '2025-06-05 11:10:18', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1115, 1, '2025-06-05 11:10:22', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1116, 1, '2025-06-05 11:10:38', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1117, 1, '2025-06-05 11:10:42', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1118, 1, '2025-06-05 11:10:46', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1119, 1, '2025-06-05 11:10:54', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1120, 1, '2025-06-05 11:11:15', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1121, 1, '2025-06-05 11:11:23', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1122, 1, '2025-06-05 11:11:27', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1123, 1, '2025-06-05 11:11:31', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1124, 1, '2025-06-05 11:11:39', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1125, 1, '2025-06-05 11:11:43', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1126, 1, '2025-06-05 11:11:51', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1127, 1, '2025-06-05 11:11:55', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1128, 1, '2025-06-05 11:12:03', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1129, 1, '2025-06-05 11:12:12', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1130, 1, '2025-06-05 11:12:56', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1131, 1, '2025-06-05 11:13:53', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1132, 1, '2025-06-05 11:14:21', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1133, 1, '2025-06-05 11:14:29', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1134, 1, '2025-06-05 11:14:37', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1135, 1, '2025-06-05 11:14:41', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1136, 1, '2025-06-05 11:14:46', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1137, 1, '2025-06-05 11:15:06', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1138, 1, '2025-06-05 11:15:14', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1139, 1, '2025-06-05 11:15:26', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1140, 1, '2025-06-05 11:15:42', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1141, 1, '2025-06-05 11:15:46', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1142, 1, '2025-06-05 11:15:58', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1143, 1, '2025-06-05 11:16:19', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1144, 1, '2025-06-05 11:16:23', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1145, 1, '2025-06-05 11:16:35', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1146, 1, '2025-06-05 11:17:44', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1147, 1, '2025-06-05 11:17:48', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1148, 1, '2025-06-05 11:17:52', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1149, 1, '2025-06-05 11:18:08', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1150, 1, '2025-06-05 11:18:16', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1151, 1, '2025-06-05 11:18:24', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1152, 1, '2025-06-05 11:18:32', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1153, 1, '2025-06-05 11:18:36', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1154, 1, '2025-06-05 11:18:49', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1155, 1, '2025-06-05 11:18:53', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1156, 1, '2025-06-05 11:18:57', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1157, 1, '2025-06-05 11:19:29', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1158, 1, '2025-06-05 11:19:33', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1159, 1, '2025-06-05 11:19:41', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1160, 1, '2025-06-05 11:19:49', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1161, 1, '2025-06-05 11:20:09', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1162, 1, '2025-06-05 11:20:14', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1163, 1, '2025-06-05 11:20:18', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1164, 1, '2025-06-05 11:20:22', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1165, 1, '2025-06-05 11:20:42', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1166, 1, '2025-06-05 11:20:54', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1167, 1, '2025-06-05 11:20:58', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1168, 1, '2025-06-05 11:21:06', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1169, 1, '2025-06-05 11:21:35', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1170, 1, '2025-06-05 11:22:11', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1171, 1, '2025-06-05 11:22:23', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1172, 1, '2025-06-05 11:22:31', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1173, 1, '2025-06-05 11:22:43', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1174, 1, '2025-06-05 11:23:12', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1175, 1, '2025-06-05 11:23:24', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1176, 1, '2025-06-05 11:23:28', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1177, 1, '2025-06-05 11:23:48', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1178, 1, '2025-06-05 11:23:56', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1179, 1, '2025-06-05 11:24:08', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1180, 1, '2025-06-05 11:24:17', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1181, 1, '2025-06-05 11:24:41', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1182, 1, '2025-06-05 11:24:45', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1183, 1, '2025-06-05 11:25:01', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1184, 1, '2025-06-05 11:25:05', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1185, 1, '2025-06-05 11:25:13', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1186, 1, '2025-06-05 11:25:17', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1187, 1, '2025-06-05 11:25:42', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1188, 1, '2025-06-05 11:25:46', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1189, 1, '2025-06-05 11:26:06', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1190, 1, '2025-06-05 11:26:14', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1191, 1, '2025-06-05 11:26:22', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1192, 1, '2025-06-05 11:26:30', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1193, 1, '2025-06-05 11:26:34', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1194, 1, '2025-06-05 11:26:39', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1195, 1, '2025-06-05 11:26:51', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1196, 1, '2025-06-05 11:26:55', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1197, 1, '2025-06-05 11:26:59', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1198, 1, '2025-06-05 11:27:35', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1199, 1, '2025-06-05 11:27:51', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1200, 1, '2025-06-05 11:27:55', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1201, 1, '2025-06-05 11:28:12', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1202, 1, '2025-06-05 11:28:48', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1203, 1, '2025-06-05 11:29:08', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1204, 1, '2025-06-05 11:29:21', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1205, 1, '2025-06-05 11:29:29', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1206, 1, '2025-06-05 11:29:49', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1207, 1, '2025-06-05 11:29:53', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1208, 1, '2025-06-05 11:30:34', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1209, 1, '2025-06-05 11:30:46', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1210, 1, '2025-06-05 11:30:58', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1211, 1, '2025-06-05 11:31:18', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1212, 1, '2025-06-05 11:31:31', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1213, 1, '2025-06-05 11:31:35', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1214, 1, '2025-06-05 11:31:55', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1215, 1, '2025-06-05 11:32:03', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1216, 1, '2025-06-05 11:32:35', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1217, 1, '2025-06-05 11:32:40', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1218, 1, '2025-06-05 11:32:44', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1219, 1, '2025-06-05 11:33:04', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1220, 1, '2025-06-05 11:33:28', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1221, 1, '2025-06-05 11:33:52', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1222, 1, '2025-06-05 11:34:09', 2, '发现动物入侵', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1223, 1, '2025-06-05 11:34:13', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1224, 1, '2025-06-05 11:34:21', 1, '发现病虫害', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1225, 1, '2025-06-05 11:34:29', 3, '发现火灾', NULL, NULL, NULL);
INSERT INTO `monitor_logs` VALUES (1226, 1, '2025-06-05 11:34:37', 3, '发现火灾', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for pest_proposal
-- ----------------------------
DROP TABLE IF EXISTS `pest_proposal`;
CREATE TABLE `pest_proposal`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '害虫建议模型主键ID',
  `pest` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '害虫名称',
  `proposal` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '针对害虫的建议措施',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '记录创建的时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '记录最后更新的时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '记录软删除的时间戳（如果是未删除状态则为NULL）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '存储关于害虫的建议措施' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pest_proposal
-- ----------------------------
INSERT INTO `pest_proposal` VALUES (1, '稻瘟病', '为了有效防治稻瘟病，建议选择抗病品种，播种前对种子进行消毒处理，如使用药剂浸种；在水稻生长期间，合理施肥，避免偏施氮肥，增施磷、钾肥以增强植株抗病能力，同时注意田间排水和通风透光；在发病初期及时使用合适的药剂进行喷雾防治，如三环唑、稻瘟灵等，按照说明书推荐的剂量和方法施用，以减少病害发生和蔓延。', '2025-04-23 22:52:56', '2025-04-23 22:52:58', NULL);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `Id` bigint(20) NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `HeadImg` int(11) NULL DEFAULT NULL,
  `last_login_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `last_login_ipAddr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `createdTime` datetime NULL DEFAULT NULL,
  `isSuper` tinyint(4) NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2025-04-23 07:04:09', '2025-04-23 07:04:15', '13800000001', 1, '127.0.0.1', '', '2025-04-23 17:02:12', 0);
INSERT INTO `user` VALUES (NULL, NULL, NULL, '13800000002', 1, '127.0.0.1', '', '2025-04-24 15:53:30', 0);
INSERT INTO `user` VALUES (NULL, NULL, NULL, '13800000003', 1, '127.0.0.1', '', '2025-04-24 15:55:45', 0);
INSERT INTO `user` VALUES (NULL, NULL, NULL, '13800000000', 1, '127.0.0.1', '', '2025-05-31 20:27:32', 1);

-- ----------------------------
-- Table structure for userToken
-- ----------------------------
DROP TABLE IF EXISTS `userToken`;
CREATE TABLE `userToken`  (
  `Id` bigint(20) NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `expires_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `client_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of userToken
-- ----------------------------
INSERT INTO `userToken` VALUES (NULL, '2025-04-23 17:07:30', NULL, '13800000001', 'your-jwt-token-example', '2025-05-23 17:07:30', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-23 21:33:50', NULL, '13800000001', 'your-jwt-token-example', '2025-05-23 21:33:50', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-24 00:16:31', NULL, '13800000001', 'your-jwt-token-example', '2025-05-24 00:16:31', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-24 07:29:36', NULL, '13800000001', 'your-jwt-token-example', '2025-05-24 07:29:36', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-24 15:51:55', NULL, '13800000001', 'your-jwt-token-example', '2025-05-24 15:51:55', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-24 16:41:25', NULL, '13800000001', 'your-jwt-token-example', '2025-05-24 16:41:25', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-24 16:56:16', NULL, '13800000001', 'your-jwt-token-example', '2025-05-24 16:56:16', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-24 17:33:52', NULL, '13800000001', 'your-jwt-token-example', '2025-05-24 17:33:52', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-24 17:37:13', NULL, '13800000001', 'your-jwt-token-example', '2025-05-24 17:37:13', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-25 00:01:27', NULL, '13800000001', 'your-jwt-token-example', '2025-05-25 00:01:27', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-25 10:47:51', NULL, '13800000001', 'your-jwt-token-example', '2025-05-25 10:47:51', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-28 13:10:58', NULL, '13800000001', 'your-jwt-token-example', '2025-05-28 13:10:58', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-28 13:16:41', NULL, '13800000001', 'your-jwt-token-example', '2025-05-28 13:16:41', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-30 13:50:37', NULL, '13800000001', 'your-jwt-token-example', '2025-05-30 13:50:37', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-30 13:58:43', NULL, '13800000001', 'your-jwt-token-example', '2025-05-30 13:58:43', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-30 14:00:10', NULL, '13800000001', 'your-jwt-token-example', '2025-05-30 14:00:10', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-30 14:03:16', NULL, '13800000001', 'your-jwt-token-example', '2025-05-30 14:03:16', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-30 14:04:48', NULL, '13800000001', 'your-jwt-token-example', '2025-05-30 14:04:48', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-04-30 15:18:33', NULL, '13800000001', 'your-jwt-token-example', '2025-05-30 15:18:33', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-05-31 20:33:49', NULL, '13800000000', 'your-jwt-token-example', '2025-06-30 20:33:49', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-05-31 20:39:06', NULL, '13800000000', 'your-jwt-token-example', '2025-06-30 20:39:06', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-05-31 21:05:59', NULL, '13800000000', 'your-jwt-token-example', '2025-06-30 21:05:59', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-05-31 21:13:07', NULL, '13800000000', 'your-jwt-token-example', '2025-06-30 21:13:07', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-05-31 21:15:41', NULL, '13800000000', 'your-jwt-token-example', '2025-06-30 21:15:41', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-05-31 21:30:55', NULL, '13800000000', 'your-jwt-token-example', '2025-06-30 21:30:55', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-05-31 21:54:47', NULL, '13800000000', 'your-jwt-token-example', '2025-06-30 21:54:47', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-01 00:59:44', NULL, '13800000000', 'your-jwt-token-example', '2025-07-01 00:59:44', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-01 14:06:04', NULL, '13800000000', 'your-jwt-token-example', '2025-07-01 14:06:04', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-01 14:14:47', NULL, '13800000000', 'your-jwt-token-example', '2025-07-01 14:14:47', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-01 14:42:22', NULL, '13800000000', 'your-jwt-token-example', '2025-07-01 14:42:22', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-04 17:18:22', NULL, '13800000001', 'your-jwt-token-example', '2025-07-04 17:18:22', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-04 21:57:33', NULL, '13800000001', 'your-jwt-token-example', '2025-07-04 21:57:33', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-04 21:59:42', NULL, '13800000001', 'your-jwt-token-example', '2025-07-04 21:59:42', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-04 22:12:55', NULL, '13800000001', 'your-jwt-token-example', '2025-07-04 22:12:55', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-08 20:37:09', NULL, '13800000000', 'your-jwt-token-example', '2025-07-08 20:37:09', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-09 23:21:30', NULL, '13800000000', 'your-jwt-token-example', '2025-07-09 23:21:30', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-09 23:28:46', NULL, '13800000000', 'your-jwt-token-example', '2025-07-09 23:28:46', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-09 23:38:00', NULL, '13800000000', 'your-jwt-token-example', '2025-07-09 23:38:00', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-10 12:21:08', NULL, '13800000000', 'your-jwt-token-example', '2025-07-10 12:21:08', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-10 12:24:22', NULL, '13800000000', 'your-jwt-token-example', '2025-07-10 12:24:22', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-10 12:32:56', NULL, '13800000000', 'your-jwt-token-example', '2025-07-10 12:32:56', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-10 12:50:39', NULL, '13800000000', 'your-jwt-token-example', '2025-07-10 12:50:39', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-10 13:07:23', NULL, '13800000000', 'your-jwt-token-example', '2025-07-10 13:07:23', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-10 13:15:06', NULL, '13800000000', 'your-jwt-token-example', '2025-07-10 13:15:06', '127.0.0.1');
INSERT INTO `userToken` VALUES (NULL, '2025-06-10 13:20:59', NULL, '13800000000', 'your-jwt-token-example', '2025-07-10 13:20:59', '127.0.0.1');

SET FOREIGN_KEY_CHECKS = 1;
