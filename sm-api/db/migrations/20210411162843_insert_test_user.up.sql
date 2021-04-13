SET NAMES utf8 ;

LOCK TABLES `users` WRITE;

INSERT INTO `users` VALUES (1,CURDATE(),CURDATE(),NULL,'test123@fake.com','test123@fake.com','$2a$14$tfakwOfXgRJEq5lCpy6oUu6exCwHapIWusDES4LfhMr9pQiC3eF1u');

UNLOCK TABLES;

