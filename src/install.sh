#! /bin/bash

echo "Installation du programme" 
mkdir $HOME/.dutyManager
go build -o dutyManager main.go
mv ./dutyManager $HOME/.dutyManager/
mv db.db $HOME/.dutyManager/
sudo echo "export PATH=$PATH:$HOME/.dutyManager/" > $HOME/.profile
echo $HOME/.profile
