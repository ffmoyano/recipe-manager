create table user
(
    id         int primary key auto_increment,
    name       varchar(20),
    password   varchar(40),
    email      varchar(60),
    created_at datetime default NOW(),
    verified   tinyint(1) default false
);

create table role
(
    id   int primary key auto_increment,
    name varchar(10)
);

create table user_role
(
    user_id int,
    role_id int,
    foreign key (user_id) references user (id),
    foreign key (role_id) references role (id)
);

insert into role (name) values ('ADMIN');
insert into role (name) values ('USER');