# Chieftan
**A task automation framework driven through an HTTP API**

Chieftan was designed to simplify the process of creating, managing and triggering tasks
on your infrastructure. It is driven through an HTTP API, making it trivially easy to
automate and consume - while the simple API ensures you won't be left reading docs.

Its primary goal is to provide a data-defined task framework, enabling rich interactions
with other systems such as build and deployment tools.

## Getting Started

To start your Chieftan server, all you need to do is run `npm start`. You can customize
other aspects of your installation through the use of environment variables.

```sh
git clone $CHIEFTAN_REPO chief_server
cd chief_server
npm install
npm run build
npm start
```

### Configuration

 - **PORT** can be used to set the port that Chieftan listens on, it defaults to port `80`.
 - **MONGODB_URL** can be used to set the MongoDB database URL, it defaults to `mongodb://localhost/chief`.