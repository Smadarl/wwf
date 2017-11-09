CREATE TABLE Game ( GameID integer PRIMARY KEY AUTOINCREMENT NOT NULL,
MaxLength integer NOT NULL,
MaxRecurrance integer NOT NULL,
Started integer NOT NULL,
StartedBy integer NOT NULL,
Turn integer NOT NULL,
Status varchar(12) NOT NULL );

CREATE TABLE Game_Player ( GameID integer NOT NULL,
PlayerID integer NOT NULL,
Word varchar(40) NOT NULL,
LastTime integer DEFAULT 0 );

CREATE TABLE Player (PlayerID integer PRIMARY KEY AUTOINCREMENT NOT NULL,
Username varchar(120) NOT NULL,
Password varchar(120) NOT NULL,
Created integer NOT NULL,
LastConnect integer DEFAULT 0 );

CREATE TABLE Friend (PlayerID integer NOT NULL, 
FriendID integer NOT NULL,
PRIMARY KEY (PlayerID, FriendID));

CREATE TABLE PlayerMove (GameID integer NOT NULL,
PlayerID integer NOT NULL,
Guess varchar(120) NOT NULL,
AtTime integer NOT NULL,
Result integer NOT NULL );