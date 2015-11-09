# nagflux
A connector which copies performancedata from Nagios/Icinga to InfluxDB

##Install
```
go get -u github.com/Downlord/nagflux
go build github.com/Downlord/nagflux
```

##Start
If the configfile is in the same folder as the executable:
```
./nagflux
```
else:
```
./nagflux -configPath=/path/to/config.gcfg
```
