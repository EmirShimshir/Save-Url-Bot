build:
	docker build -t save-url-bot .

run:
	docker run --name save-url-bot save-url-bot

clean:
	docker stop save-url-bot || true
	docker rm save-url-bot || true
	docker rmi save-url-bot || true
