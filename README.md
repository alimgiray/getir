# getir

### Before running the application

- Create a `.env` file in root folder of the project. For default values you can rename `.env.example` into `.env`
- Use `go mod download` to install all required packages.

### How to run?

- For development environment, use `make dev` command.
- For production, use `make all` command.

### Running Tests

- Run `make test` command.
- You can use `example.requests` file for manuel testing. It is also a good place to examine endpoints.
- For mongo endpoints use `/remote` and for in-memory endpoints use `/in-memory`.

### Testing on live

- You can use [getir.giray.io](https://getir.giray.io) for live environment.

#### If you want to use docker

- You still need to create `.env` file
- Run `docker build --tag getir-alimgiray .` to build the image
- Run `docker run -p 8080:8080 getir-alimgiray` to run the container

#### Architecture & Patterns

- Domain related packages placed under `/internal` folder.
