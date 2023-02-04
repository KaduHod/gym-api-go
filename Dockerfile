FROM golang:bullseye

WORKDIR /usr/src/app

COPY . ./usr/src/app

COPY --from=node:lts-bullseye /usr/local/bin/node /usr/local/bin/node

COPY --from=node:lts-bullseye /usr/local/lib/node_modules /usr/local/lib/node_modules

COPY go.mod go.sum ./

RUN ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npm

EXPOSE 3000

RUN go mod download

RUN npm i -g nodemon

CMD ["nodemon","--exec","go","run","main.go","--signal","SIGTERM"]