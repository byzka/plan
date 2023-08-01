
create table search
(
    type varchar(30) not null,
    name varchar(30) not null,
    local varchar(30) not null,
    pric varchar(30) not null
)

create table user
(
    id int auto_increment primary key,
    email varchar(30) not null,
    password varchar(40) not null,
    head  varchar(100) null,
    phone_number int
)

create table setting
(
    id int auto_increment primary key,
    local varchar(30) not null ,
    enviroment varchar(30) not null
)
