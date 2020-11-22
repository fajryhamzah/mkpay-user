create table "user" (
    "id" bigserial not null constraint user_pk primary key,
    "code" uuid not null,
    "email" text not null,
    "password" text not null,
    "phone_number" VARCHAR(13) not null,
    "phone_verified_at" timestamp,
    "user_type" text not null,
    "active" boolean default true not null,
    "deleted_at" timestamp
);
create unique index user_email_uindex on "user" (email);
create unique index user_id_uindex on "user" (id);
create unique index phone_number_uindex on "user" (phone_number);
create unique index code_uindex on "user" (code);