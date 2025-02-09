-- +goose Up
INSERT INTO currency_magnifier (currency, magnifier) VALUES
  ('USD', 100),
  ('EUR', 100),
  ('GBP', 100),
  ('JPY', 1),
  ('AUD', 100),
  ('CAD', 100),
  ('CHF', 100),
  ('CNY', 100),
  ('HKD', 100),
  ('NZD', 100),
  ('SEK', 100),
  ('KRW', 1),
  ('SGD', 100),
  ('NOK', 100),
  ('MXN', 100),
  ('INR', 100),
  ('RUB', 100),
  ('ZAR', 100),
  ('TRY', 100),
  ('IDR', 1),
  ('THB', 100),
  ('MYR', 100),
  ('VND', 1),
  ('PHP', 100);

-- +goose Down
DELETE FROM currency_magnifier 
WHERE currency IN (
  'USD','EUR','GBP','JPY','AUD','CAD','CHF','CNY','HKD','NZD',
  'SEK','KRW','SGD','NOK','MXN','INR','RUB','ZAR','TRY',
  'IDR','THB','MYR','VND','PHP'
);