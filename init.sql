CREATE DATABASE kitchen_test;

\connect kitchen_test;

CREATE TABLE IF NOT EXISTS dishes (
  id SERIAL PRIMARY KEY,
  restaurant INTEGER,
  name TEXT,
  description TEXT,
  price DOUBLE PRECISION,
  image TEXT,
  protein DOUBLE PRECISION,
  fat DOUBLE PRECISION,
  carbs DOUBLE PRECISION,
  calories DOUBLE PRECISION
);

CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		restaurant INTEGER NOT NULL,
		items JSONB NOT NULL,
		order_time TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		status TEXT DEFAULT 'pending',
		payment_status TEXT DEFAULT 'unpaid'
);