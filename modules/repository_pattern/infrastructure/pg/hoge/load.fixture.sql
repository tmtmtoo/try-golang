insert into hoge (id, value) values
  ('00000000-0000-0000-0000-000000000001', 'unprocessed'),
  ('00000000-0000-0000-0000-000000000002', 'cancelled'),
  ('00000000-0000-0000-0000-000000000003', 'processed');

insert into cancelled_hoge (id, reason) values
  ('00000000-0000-0000-0000-000000000002', 'reason_text');

insert into piyo (id, value) values
  ('00000000-0000-0000-0000-000000000001', 'piyopiyo');

insert into processed_hoge (id, piyo_id) values
  ('00000000-0000-0000-0000-000000000003', '00000000-0000-0000-0000-000000000001');
