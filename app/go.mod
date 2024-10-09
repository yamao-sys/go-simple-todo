module app

require app/config v1.2.3

require gopkg.in/go-ini/ini.v1 v1.67.0 // indirect

replace app/config => ./config

go 1.23.1
