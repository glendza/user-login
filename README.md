# User Login Code Challenge

Using a JavaScript framework of your choice (preferably [React](https://reactjs.org/)), create a simple login screen that allows users to enter their username and password and submit the login form to a backend process.

Using [GoLang](https://go.dev/), create a backend that processes the login information and checks if the username and password are valid. If the login information is valid, the backend should return a success message to the user. If the login information is invalid, the backend should return an error message to the user.


## Instructions
1. Click "Use this template" to create a copy of this repository in your personal github account.  
1. Update the README in your new repo with:
    * a `How-To` section containing any instructions needed to execute your program.
    * an `Assumptions` section containing documentation on any assumptions made while interpreting the requirements.
1. Send an email to Scoir (code_challenge@scoir.com) with a link to your newly created repo containing the completed exercise (preferably no later than one day before your next interview).

## Expectations
1. This exercise is meant to drive a conversation between you and Scoir's hiring team.  
1. Please invest only enough time needed to demonstrate your approach to problem solving and code design.  
1. Within reason, treat your solution as if it would become a production system.
1. If you have any questions, feel free to contact us at code_challenge@scoir.com

## How-To

### Cloning the repo

`cd` into a wanted directory and do the:

```sh
git clone https://github.com/glendza/user-login.git user-login-test-challenge
```

### Running back-end

From the repo root:

```sh
# Navigate to back-end app directory
cd server
# Install dependencies
go get .
# Run back-end
go run *.go
```

### Running front-end

From the repo root:

```sh
# Navigate to front-end app directory
cd client
# Create the environment file
cp .env.example .env
# Install dependencies
npm install
# Run front-end
npm run start
```

### Running dockerized

From the repo root:

```sh
docker-compose build --no-cache && docker-compose up -d
```

The app will be available on port 3000.

## Assumptions & shortcomings
- The app was developed and tested on Linux. Should make no significant difference for MacOS or any other Unix-based OS, but for windows, the setup is probably different;
- Since the solution should be treated as it would become a production system, special attention was given to make the solution scalable;
- Since it's just a showcase, there is no DB/ORM installed for the time being;
- The app itself is crude, no attention whatsoever was given to the styling of the actual layout;
- Docker setup is not optimized;
- Barely any attention was given to environment separation.
