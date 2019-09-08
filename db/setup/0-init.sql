CREATE TABLE shortened_urls (
  id bigint not null generated always as identity,
  short_code varchar(6) not null,
  long_url varchar(500) not null
);
CREATE INDEX short_code_index ON shortened_urls (short_code);

CREATE TABLE access_logs (
  id bigint not null generated always as identity,
  shortened_url_id bigint not null,
  access_date_time timestamp not null
);