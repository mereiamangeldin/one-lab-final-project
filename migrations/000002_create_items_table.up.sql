create table items(
   id bigserial primary key,
   name varchar(255) not null,
   description varchar(255) not null,
   cost float not null
)