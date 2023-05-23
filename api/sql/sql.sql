CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(30) not null,
    nick varchar(30) not null unique,
    email varchar(30) not null unique,
    passwrd varchar(30) not null unique,
    createdDate timestamp default current_timestamp()
) ENGINE=INNODB;