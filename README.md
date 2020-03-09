# Gontainer

## Download netsetgo to create bridge and veth pair
```bash
wget https://github.com/teddyking/netsetgo/releases/download/0.0.1/netsetgo
sudo mv netsetgo /usr/local/bin/
sudo chown root:root /usr/local/bin/netsetgo
sudo chmod 4755 /usr/local/bin/netsetgo
```

## Establish internet connectivity
```bash
sudo iptables -tnat -N netsetgo
sudo iptables -tnat -A PREROUTING -m addrtype --dst-type LOCAL -j netsetgo
sudo iptables -tnat -A OUTPUT ! -d 127.0.0.0/8 -m addrtype --dst-type LOCAL -j netsetgo
sudo iptables -tnat -A POSTROUTING -s 10.10.10.0/24 ! -o brg0 -j MASQUERADE
sudo iptables -tnat -A netsetgo -i brg0 -j RETURN
```