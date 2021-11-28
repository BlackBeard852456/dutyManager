#! /bin/bash

echo "Installation du programme" 
mkdir $HOME/.dutyManager
sudo go build -o dutyManager main.go
sudo mv ./dutyManager $HOME/.dutyManager/
sudo mv db.db $HOME/.dutyManager/
sudo echo "export PATH=$PATH:$HOME/.dutyManager/" > $HOME/.profile
sudo echo $HOME/.profile
