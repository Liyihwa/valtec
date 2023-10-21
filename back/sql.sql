-- MySQL dump 10.13  Distrib 5.7.42, for Linux (x86_64)
--
-- Host: localhost    Database: valtec
-- ------------------------------------------------------
-- Server version	5.7.42-0ubuntu0.18.04.1

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
-- Table structure for table `HERO`
--
DROP DATABASE IF EXISTS `valtec`;
CREATE DATABASE `valtec`;
USE `valtec`;

DROP TABLE IF EXISTS `HERO`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `HERO` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(31) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `c` varchar(255) DEFAULT NULL,
  `q` varchar(255) DEFAULT NULL,
  `e` varchar(255) DEFAULT NULL,
  `x` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `HERO`
--

LOCK TABLES `HERO` WRITE;
/*!40000 ALTER TABLE `HERO` DISABLE KEYS */;
INSERT INTO `HERO` VALUES (2,'sova','https://valtec.oss-cn-chengdu.aliyuncs.com/sova.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/sova_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/sova_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/sova_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/sova_x.webp'),(3,'astra','https://valtec.oss-cn-chengdu.aliyuncs.com/astra.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/astra_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/astra_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/astra_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/astra_x.webp'),(4,'breach','https://valtec.oss-cn-chengdu.aliyuncs.com/breach.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/breach_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/breach_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/breach_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/breach_x.webp'),(5,'brimstone','https://valtec.oss-cn-chengdu.aliyuncs.com/brimstone.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/brimstone_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/brimstone_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/brimstone_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/brimstone_x.webp'),(6,'chamber','https://valtec.oss-cn-chengdu.aliyuncs.com/chamber.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/chamber_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/chamber_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/chamber_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/chamber_x.webp'),(7,'cypher','https://valtec.oss-cn-chengdu.aliyuncs.com/cypher.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/cypher_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/cypher_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/cypher_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/cypher_x.webp'),(8,'deadlock','https://valtec.oss-cn-chengdu.aliyuncs.com/deadlock.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/deadlock_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/deadlock_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/deadlock_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/deadlock_x.webp'),(9,'fade','https://valtec.oss-cn-chengdu.aliyuncs.com/fade.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/fade_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/fade_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/fade_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/fade_x.webp'),(10,'gekko','https://valtec.oss-cn-chengdu.aliyuncs.com/gekko.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/gekko_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/gekko_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/gekko_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/gekko_x.webp'),(11,'harbor','https://valtec.oss-cn-chengdu.aliyuncs.com/harbor.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/harbor_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/harbor_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/harbor_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/harbor_x.webp'),(12,'jett','https://valtec.oss-cn-chengdu.aliyuncs.com/jett.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/jett_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/jett_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/jett_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/jett_x.webp'),(13,'kayo','https://valtec.oss-cn-chengdu.aliyuncs.com/kayo.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/kayo_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/kayo_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/kayo_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/kayo_x.webp'),(14,'killjoy','https://valtec.oss-cn-chengdu.aliyuncs.com/killjoy.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/killjoy_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/killjoy_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/killjoy_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/killjoy_x.webp'),(15,'neon','https://valtec.oss-cn-chengdu.aliyuncs.com/neon.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/neon_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/neon_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/neon_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/neon_x.webp'),(16,'omen','https://valtec.oss-cn-chengdu.aliyuncs.com/omen.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/omen_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/omen_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/omen_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/omen_x.webp'),(17,'phoenix','https://valtec.oss-cn-chengdu.aliyuncs.com/phoenix.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/phoenix_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/phoenix_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/phoenix_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/phoenix_x.webp'),(18,'raze','https://valtec.oss-cn-chengdu.aliyuncs.com/raze.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/raze_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/raze_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/raze_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/raze_x.webp'),(19,'reyna','https://valtec.oss-cn-chengdu.aliyuncs.com/reyna.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/reyna_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/reyna_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/reyna_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/reyna_x.webp'),(20,'sage','https://valtec.oss-cn-chengdu.aliyuncs.com/sage.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/sage_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/sage_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/sage_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/sage_x.webp'),(21,'skye','https://valtec.oss-cn-chengdu.aliyuncs.com/skye.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/skye_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/skye_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/skye_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/skye_x.webp'),(22,'viper','https://valtec.oss-cn-chengdu.aliyuncs.com/viper.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/viper_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/viper_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/viper_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/viper_x.webp'),(23,'yoru','https://valtec.oss-cn-chengdu.aliyuncs.com/yoru.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/yoru_c.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/yoru_q.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/yoru_e.webp','https://valtec.oss-cn-chengdu.aliyuncs.com/yoru_x.webp');
/*!40000 ALTER TABLE `HERO` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `MAP`
--

DROP TABLE IF EXISTS `MAP`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `MAP` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(31) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `MAP`
--

LOCK TABLES `MAP` WRITE;
/*!40000 ALTER TABLE `MAP` DISABLE KEYS */;
INSERT INTO `MAP` VALUES (1,'莲华古城','https://valtec.oss-cn-chengdu.aliyuncs.com/莲华古城平面.jpg','https://valtec.oss-cn-chengdu.aliyuncs.com/莲华古城.jpg'),(2,'霓虹町','https://valtec.oss-cn-chengdu.aliyuncs.com/霓虹町平面.jpg','https://valtec.oss-cn-chengdu.aliyuncs.com/霓虹町.jpg'),(4,'微风岛屿','https://valtec.oss-cn-chengdu.aliyuncs.com/微风岛屿平面.jpg','https://valtec.oss-cn-chengdu.aliyuncs.com/微风岛屿.jpg'),(5,'亚海悬城','https://valtec.oss-cn-chengdu.aliyuncs.com/亚海悬城平面.jpg','https://valtec.oss-cn-chengdu.aliyuncs.com/亚海悬城.jpg'),(6,'源工重镇','https://valtec.oss-cn-chengdu.aliyuncs.com/源工重镇平面.jpg','https://valtec.oss-cn-chengdu.aliyuncs.com/源工重镇.jpg');
/*!40000 ALTER TABLE `MAP` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `POSITION`
--

DROP TABLE IF EXISTS `POSITION`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `POSITION` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `hero_id` int(11) DEFAULT NULL,
  `skill` varchar(255) DEFAULT NULL,
  `map_id` int(11) DEFAULT NULL,
  `stand_x` float DEFAULT NULL,
  `stand_y` float DEFAULT NULL,
  `put_x` float DEFAULT NULL,
  `put_y` float DEFAULT NULL,
  `like` int(11) DEFAULT NULL,
  `dislike` int(11) DEFAULT NULL,
  `description` text CHARACTER SET utf8,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `POSITION`
--

LOCK TABLES `POSITION` WRITE;
/*!40000 ALTER TABLE `POSITION` DISABLE KEYS */;
INSERT INTO `POSITION` VALUES (17,NULL,0,'',0,0,0,0,0,1,1,'qaaaa'),(18,NULL,0,'',0,0,0,0,0,1,1,'df'),(19,NULL,2,'',1,408.886,651.436,615.401,766.052,1,1,'nono'),(20,NULL,2,'',1,265.248,529.745,419.366,652.851,1,1,'nono'),(21,NULL,2,'c',1,310.494,715.111,574.898,780.202,1,1,'nono'),(22,NULL,2,'c',1,123.856,461.824,422.194,426.449,1,1,'nono'),(23,NULL,2,'c',1,153.548,577.855,289.285,729.261,1,1,'nono'),(24,NULL,2,'c',1,244.806,575.025,411.715,621.721,1,1,'nono'),(25,NULL,2,'c',1,396.156,696.716,280.168,541.065,1,1,'nono'),(26,NULL,2,'c',1,326.407,198.906,559.977,219.126,1,1,'nono'),(27,NULL,2,'c',1,577.21,791.522,440.005,753.317,1,1,'nono'),(28,NULL,2,'c',1,480.656,144.09,674.735,254.058,1,1,'213123ikldsfhjklasdfklasdlfk;asdj;flasdklfjlasd;jfasdfasdfasdfasdfasdfasdfasdfasd'),(29,NULL,2,'c',1,529.117,242.497,892.64,589.175,1,1,'打算打开了飞机纳斯达克了；放假你快乐；阿斯顿放假了；阿斯顿飞机了； 阿斯顿加分了； 阿斯顿加哦；if家i哦怕是的；加分i哦怕阿斯顿急哦怕；放假哦i阿斯顿；加分i哦怕叫阿水都佩服阿水淀粉阿水淀粉阿水淀粉阿水淀粉阿水淀粉阿水淀粉阿水淀粉阿斯顿'),(30,NULL,2,'c',1,520.631,428.139,437.176,172.166,1,1,'<p>发生的款礼服健康啦；是点击发送到家分快乐；阿斯顿加疯狂拉屎的家乐福就阿萨德了；加法快乐；阿斯顿加了；分就阿斯顿款礼服克拉；史蒂夫；阿斯顿加分了卡上的纠纷</p><p><img src=\"https://valtec.oss-cn-chengdu.aliyuncs.com/亚海悬城.jpg\" alt=\"\" data-href=\"\" style=\"\"/></p>'),(31,NULL,6,'c',1,268.694,319.138,476.793,314.063,1,1,'<p>weqwewq</p>'),(32,NULL,14,'c',1,341.08,709.655,539.107,531.465,1,1,'<p>水淀粉快乐风男裤拉屎的奶粉；卢卡斯的；分了卡斯蒂略； 分快乐；阿斯顿非法所得阿水淀粉</p><p><img src=\"https://valtec.oss-cn-chengdu.aliyuncs.com/亚海悬城平面.jpg\" alt=\"\" data-href=\"\" style=\"\"/></p>');
/*!40000 ALTER TABLE `POSITION` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `USER`
--

DROP TABLE IF EXISTS `USER`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `USER` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `password` binary(32) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `captcha` char(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `USER`
--

LOCK TABLES `USER` WRITE;
/*!40000 ALTER TABLE `USER` DISABLE KEYS */;
/*!40000 ALTER TABLE `USER` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-10-21 13:44:53
