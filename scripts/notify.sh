#!/bin/bash

cat - |
    awk -F ',' '{printf("dunstify --timeout 0 \"Azure\" '"'"'<a href=\"%s\">%s</a>'"'"'\n"), $3, $2}' |
    bash
