CREATE TABLE products (
	id bigserial primary KEY not null,
  uid bigserial not null,
	title TEXT not null,
	price int not null,
  amount int not null,
  num_unpaid int not null DEFAULT(0),
	description TEXT not null,
  image_name TEXT not null,
	created_at TIMESTAMP not null DEFAULT(now()),
	updated_at TIMESTAMP not null DEFAULT(now())
);
