create type gender_user as enum ('seller', 'customer');
create type shipping_status as enum ('packing', 'ongoing', 'transite' ,'accepted');

CREATE TABLE categories (
  id varchar(255) primary key,
  name varchar(200) not null
);

CREATE TABLE products (
    id varchar(255) primary key,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price numeric(15) NOT NULL,
    color VARCHAR(100) NOT NULL,
    size VARCHAR(100) NOT NULL,
    stock INT NOT NULL,
    id_category varchar(255) REFERENCES categories ON DELETE CASCADE,
    foreign key (id_category) REFERENCES categories(id),
    id_user varchar(255) REFERENCES users ON DELETE CASCADE,
    foreign key (id_user) REFERENCES users(id),
    condition status_condition not null,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE images_product (
  id varchar(255) primary key,
  name VARCHAR(255) NOT NULL,
  url VARCHAR(255) NOT NULL,
  product_id varchar(255) REFERENCES products ON DELETE CASCADE,
  foreign key (product_id) REFERENCES products(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create table users (
  id varchar(255) primary key,
  name varchar(255) NOT NULL,
  email varchar(100) NOT NULL unique,
  password text NOT NULL,
  birth_date date,
  store_name varchar(255),
  description text,
  role varchar(100) NOT NULL DEFAULT('seller'),
  phone varchar(255),
  address text,
  photo varchar(255) DEFAULT('photodefault.jpg')
);

/* Table User address */

create table shipping_addresses (
  id varchar(255) primary key,
  id_user varchar(255) REFERENCES users ON DELETE CASCADE,
  foreign key(id_user) REFERENCES users(id),
  description TEXT,
  is_active boolean,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create table carts (
  id varchar(255) primary key,
  id_user varchar(255) REFERENCES users ON DELETE CASCADE,
  foreign key(id_user) REFERENCES users(id),
  id_product varchar(255) REFERENCES products ON DELETE CASCADE,
  foreign key(id_product) REFERENCES products(id), 
  description TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
);

create table shipping (
  id varchar(255) primary key,
  status shipping_status,
  description TEXT,
  id_user varchar(255) REFERENCES users ON DELETE CASCADE,
  foreign key(id_user) REFERENCES users(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create table user_orders (
  id varchar(255) primary key,
  id_carts varchar(255) REFERENCES carts ON DELETE CASCADE,
  foreign key(id_carts) REFERENCES carts(id),
  id_shipping varchar(255) REFERENCES shipping ON DELETE CASCADE,
  foreign key(id_shipping) REFERENCES shipping(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)

