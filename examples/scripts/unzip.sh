#!/bin/bash

# Unzip the file
unzip ./.protobuffed/armortal/main.zip -d ./.protobuffed/armortal
mv ./.protobuffed/armortal/protobuffed-main/* ./.protobuffed/armortal
rm -rf ./.protobuffed/armortal/protobuffed-main ./.protobuffed/armortal/main.zip