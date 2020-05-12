/* Dependencies */
const os = require('os');
const _ = require('lodash')
const winston = require('winston');
const { getValue } = require('../utils');
const { Rollbar } = require('./winston-rollbar');

const { format } = winston

const log_config = getValue('logs')
const log_files = log_config.files

const winston_transports = log_files.map( log  => new winston.transports.File( log ) );

/* Print to console */
if (log_config.console === true) {
    winston_transports.push(  new winston.transports.Console() )
}

/* Log to rollbar  */
if ( log_config.rollbar === true ) {
  const r = new Rollbar({
     rollbarConfig: {
       accessToken: process.env["OTIS_ROLLBAR_ACCESS_TOKEN"],
       environment: process.env["OTIS_ENV"],
       reportLevel: 'error'
    },
  })
  winston_transports.push( r )
}

const get_winston_format = function() {
   if (log_config.pretty_print === true) {
        return format.combine(
         format.timestamp(),
         format.json(),
         format.prettyPrint()
        )
   }
   else {
        return format.combine(
        format.timestamp(),
        format.json()
     )
   }
}

const winston_conf = {
   level: 'info',
   format: get_winston_format(),
   meta: true,
   defaultMeta: { service: 'api-gateway' },
   transports: winston_transports
}

const get_log_message = function( req, res ) {
    return  '[api-gateway:' + req.hostname + '] ' + req.method + "/" + req.httpVersion
    + " " + req.originalUrl + " (" + res.statusCode + ")"
}

const get_meta_log_body = function( req, res ) {
    const meta = {}
    meta['memory_usage'] = process.memoryUsage();
    meta['total_mem'] = os.totalmem();
    meta['pid'] = process.pid ? process.pid : '-';
    meta['gid'] = process.gid ? process.gid : '-';
    meta['argv'] = process.argv ? process.argv : '-';
    meta['platform'] = process.platform ? process.platform : '-';
    meta['total_unused_mem'] = os.freemem();
    return meta;
}

const get_err_log_body = function( err ) {
    const err_log_body = {};
    err_log_body['message'] = err.message || 'UNKNOWN';
    err_log_body['stack'] = err.stack
    err_log_body['error'] = err;
    return err_log_body;
}

const get_log_body = function( time, req, res, err ) {
    const req_log_body = {}
    const res_log_body = {}
    const log_body = {}

    req_log_body['type'] = 'http request'
    res_log_body['type'] = 'http response'

    req_log_body['headers'] = req.headers
    res_log_body['headers'] = res.getHeaders()

    req_log_body['url'] = req.originalUrl;


    req_log_body['host'] = req.hostname
    req_log_body['ip'] = req.ip

    const req_body = req.body
    const res_body = res.locals.body

    if (req_body)
        req_log_body['body'] = JSON.stringify(req_body);
    if (res_body)
        res_log_body['body'] = JSON.stringify(res_body);

    const req_cookies = req.cookies;
    const res_cookies = res.cookies;

    if (req_cookies)
        req_log_body['cookies'] = req_cookies
    if (res_cookies)
        res_log_body['cookies'] = res_cookies


    log_body['request'] = req_log_body;
    log_body['response'] = res_log_body
    log_body['message'] = get_log_message( req, res )
    log_body['meta'] = get_meta_log_body( req, res )

    if ( err ) {
        const err_log_body = get_err_log_body( err )

        log_body['error'] = err_log_body
        log_body['message'] += (' - ' + err_log_body['message'])
    }

    if ( time ) {
      log_body['response_time'] = time.toFixed(3)
      if ( time > 3500 ) {
        res.locals.logLevel = 'warning'
      }
    }

    log_body['level'] = res.locals.logLevel || 'info'

    return log_body
}


const logger = winston.createLogger( winston_conf );


module.exports = {
    logger,
    get_log_body
}