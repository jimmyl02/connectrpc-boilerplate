# connectrpc-boilerplate

this repository is a boilerplate example of using connectrpc. it includes a golang backend microservices architecture built with bazel and a frontend architecture built with turborepo. protobuf is used as the schema format shared between all servies

## usage

### backend

#### running the backend service

`bazelisk run //cmd/service-a`

#### adding an external import

run the following:

1. `bazelisk run @rules_go//go -- get -u <package>`
2. `bazelisk run @rules_go//go -- mod tidy -e`
3. `bazelisk run //:gazelle`
4. `bazelisk mod tidy`
5. reload your vscode language server `Go: Restart Language Server`

#### adding an internal import

use the import automatically within golang then run `bazelisk run //:gazelle` and reload the language server

### frontend

#### creating a new vite template

1. `pnpm create vite@latest apps/frontend --template=react-ts`
2. `pnpm install`

#### running dev server

`pnpm run dev --filter=<package>`

#### installing dependencies

dependencies should be installed where they are used so that everything has what it needs

`pnpm install <package> --filter=<package>`

## references

backend - https://github.com/jimmyl02/bazel-playground-connectrpc
frontend - https://dev.to/joemckenney/building-a-modern-grpc-powered-microservice-using-nodejs-typescript-and-connect-51a9
