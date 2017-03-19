CREATE DATABASE  IF NOT EXISTS `iman_db` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `iman_db`;
-- MySQL dump 10.13  Distrib 5.7.12, for linux-glibc2.5 (x86_64)
--
-- Host: localhost    Database: iman_db
-- ------------------------------------------------------
-- Server version	5.7.17-0ubuntu0.16.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `articles`
--

DROP TABLE IF EXISTS `articles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `articles` (
  `id_articles` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(250) CHARACTER SET utf8 NOT NULL,
  `content` longtext CHARACTER SET utf8 NOT NULL,
  `thumbnail` varchar(250) CHARACTER SET utf8 NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id_articles`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `articles`
--

LOCK TABLES `articles` WRITE;
/*!40000 ALTER TABLE `articles` DISABLE KEYS */;
INSERT INTO `articles` VALUES (1,'Kala Senja Sejuk','Ketika malam tiba ,saya menangis','thumb.png',NULL,NULL),(21,'\"Makan Gambaran\"','Oke seepe','','2017-03-16 16:38:17','2017-03-16 16:38:17'),(24,'Makan Kue Yang Enak','Hello Dunia Indah','oke.png','2017-03-17 02:19:07','2017-03-17 02:19:07'),(25,'Aduh Sayang Sekali','LappetToba','lvchat1.png','2017-03-17 02:20:11','2017-03-17 02:20:11');
/*!40000 ALTER TABLE `articles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'iman_db'
--

--
-- Dumping routines for database 'iman_db'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-03-17 10:55:50
