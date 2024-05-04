create table items
(
    id              uuid primary key not null ,
    item_name       varchar,
    cost            decimal,
    price           decimal,
    sort            integer,
    created_at      timestamp default now(),
    deleted_at      timestamp,
    updated_at      timestamp,
    updated_by uuid references users(id),
    created_by uuid references users(id),
    deleted_by uuid references users(id)
);

