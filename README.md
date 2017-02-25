# Extension of [kylemanna/docker-openvpn](https://github.com/kylemanna/docker-openvpn)
# OpenVPN server with included mysql client, tools to compile/run radius plugin
The image was originally created to download config from Google Cloud Storage on start, but can download config via bash commands on start.

A shellscript to check udp server health is included. It works by sending bytes to the server (Currently only on port 50005) and checking if it gets a reply. If the server replies, a netcat tcp listener is started on 20002. If the server stops responding, the listener is killed (made for haproxy healthchecks).

### Note that udphealthcheck is experemental. It saturates the log quite dramatically as it sends packets to the server.

## Usage

* Use [kylemanna/docker-openvpn](https://github.com/kylemanna/docker-openvpn) README to setup OpenVPN configuration if needed and get started.

* [Do not forget to give container additional permissions eg. `NET_ADMIN`]

### Google Cloud Storage

When running, specify the following environmental variables:

* `GAUTH_CREATOR` - A chain of bash commands to create/download/otherwise get .json file for `GOOGLE_APPLICATION_CREDENTIALS`

        -e "GAUTH_CREATOR=echo hi > /GCS.json&&echo bye >> /GCS.json"

* `GOOGLE_APPLICATION_CREDENTIALS` - path to the GOOGLE_APPLICATION_CREDENTIALS .json file (needed for auth. to google)

        -e "GOOGLE_APPLICATION_CREDENTIALS=/GCS.json"
        
* `GOOGLE_CLOUD_PROJECT` - your google cloud project name (full id)

        -e "GOOGLE_CLOUD_PROJECT=google-cloud-test-123"
        
* `GCS_FULL_PATH` - full path to config .zip to download (BUCKET:PATH) [any *.zip is extracted to /etc/openvpn]

        -e "GCS_FULL_PATH=my-buck:conf.zip"
        
* `RUN_UDP_HEALTHCHECK` - if =1 then udphealthcheck runs 

        -e "RUN_UDP_HEALTHCHECK=0"

### Custom

* Use GAUTH_CREATOR to ether download a *.zip of OpenVPN config files into working dir or to put the config files into /etc/openvpn (Put your bash command chain there) 

* Ignore `GOOGLE_APPLICATION_CREDENTIALS`,`GOOGLE_CLOUD_PROJECT`,`GCS_FULL_PATH`

