create table balance(
    user_id integer primary key,
    amount float not null,
    foreign key (user_id) references users(id)
)