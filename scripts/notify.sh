#!/bin/bash

cat - |
    awk -F ',' '{printf("dunstify \"Azure\" '"'"'<a href=\"%s\">%s</a>'"'"'\n"), $3, $2}' |
    bash
