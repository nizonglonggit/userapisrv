create table role_permission
(
    role_id int,
    permission_id int,
    CreatedAt int,
    DeletedAt int,
    UpdatedAt int,
    primary key (role_id, permission_id)
);

create table user_role
(
    user_id int,
    role_id int,
    CreatedAt int,
    DeletedAt int,
    UpdatedAt int,
    primary key (user_id, role_id)
);