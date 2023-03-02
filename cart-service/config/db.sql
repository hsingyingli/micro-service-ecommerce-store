CREATE TABLE products (
	id bigint primary KEY not null,
  uid bigint not null,
	title TEXT not null,
	price int not null,
  amount int not null,
  image_data BYTEA not null,
  image_name TEXT not null,
  image_type TEXT not null,
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

CREATE TABLE carts (
  id bigserial PRIMARY KEY NOT null,
  uid bigint NOT null,
  pid bigint NOT null,
  amount int not null,
  created_at TIMESTAMP not null DEFAULT(now()),
  updated_at TIMESTAMP not null DEFAULT(now())
);

ALTER TABLE "carts" 
ADD FOREIGN KEY ("uid") 
REFERENCES "users" ("id")
ON DELETE CASCADE;

ALTER TABLE "carts"
ADD FOREIGN KEY ("pid")
REFERENCES "products" ("id")
ON DELETE CASCADE;

ALTER TABLE "products"
ADD FOREIGN KEY ("uid")
REFERENCES "users" ("id")
ON DELETE CASCADE;
