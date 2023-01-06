SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
    `order_id` varchar(255) NOT NULL,
    `uuid` varchar(255) DEFAULT NULL,
    `from` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `to` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `contract_address` varchar(255) NOT NULL,
    `tx_hash` varchar(255) DEFAULT NULL,
    `status` tinyint DEFAULT '0',
    `notify_status` tinyint DEFAULT '0',
    `create_time` bigint DEFAULT NULL,
    `pay_time` bigint DEFAULT NULL,
    `amount` varchar(255) DEFAULT NULL,
    `token_id` bigint DEFAULT NULL,
    `cb` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    PRIMARY KEY (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

SET FOREIGN_KEY_CHECKS = 1;


SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;


-- ----------------------------
-- Table structure for nft_owner
-- ----------------------------
DROP TABLE IF EXISTS `nft_owner`;
CREATE TABLE `nft_owner` (
    `id` varchar(255) NOT NULL,
    `owner_address` varchar(255) CHARACTER SET utf8mb4 NOT NULL,
    `contract_address` varchar(255) NOT NULL,
    `token_id` bigint NOT NULL,
    `update_time` bigint NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api_key
-- ----------------------------
DROP TABLE IF EXISTS `api_key`;
CREATE TABLE `api_key` (
    `id` varchar(255) NOT NULL,
    `api_key` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

SET FOREIGN_KEY_CHECKS = 1;

