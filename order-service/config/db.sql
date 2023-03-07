CREATE TABLE products (
	id bigint primary KEY not null,
  uid bigint not null,
	title TEXT not null,
	price int not null,
  amount int not null,
  image_name TEXT not null,
	created_at TIMESTAMP not null DEFAULT(now()),
	updated_at TIMESTAMP not null DEFAULT(now())
);

CREATE TABLE users (
  id bigint PRIMARY KEY NOT NULL,
  username varchar(50) not null,
  email varchar(100) not null unique,
  created_at TIMESTAMP not null DEFAULT(now()),
  updated_at TIMESTAMP not null DEFAULT(now())
);

CREATE TABLE orders (
  id bigserial PRIMARY KEY NOT null,
  uid bigint NOT null,
  status varchar(10) not null DEFAULT 'WAIT',
  created_at TIMESTAMP not null DEFAULT(now()),
  updated_at TIMESTAMP not null DEFAULT(now())
);

CREATE TABLE order_items (
  id bigserial PRIMARY KEY NOT null,
  oid bigint NOT NULL,
  pid bigint NOT NULL,
  amount int not null
  created_at TIMESTAMP not null DEFAULT(now()),
  updated_at TIMESTAMP not null DEFAULT(now())
);

ALTER TABLE "order_items"
ADD FOREIGN KEY ("oid")
REFERENCES "orders" ("id")
ON DELETE CASCADE;

ALTER TABLE "order_items"
ADD FOREIGN KEY ("pid")
REFERENCES "products" ("id")

ALTER TABLE "orders" 
ADD FOREIGN KEY ("uid") 
REFERENCES "users" ("id")
ON DELETE CASCADE;

ALTER TABLE "orders"
ADD FOREIGN KEY ("pid")
REFERENCES "products" ("id")
ON DELETE CASCADE;

ALTER TABLE "products"
ADD FOREIGN KEY ("uid")
REFERENCES "users" ("id")
ON DELETE CASCADE;
