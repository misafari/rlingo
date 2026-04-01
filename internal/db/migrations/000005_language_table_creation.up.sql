CREATE TABLE language
(
    code        VARCHAR(10) PRIMARY KEY NOT NULL,
    name        VARCHAR(50)             NOT NULL,
    native_name VARCHAR(100)            NOT NULL,
    rtl         BOOLEAN                 NOT NULL DEFAULT FALSE
)