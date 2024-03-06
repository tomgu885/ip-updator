```bash
iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 10028 -j DNAT --to-destination 220.178.185.54:10028
iptables -A FORWARD -i eth0 -p tcp --dport 10028 -j ACCEPT

iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 10024 -j DNAT --to-destination 114.235.208.124:10024
iptables -A FORWARD -i eth0 -p tcp --dport 10024 -j ACCEPT

iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 10020 -j DNAT --to-destination 222.79.58.9:10020
iptables -A FORWARD -i eth0 -p tcp --dport 10020 -j ACCEPT

iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 10100 -j DNAT --to-destination 222.79.59.199:10100
iptables -A FORWARD -i eth0 -p tcp --dport 10100 -j ACCEPT

iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 10100 -j DNAT --to-destination 222.79.59.199:10100
iptables -A FORWARD -i eth0 -p tcp --dport 10100 -j ACCEPT

iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 10022 -j DNAT --to-destination 121.228.255.121:10022
iptables -A FORWARD -i eth0 -p tcp --dport 10022 -j ACCEPT

iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 10143 -j DNAT --to-destination 27.159.184.79:10143
iptables -A FORWARD -i eth0 -p tcp --dport 10143 -j ACCEPT

iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 10145 -j DNAT --to-destination 121.227.26.24:10145
iptables -A FORWARD -i eth0 -p tcp --dport 10145 -j ACCEPT
```

## delete

```bash
#iptables -t nat -v -L PREROUTING -n --line-number
#iptables -t nat -D PREROUTING {rule-number-here}
```

我想要个福建福州的
苏州