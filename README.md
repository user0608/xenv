## xenv

`xenv` is a utility for parsing environment variable files (.env) and printing their contents to standard output.

### Features
- Parse .env files with ease.
- Print environment variables in a readable format.
- Useful for debugging and environment configuration management.

### Installation

To install `xenv`, run:

```sh
go install github.com/user0608/xenv@latest
```

### Usage

Create a `.env` file:

```env
DATABASE_URL=postgres://user:password@localhost:5432/dbname
SECRET_KEY=your-secret-key
DEBUG=true
```

Run `xenv` to parse and print the environment variables:

```sh
xenv
```

Output:

```
DATABASE_URL=postgres://user:password@localhost:5432/dbname
SECRET_KEY=your-secret-key
DEBUG=true
```

#### Custom File Usage

You can also specify a custom .env file as an argument:

```sh
xenv .production.env
```

#### Integration with Makefile
-  Method 1: Using installed xnv
    You can integrate `xenv` with your Makefile like this:

    ```makefile
    run:
        env $(shell xenv) go run main.go
    ```

    This command will run your Go application with the environment variables parsed by `xenv`.

-  Method 2: Without installing xnv
    You can integrate `xenv` with your Makefile like this:

    ```makefile
        variables := $(shell go run github.com/user0608/xenv@latest)
        run:
            env $(variables) go run main.go
    ```

    This command will run your Go application with the environment variables parsed by `xenv`.