curl --tlsv1.2 --silent --insecure --connect-timeout 10 https://${rmt_server_name}/rmt.crt --output /etc/pki/trust/anchors/rmt-server.pem && /usr/sbin/update-ca-certificates &> /dev/null
SUSEConnect --url https://${rmt_server_name}
SUSEConnect -p sle-module-containers/15.1/x86_64
SUSEConnect -p caasp/5/x86_64
