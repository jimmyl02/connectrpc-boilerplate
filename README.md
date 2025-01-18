# connectrpc-boilerplate

this repository is a boilerplate example of using connectrpc within a monorepo. it includes a golang backend microservices architecture built with bazel and a frontend architecture built with turborepo. protobuf is used as the schema format shared between all services.

we target two different build systems as bazel is well supported for primarily backend languages but doesn't integrate super well with the javascript world. it's a bit unfortunate for the divergence but turborepo supports javascript much better than bazel does.

## usage

### backend

#### running the backend service

`bazelisk run //services/service-a`

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

#### installing external dependencies

dependencies should be installed where they are used so that everything has what it needs

`pnpm install <package> --filter=<package>`

#### installing internal dependencies

add the package name to the `package.json` dependencies of the application

```
  "dependencies": {
    ...
    "@repo/gen-api": "workspace:*",
    ...
  },
```

then insure it's installed by running `pnpm install`

#### generating connectrpc clients

in the root run `pnpm generate` to re-generate the typescript types.

if we add a new service we want to expose to the frontend, we have to export the service client within the `packages/gen-api` workspace. we do this by first running the generate command and then adding the service export to `index.ts` via `export * from "./service_a/v1/service_a_pb";` which allows us to export the service from the package.

### protobuf

in general, you should be able to modify the protobuf files and they will be automatically picked up and regenerated. once modified, if a new import was added make sure to run `bazelisk run //:gazelle`!

the `buf` vscode extension is recommended as it will automatically run linting and catch issues early. this helps prevent the "silent" errors that can happen when protobuf fails to compile as it might seem they just aren't picked up by the language servers.

reminder: frontend generation of conenctrpc clients is not automatic! you must run `pnpm generate` to re-generate the types.

### FAQ

#### question: my language server is complaining something isn't found!

because bazel does generations behind the scenes, it's common that the language server won't pick up new changes without forcefully reloading it. most issues can be resolved in vscode by running `ctrl+shift+p` and selecting `Restart language server` for the appropriate language

## references

backend - https://github.com/jimmyl02/bazel-playground-connectrpc
frontend - https://dev.to/joemckenney/building-a-modern-grpc-powered-microservice-using-nodejs-typescript-and-connect-51a9
