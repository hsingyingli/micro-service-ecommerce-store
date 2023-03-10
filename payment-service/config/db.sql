CREATE TABLE orders (
  id BIGINT NOT NULL PRIMARY KEY,
  uid BIGINT NOT NULL,
  price BIGINT NOT NULL,
  created_at  TIMESTAMP NOT NULL DEFAULT(now()),
  updated_at  TIMESTAMP NOT NULL DEFAULT(now())
);


CREATE TABLE payments (
  id BIGSERIAL NOT NULL PRIMARY KEY,
  uid BIGINT NOT NULL,
  oid BIGINT NOT NULL,
  created_at  TIMESTAMP NOT NULL DEFAULT(now()),
  updated_at  TIMESTAMP NOT NULL DEFAULT(now())
);

-- ALTER TABLE "payments"
-- ADD FOREIGN KEY ("oid")
-- REFERENCES "orders" ("id");
