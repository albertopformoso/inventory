CREATE DATABASE IF NOT EXISTS inventory_db;
USE inventory_db;

CREATE TABLE IF NOT EXISTS user (
    id int not null auto_increment,
    email varchar(255) not null,
    name varchar(255) not null,
    password varchar(255) not null,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS product (
    id int not null auto_increment,
    name varchar(255) not null,
    description varchar(255) not null,
    price float not null,
    created_by int not null,
    primary key (id),
    foreign key (created_by) references user(id)
);

CREATE TABLE IF NOT EXISTS role (
    id int not null auto_increment,
    name varchar(255) not null,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS user_role (
    id int not null auto_increment,
    user_id int not null,
    role_id int not null,
    primary key (id),
    foreign key (user_id) references user(id),
    foreign key (role_id) references role(id)
);

INSERT INTO role (id, name) VALUES (1, 'admin');
INSERT INTO role (id, name) VALUES (2, 'seller');
INSERT INTO role (id, name) VALUES (3, 'customer');
