setup:
	mkdir db-data; \
	mkdir db-data/postgres; \
	mkdir db-data/mongo; \
	mkdir db-data/rabbitmq; \
	openssl genrsa -out private-key.pem 4096; \
	cp -f private-key.pem authentication-service; \
	cp -f private-key.pem gateway-service