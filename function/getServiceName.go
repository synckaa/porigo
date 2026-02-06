package function

var commonPorts = map[int]string{
	21:    "FTP",
	22:    "SSH",
	23:    "Telnet",
	25:    "SMTP",
	53:    "DNS",
	80:    "HTTP",
	110:   "POP3",
	143:   "IMAP",
	443:   "HTTPS",
	445:   "SMB",
	1433:  "MSSQL",
	3306:  "MySQL",
	3389:  "RDP",
	5432:  "PostgreSQL",
	6379:  "Redis",
	8080:  "HTTP-Proxy",
	27017: "MongoDB",
}

func getServiceName(port int) string {
	if name, ok := commonPorts[port]; ok {
		return name
	}
	return "unknown"
}
