create table if not exists storage
(
    id           serial not null primary key,
    name         text   not null,
    code         text   not null unique,
    is_available bool   not null default true
);

create table if not exists product
(
    id   serial  not null primary key,
    name text    not null,
    size integer not null check (size > 0),
    code text    not null unique
);

create table if not exists storage_product
(
    id              serial  not null primary key,
    product_id      integer not null,
    storage_id      integer not null,
    count           integer not null check (count >= 0),

    CONSTRAINT fk_storage
        FOREIGN KEY (storage_id)
            REFERENCES storage (id),
    CONSTRAINT fk_product
        FOREIGN KEY (product_id)
            REFERENCES product (id)
);

create table if not exists reservation
(
    id                 serial  not null primary key,
    request            text    null,
    storage_product_id integer not null,
    reserve_size       integer not null check (reserve_size > 0),
    is_active          bool    not null default true,

    CONSTRAINT fk_storage_product
        FOREIGN KEY (storage_product_id)
            REFERENCES storage_product (id)
);