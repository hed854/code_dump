#!/bin/bash
# usage: ./cachescan.sh domains.txt

declare -a WEBSITES
declare -a HEADERS
readarray WEBSITES < $1
HEADERS=(date: cache-control: etag: last-modified: expires: pragma:)

for i in "${WEBSITES[@]}"
do

	echo "➡️ ${i}"
	tmpfile="/tmp/scan_"$(echo "${i}" | md5sum | cut -d ' ' -f1)
	http -h ${i} --follow > ${tmpfile}
	for j in ${HEADERS[@]}
	do
		grep -i ${j} ${tmpfile}
	done
	rm -rf ${tmpfile}
	echo ""
done

