create table users(
    id serial not null unique,
    name varchar not null,
    surname varchar,
    username varchar,
    password varchar not null,
    balance float,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

create table categories(
    id serial not null unique,
    name varchar
);
create table products(
    id serial not null unique,
    name varchar not null,
    description varchar,
    price float,
    img varchar,
    liked boolean,
    category_id integer not null references categories(id) ON DELETE CASCADE
);

create table favorites(
    id serial not null unique,
    liked boolean,
    user_id integer not null references users(id) ON DELETE CASCADE,
    product_id integer not null references products(id) ON DELETE CASCADE
);


create table transactions(
    id serial not null unique,
    user_id integer not null references users(id) ON DELETE CASCADE,
    product_id integer not null references products(id) ON DELETE CASCADE,
    amount float,
    taken_at timestamp
);

insert into categories(name) values ('спортивные костюмы');
insert into categories(name) values ('рашгарды');
insert into categories(name) values ('футболки');
insert into categories(name) values ('кроссовки');

insert into products(name, description, price, img, liked, category_id) values ('Костюм спортивный', 'EA7', 77850, 'https://a.lmcdn.ru/img600x866/E/A/EA002EMFXFD7_9676281_2_v2_2x.jpg', false, 1);

insert into products(name, description, price, img, liked, category_id) values ('hardcore', 'EA7', 77850, 'https://a.lmcdn.ru/img600x866/R/T/RTLACF321801_18727486_2_v1_2x.jpg', false, 2);

insert into products(name, description, price, img, liked, category_id) values ('uniqlo', 'EA7', 77850, 'https://a.lmcdn.ru/img600x866/R/T/RTLACH216301_18851409_2_v1.jpg', false, 3);

insert into products(name, description, price, img, liked, category_id) values ('patrol', 'EA7', 77850, 'https://a.lmcdn.ru/img600x866/R/T/RTLACL778801_19373957_4_v1.jpg', false, 4);







