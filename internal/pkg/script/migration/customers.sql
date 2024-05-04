create table customers
(
    id              uuid primary key not null ,
    customer_name   varchar not null,
    balanse         decimal,
    created_at      timestamp default now(),
    deleted_at      timestamp,
    updated_at      timestamp,
    updated_by      uuid references users(id),
    created_by      uuid references users(id),
    deleted_by      uuid references users(id)
);

