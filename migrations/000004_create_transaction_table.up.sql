create table transactions(
    id bigserial primary key,
    user_id integer not null,
    item_id integer not null,
    amount float not null,
    created_at timestamp
);
