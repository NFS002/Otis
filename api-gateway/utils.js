/* Utility functions */
const { loadPackageDefinition } = require("grpc")
const config = require("./gateway-config")[process.env.NODE_ENV || "development"]
const protoLoader = require("@grpc/proto-loader")

/**
 * A function to take a string written in dot notation style, and use it to
 * find a nested object property inside of an object.
 */
function getValue (key) {
	var parts = key.split("::")
	var length = parts.length
	var property = config

	for (var i = 0; i < length; i++)
		property = property[parts[i]]

	return property
}

function getService (serviceName) {
	const services = config.services
	for (var service of services)
		if (service.name === serviceName) return service

	throw Error(`Service ${serviceName} not found`)
}

function isEmpty (obj) {
	for (var i in obj)
		if (Object.hasOwnProperty.call(obj, i))
			return false

	return true
}

function getReqData (req) {
	switch (req.method) {
	    case "DELETE":
	    case "GET":
		    return req.query

	    case "POST":
	    default:
		    return req.body
	}
}


function wrapFuncInMiddleware (func, validators) {
	return async function wrap (req, res, next) {
		try {
		    if (validators && validators.length)
		        for (var validate of validators)
		            await validate(getReqData(req))
		    func(req, res).catch(e => next(e))
		} catch (e) {
			next(e)
		}
	}
}

/* Load a grpc protobuf pacakge */
function getProtoPackage (protoPath, packageName, dirs = [".", "proto"]) {
	const packageDefinition = protoLoader.loadSync(protoPath, {
		keepCase: false,
		longs: String,
		enums: String,
		defaults: true,
		oneofs: true,
		includeDirs: dirs
	})

	return loadPackageDefinition(packageDefinition)[packageName]
}

class ServiceError extends Error {
	constructor (message, ...args) {
		super(message, ...args)
	}
}

/* Throws a generic server error if the given property exists on the given object */
function throwForProperty (obj, prop, message = "Request not executed") {
	if (Object.hasOwnProperty.call(obj, prop) && obj.executed === false)
		throw new ServiceError(message)
}

module.exports = {
	getReqData,
	getProtoPackage,
	getService,
	getValue,
	isEmpty,
	wrapFuncInMiddleware,
	throwForProperty
}
