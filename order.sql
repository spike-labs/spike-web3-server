/*
 Navicat Premium Data Transfer

 Source Server         : mysql_1
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3309
 Source Schema         : spike_frame_server

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 15/09/2022 17:05:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `order_id` varchar(255) NOT NULL,
  `uuid` varchar(255) DEFAULT NULL,
  `from` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `to` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `contract_address` varchar(255) NOT NULL,
  `tx_hash` varchar(255) DEFAULT NULL,
  `status` tinyint DEFAULT '0',
  `notify_status` tinyint DEFAULT '0',
  `create_time` bigint DEFAULT NULL,
  `pay_time` bigint DEFAULT NULL,
  `amount` varchar(255) DEFAULT NULL,
  `token_id` bigint DEFAULT NULL,
  `cb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
