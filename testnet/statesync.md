-Stop Neutron background service and reset database:
sudo systemctl stop neutrond
neutrond tendermint unsafe-reset-all --home $HOME/.neutrond
-Use variables to connect to RPC node and set the number and the hash of the block for state sync:
PEER="7d1376c820cab7fecee27a5902405be4559ffc84@65.108.241.166:26656"
SNAP="http://65.108.241.166:26657"
LATEST_HEIGHT=$(curl -s $SNAP/block | jq -r .result.block.header.height)
TRUST_HEIGHT=$((LATEST_HEIGHT - 1000))
TRUST_HASH=$(curl -s "$SNAP/block?height=$TRUST_HEIGHT" | jq -r .result.block_id.hash)

sed -i.bak -E "s|^(enable[[:space:]]+=[[:space:]]+).*$|\1true| ; \
s|^(rpc_servers[[:space:]]+=[[:space:]]+).*$|\1\"$SNAP,$SNAP\"| ; \
s|^(trust_height[[:space:]]+=[[:space:]]+).*$|\1$TRUST_HEIGHT| ; \
s|^(trust_hash[[:space:]]+=[[:space:]]+).*$|\1\"$TRUST_HASH\"|" $HOME/.neutrond/config/config.toml
sed -i.bak -e "s/^persistent_peers *=.*/persistent_peers = \"$PEER\"/" $HOME/.neutrond/config/config.toml
-Run the node 
sudo systemctl restart neutrond.service
-logs :
sudo journalctl -u neutrond.service -f -o cat
