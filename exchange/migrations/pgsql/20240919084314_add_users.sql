-- +goose Up
-- +goose StatementBegin

-- Создание таблицы юзеров
CREATE TABLE ssp.users (
                           id varchar NOT NULL,
                           last_name varchar NOT NULL,
                           first_name varchar NOT NULL,
                           email varchar NOT NULL,
                           password_hash bytea NOT NULL,
                           password_salt bytea NOT NULL,
                           author_id varchar NULL,
                           last_login_at timestamptz NULL,
                           is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
                           CONSTRAINT users_pk PRIMARY KEY (id),
                           CONSTRAINT users_users_fk FOREIGN KEY (author_id) REFERENCES ssp.users(id)
);

-- Создание таблицы разрешений пользователя
CREATE TABLE ssp.permissions (
                           id int4 NOT NULL,
                           description varchar NOT NULL,
                           CONSTRAINT permissions_pk PRIMARY KEY (id)
);

-- Создание таблицы связки таблицы юзеров и разрешениями
CREATE TABLE ssp.permissions_to_users (
                            permission_id int4 NOT NULL,
                            user_id varchar NOT NULL,
                            CONSTRAINT permissions_to_users_pk PRIMARY KEY (permission_id,user_id),
                            CONSTRAINT permissions_to_users_permissions_fk FOREIGN KEY (permission_id) REFERENCES ssp.permissions(id),
                            CONSTRAINT permissions_to_users_users_fk FOREIGN KEY (user_id) REFERENCES ssp.users(id)
);

-- Добавляем данные в таблицу permissions
INSERT INTO ssp.permissions (id, description) VALUES(1, 'ROOT');
INSERT INTO ssp.permissions (id, description) VALUES(2, 'ADMIN');
INSERT INTO ssp.permissions (id, description) VALUES(3, 'USER_CREATE');
INSERT INTO ssp.permissions (id, description) VALUES(4, 'USER_UPDATE');
INSERT INTO ssp.permissions (id, description) VALUES(5, 'PERMISSIONS_MANAGING');
INSERT INTO ssp.permissions (id, description) VALUES(6, 'USER_DELETE');
INSERT INTO ssp.permissions (id, description) VALUES(7, 'USER_REMOVE');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE ssp.users;
DROP TABLE ssp.permissions;
DROP TABLE ssp.permissions_to_users;
DELETE FROM ssp.permissions WHERE id IN (1,2,3,4,5,6,7);
-- +goose StatementEnd
