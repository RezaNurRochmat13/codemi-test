-- MySQL dump 10.13  Distrib 5.7.17, for macos10.12 (x86_64)
--
-- Host: 127.0.0.1    Database: codemi
-- ------------------------------------------------------
-- Server version	5.7.21

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
-- Table structure for table `classrooms`
--

DROP TABLE IF EXISTS `classrooms`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `classrooms` (
  `id_classroom` varchar(50) NOT NULL,
  `uuid_classroom` varchar(50) DEFAULT NULL,
  `uuid_participants` varchar(50) DEFAULT NULL,
  `classroom_name` varchar(20) DEFAULT NULL,
  `classroom_time` datetime DEFAULT NULL,
  `room` varchar(10) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id_classroom`),
  KEY `fk_participants_idx` (`uuid_participants`),
  CONSTRAINT `fk_participants` FOREIGN KEY (`uuid_participants`) REFERENCES `participants` (`uuid_participants`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `classrooms`
--

LOCK TABLES `classrooms` WRITE;
/*!40000 ALTER TABLE `classrooms` DISABLE KEYS */;
INSERT INTO `classrooms` VALUES ('1aeab32c-e783-4a5b-ab92-7f1d7fde534d','246ff729-1654-4ede-94e3-cc5728fa0ac6','13d341af-8496-4ebc-9436-1752e994787b','Belajar ORM','2018-10-10 00:00:00','B1','2018-10-27 13:46:12','2018-10-27 13:46:12',NULL),('c94da4e8-21cf-4f77-b40d-cb80c468c307','85daa96b-4c05-4e68-9e29-bab77ce0d5d2','b62ecff0-14b1-42cb-81a0-9aa08e9ebf13','Belajar Gorm','2018-09-09 00:00:00','R1',NULL,NULL,NULL);
/*!40000 ALTER TABLE `classrooms` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `participants`
--

DROP TABLE IF EXISTS `participants`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `participants` (
  `uuid_participants` varchar(50) NOT NULL,
  `participants_name` varchar(45) DEFAULT NULL,
  `participants_address` varchar(45) DEFAULT NULL,
  `participants_birthday` datetime DEFAULT NULL,
  `participants_age` int(10) DEFAULT NULL,
  `participants_gender` char(5) DEFAULT NULL,
  `participants_phone` varchar(20) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`uuid_participants`),
  KEY `uuid_participants` (`uuid_participants`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `participants`
--

LOCK TABLES `participants` WRITE;
/*!40000 ALTER TABLE `participants` DISABLE KEYS */;
INSERT INTO `participants` VALUES ('13d341af-8496-4ebc-9436-1752e994787b','Reja','Pakem','1992-12-13 00:00:00',22,'Male','85567756',NULL,NULL,NULL),('20c0dce8-1ce6-4257-ab2d-e8705ea7c1fb','Andreas A','Beran Tridadi','1992-08-09 00:00:00',25,'Male','086787668','2018-10-27 14:37:13','2018-10-27 14:37:13',NULL),('b62ecff0-14b1-42cb-81a0-9aa08e9ebf13','Gema','Jombang','1992-12-13 00:00:00',22,'Male','85567756',NULL,NULL,NULL);
/*!40000 ALTER TABLE `participants` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-10-27 14:51:05
