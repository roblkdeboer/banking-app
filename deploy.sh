#!/bin/bash
sshcmd="ssh -t roblkdeboer@banking.roblkdeboer.com"
$sshcmd screen -S "deployment" /home/roblkdeboer/banking-app/prod_deploy.sh