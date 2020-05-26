const { readFileSync } = require("fs")
const { join } = require("path")
const grpc = require("grpc")
const delay = require("delay")
const stream = require("stream")
const { isEmpty } = require("./utils")

// Properties - Symbols
const maxConns = Symbol("MaxConns")
const prefix = Symbol("Prefix")
const urls = Symbol("URLs")
const client = Symbol("Client")
const connPool = Symbol("ConnPool")

// Methods - Symbols
const genCreds = Symbol("GenCreds")
const createNewConn = Symbol("CreateNewConn")
const getFreeConn = Symbol("GetFreeConn")
const findFreeConn = Symbol("FindFreeConn")
const reserveConn = Symbol("ReserveConn")
const releaseConn = Symbol("ReleaseConn")
const changeConnStatus = Symbol("ChangeConnStatus")
const initializeRPCs = Symbol("InitializeRPCs")

// Connection Status
const CONN_STATUS = {
	FREE: 0,
	BUSY: 1
}

/**
 * @typedef {object} ConnObj
 * @property {number} id
 * @property {object} conn
 */

module.exports = class GRPCClientPool {
	constructor ({ grpcPkg, serviceName, urls: serverURLs, maxConnections, timeout = 5000, rpcPrefix = "_", poolInterval = 200, tlsConf = {} } = {}) {
		if (!serviceName) throw new Error("option.serviceName is a required field")

		if (!grpcPkg) throw new Error("option.grpcPkg is a required field")

		if (!serverURLs || !serverURLs.length) throw new Error("option.url is a required field")


		this.tlsConf = tlsConf

		// Max Client connections to Server
		this[maxConns] = maxConnections || urls.length * 2

		// Prefix for GRPC Methods
		this[prefix] = rpcPrefix

		// Connection Ids
		this.connCount = 0

		// Index of server url to try next
		this.connIndex = 0

		// Free-Client Check Interval
		this.poolInterval = poolInterval

		// gRPC-Server URL
		this[urls] = serverURLs

		// gRPC Client Channel
		this[client] = grpcPkg[serviceName]

		/* Connection timeout */
		this.timeout = timeout

		// Connection Pool Buffer
		this[connPool] = {
			[CONN_STATUS.FREE]: {},
			[CONN_STATUS.BUSY]: {}
		}

		// Create a first Client
		this[createNewConn]()

		// Initialize RPC Methods by using the First Created Client
		this[initializeRPCs](this[findFreeConn]())
	}

	/**
	 * Generates a set of TLS credentials using the global config,
	 * to use in the next gRPC connection
	 *
	 */
	[genCreds] () {
	    var creds = {}
	    var opts = {}
	    if (!isEmpty(this.tlsConf) && Object.hasOwnProperty.call(this.tlsConf, "use_tls") && this.tlsConf.use_tls === true) {
	        var rootDir = this.tlsConf.root_dir
	        var rootCa = readFileSync(join(rootDir, this.tlsConf.root_ca))
	        var privateKey = readFileSync(join(rootDir, this.tlsConf.private_key))
	        var certChain = readFileSync(join(rootDir, this.tlsConf.cert_chain))
			if (Object.hasOwnProperty.call(this.tlsConf, "domain_override"))
				opts = {
				    "grpc.ssl_target_name_override": this.tlsConf.domain_override
				}

			creds = grpc.credentials.createSsl(rootCa, privateKey, certChain)
	    } else {
	        creds = grpc.credentials.createInsecure()
	    }
		return [creds, opts]
	}

	/**
	 * Creates a New Connection and Adds it to the pool in FREE status
	 */
	[createNewConn] () {
		const newConnId = ++this.connCount
		const nextUrl = this[urls][this.connIndex]
		this.connIndex = (this.connIndex + 1) % this[urls].length
		var allCreds = this[genCreds]()
		this[connPool][CONN_STATUS.FREE][newConnId] = {
			conn: new this[client](nextUrl, allCreds[0], allCreds[1]),
			id: newConnId
		}
	}

	/**
	 * Finds/Waits for a FREE connection
	 */
	async [getFreeConn] () {
		// Minute delay for handling recursive calls
		await delay(Math.random() * 5)

		const freeConnObj = this[findFreeConn]()
		if (freeConnObj) return freeConnObj

		// if number of connections < Max Allowed Connections, then Create a New Connection
		if (this.connCount < this[maxConns])
			this[createNewConn]()
		 else
			await delay(this.poolInterval + (Math.random() * 10))

		return this[getFreeConn]()
	}

	/**
	 * Returns the first FREE connection if exists, else returns undefined
	 */
	[findFreeConn] () {
		return Object.values(this[connPool][CONN_STATUS.FREE])[0]
	}

	/**
	 * Changes the Connection Status
	 * @param {ConnObj} connObj
	 * @param {number} newStatus
	 */
	[changeConnStatus] (connObj, newStatus) {
		// Converts 0->1(FREE->BUSY) & 1->0(BUSY->FREE) for changing status
		const currStatus = newStatus ^ 1

		// Add the ConnObj to the NewStatus
		this[connPool][newStatus][connObj.id] = connObj

		// Remove the ConnObj from CurrentStatus
		delete this[connPool][currStatus][connObj.id]
	}

	/**
	 * Changes the status of the given ConnObj to BUSY
	 * @param {ConnObj} connObj
	 */
	[reserveConn] (connObj) {
		this[changeConnStatus](connObj, CONN_STATUS.BUSY)
	}

	/**
	 * Changes the status of the given ConnObj to FREE
	 * @param {ConnObj} connObj
	 */
	[releaseConn] (connObj) {
		this[changeConnStatus](connObj, CONN_STATUS.FREE)
	}

	/**
	 * Adds Methods from protoBuf file to `this` instance object
	 * @param {ConnObj} connObj
	 */
	[initializeRPCs] (connObj) {
        for (const rpc in connObj.conn) { // eslint-disable-line
			if (rpc.match(/^[a-zA-Z]+$/))
				// Creating Method on `this` instance => prefix + rpc_method
				this[`${this[prefix]}${rpc}`] = async (data, cb) => {
					const freeConnObj = await this[getFreeConn]()

					// Reserve a FREE Connection on obtaining one
					this[reserveConn](freeConnObj)

					return new Promise((resolve, reject) => {
						// To avoid Duplicate resolving of Promise
						let resolved = false
						const deadline = new Date(Date.now() + this.timeout)
						const response = freeConnObj.conn[rpc](data, { deadline: deadline }, (err, result) => {
							// Release the connection after the request is Done
							this[releaseConn](freeConnObj)

							cb && cb(err, result)
							return !resolved && (err ? reject(err) : resolve(result))
						})
						if (response instanceof stream.Readable || response instanceof stream.Writable) {
							response.on && response.on("end", () => this[releaseConn](freeConnObj))
							resolved = true
							resolve(response)
						}
					})
				}
		}
	}
}
