create table user_user (
    id serial ,
    nick_name text not null default '',
    pass_word text not null,
    gender int not null default 0, -- 0保密，1男，2女
    phone text,
    email text,
    head_portrait_url text not null,
    reg_date date not null,
    signature text,
    introduction text,
    description text,
    status int not null default 0, -- 0未激活，1正常，2封禁
    Created_At timestamp ,
    Deleted_At timestamp,
    Updated_At timestamp
)