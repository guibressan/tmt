#!/usr/bin/env bash
####################

set -e

####################

readonly RELDIR="$(dirname ${0})"
readonly INTERNAL="${PWD}/${RELDIR}/../internal"

####################

eprintln() {
	! [ -z "${1}" ] || eprintln "expected: msg"
	printf "${1}" 1>&2
	return 1
}

fetch_godotenv() {
	local dest="${INTERNAL}/godotenv"
	local tmpdest="/tmp/tmp_godotenv"
	[ -e "${dest}" ] && return 0 || true
	! [ -e "${tmpdest}" ] || rm -rf "${tmpdest}"
	git clone https://github.com/joho/godotenv "${tmpdest}"
	cd "${tmpdest}"
	git checkout '3a7a19020151b45a29896c9142723efe5b11a061'
	rm -rf .git .github cmd go.mod autoload
	mv "${tmpdest}" "${dest}"
}

fetch_bstd() {
	local dest="${INTERNAL}/bstd"
	local tmpdest="/tmp/tmp_bstd"
	[ -e "${dest}" ] && return 0 || true
	! [ -e "${tmpdest}" ] || rm -rf "${tmpdest}"
	git clone git@master:repos/private/bstd.git "${tmpdest}"
	cd "${tmpdest}"
	git checkout 'v0.5.0'
	mkdir "${dest}"
	find . -type f -name '*test.go' -exec rm {} \;
	cp -a util cli stackerr "${dest}/"
	rm -rf ${tmpdest}
}

fetch_schema() {
	local dest="${INTERNAL}/schema"
	local tmpdest="/tmp/tmp_schema"
	[ -e "${dest}" ] && return 0 || true
	! [ -e "${tmpdest}" ] || rm -rf "${tmpdest}"
	git clone https://github.com/gorilla/schema "${tmpdest}"
	cd "${tmpdest}"
	git checkout 'cd59f2f12cbdfa9c06aa63e425d1fe4a806967ff'
	mkdir "${dest}"
	find . -type f -name '*test.go' -exec rm {} \;
	cp -a *.go LICENSE "${dest}/"
	rm -rf ${tmpdest}
}

fetch_captcha() {
	local dest="${INTERNAL}/captcha"
	local tmpdest="/tmp/tmp_captcha"
	[ -e "${dest}" ] && return 0 || true
	! [ -e "${tmpdest}" ] || rm -rf "${tmpdest}"
	git clone https://github.com/dchest/captcha "${tmpdest}"
	cd "${tmpdest}"
	git checkout '6f8a3fce31a3a5f25edd18555c6f38eff73275f6'
	mkdir "${dest}"
	find . -type f -name '*test.go' -exec rm {} \;
	cp -a *.go LICENSE "${dest}/"
	rm -rf ${tmpdest}
}

fetch() {
	fetch_godotenv
	fetch_bstd
	fetch_schema
	fetch_captcha
}

####################

case ${1} in
	fetch) fetch ;;
esac

