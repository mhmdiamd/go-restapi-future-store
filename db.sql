CREATE TABLE category (
  id varchar(255) primary key,
  name varchar(200) not null
);

CREATE TABLE product (
  id varchar(255) primary key,
  product_name varchar(200) not null,
  stock int NOT NULL,
  price int NOT NULL,
  description TEXT NOT NULL,
  id_category varchar(200) NOT NULL ,
  photo varchar(200) NOT NULL ,
  FOREIGN KEY (id_category)
      REFERENCES category (id)
);