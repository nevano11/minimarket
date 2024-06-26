create table if not exists "user"
(
    id              serial    not null primary key,
    login           text      not null unique CHECK (char_length(login) <= 20 AND char_length(login) > 3),
    password_hash   text      not null,
    token           text      null,
    expiration_date timestamp null
);

create table if not exists post
(
    id          serial    not null primary key,
    user_id     integer   not null,
    name        text      not null CHECK (char_length(name) <= 20 AND char_length(name) > 3),
    description text      not null,
    img         text      null,
    price       integer   not null,
    created_at  timestamp not null,

    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES "user" (id)
);