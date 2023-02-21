-- +migrate Up
create extension if not exists pg_trgm;

create table account
(
    id           uuid primary key     default gen_random_uuid(),
    email        text        not null default '',
    display_name text        not null default '',
    custom_url   text                 default null unique,
    is_superuser bool        not null default false,
    created_at   timestamptz not null default now(),
    updated_at   timestamptz not null default now()
);

create index account_email_trgm_gin on account using gin (email gin_trgm_ops);
create index account_display_name_trgm_gin on account using gin (display_name gin_trgm_ops);
create unique index on account (email);
create index on account (custom_url);

create table auth_provider
(
    id                  uuid primary key     default gen_random_uuid(),
    account_id          uuid        not null references account (id) on delete cascade,
    provider            text        not null default '',
    email               text        not null default '',
    name                text        not null default '',
    first_name          text        not null default '',
    last_name           text        not null default '',
    nick_name           text        not null default '',
    description         text        not null default '',
    provider_user_id    text        not null default '',
    avatar_url          text        not null default '',
    location            text        not null default '',
    access_token        text        not null default '',
    access_token_secret text        not null default '',
    refresh_token       text        not null default '',
    expires_at          timestamptz not null default now(),
    id_token            text        not null default ''
);

create index on auth_provider (provider);
create unique index on auth_provider (account_id);
create unique index on auth_provider (provider, provider_user_id);

create table invitation
(
    id                    uuid primary key                  default gen_random_uuid(),
    email                 text                     not null,
    active                bool                     not null default false,
    created_by_account_id uuid                     null references account (id) on delete set null,
    expires_at            timestamp with time zone not null,
    created_at            timestamp with time zone not null default now(),
    updated_at            timestamp with time zone not null default now()
);

create index invitation_trgm_gin on invitation using gin (email gin_trgm_ops);
create index on invitation (email);

create table session
(
    id         uuid primary key     default gen_random_uuid(),
    account_id uuid        not null references account (id) on delete cascade,
    ip_address text        not null,
    expires_at timestamptz not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

create index on session (account_id);

create table file_group
(
    id               uuid primary key     default gen_random_uuid(),
    account_id       uuid        not null references account (id) on delete cascade,
    name             text        not null,
    slug             text        not null,
    show_public_list bool        not null default false,
    created_at       timestamptz not null default now(),
    updated_at       timestamptz not null default now()
);

create unique index on file_group (account_id, slug);

create table file
(
    id                 uuid primary key     default gen_random_uuid(),
    account_id         uuid        not null references account (id) on delete cascade,
    name               text        not null,
    slug               text        not null,
    short_id           text        not null,
    download_path      text        not null default '',
    size               bigint      not null,
    mime_type          text        not null,
    password_protected bool        not null default false,
    file_group_id      uuid        null references file_group (id) on delete set null,
    created_at         timestamptz not null default now(),
    updated_at         timestamptz not null default now()
);

create index file_trgm_gin on file using gin (name gin_trgm_ops);
create index on file (account_id);
create index on file (account_id, slug);
create index on file (slug);
create index on file (short_id);
create unique index on file (account_id, file_group_id, slug);

create table file_password
(
    id            uuid primary key     default gen_random_uuid(),
    file_id       uuid        not null references file (id) on delete cascade,
    password_hash text        not null,
    created_at    timestamptz not null default now(),
    updated_at    timestamptz not null default now()
);

create index on file_password (file_id);

-- +migrate Down
drop table file_password;
drop table file;
drop table file_group;
drop table invitation;
drop table session;
drop table auth_provider;
drop table account;