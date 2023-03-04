init_db:
	docker exec -i auth-db psql -U test -b test < ./auth-service/config/db.sql
	docker exec -i product-db psql -U test -b test < ./product-service/config/db.sql
	docker exec -i cart-db psql -U test -b test < ./cart-service/config/db.sql
	docker exec -i order-db psql -U test -b test < ./order-service/config/db.sql
	docker exec -i payment-db psql -U test -b test < ./payment-service/config/db.sql

.PHONY: init_db
