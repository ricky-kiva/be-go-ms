-- DDL create
create table todos 
(
id int,
title varchar(30) not null,
completion boolean,
description text,
created_at date,
updated_at date
);

-- DDL alter table
alter table todos add primary key (id);

-- DDL drop column
alter table todos drop column description;

-- DDL alter constraints
alter table todos drop constraint todos_pkey;

-- DDL drop table
drop table todos

-- DML insert
insert into todos values 
(
2, 
'The Witcher', 
false, 
'Geralt of Rivia, a solitary monster hunter, struggles to find his place in a world where people often prove more wicked than beasts.',
now(),
now()
);

-- DML select
select * from todos;

-- DML update
update todos 
set title='The Witcher 2: Assassins of..'
where id = 2;

-- DML delete
delete from todos
where id = 2;