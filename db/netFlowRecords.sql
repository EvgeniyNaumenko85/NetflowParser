-- MySQL Script generated by MySQL Workbench
-- Sat Jul 22 15:27:20 2023
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema netflowdb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema netflowdb
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `netflowdb` DEFAULT CHARACTER SET utf8 ;
USE `netflowdb` ;

-- -----------------------------------------------------
-- Table `netflowdb`.`record`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `netflowdb`.`record` (
  `id` BIGINT(19) NOT NULL AUTO_INCREMENT,
  `source` VARCHAR(15) NOT NULL,
  `destination` VARCHAR(15) NOT NULL,
  `account_id` BIGINT(19) NOT NULL,
  `tclass` BIGINT(19) NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
