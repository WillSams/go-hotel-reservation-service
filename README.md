# GoLang Example - Hotel Reservation Service

This is still a work-in-progress, the tests are broken.

Table of Contents

- [Overview](#overview)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Running the service](#running-the-service)
  - [Debugging](#debugging)
- [Running the tests](#running-the-tests)s
- [The Case for Why We Should Switch to Go For Our Backends](./docs/case_for_go.md)

## Overview

This is a simple example of a hotel reservation service written in Go. It is aGraphQL API that allows you to create and list reservations as well as the ability to list available rooms for a given date range.

My first attempt at a Golang application.  I may have done some unconvential things here.  

## Prerequisites

To run the service, you will need to install the following tools.

- [Go Lang](https://golang.org/)
- [NodeJS](https://nodejs.org/en/).  Used for the Serverless Framework for local development and for migrations.
- [nvm](https://github.com/nvm-sh/nvm).  Used to manage NodeJS versions.  However, this is optional but I've included an `.nvmrc` file in this repository just in case.
- [Direnv](https://direnv.net/).  Used to manage environment variables.
- [Docker](https://www.docker.com/).  For deploying to AWS Lambda. Actually, to ECR and then to Lambda. With that said, you can skip this if you don't want to deploy.  Alternatively, you can use the full extent of the Serverless Framework but that's out of the scope of this README.

## Getting Started

First things first, we'll need to set up our environment variables.  If you've never used Direnv before, you'll understand it's use soon enough.

```bash
cp .envrc.example .envrc
direnv allow
```

### Install Node Packages

Execute the following within your terminal:

```bash
 # To eliminate any issues, install/use the version listed in .nvmrc.  
 # If you need to install the listed version of Node, execute `nvm install <version-listed-in-.nvmrc>`
nvm use            

cd ./db
npm i  # Install the packages needed for migrations/seeding, knex and pg
```

### Create the database

Let's create and seed the databases and our Reservations and Rooms tables:

```bash
# Create the databases and seed them
NODE_ENV=development | ./create_db.sh && npm run refresh
NODE_ENV=test | ./create_db.sh && npm run refresh
```

Now, navigate back to the root of the project via `cd ..` and execute the following to install the Go packages needed for the service:

```bash
go mod tidy
```

Plus, you'll need to install the following packages **globally*:

```bash
go install github.com/onsi/ginkgo/v2/ginkgo         # install the Ginkgo BDD testing framework
go install github.com/go-delve/delve/cmd/dlv@latest # install the Delve debugger
```

I should explain what *globally* means in regards to your Go installation.  When you install a package globally, it will be installed in your `$GOPATH/bin` directory.  In my case, it's `/Users/username/go/bin`.  You can check your `$GOPATH` by executing `go env GOPATH`.  If you don't have a `$GOPATH` set, it will default to `$HOME/go`.  You can set your `$GOPATH` by executing `export GOPATH=/path/to/your/go/path`.  You can also add this to your `.bash_profile` or `.zshrc` file.  For instance, I have the following in my `.bashrc` file:

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

## Running the service

To run the Lambda function locally, we'll use the Serverless Framework's [go](https://github.com/mthenw/serverless-go-plugin) and [offline](https://github.com/dherault/serverless-offline) plugins. We'll install NPM packages locally to do so:

```bash
nvm use          # optional; you can just use the version of Node listed in .nvmrc
npm i -g serverless serverless-go-plugin serverless-offline

# You'll need to modify the serverless.yml file to use the environment variables in your .envrc file
cp serverless.yml.example serverless.yml
```

Subsequently, we can run the service locally by executing `make run` in the root of the project.  Build the GraphQL Playground handler and start the service on ports 8081 and 8080, respectively, or whatever ports the *PLAYGROUD_PORT* and *API_PORT* environment variable are set to.

```cli
curl http://localhost:$API_PORT/development/api \
  -H 'Content-Type: application/json' \
  -d 'query GetAllReservations {\n reservations {\n Id\n RoomId\n }\n}'
```

Viola!  Again, you can also acces the non-Lambda function GraphQL playground at [http://localhost:$PLAYGROUND_PORT/playground](http://localhost:$PLAYGROUND_PORT/playground).  

In summary, to run the Lambda function and the GraphQL playground locally, execute the following:

```bash
docker-compose up -d
make run
```

### Debugging

You can painlessly debug your service using [Delve](https://github.com/go-delve/delve) and it works in VS Code as well.  

Delve is a debugger for the Go programming language. It provides GDB-like command-line debugging experience and is much more powerful than the standard Go debugger. To install it, run the following command:

```cli
go install github.com/go-delve/delve/cmd/dlv@latest
```

If you are using VS Code, install the [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go) extension and add the following configuration to your `launch.json` file:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Code",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "port": 2345,
            "host": "127.0.0.1",
            "program": "${workspaceRoot}",
            "env": {},
            "args": [],
            "showLog": true
        },
        {
            "name": "Test Current File",
            "type": "go",
            "request": "launch",
            "mode": "test",
            "port": 2345,
            "host": "127.0.0.1",
            "program": "${file}",
            "env": {},
            "args": [],
            "showLog": true
        }       
    ]
}
```

You can then start the debugger by pressing `F5` or by clicking on the `Debug` button in the VS Code sidebar.  To make it easier to debug, there is a `api/debug.go` file containing functions you can use in main to debug the service.  For example, you can change the following in your `lambdafunc/main.go` file:
  
  ```go
  func main() {
    // lambda.Start(api.GraphQlApiHandler)

    api.DebubGraphQlApiHandler()
  }
  ```

You can then set breakpoints in VS Code and debug the service with ease.

## Running the tests

This project contains BDD style tests with the help of [Ginkgo v2](https://onsi.github.io/ginkgo/). You will need to have Ginkgo installed, something you should have achieved when you followed the [Getting Started](#getting-started) step.  To run the tests, execute the following:

```bash
go test ./specs
```

## Deploying the service

There are multiple options to deploy the Lambda functions.  You can use the Serverless Framework, AWS SAM, AWS CLI, push Docker containers to ECR, or use a custom GitHub Action.  Using Docker containers may simplify things but it may lengthen cold start times and add additional costs.  Going the GitHub Action route is a more cost-effective route to build and deploy Lambda functions to AWS.  Using a GitHub Action to build and deploy your Lambda functions to AWS can be a more cost-effective approach, as it can leverage the existing infrastructure of your GitHub repository and doesn't require additional resources to be provisioned. This approach can also be more flexible and customizable, as you can tailor the deployment workflow to meet your specific needs.  See the [deployment action workflow](.github/workflows/deployment.yml) for more details of how you would deploy the service.  For the example workflow, you'll need to add the following secrets to your GitHub repository: *AWS_ACCESS_KEY_ID* and *AWS_SECRET_ACCESS_KEY*.  

### Other Resources

- [Go Lang](https://golang.org/)
- [Serverless Framework](https://www.serverless.com/)
- [HowToGraphQL Go Lang Tutorial](https://www.howtographql.com/graphql-go/0-introduction/)
