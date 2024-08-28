package portsrv

var Service = "Unknown"

func PortSrv(port int) {
    switch port {
case 80:
    Service = "http"
case 53:
    Service = "domain"
case 443:
    Service = "https"
case 22:
    Service = "ssh"
case 21:
    Service = "ftp"
case 23:
    Service = "telnet"
case 25:
    Service = "smtp"
case 123:
    Service = "ntp"
case 2086:
    Service = "cPanel(non-SSL)"
case 2087:
    Service = "cPanel(SSL)"
case 2052:
    Service = "cloudflare(http-alt)"
case 2053:
    Service = "cloudflare(https-alt)"
case 2096:
    Service = "cPanel-Webmail(SSL)"
case 2095:
    Service = "cPanel-Webmail(non-SSL)"
case 445:
    Service = "smb"
case 67:
    Service = "dhcp"
case 110:
    Service = "pop3"
case 143:
    Service = "imap"
case 465:
    Service = "smtps"
case 990:
    Service = "ftps"
case 389:
    Service = "ldap"
case 636:
    Service = "ldaps"
case 3306:
    Service = "mysql"
case 5432:
    Service = "postgresql"
case 3389:
    Service = "rdp"
case 6379:
    Service = "redis"
case 27017:
    Service = "mongodb"
case 9200:
    Service = "elasticsearch"
case 11211:
    Service = "memcached"
case 161:
    Service = "snmp"
case 179:
    Service = "bgp"
case 88:
    Service = "kerberos"
case 500:
    Service = "ike"
case 5060:
    Service = "sip"
case 1883:
    Service = "mqtt"
case 5900:
    Service = "vnc"
case 69:
    Service = "tftp"
case 5222:
    Service = "xmpp"
case 9418:
    Service = "git"
case 873:
    Service = "rsync"
case 2049:
    Service = "nfs"
case 3128:
    Service = "squid"
case 5901:
    Service = "vnc"
case 3000:
    Service = "node"
case 7000:
    Service = "afp"
case 9000:
    Service = "php-fpm"
case 3690:
    Service = "svn"
case 4444:
    Service = "metasploit"
case 5061:
    Service = "sips"
case 7080:
    Service = "http-proxy"
case 8443:
    Service = "https-alt"
case 2022:
    Service = "bitbucket"
case 25565:
    Service = "minecraft"
case 6667:
    Service = "irc"
case 10000:
    Service = "webmin"
case 49152:
    Service = "dynamic"
case 5000:
    Service = "upnp"
case 6000:
    Service = "x11"
case 8000:
    Service = "http-alt"
case 8081:
    Service = "http-alt"
case 8888:
    Service = "http-alt"
case 7001:
    Service = "weblogic"
case 4848:
    Service = "glassfish"
case 5433:
    Service = "postgresql"
case 8880:
    Service = "http-alt"
case 6378:
    Service = "redis"
case 9300:
    Service = "elasticsearch"
case 9201:
    Service = "elasticsearch"
case 2181:
    Service = "zookeeper"
case 11214:
    Service = "memcached"
case 8082:
    Service = "http-alt"
case 27018:
    Service = "mongodb"
case 50000:
    Service = "sap"
case 2048:
    Service = "oracle"
case 9500:
    Service = "apache"
case 50030:
    Service = "hadoop"
case 50070:
    Service = "hadoop"
case 50090:
    Service = "hadoop"
case 61616:
    Service = "activemq"
case 8181:
    Service = "websphere"
case 9090:
    Service = "websphere"
case 7474:
    Service = "neo4j"
case 6001:
    Service = "x11"
case 28017:
    Service = "mongodb"
case 9202:
    Service = "elasticsearch"
case 9301:
    Service = "elasticsearch"
case 5001:
    Service = "kafka"
case 6002:
    Service = "x11"
case 3700:
    Service = "cassandra"
case 8444:
    Service = "solr"
case 5902:
    Service = "vnc"
case 8085:
    Service = "http-alt"
case 8086:
    Service = "influxdb"
case 8090:
    Service = "webadmin"
case 8087:
    Service = "freepbx"
case 9091:
    Service = "webmin"
case 3127:
    Service = "backdoor"
case 2222:
    Service = "ssh"
case 37015:
    Service = "cassandra"
case 37777:
    Service = "torrent"
case 15555:
    Service = "backdoor"
case 31000:
    Service = "backdoor"
case 49153:
    Service = "dynamic"
case 49154:
    Service = "dynamic"
case 49155:
    Service = "dynamic"
case 49156:
    Service = "dynamic"
case 49157:
    Service = "dynamic"
case 5431:
    Service = "db2"
case 2181:
    Service = "zookeeper"
case 16384:
    Service = "epmap"
case 2266:
    Service = "ebgp"
case 3846:
    Service = "bmc"
case 3715:
    Service = "dawn"
case 3701:
    Service = "smtp"
case 5480:
    Service = "admin"
case 6464:
    Service = "flash"
case 8088:
    Service = "http-alt"
case 9092:
    Service = "kafka"
case 8445:
    Service = "apache"
case 8008:
    Service = "http-alt"
case 6100:
    Service = "x11"
case 4445:
    Service = "backdoor"
case 5353:
    Service = "mdns"
case 6881:
    Service = "torrent"
case 33060:
    Service = "mysqlx"
case 20048:
    Service = "nfs3"
case 135:
    Service = "msrpc"
case 162:
    Service = "snmptrap"
case 5431:
    Service = "db2"
case 8001:
    Service = "http-alt"
case 7547:
    Service = "cwmp"
case 50001:
    Service = "sap"
case 5701:
    Service = "hazelcast"
case 1514:
    Service = "syslog"
case 8002:
    Service = "http-alt"
case 7575:
    Service = "couchbase"
case 6882:
    Service = "torrent"
case 1099:
    Service = "rmi"
case 2017:
    Service = "robo"
case 8099:
    Service = "bigfix"
case 9929:
    Service = "rsync"
case 31337:
    Service = "back-orifice"
default:
    Service = "unknown"

    }
}
