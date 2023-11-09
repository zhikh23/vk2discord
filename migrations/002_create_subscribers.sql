CREATE TABLE subscribers (
    domain     VARCHAR NOT NULL,
    channel_id BIGINT NOT NULL,

    UNIQUE (domain, channel_id)
);

---- create above / drop below ----

DROP TABLE subscribers;
