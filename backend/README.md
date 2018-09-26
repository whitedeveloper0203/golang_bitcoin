##ToDo:
-ASAP review method of connecting to db (can we open the connection only once in the main func, exposing its methods cross-package? Is that safe?)
-build the demo bet engine that freezes users bet amount, and performs appropriate calculations (unfreeze balance and grant profit of expiry or burn users bet amount.) depending on outcome at end of expiry
-should outcome be based on random numbers or should actually just we hit cryptocompares api for this?

###Im Using postgres.app with the postico manager to spin up our db.
Start a server on localhost with postgres.app or similar and create a db called "userinfo_db", do this inside -
```bash
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    auth0_uuid character varying(50) NOT NULL UNIQUE,
    email character varying(255) NOT NULL UNIQUE,
    email_verified boolean NOT NULL
);

CREATE UNIQUE INDEX users_pkey ON users(id int4_ops);
CREATE UNIQUE INDEX users_auth0_uuid_key ON users(auth0_uuid text_ops);
CREATE UNIQUE INDEX users_email_key ON users(email text_ops);
```

Run the app -

```bash
cd backend
go build
./backend
```

### Run the Application

```bash
npm start
```

The application will be served at `http://localhost:3000`.

### Run the Application With Docker

In order to run the with docker you need to have `docker` installed.

Execute in command line `sh exec.sh` to run the Docker in unix, or `.\exec.ps1` to run the Docker in Windows.
