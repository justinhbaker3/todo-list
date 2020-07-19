CREATE DATABASE main;
USE main;
CREATE TABLE lists (
    id int not null auto_increment, title varchar(20), primary key (id)
);
CREATE TABLE items (
    title varchar(20), description varchar(50), list int, opened TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    foreign key (list) references lists(id) on delete cascade,
    primary key (title, list)
);
