#!/bin/bash
alias ls='ls --color=auto'
alias du='du -h'
alias cls='clear;reset'

PS1='[\u@\h \W]\$ '

GOBIN=/loki/bin
GOPATH=/loki
PATH=$GOBIN:$PATH
