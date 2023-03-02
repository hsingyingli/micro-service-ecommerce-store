init_db:
	docker exec -i auth-db psql -U test -b test < ./auth-service/config/db.sql
	docker exec -i product-db psql -U test -b test < ./product-service/config/db.sql
	docker exec -i cart-db psql -U test -b test < ./cart-service/config/db.sql

.PHONY: init_db
