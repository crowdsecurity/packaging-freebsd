# freebsd-crowdsec

This repository contains the development version of the FreeBSD ports for crowdsec and bouncers.

It can be used as an overlay for the upstream ports repository in a poudriere environment:


```
# poudriere jail -c -j 132 -v 13.2-RELEASE -a amd64
[...]
# poudriere ports -c -m git+https -U https://git.freebsd.org/ports.git -B main -p freebsd
[...]
# poudriere ports -c -m git+https -U https://github.com/crowdsecurity/packaging-freebsd.git -B main -p freebsd_crowdsec
[...]
# cat <<EOT >/root/port_freebsd
security/crowdsec
security/crowdsec-firewall-bouncer
security/crowdsec-blocklist-mirror
EOT
# poudriere bulk -J 4 -j 132 -p freebsd -O freebsd_crowdsec -f /root/port_freebsd
[...]
```

For the overlay feature you may need the poudriere-devel version (3.3.99-20220831 at the time of writing).

