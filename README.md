# Chieftan
**A task automation framework driven through an HTTP API**

Chieftan was designed to simplify the process of creating, managing and triggering tasks
on your infrastructure. It is driven through an HTTP API, making it trivially easy to
automate and consume - while the simple API ensures you won't be left reading docs.

Its primary goal is to provide a data-defined task framework, enabling rich interactions
with other systems such as build and deployment tools.

## Administration
The `chieftan` executable provides a command line interface which allows you to interact
with the database to perform various administrative functions such as creating users,
generating access tokens etc.

This functionality is all available through the API as well, however the command line
offers a convenient way to bootstrap your server.

**CAUTION** Changes made through the command line interface will not be logged to the
standard Audit Log provided by Chieftan as there is no way to determine their source.
As the command line is required to be run on the same server as your database, we assume
this will not be an issue, however you should put things in place to prevent unauthorized
access to your data nodes nonetheless. 

### Users

#### Creating a User

```sh
chieftan user create --admin "Benjamin Pannell" admin@sierrasoftworks.com
```

#### Get a User's Details

```sh
chieftan user info admin@sierrasoftworks.com
```

#### Removing a User

```sh
chieftan user remove admin@sierrasoftworks.com
```

### Permissions

#### Changing a User's Permissions

```sh
chieftan permissions set admin@sierrasoftworks.com project/:project project/:project/admin admin admin/users
```

#### Granting a User Permissions

```sh
chieftan permissions add admin@sierrasoftworks.com admin admin/users
```

#### Removing a User's Permissions

```sh
chieftan permissions remove admin@sierrasoftworks.com project/:project
```

### Access Tokens
#### Creating an Access Token

```sh
chieftan token create admin@sierrasoftworks.com
```

#### Getting a User's Access Tokens

```sh
chieftan token list admin@sierrasoftworks.com
```

#### Revoking an Access Token

```sh
chieftan token remove 54e81577954376dac6ed9d8b134a790b
```

#### Revoking a User's Access Tokens

```sh
chieftan token remove --user admin@sierrasoftworks.com
```

#### Revoking All Access Tokens

```sh
chieftan token remove --global
```

## Development
Development is conducted using standard Go and `gvt` for dependency management. Tests
are written using `gocheck` to help keep things a bit more succinct.

You are expected to have a MongoDB database available at `$MONGODB_URL` or
`mongodb://localhost/chieftan` for all tests and executions.

### Environment Setup
You just need to restore the correct versions of the various dependencies using `gvt`.

```sh
go get -u github.com/FiloSottile/gvt
gvt restore
```

### Automated Tests
You can run the automated test suite by executing the following:

```sh
goconvey
```

Alternatively, you may make use of the standard `gotest` binary if you wish to run the
tests on a headless server.

```sh
go test ./ ./api ./executors ./models ./tasks ./tools ./utils
```

### Building
Creating a build requires you to run `go build`. You can provide additional linker
flags to embed version information in the binary, specifically the git commit and
semantic version of the release.

```sh
go build -o chieftan "-X main.version=$(git describe --abbrev=0 --tags)" "-X main.commit=$(git rev-parse HEAD)"
```