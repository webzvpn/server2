#!/bin/ash

while true

do

OUT="$(echo -e "\x38\x01\x00\x00\x00\x00\x00\x00\x00" | timeout -t 1 nc -u uk2.webzvpn.ru 50005 | cat )"

sleep 1s

if [ "${OUT}" != '' ]

then
	if ! pgrep -x "nc" > /dev/null
	then
	    nc -lk -p 20002 -e cat &
    	    NCPID=$!
	fi
else
        if pgrep -x "nc" > /dev/null
        then
	    kill -9 "${NCPID}"
        fi

fi

done
