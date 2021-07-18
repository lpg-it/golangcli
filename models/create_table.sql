--
DROP TABLE IF EXISTS `attack_plan`;
CREATE TABLE `attack_plan` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `attack_plan_id` bigint(20) unique NOT NULL,
                        `attack_plan_name` varchar(128) unique COLLATE utf8mb4_general_ci NOT NULL,
                        `attack_plan_type` int(10) unsigned NOT NULL DEFAULT '0',
                        `attack_plan_content` longtext COLLATE utf8mb4_general_ci NOT NULL,
                        `attack_plan_use_count` int(20) NULL DEFAULT '0',
                        `attack_plan_use_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `attack_plan_finish_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `attack_plan_status` int(4) NULL DEFAULT '0',
                        `attack_plan_remark` longtext COLLATE utf8mb4_general_ci NULL,

                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;