#!/bin/bash

for i in {1..100};
do 
echo ''
curl -s --data '{"method":"eth_getBlockByNumber","params":["latest",true],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST https://volta-rpc.energyweb.org | jq .result.extraData | xxd -p -r ; sleep 3; 
done

