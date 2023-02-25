CREATE TABLE products (
	id bigserial primary KEY not null,
  uid bigserial not null,
	title TEXT not null,
	price int not null,
  amount int not null,
	description TEXT not null,
  image_data BYTEA not null,
  image_name TEXT not null,
  image_type TEXT not null,
	created_at TIMESTAMP not null DEFAULT(now()),
	updated_at TIMESTAMP not null DEFAULT(now())
);
