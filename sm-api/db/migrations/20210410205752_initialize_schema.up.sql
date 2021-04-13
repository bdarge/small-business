CREATE DATABASE  IF NOT EXISTS `profile` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `profile`;
-- MySQL dump 10.13  Distrib 8.0.13, for macos10.14 (x86_64)
--
-- Host: localhost    Database: profile
-- ------------------------------------------------------
-- Server version	8.0.23

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `customers`
--

DROP TABLE IF EXISTS `customers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET character_set_client = utf8mb4 ;
CREATE TABLE `customers` (
                             `id` bigint NOT NULL AUTO_INCREMENT,
                             `created_at` datetime(3) DEFAULT NULL,
                             `updated_at` datetime(3) DEFAULT NULL,
                             `deleted_at` datetime(3) DEFAULT NULL,
                             `email` longtext,
                             `name` longtext,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `order_items`
--

DROP TABLE IF EXISTS `order_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET character_set_client = utf8mb4 ;
CREATE TABLE `order_items` (
                               `id` bigint NOT NULL AUTO_INCREMENT,
                               `created_at` datetime(3) DEFAULT NULL,
                               `updated_at` datetime(3) DEFAULT NULL,
                               `deleted_at` datetime(3) DEFAULT NULL,
                               `description` longtext,
                               `qty` longtext,
                               `unit` longtext,
                               `unit_price` longtext,
                               `order_id` bigint DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               KEY `fk_orders_items` (`order_id`),
                               CONSTRAINT `fk_orders_items` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET character_set_client = utf8mb4 ;
CREATE TABLE `orders` (
                          `id` bigint NOT NULL AUTO_INCREMENT,
                          `created_at` datetime(3) DEFAULT NULL,
                          `updated_at` datetime(3) DEFAULT NULL,
                          `deleted_at` datetime(3) DEFAULT NULL,
                          `description` longtext,
                          `delivery_date` longtext,
                          `invoice_number` longtext,
                          `currency` longtext,
                          `order_number` longtext,
                          `created_by` bigint DEFAULT NULL,
                          `customer_id` bigint DEFAULT NULL,
                          PRIMARY KEY (`id`),
                          KEY `fk_users_orders` (`created_by`),
                          KEY `fk_customers_orders` (`customer_id`),
                          CONSTRAINT `fk_customers_orders` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`),
                          CONSTRAINT `fk_users_orders` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `quote_items`
--

DROP TABLE IF EXISTS `quote_items`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET character_set_client = utf8mb4 ;
CREATE TABLE `quote_items` (
                               `id` bigint NOT NULL AUTO_INCREMENT,
                               `created_at` datetime(3) DEFAULT NULL,
                               `updated_at` datetime(3) DEFAULT NULL,
                               `deleted_at` datetime(3) DEFAULT NULL,
                               `description` longtext,
                               `qty` longtext,
                               `unit` longtext,
                               `unit_price` longtext,
                               `quote_id` bigint DEFAULT NULL,
                               PRIMARY KEY (`id`),
                               KEY `fk_quotes_items` (`quote_id`),
                               CONSTRAINT `fk_quotes_items` FOREIGN KEY (`quote_id`) REFERENCES `quotes` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `quotes`
--

DROP TABLE IF EXISTS `quotes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET character_set_client = utf8mb4 ;
CREATE TABLE `quotes` (
                          `id` bigint NOT NULL AUTO_INCREMENT,
                          `created_at` datetime(3) DEFAULT NULL,
                          `updated_at` datetime(3) DEFAULT NULL,
                          `deleted_at` datetime(3) DEFAULT NULL,
                          `description` longtext,
                          `quote_number` longtext,
                          `created_by` bigint DEFAULT NULL,
                          `customer_id` bigint DEFAULT NULL,
                          PRIMARY KEY (`id`),
                          KEY `fk_users_quotes` (`created_by`),
                          KEY `fk_customers_quotes` (`customer_id`),
                          CONSTRAINT `fk_customers_quotes` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`),
                          CONSTRAINT `fk_users_quotes` FOREIGN KEY (`created_by`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
SET character_set_client = utf8mb4 ;
CREATE TABLE `users` (
                         `id` bigint NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) DEFAULT NULL,
                         `updated_at` datetime(3) DEFAULT NULL,
                         `deleted_at` datetime(3) DEFAULT NULL,
                         `email` longtext,
                         `userName` longtext,
                         `password` longtext,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-04-10 11:08:41
