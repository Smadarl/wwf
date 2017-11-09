

CREATE TABLE `Game` (
  `GameID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `MaxLength` tinyint(4) NOT NULL,
  `MaxRecurrance` tinyint(4) NOT NULL,
  `Started` datetime NOT NULL,
  `StartedBy` int(10) unsigned NOT NULL,
  `Turn` int(10) unsigned NOT NULL,
  `Status` enum('Started','Playing','Cancelled','Finished') NOT NULL,
PRIMARY KEY `PRIMARY` (`GameID`)
) ENGINE=InnoDB;

# Reading .frm file for GameMoves.frm:
# CREATE VIEW Statement:

select `g`.`GameID` AS `GameID`,`p`.`Username` AS `Username`,`pm`.`Guess` AS `Guess`,`pm`.`AtTime` AS `AtTime`,`pm`.`Result` AS `Result` from ((`wwf`.`Game` `g` join `wwf`.`PlayerMove` `pm` on((`pm`.`GameID` = `g`.`GameID`))) join `wwf`.`Player` `p` on((`p`.`PlayerID` = `pm`.`PlayerID`)))

CREATE TABLE `Game_Player` (
  `GameID` int(11) NOT NULL,
  `PlayerID` int(11) NOT NULL,
  `Word` varchar(120) NOT NULL,
  `LastTime` datetime DEFAULT NULL,
PRIMARY KEY `PRIMARY` (`GameID`,`PlayerID`)
) ENGINE=InnoDB;

# Reading .frm file for MasterTable.frm:
# The .frm file is a TABLE.
# CREATE TABLE Statement:

CREATE TABLE `MasterTable` (
  `masterTableId` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `rowName` varchar(120) NOT NULL,
  `created` datetime NOT NULL,
  `anotherId` int(11) unsigned NOT NULL,
PRIMARY KEY `PRIMARY` (`masterTableId`),
UNIQUE KEY `rowName` (`rowName`),
KEY `anotherId` (`anotherId`)
) ENGINE=InnoDB;

# Reading .frm file for Player.frm:
# The .frm file is a TABLE.
# CREATE TABLE Statement:

CREATE TABLE `Player` (
  `PlayerID` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `Username` varchar(120) NOT NULL,
  `Created` datetime NOT NULL,
  `LastConnect` datetime DEFAULT NULL,
PRIMARY KEY `PRIMARY` (`PlayerID`)
) ENGINE=InnoDB;

# Reading .frm file for PlayerMove.frm:
# The .frm file is a TABLE.
# CREATE TABLE Statement:

CREATE TABLE `PlayerMove` (
  `GameID` int(10) unsigned NOT NULL,
  `PlayerID` int(10) unsigned NOT NULL,
  `Guess` varchar(120) NOT NULL,
  `AtTime` datetime NOT NULL,
  `Result` tinyint(4) NOT NULL,
PRIMARY KEY `PRIMARY` (`GameID`,`PlayerID`)
) ENGINE=InnoDB;

# Reading .frm file for SlaveTable.frm:
# The .frm file is a TABLE.
# CREATE TABLE Statement:

CREATE TABLE `SlaveTable` (
  `slaveTableId` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `slaveName` varchar(120) NOT NULL,
  `masterId` int(11) unsigned NOT NULL,
PRIMARY KEY `PRIMARY` (`slaveTableId`),
UNIQUE KEY `slaveName` (`slaveName`),
KEY `foreignId` (`masterId`)
) ENGINE=InnoDB;

