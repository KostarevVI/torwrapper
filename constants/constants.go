package constants

const (
	BACKUP_TORRC_CMD = `sudo cp -p /etc/tor/torrc /etc/tor/torrc.torwrapper.bak`

	TORRC_CONFIG = `
VirtualAddrNetwork 10.0.0.0/10
AutomapHostsOnResolve 1
TransPort 9040
DNSPort 53
ControlPort 9051
`

	BACKUP_RESOLV_CONV_CMD = `sudo cp -p /etc/resolv.conf /etc/resolv.conf.torwrapper.bak`

	RESOLV_CONV_CONFIG = `
nameserver 127.0.0.1
`

	BACKUP_IPTABLES_RULES_CMD = `
iptables-save > /etc/iptables/rules.v4
ip6tables-save > /etc/iptables/rules.v6
`
	CLEAR_IPTABLES_RULES = `
iptables -P INPUT ACCEPT
iptables -P FORWARD ACCEPT
iptables -P OUTPUT ACCEPT
iptables -t nat -F
iptables -t mangle -F
iptables -F
iptables -X
ip6tables -F
`

	APPLY_TORWRAPPER_IPTABLES_RULES = `
NON_TOR="192.168.1.0/24 192.168.0.0/24"
TOR_UID=%s
DNS_PORT="53"
TRANS_PORT="9040"

iptables -t nat -A OUTPUT -m owner --uid-owner $TOR_UID -j RETURN
iptables -t nat -A OUTPUT -p udp --dport 53 -j REDIRECT --to-ports $DNS_PORT
for NET in $NON_TOR 127.0.0.0/9 127.128.0.0/10; do
 iptables -t nat -A OUTPUT -d $NET -j RETURN
done
iptables -t nat -A OUTPUT -p tcp --syn -j REDIRECT --to-ports $TRANS_PORT

iptables -A OUTPUT -m state --state ESTABLISHED,RELATED -j ACCEPT
for NET in $NON_TOR 127.0.0.0/8; do
 iptables -A OUTPUT -d $NET -j ACCEPT
done
iptables -A OUTPUT -m owner --uid-owner $TOR_UID -j ACCEPT

iptables -A OUTPUT -j REJECT
ip6tables -A OUTPUT -j REJECT
`

	RESTORE_IPTABLE_RULES_CMD = `
iptables-restore < /etc/iptables/rules.v4
ip6tables-restore < /etc/iptables/rules.v6
`

	RESTORE_RESOLV_CONV_CMD = `sudo cp -p /etc/resolv.conf.torwrapper.bak /etc/resolv.conf`

	RESTORE_TORRC_CMD = `sudo cp -p /etc/tor/torrc.torwrapper.bak /etc/tor/torrc`

	DOWNLOAD_BRIDGES_CMD = `
sudo wget https://github.com/torwrapper/bridges.txt/raw/master/bridges.txt -O /etc/tor/bridges.txt
`

	CHECK_TOR_CONNECTION_CMD = `wget -qO- https://check.torproject.org | grep -Po "(?<=strong>)[\\d\\.]+(?=</strong)"`
)
