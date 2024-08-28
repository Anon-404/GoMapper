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
    default:
        Service = "Unknown"
    }
}
