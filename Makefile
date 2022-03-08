run-dev:
	docker-compose down
	docker build -f container/todo/Dockerfile . -t todo
	docker-compose up

insert-mock-data:
	./scripts/filldata.sh $(BASE_URL)

build:
	docker build -f container/todo/Dockerfile . -t todo  

clean: 
	docker-compose down

push-image: 
	docker build -f container/todo/Dockerfile . -t todo  
	docker tag todo eu.gcr.io/idyllic-silicon-343409/todo  
	docker push eu.gcr.io/idyllic-silicon-343409/todo