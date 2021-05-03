# Go(lang) tools for exploiting warm/cold boot in iCE40 FPGAs (root)

fmt-all: fmt-lib fmt-cli fmt-api

fmt-api:
	cd api && gofmt -s -w .

fmt-cli:
	cd cli && gofmt -s -w .

fmt-lib:
	cd lib && gofmt -s -w .

prod-app:
	cd app && yarn build

dev-serve:
	cd api && go run main.go -v -p 8080 -d ../app/dist

dev:
	cd app && \
	yarn build && \
	cd ../doc && hugo && mv public ../app/dist/doc && \
	cd ../api && go run main.go -v -p 8080 -d ../app/dist
