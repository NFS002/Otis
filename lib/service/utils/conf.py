"""Load and read config file"""
import os
import math
import grpc


def get_value(conf, key="", default_value="raise"):
    """ Return a value or dictionary of values from the passed configuration
    object (i.e from service-config.py), indexed by the given key"""
    if key == '':
        return conf
    parts = key.split('::')
    length = len(parts)
    for i in range(length):
        key = parts[i]
        if key not in conf and default_value != "raise":
            return default_value
        conf = conf[key]
    return conf


def gen_creds(service_config):
    """Generate a set of TLS credentials to use in a gRPC server """
    tls_conf = get_value(service_config, 'tls')
    if tls_conf is not None and 'use_tls' in tls_conf and tls_conf['use_tls'] is True:
        certs_dir = tls_conf["root_dir"]
        root_ca = open(os.path.join(certs_dir, tls_conf["root_ca"]), 'rb').read()
        private_key = open(os.path.join(certs_dir, tls_conf["private_key"]), 'rb').read()
        cert_chain = open(os.path.join(certs_dir, tls_conf["cert_chain"]), 'rb').read()
        verify_client = tls_conf.get("verify_client", True)
        creds = grpc.ssl_server_credentials(
            [(private_key, cert_chain)],
            root_certificates=root_ca, require_client_auth=verify_client,
        )

        return creds

    return None


def get_distance(start, end):
    """Distance between two points."""
    coord_factor = 10000000.0
    lat_1 = start.latitude / coord_factor
    lat_2 = end.latitude / coord_factor
    lon_1 = start.longitude / coord_factor
    lon_2 = end.longitude / coord_factor
    lat_rad_1 = math.radians(lat_1)
    lat_rad_2 = math.radians(lat_2)
    delta_lat_rad = math.radians(lat_2 - lat_1)
    delta_lon_rad = math.radians(lon_2 - lon_1)

    # Formula is based on http://mathforum.org/library/drmath/view/51879.html
    a = (pow(math.sin(delta_lat_rad / 2), 2) +
         (math.cos(lat_rad_1) * math.cos(lat_rad_2) *
          pow(math.sin(delta_lon_rad / 2), 2)))
    c = 2 * math.atan2(math.sqrt(a), math.sqrt(1 - a))
    R = 6371000
    # metres
    return R * c
