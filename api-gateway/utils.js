/* Utility functions */
const { loadPackageDefinition, credentials } =  require('grpc');
const config = require('./gateway-config')[process.env.NODE_ENV || 'development'];
const protoLoader = require('@grpc/proto-loader');


/**
 * A function to take a string written in dot notation style, and use it to
 * find a nested object property inside of an object.
 */
function getValue( key ) {
    var parts = key.split( "::" ),
        length = parts.length,
        property = config;

    for ( var i = 0; i < length; i++ ) {
        property = property[parts[i]];
    }

    return property;
}

function getService( service_name ) {
    const services = config.services
    for ( service of services ) {
        if ( service.name == service_name )
            return service
    }
    throw Error(`Service ${service_name} not found`)
}


function wrapFuncInMiddleware( func ) {
    return async function wrap( req, res, next ) {
        try {
            await func(req, res, next)
        }
        catch(e) {
            next(e)
        }
    }
}

/* Load a grpc protobuf pacakge */
function getProtoPackage( proto_path, package_name, dirs = ['proto'] ) {

    const packageDefinition = protoLoader.loadSync(proto_path, {
        keepCase: false,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true,
        includeDirs: dirs
    });

    return loadPackageDefinition(packageDefinition)[package_name]
}

module.exports = {
    getProtoPackage,
    getService,
    getValue,
    wrapFuncInMiddleware
}