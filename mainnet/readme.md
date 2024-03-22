
```sudo apt update && sudo apt upgrade -y
sudo apt install curl tar wget clang pkg-config libssl-dev jq build-essential bsdmainutils git make ncdu gcc git jq chrony liblz4-tool -y

GO 
sudo rm -rf /usr/local/go
curl -Ls https://go.dev/dl/go1.21.8.linux-amd64.tar.gz | sudo tar -xzf - -C /usr/local
eval $(echo 'export PATH=$PATH:/usr/local/go/bin' | sudo tee /etc/profile.d/golang.sh)
eval $(echo 'export PATH=$PATH:$HOME/go/bin' | tee -a $HOME/.profile)

Build Binaries

git clone https://github.com/akiroinu/akiro.git
cd akiro
git checkout v0.2
make install

akirod init VALIDATOR --chain-id akiro-1
akirod config chain-id akiro-1


Download Genesis
wget -O $HOME/.akiro/config/genesis.json https://raw.githubusercontent.com/akiroinu/akiro/main/mainnet/genesis.json

PEER
ae1a5b5a94888d18a08fc64a9343b4450cc5cedc@213.199.36.82:26656

Create a service file

sudo tee /etc/systemd/system/akirod.service > /dev/null <<EOF
[Unit]
Description=akirod
After=network-online.target

[Service]
User=$USER
ExecStart=$(which akirod) start
Restart=on-failure
RestartSec=3
LimitNOFILE=65535

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl daemon-reload
sudo systemctl enable akirod
sudo systemctl restart akirod && sudo journalctl -u akirod -f -o cat```



