CREATE TABLE roles(
    id bigserial not null primary key,
    name varchar(30) unique
);

CREATE TABLE users(
    id bigserial not null primary key,
    login varchar(50) not null unique,
    password varchar(255),
    email varchar(255) not null unique,
    display_name varchar(255),
    contact_info varchar(255),
    is_active boolean DEFAULT false,
    role_id bigint not null,
    CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES roles(id) on delete cascade
);

CREATE TABLE permissions(
    id bigserial not null primary key,
    req_path varchar(255),
    req_method varchar(255),
    req_server_id int,
    CONSTRAINT un_unique_request UNIQUE (req_path, req_method)
);

CREATE TABLE roles_permissions(
    roles_id bigint,
    permissions_id bigint,
    CONSTRAINT fk_roles_id FOREIGN KEY (roles_id) REFERENCES roles(id) on delete cascade,
    CONSTRAINT fk_permissions_id FOREIGN KEY (permissions_id) REFERENCES permissions(id) on delete cascade,
    CONSTRAINT un_unique_pair UNIQUE (roles_id, permissions_id)
);