#!/usr/bin/env bash
####################

set -e

####################

readonly RELDIR="$(dirname ${0})"
readonly OUTDIR="${PWD}/${RELDIR}/../out"

####################

eprintln(){
	! [ -z "${1}" ] || eprintln 'eprintln err: undefined message'
	printf "${1}\n" 1>&2
	return 1
}

dist() {
	! [ -z "${1}" ] && ! [ -z "${2}" ] && ! [ -z "${3}" ] && ! [ -z "${4}" ] \
	|| eprintln 'expected: <maindir> <target_name> <goos> <goarch>'
	! [ -e "${OUTDIR}/${2}-${3}-${4}" ] || rm -r "${OUTDIR}/${2}-${3}-${4}"
	local startdir="$(pwd)"
	mkdir -p ${OUTDIR}/${2}-${3}-${4}
	GOOS=${3} GOARCH=${4} go build \
		-ldflags '-s -w' \
		-o "${OUTDIR}/${2}-${3}-${4}/${2}" \
		${1}
	cp -a \
		${RELDIR}/../.env.example \
		${RELDIR}/control.sh \
		${OUTDIR}/${2}-${3}-${4}/
	cd ${OUTDIR}
	tar -czf ${2}-${3}-${4}.tar.gz ${2}-${3}-${4}
	cd "${startdir}"
}

dist_all() {
	dist ${RELDIR}/../cmd/tmt/main.go tmt linux amd64
	dist ${RELDIR}/../cmd/tmt/main.go tmt darwin arm64
}

####################

dist_all


