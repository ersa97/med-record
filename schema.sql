CREATE TABLE users (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    roleid BIGINT NOT NULL,
    username text NOT NULL,
    password_hash text NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY (id),
    FOREIGN KEY (roleid) REFERENCES roles(id)
);

CREATE TABLE roles (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    rolename text NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY (id)
);

CREATE TABLE rolepermissions (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    roleid BIGINT NOT NULL,
    permission_key varchar(50),
    created_at datetime NOT NULL,
    updated_at datetime,
    deleted_at datetime,
    PRIMARY KEY (id),
    FOREIGN KEY (roleid) REFERENCES roles(id)
);