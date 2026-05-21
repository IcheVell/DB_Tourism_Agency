CREATE DOMAIN non_empty_varchar_255 AS VARCHAR(255)
    CHECK (btrim(VALUE) <> '');

CREATE TABLE permissions (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    code non_empty_varchar_255 NOT NULL UNIQUE,
    description VARCHAR(255)
);

CREATE TABLE roles (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name non_empty_varchar_255 NOT NULL UNIQUE,
    description VARCHAR(255)
);


CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    login non_empty_varchar_255 NOT NULL UNIQUE,
    password_hash non_empty_varchar_255 NOT NULL,
    email non_empty_varchar_255 NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT chk_users_login
        CHECK (login ~ '^[a-zA-Z][a-zA-Z0-9_]{2,49}$'),

    CONSTRAINT chk_users_email
        CHECK (email ~* '^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}$')
);

CREATE TABLE role_permissions (
    role_id BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,
    CONSTRAINT pk_role_permissions PRIMARY KEY(role_id, permission_id),

    CONSTRAINT fk_role_permissions_role_id
        FOREIGN KEY (role_id)
        REFERENCES roles(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_role_permissions_permission_id
        FOREIGN KEY (permission_id)
        REFERENCES permissions(id)
        ON DELETE CASCADE
);

CREATE TABLE user_roles (
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    CONSTRAINT pk_user_roles PRIMARY KEY(user_id, role_id),

    CONSTRAINT fk_user_roles_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id),

    CONSTRAINT fk_user_roles_role_id
        FOREIGN KEY (role_id)
        REFERENCES roles(id)
);