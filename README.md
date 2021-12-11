# webdav-server
a local webdav server written in go
use nplayer as client-player, the default port is 7212, can modify it in config.json
in config.json, we can define the nplayer webdav prefix with key: "prefix" as follows.
and key "dir" means the file-dir in local pc.

![](https://jaroffertree.oss-cn-hongkong.aliyuncs.com/QQ%E5%9B%BE%E7%89%8720211211160942.png)

e.g. 
movies stores in "h:/movie", then we set "prefix" as "movie" and "dir" as "h:/movie", then we fill "movie" in nplayer - new server- webdav - path
it works.
