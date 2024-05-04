create table users
(
    id         uuid primary key not null ,
    first_name varchar,
    last_name  varchar,
    username   varchar (150),
    password   varchar,
    status     user_statuses,
    gmail      varchar,
    created_at timestamp default now(),
    deleted_at timestamp,
    updated_at timestamp,
    updated_by uuid references users(id),
    created_by uuid references users(id),
    deleted_by uuid references users(id)
);

CREATE TYPE user_statuses AS ENUM ('at_work', 'off_work');

