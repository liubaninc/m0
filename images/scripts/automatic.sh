#!/bin/sh

set -ex
uid=a$(date +%s)
echo ${AUTOMATIC_MNEMONIC-key erupt service six thing spy noise heart giggle year oil fuel rival drop goat deal moral require knee pact bind brain word nuclear} | m0d keys add $uid --recover
m0d automatic --from $uid