#!/bin/bash

peer channel create -c mychannel -f channel.tx --outputBlock new_block.pb -o orderer1:7050 --tls --cafile $CORE_PEER_TLS_ROOTCERT_FILE