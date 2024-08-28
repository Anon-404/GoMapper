package portsrv

var Service string

var portServices = map[int]string{
    80: "http",
    53: "domain",
    443: "https",
    22: "ssh",
    21: "ftp",
    23: "telnet",
    25: "smtp",
    123: "ntp",
    2086: "cPanel(non-SSL)",
    2087: "cPanel(SSL)",
    2052: "cloudflare(http-alt)",
    2053: "cloudflare(https-alt)",
    2096: "cPanel-Webmail(SSL)",
    2095: "cPanel-Webmail(non-SSL)",
    445: "smb",
    67: "dhcp",
    110: "pop3",
    143: "imap",
    465: "smtps",
    990: "ftps",
    389: "ldap",
    636: "ldaps",
    3306: "mysql",
    5432: "postgresql",
    3389: "rdp",
    6379: "redis",
    27017: "mongodb",
    9200: "elasticsearch",
    11211: "memcached",
    161: "snmp",
    179: "bgp",
    88: "kerberos",
    500: "ike",
    5060: "sip",
    1883: "mqtt",
    5900: "vnc",
    69: "tftp",
    5222: "xmpp",
    9418: "git",
    873: "rsync",
    2049: "nfs",
    3128: "squid",
    5901: "vnc",
    3000: "node",
    7000: "afp",
    9000: "php-fpm",
    3690: "svn",
    4444: "metasploit",
    5061: "sips",
    7080: "http-proxy",
    8443: "https-alt",
    2022: "bitbucket",
    25565: "minecraft",
    6667: "irc",
    10000: "webmin",
    49152: "dynamic",
    5000: "upnp",
    6000: "x11",
    8000: "http-alt",
    8081: "http-alt",
    8888: "http-alt",
    7001: "weblogic",
    4848: "glassfish",
    5433: "postgresql",
    8880: "http-alt",
    6378: "redis",
    9300: "elasticsearch",
    9201: "elasticsearch",
    2181: "zookeeper",
    11214: "memcached",
    8082: "http-alt",
    27018: "mongodb",
    50000: "sap",
    2048: "oracle",
    9500: "apache",
    50030: "hadoop",
    50070: "hadoop",
    50090: "hadoop",
    61616: "activemq",
    8181: "websphere",
    9090: "websphere",
    7474: "neo4j",
    6001: "x11",
    28017: "mongodb",
    9202: "elasticsearch",
    9301: "elasticsearch",
    5001: "kafka",
    6002: "x11",
    3700: "cassandra",
    8444: "solr",
    5902: "vnc",
    8085: "http-alt",
    8086: "influxdb",
    8090: "webadmin",
    8087: "freepbx",
    9091: "webmin",
    3127: "backdoor",
    2222: "ssh",
    37015: "cassandra",
    37777: "torrent",
    15555: "backdoor",
    31000: "backdoor",
    49153: "dynamic",
    49154: "dynamic",
    49155: "dynamic",
    49156: "dynamic",
    49157: "dynamic",
    5431: "db2",
    16384: "epmap",
    2266: "ebgp",
    3846: "bmc",
    3715: "dawn",
    3701: "smtp",
    5480: "admin",
    6464: "flash",
    8088: "http-alt",
    9092: "kafka",
    8445: "apache",
    8008: "http-alt",
    6100: "x11",
    4445: "backdoor",
    5353: "mdns",
    6881: "torrent",
    33060: "mysqlx",
    20048: "nfs3",
    135: "msrpc",
    162: "snmptrap",
    7547: "cwmp",
    5701: "hazelcast",
    1514: "syslog",
    7575: "couchbase",
    6882: "torrent",
    1099: "rmi",
    2017: "robo",
    8099: "bigfix",
    9929: "rsync",
    31337: "back-orifice",
}

func PortSrv(port int) {
    if service, ok := portServices[port]; ok {
        Service = service
    } else {
        Service = "unknown"
    }
}
