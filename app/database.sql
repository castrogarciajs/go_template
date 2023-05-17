CREATE DATABASE mygo;

SHOW DATABASES;


CREATE TABLE gomi(
    id INT AUTO_INCREMENT,
    name VARCHAR(255),
    email VARCHAR(255),
    PRIMARY KEY (id)
);

SHOW TABLES;

INSERT INTO gomi (name, email) 
VALUES ('Jhon Doe', 'jhondoe@example.com');