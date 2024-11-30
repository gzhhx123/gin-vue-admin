/*
 Navicat Premium Data Transfer

 Source Server         : AAA公司PC
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : bagique

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 30/11/2024 14:37:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bagique_brands
-- ----------------------------
DROP TABLE IF EXISTS `bagique_brands`;
CREATE TABLE `bagique_brands`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `brand_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '品牌名称',
  `brand_short_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '品牌简称',
  `brand_logo` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '品牌logo',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_bagique_brands_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bagique_brands
-- ----------------------------
INSERT INTO `bagique_brands` VALUES (1, '2024-11-05 21:02:14.856', '2024-11-11 17:09:27.814', NULL, '恐龙', 'konglong', 'uploads/file/ae0a9abcd8364ab148486157cb90cbb7_20241105210208.png', '测试');
INSERT INTO `bagique_brands` VALUES (2, '2024-11-06 21:19:53.097', '2024-11-11 17:08:36.229', NULL, '测试品牌', 'test', '', '');
INSERT INTO `bagique_brands` VALUES (3, '2024-11-11 11:53:10.299', '2024-11-21 22:06:26.393', NULL, '测试', '213', '', '');

-- ----------------------------
-- Table structure for bagique_companies
-- ----------------------------
DROP TABLE IF EXISTS `bagique_companies`;
CREATE TABLE `bagique_companies`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `company_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公司名称',
  `company_short_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公司简称',
  `company_logo` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公司logo',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_bagique_companies_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bagique_companies
-- ----------------------------
INSERT INTO `bagique_companies` VALUES (1, '2024-11-06 21:19:29.024', '2024-11-27 01:59:50.973', NULL, 'CO公司', 'coco', 'uploads/file/ae0a9abcd8364ab148486157cb90cbb7_20241105210208.png', '');
INSERT INTO `bagique_companies` VALUES (3, '2024-11-26 13:53:08.997', '2024-11-27 02:00:01.062', NULL, '蜜雪公司', '雪人', '', '');

-- ----------------------------
-- Table structure for bagique_evaluate_prices
-- ----------------------------
DROP TABLE IF EXISTS `bagique_evaluate_prices`;
CREATE TABLE `bagique_evaluate_prices`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `company_id` bigint NULL DEFAULT NULL COMMENT '公司id',
  `evaluate_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '估价记录id',
  `price` double NULL DEFAULT 0 COMMENT '估价',
  `fee` double NULL DEFAULT 0 COMMENT '估价成本',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_bagique_evaluate_prices_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bagique_evaluate_prices
-- ----------------------------
INSERT INTO `bagique_evaluate_prices` VALUES (1, '2024-11-26 21:52:35.390', '2024-11-26 21:52:54.098', '2024-11-26 21:53:02.792', 1, 7, 2, 1, '3');
INSERT INTO `bagique_evaluate_prices` VALUES (2, '2024-11-26 21:52:54.099', '2024-11-30 11:47:04.055', NULL, 3, 7, 6, 4, '价格低了，还能观望一下');
INSERT INTO `bagique_evaluate_prices` VALUES (3, '2024-11-26 21:53:16.324', '2024-11-30 11:47:04.056', NULL, 1, 7, 7, 7, '');
INSERT INTO `bagique_evaluate_prices` VALUES (4, '2024-11-27 00:38:21.928', '2024-11-27 00:38:21.928', NULL, 3, 8, 2, 22, '');

-- ----------------------------
-- Table structure for bagique_evaluates
-- ----------------------------
DROP TABLE IF EXISTS `bagique_evaluates`;
CREATE TABLE `bagique_evaluates`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `product_id` bigint NULL DEFAULT NULL COMMENT '产品id',
  `seller_id` bigint NULL DEFAULT NULL COMMENT '卖家id',
  `evaluate_pics` json NULL COMMENT '细节图',
  `status` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '估价状态',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `rate` double NULL DEFAULT 0 COMMENT '汇率',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_bagique_evaluates_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bagique_evaluates
-- ----------------------------
INSERT INTO `bagique_evaluates` VALUES (7, '2024-11-26 21:52:35.390', '2024-11-30 14:04:30.890', NULL, 15, 1, '[\"uploads/file/ae0a9abcd8364ab148486157cb90cbb7_20241105210208.png\"]', 'FINISH', '', 7.24);
INSERT INTO `bagique_evaluates` VALUES (8, '2024-11-27 00:38:21.925', '2024-11-30 14:16:59.762', NULL, 20, 1, '[\"uploads/file/ae0a9abcd8364ab148486157cb90cbb7_20241105210208.png\"]', 'WAITING', '', 7.25);

-- ----------------------------
-- Table structure for bagique_products
-- ----------------------------
DROP TABLE IF EXISTS `bagique_products`;
CREATE TABLE `bagique_products`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `brand_id` bigint NULL DEFAULT NULL COMMENT '品牌id',
  `product_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产品名称',
  `product_sku` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '产品sku',
  `refer_price_min` double NULL DEFAULT 0 COMMENT '参考价（Min）',
  `refer_price_max` double NULL DEFAULT 0 COMMENT '参考价（Max）',
  `refer_pics` json NULL COMMENT '参考图片',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `rate` double NULL DEFAULT 0 COMMENT '汇率',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_bagique_products_product_sku`(`product_sku` ASC) USING BTREE,
  INDEX `idx_bagique_products_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bagique_products
-- ----------------------------
INSERT INTO `bagique_products` VALUES (15, '2024-11-01 17:27:35.000', '2024-11-21 17:26:58.776', NULL, 1, '111', '111', 10, 20, '[]', '', 7.23);
INSERT INTO `bagique_products` VALUES (20, '2024-11-20 17:27:35.000', '2024-11-21 17:27:35.704', NULL, 3, 'aaa', 'aaa', 1, 2, '[\"uploads/file/ae0a9abcd8364ab148486157cb90cbb7_20241105210208.png\"]', 'bb', 7.24);
INSERT INTO `bagique_products` VALUES (21, '2024-11-22 00:08:55.637', '2024-11-22 00:08:55.637', NULL, 1, 'aaa', '22232', 0, 0, '[]', '', 7.24);

-- ----------------------------
-- Table structure for bagique_purchases
-- ----------------------------
DROP TABLE IF EXISTS `bagique_purchases`;
CREATE TABLE `bagique_purchases`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `evaluate_price_id` bigint NULL DEFAULT NULL COMMENT '估价公司估价id',
  `status` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '采购状态',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_bagique_purchases_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bagique_purchases
-- ----------------------------

-- ----------------------------
-- Table structure for bagique_sellers
-- ----------------------------
DROP TABLE IF EXISTS `bagique_sellers`;
CREATE TABLE `bagique_sellers`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `seller_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '卖家名称',
  `seller_platform_code` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '卖家平台id',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_bagique_sellers_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bagique_sellers
-- ----------------------------
INSERT INTO `bagique_sellers` VALUES (1, '2024-11-06 21:36:28.565', '2024-11-21 23:41:37.873', NULL, 'AAA', '', '');

-- ----------------------------
-- Table structure for bagique_track_companies
-- ----------------------------
DROP TABLE IF EXISTS `bagique_track_companies`;
CREATE TABLE `bagique_track_companies`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `company_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公司名称',
  `company_short_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公司简称',
  `company_logo` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '公司logo',
  `company_url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '查询链接',
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_bagique_track_companies_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bagique_track_companies
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
