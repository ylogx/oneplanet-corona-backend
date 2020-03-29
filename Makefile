GIT_TAG=$$(git describe --always)

all: test

run:
	cd api && go run .

server:
	# Install air from https://github.com/cosmtrek/air i.e. curl -fLo air https://git.io/darwin_air
	cd api && air

test:
	cd api && go test -v ./...

GCR_IMG_NAME=us.gcr.io/public-projects-2603/oneplanet/corona/backend

docker_login:
	yes | gcloud auth configure-docker us.gcr.io --project public-projects-2603

build_gcr_docker:
	docker build -f Dockerfile -t ${GCR_IMG_NAME}:${GIT_TAG} .
	docker tag ${GCR_IMG_NAME}:${GIT_TAG} ${GCR_IMG_NAME}:latest

push_gcr_docker: build_gcr_docker docker_login
	docker push ${GCR_IMG_NAME}:${GIT_TAG}
	docker push ${GCR_IMG_NAME}:latest

gcr_deploy:
	gcloud run deploy --project public-projects-2603 --region us-central1 --platform=managed --image="${GCR_IMG_NAME}:${GIT_TAG}" oneplanet-corona-backend

gcr_pipeline: build_gcr_docker push_gcr_docker gcr_deploy
