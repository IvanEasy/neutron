# RPC node & State sync endpoint for Neutron Quark Testnet by JOHNEE
# RPC node

RPC endpoint with "default" pruning: [http://38.242.235.176:26657/ ](http://65.108.241.166:26657/)

# State sync guide


Stop and reset database:
```
sudo systemctl stop neutrond.service
neutrond tendermint unsafe-reset-all --home $HOME/.neutrond
```
Use this variables to connect to RPC node and set the number and the hash of the block for state sync:
```
RPC_PEER="7d1376c820cab7fecee27a5902405be4559ffc84@65.108.241.166:26656"
SNAP_RPC="http://65.108.241.166:26657"
LATEST_HEIGHT=$(curl -s $SNAP_RPC/block | jq -r .result.block.header.height)
TRUST_HEIGHT=$((LATEST_HEIGHT - 1000))
TRUST_HASH=$(curl -s "$SNAP_RPC/block?height=$TRUST_HEIGHT" | jq -r .result.block_id.hash)

sed -i.bak -E "s|^(enable[[:space:]]+=[[:space:]]+).*$|\1true| ; \
s|^(rpc_servers[[:space:]]+=[[:space:]]+).*$|\1\"$SNAP_RPC,$SNAP_RPC\"| ; \
s|^(trust_height[[:space:]]+=[[:space:]]+).*$|\1$TRUST_HEIGHT| ; \
s|^(trust_hash[[:space:]]+=[[:space:]]+).*$|\1\"$TRUST_HASH\"|" $HOME/.neutrond/config/config.toml
sed -i.bak -e "s/^persistent_peers *=.*/persistent_peers = \"$RPC_PEER\"/" $HOME/.neutrond/config/config.toml
```
Start your node
```
sudo systemctl restart neutrond.service
```
Watch your logs 
```
sudo journalctl -u neutrond.service -f -o cat
