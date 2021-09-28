-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table locations(
    id integer not null primary key,
    name text not null
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table locations;