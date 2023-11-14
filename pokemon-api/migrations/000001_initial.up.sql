create table if not exists pokemon (
    id integer primary key,
    name varchar not null,
    types varchar[] not null,
    created_at timestamptz not null,
    updated_at timestamptz not null
);
