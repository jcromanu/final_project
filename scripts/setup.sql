CREATE TABLE IF NOT EXISTS USER (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    pwd_hash VARCHAR(10) NOT NULL,
    username VARCHAR(30) NOT NULL , 
    age INT NOT NULL , 
    additional_information VARCHAR(40) NOT NULL ,
    parent VARCHAR(40) 
);
