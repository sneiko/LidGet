create table users
(
    id            serial primary key,
    first_name    text                    not null,
    last_name     text                    not null,
    country_code  text                    not null,
    password_hash text                    not null,
    is_active     boolean   default true  not null,
    created_at    timestamp default now() not null,
    updated_on    timestamp,
    deleted_on    timestamp
);

create table requests
(
    id           serial primary key,
    contact_link text                    not null,
    text         text                    not null,
    created_at   timestamp default now() not null,
    updated_at   timestamp,
    deleted_at   timestamp
);

create table category
(
    id   serial primary key,
    name text not null
);

create table category_patterns
(
    id          serial primary key,
    pattern     text not null,
    category_id int  not null,
    constraint fk_category_pattern_category_id foreign key (category_id) references category
);

create table request_category
(
    id          serial primary key,
    request_id  int not null,
    category_id int not null,
    constraint fk_request_category_request_id foreign key (request_id) references requests,
    constraint fk_request_category_category_id foreign key (category_id) references category
);

