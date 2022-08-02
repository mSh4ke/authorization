

CREATE TABLE IF NOT EXISTS roles(
id bigserial not null primary key,
name varchar(30) unique
);

CREATE TABLE IF NOT EXISTS users(
id bigserial not null primary key,
login varchar(50) not null unique,
password varchar(255),
role_id bigint not null,
CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES roles(id) on delete cascade
);
CREATE TABLE IF NOT EXISTS permisions(
id bigserial not null primary key,
name varchar(255) unique
);

CREATE TABLE IF NOT EXISTS roles_permisions(
roles_id bigint,
permisions_id bigint,
CONSTRAINT fk_roles_id FOREIGN KEY (roles_id) REFERENCES roles(id) on delete cascade,
CONSTRAINT fk_permisions_id FOREIGN KEY (permisions_id) REFERENCES permisions(id) on delete cascade
PRIMARY KEY (roles_id, permisions_id)
);