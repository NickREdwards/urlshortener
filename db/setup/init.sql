CREATE TABLE shortened_urls (
  id bigint not null generated always as identity,
  short_code varchar(6) not null,
  long_url varchar(500) not null
);

CREATE INDEX short_code_index ON shortened_urls (short_code);