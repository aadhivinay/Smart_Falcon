#!/bin/bash

peer chaincode deploy -n account_manager -v 1.0 -c '{"Args":["init"]}' -o orderer1:7050 --tls --cafile $CORE_PEER_TLS_ROOTCERT_FILE