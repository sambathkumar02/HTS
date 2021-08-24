#!/bin/bash
#Binary Hts file
sudo cp hts /bin/hts

#Directory for config files
sudo mkdir /etc/HTS
sudo cp -r Static /etc/HTS
sudo touch /etc/HTS/hts.log
sudo cp config.json /etc/HTS
echo "HTS Installed sucessfully"
echo "Try hts for Help!"


