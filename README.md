### Authenthication, login and register module made in golang - by RGMA91

This service can:
* Register an user, and save the password after hashing and salting (encrypted)
* Login user with email and password and return a JWT Token
* Protect paths by verifying the JWT returned in the login

To try it, just run main.go file with go:

	go run main.go

---
Note: the database connector is configured to connect to a PostgreSQL database, with an "account" table needed with the following structure:


	CREATE TABLE public.account (
		account_id serial4 NOT NULL,
		username varchar(255) NOT NULL,
		email varchar(255) NOT NULL,
		salt text NOT NULL,
		passhash text NOT NULL,
		CONSTRAINT account_email_key UNIQUE (email),
		CONSTRAINT account_pkey PRIMARY KEY (account_id),
		CONSTRAINT account_username_key UNIQUE (username)
	);
---

The endpoints can be found in the main.go file

Routes with no protection (login and registration):
* /api/user/register (POST) -> user registration
* /api/user/login (POST) -> login

Routes that require authorization with JWT:
* /api/authenticate (GET) -> only checks if token is valid
* /api/logic/ (GET) -> example logic that prints a message if user is authenticated, is used as example for a protected path that implements some logic

Example requests
#
Register:
(POST)

	http://localhost:3000/user/register

Body:

	{
    	"username": "username",
    	"email": "test@test.com",
    	"password": "123abc"
	}
#

Login:
(POST)
	
	http://localhost:3000/user/login

Body:

	{
    	"email": "test@test.com",
    	"password": "123abc"
	}
#


Protected routes require the auth token, you can get it in the login an then use it as bearer token in the "Authorization" header