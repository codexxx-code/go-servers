-- +goose Up
-- +goose StatementBegin

CREATE TYPE ssp.permission AS ENUM ('root', 'admin', 'user_create', 'user_update', 'user_permissions_managing', 'user_delete', 'user_remove');

ALTER TABLE ssp.users
ADD COLUMN permissions ssp.permission[] default '{}'::ssp.permission[];

update ssp.users set
permissions = '{root}'::ssp.permission[]; -- У нас только один юзер, поэтому не страшно

DROP TABLE IF EXISTS ssp.permissions_to_users;
DROP TABLE IF EXISTS ssp.permissions;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
