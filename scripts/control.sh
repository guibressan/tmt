#!/usr/bin/env bash
####################

set -e

####################

readonly RELDIR="$(dirname ${0})"
readonly SVC_NAME="tmt"

####################

eprintln(){
	! [ -z "${1}" ] || eprintln 'eprintln err: undefined message'
	printf "${1}\n" 1>&2
	return 1
}


mk_systemd() {
	! [ -e "/etc/systemd/system/${SVC_NAME}.service" ] \
		|| eprintln "service ${SVC_NAME} already exists"
	local user="${USER}"
	sudo bash -c "cat << EOF > /etc/systemd/system/${SVC_NAME}.service
[Unit]
Description=${SVC_NAME} unit
After=network.target

[Service]
Environment=\"PATH=/usr/local/bin:/usr/bin:/bin:${PATH}\"
User=${user}
Type=simple
ExecStart=/bin/bash -c \"cd ${PWD}/${RELDIR}; ./${SVC_NAME}\"
Restart=no
#RestartSec=10s

[Install]
WantedBy=multi-user.target
EOF
"
	sudo systemctl enable "${SVC_NAME}".service
	printf "To start the service, run: sudo systemctl start "${SVC_NAME}".service\n"
}

rm_systemd() {
	[ -e "/etc/systemd/system/${SVC_NAME}.service" ] || return 0
	sudo systemctl stop "${SVC_NAME}".service || true
	sudo systemctl disable "${SVC_NAME}".service
	sudo rm /etc/systemd/system/"${SVC_NAME}".service
}

####################

case ${1} in
	mk-systemd) mk_systemd ;;
	rm-systemd) rm_systemd ;;
	*) eprintln 'usage: <mk-systemd|rm-systemd>' ;;
esac
