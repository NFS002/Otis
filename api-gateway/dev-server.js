const config = require('./gateway-config')[process.env.NODE_ENV || 'development'];

const { Server, ServerCredentials, loadPackageDefinition, credentials } =  require('grpc');
const protoLoader = require('@grpc/proto-loader');


const EXAMPLE_PROTO_PATH = 'example.proto';

const examplePackageDefinition = protoLoader.loadSync(EXAMPLE_PROTO_PATH, {
  keepCase: false,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
  includeDirs: ['proto', '/Users/noah/Otis/backend/' ]
});

const example = loadPackageDefinition(examplePackageDefinition).example;

function _Bite( call, callback ) {
    console.log("We have liftoff: ", call.request)
    callback( null, { "fling": JSON.stringify(call.request) } )
}

function getServer() {
  var server = new Server();
  server.addService(example.ExampleService.service, {
    _Bite: _Bite
  });
  return server;
}

var routeServer = getServer();
routeServer.bind('127.0.0.1:3005', ServerCredentials.createInsecure());
routeServer.start();
