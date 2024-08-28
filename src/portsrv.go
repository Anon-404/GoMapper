package portsrv

var Service = "Unknown"

func PortSrv(port int) {
    switch port {
    case 1:
        Service = "tcpmux"
    case 7:
        Service = "echo"
    case 9:
        Service = "discard"
    case 11:
        Service = "systat"
    case 13:
        Service = "daytime"
    case 17:
        Service = "qotd"
    case 19:
        Service = "chargen"
    case 20:
        Service = "ftp-data"
    case 21:
        Service = "ftp"
    case 22:
        Service = "ssh"
    case 23:
        Service = "telnet"
    case 25:
        Service = "smtp"
    case 37:
        Service = "time"
    case 38:
        Service = "rap"
    case 39:
        Service = "rlp"
    case 42:
        Service = "nameserver"
    case 43:
        Service = "whois"
    case 44:
        Service = "mpx"
    case 53:
        Service = "domain"
    case 67:
        Service = "dhcp"
    case 68:
        Service = "dhcpv6-client"
    case 69:
        Service = "tftp"
    case 70:
        Service = "gopher"
    case 79:
        Service = "finger"
    case 80:
        Service = "http"
    case 81:
        Service = "hosts2-ns"
    case 82:
        Service = "xfer"
    case 83:
        Service = "mit-ml-dev"
    case 84:
        Service = "ctf"
    case 85:
        Service = "mit-ml-dev"
    case 88:
        Service = "kerberos"
    case 89:
        Service = "su-mit-tg"
    case 90:
        Service = "dnsix"
    case 91:
        Service = "mit-ml-dev"
    case 92:
        Service = "supdup"
    case 93:
        Service = "dcp"
    case 94:
        Service = "objcall"
    case 95:
        Service = "supdup"
    case 97:
        Service = "timeserver"
    case 98:
        Service = "tacacs"
    case 99:
        Service = "mit-ml-dev"
    case 101:
        Service = "hostname"
    case 102:
        Service = "iso-tsap"
    case 103:
        Service = "gppitnp"
    case 104:
        Service = "acr-nema"
    case 105:
        Service = "cso"
    case 106:
        Service = "audionews"
    case 107:
        Service = "rtelnet"
    case 108:
        Service = "pop2"
    case 109:
        Service = "pop3"
    case 110:
        Service = "pop3"
    case 111:
        Service = "sunrpc"
    case 112:
        Service = "mcidas"
    case 113:
        Service = "auth"
    case 114:
        Service = "audit"
    case 115:
        Service = "sftp"
    case 116:
        Service = "ansanot"
    case 117:
        Service = "uucp"
    case 118:
        Service = "sql-net"
    case 119:
        Service = "nntp"
    case 120:
        Service = "cfdptkt"
    case 121:
        Service = "cft"
    case 122:
        Service = "erpc"
    case 123:
        Service = "ntp"
    case 124:
        Service = "buoy"
    case 125:
        Service = "cfnp"
    case 126:
        Service = "manet"
    case 127:
        Service = "caio"
    case 128:
        Service = "etcfun"
    case 129:
        Service = "smux"
    case 130:
        Service = "at"
    case 131:
        Service = "tivoli"
    case 132:
        Service = "auditd"
    case 133:
        Service = "matip"
    case 134:
        Service = "epmap"
    case 135:
        Service = "msrpc"
    case 136:
        Service = "netbios-ns"
    case 137:
        Service = "netbios-dgm"
    case 138:
        Service = "netbios-ssn"
    case 139:
        Service = "netbios-ssn"
    case 140:
        Service = "emfis-data"
    case 141:
        Service = "emfis-cntl"
    case 142:
        Service = "brf"
    case 143:
        Service = "imap"
    case 144:
        Service = "popp3"
    case 145:
        Service = "qud"
    case 146:
        Service = "ethernet"
    case 147:
        Service = "audionet"
    case 148:
        Service = "locus-map"
    case 149:
        Service = "tacacs-ds"
    case 150:
        Service = "imap3"
    case 151:
        Service = "fmp"
    case 152:
        Service = "msg-auth"
    case 153:
        Service = "telel"
    case 154:
        Service = "timed"
    case 155:
        Service = "kerberos-adm"
    case 156:
        Service = "vics"
    case 157:
        Service = "svr"
    case 158:
        Service = "x25"
    case 159:
        Service = "s-ftp"
    case 160:
        Service = "another"
    case 161:
        Service = "snmp"
    case 162:
        Service = "snmptrap"
    case 163:
        Service = "cmip"
    case 164:
        Service = "omnis"
    case 165:
        Service = "wsa"
    case 166:
        Service = "xdmcp"
    case 167:
        Service = "s-cab"
    case 168:
        Service = "link"
    case 169:
        Service = "secure"
    case 170:
        Service = "syslog"
    case 171:
        Service = "nap"
    case 172:
        Service = "lycos"
    case 173:
        Service = "edit"
    case 174:
        Service = "tcp-scan"
    case 175:
        Service = "ra"
    case 176:
        Service = "dnsix"
    case 177:
        Service = "xdmcp"
    case 178:
        Service = "nextstep"
    case 179:
        Service = "bgp"
    case 180:
        Service = "urp"
    case 181:
        Service = "ncs"
    case 182:
        Service = "ekmap"
    case 183:
        Service = "conn"
    case 184:
        Service = "cad"
    case 185:
        Service = "camac"
    case 186:
        Service = "dixie"
    case 187:
        Service = "rmon"
    case 188:
        Service = "tnmp"
    case 189:
        Service = "ascend"
    case 190:
        Service = "tic"
    case 191:
        Service = "micp"
    case 192:
        Service = "sctp"
    case 193:
        Service = "locus-map"
    case 194:
        Service = "irc"
    case 195:
        Service = "clusterm"
    case 196:
        Service = "shell"
    case 197:
        Service = "scap"
    case 198:
        Service = "remote-file"
    case 199:
        Service = "sas"
    case 200:
        Service = "csi"
    case 201:
        Service = "csp"
    case 202:
        Service = "tft"
    case 203:
        Service = "nfsd"
    case 204:
        Service = "nsr"
    case 205:
        Service = "cft"
    case 206:
        Service = "comsat"
    case 207:
        Service = "uucp-path"
    case 208:
        Service = "uucp"
    case 209:
        Service = "auth"
    case 210:
        Service = "z3950"
    case 211:
        Service = "news"
    case 212:
        Service = "ftp"
    case 213:
        Service = "ftp-data"
    case 214:
        Service = "smtp"
    case 215:
        Service = "http"
    case 216:
        Service = "smtp"
    case 217:
        Service = "sftp"
    case 218:
        Service = "pop"
    case 219:
        Service = "pop3"
    case 220:
        Service = "imap"
    case 221:
        Service = "imap"
    case 222:
        Service = "ssh"
    case 223:
        Service = "chat"
    case 224:
        Service = "emergency"
    case 225:
        Service = "unload"
    case 226:
        Service = "os-bus"
    case 227:
        Service = "res-bis"
    case 228:
        Service = "term"
    case 229:
        Service = "app"
    case 230:
        Service = "mapper"
    case 231:
        Service = "bna"
    case 232:
        Service = "pmm"
    case 233:
        Service = "sda"
    case 234:
        Service = "cim"
    case 235:
        Service = "prism"
    case 236:
        Service = "pick"
    case 237:
        Service = "portmapper"
    case 238:
        Service = "elr"
    case 239:
        Service = "portmapper"
    case 240:
        Service = "ukey"
    case 241:
        Service = "sdd"
    case 242:
        Service = "port"
    case 243:
        Service = "cdd"
    case 244:
        Service = "tcppro"
    case 245:
        Service = "dscp"
    case 246:
        Service = "tact"
    case 247:
        Service = "handshake"
    case 248:
        Service = "sample"
    case 249:
        Service = "port-xx"
    case 250:
        Service = "dmc"
    case 251:
        Service = "int"
    case 252:
        Service = "tycho"
    case 253:
        Service = "kiddie"
    case 254:
        Service = "ftp-ss"
    case 255:
        Service = "chat"
    default :
        Service = "unknown"
    }
}
