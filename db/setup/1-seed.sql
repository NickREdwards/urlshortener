INSERT INTO
  shortened_urls (short_code, long_url)
VALUES
  (
    'poiuy',
    'http://www.google.com/search?q=fjsdiugudfjjdgijsdngindfgidfgunisdunfiungdsig'
  ),
  (
    'dfgSa',
    'http://www.google.com/search?q=dfkgjsdgjkdfkgskdg'
  ),
  (
    'EIFdf',
    'http://www.google.com/search?q=orojgiortmomtrvtrmvmtrvmtrv'
  ),
  (
    '78dgA',
    'http://www.google.com/search?q=saofkioefomoweimfoimadfmlkdsmfklds'
  );
INSERT INTO
  access_logs (shortened_url_id, access_date_time)
VALUES
  (1, '2019-09-08 11:00:00.000000'), -- 6 hours ago
  (2, '2019-09-08 12:00:00.000000'), -- 5 hours ago
  (2, '2019-09-07 20:00:00.000000'), -- 21 hours ago
  (2, '2019-09-02 17:00:00.000000'), -- 6 days ago
  (2, '2019-08-31 17:00:00.000000'), -- 8 days ago
  (2, '2019-08-14 17:00:00.000000')  -- 25 days ago