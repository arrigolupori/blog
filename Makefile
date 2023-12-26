build:
	docker build -t sagewill/arrigolupori-blog . --build-arg pass=${pass}
run:
	docker run -p 8080:8080 sagewill/arrigolupori-blog:latest