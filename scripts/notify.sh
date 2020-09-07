#!/bin/bash

cat - |
    awk -F ',' '{printf("notify-send \"%s\" \"%s\"\n"), $2, $3}' |
    bash
