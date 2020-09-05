create table "user"
(
    id         bigserial            not null
        constraint user_pk
            primary key,
    code       uuid                 not null,
    email      text                 not null,
    password   text                 not null,
    user_type  smallint             not null,
    active     boolean default true not null,
    deleted_at timestamp
);

create unique index user_email_uindex
    on "user" (email);

create unique index user_id_uindex
    on "user" (id);