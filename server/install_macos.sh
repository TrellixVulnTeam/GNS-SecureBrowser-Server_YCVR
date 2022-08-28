#!/bin/bash
if [ -f "/Library/LaunchDaemons/com.s-badge.gnsservice.plist" ]; then
    echo "Found existing service will try to stop and delete it"
    launchctl stop com.s-badge.gnsservice.plist
    launchctl unload /Library/LaunchDaemons/com.s-badge.gnsservice.plist
    rm /Library/LaunchDaemons/com.s-badge.gnsservice.plist
fi

echo "installing app"
mkdir -p /Library/GNS/logs
cp gnsdeviceserver /Library/GNS/
cp GNSPrivate.pem /Library/GNS/
cp com.s-badge.gnsservice.plist /Library/LaunchDaemons/
echo "launching and starting app"
launchctl load /Library/LaunchDaemons/com.s-badge.gnsservice.plist
launchctl start com.s-badge.gnsservice.plist