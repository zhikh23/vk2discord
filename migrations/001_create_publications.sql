CREATE TABLE publications (
  domain VARCHAR NOT NULL,
  id     INTEGER NOT NULL,

  PRIMARY KEY (domain, id)
)

---- create above / drop below ----

drop table publications;
