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