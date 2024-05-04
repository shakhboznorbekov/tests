create table transactions
(
    id          uuid primary key not null ,
    qty         integer,
    price       decimal,
    amount      decimal,
    customer_id uuid references customers(id),
    item_id     uuid references items(id),
    created_at  timestamp default now(),
    deleted_at  timestamp,
    updated_at  timestamp,
    updated_by  uuid references users(id),
    created_by  uuid references users(id),
    deleted_by  uuid references users(id)
);

