CREATE TABLE transactions
(
    id       UUID PRIMARY KEY   DEFAULT gen_random_uuid(),
    user_id  INT,
    order_id VARCHAR ,
    date     TIMESTAMP NOT NULL DEFAULT (now() AT TIME ZONE 'utc'),
    cost     decimal,
    status   VARCHAR
);

