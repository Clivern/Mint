### Manual Installation

1. Update and upgrade your server packages.

```zsh
$ apt-get update
$ apt-get upgrade
```

2. You might need to reboot if the kernel got upgraded

```zsh
$ reboot
```

3. Install a MySQL server. Then create an application specific user, password and database.

```zsh
$ sudo apt install mysql-server
$ sudo mysql_secure_installation
```

4. Create the required directories and the config file.

```zsh
$ mkdir -p /srv/app/releases/1.0.0
$ cd /srv/app/releases/1.0.0
```

5. Download the [latest binary from gitlab releases page](https://gitlab.com/clivern/mint/-/releases). Move it to `/srv/app/releases/1.0.0` directory. then create the `/srv/app/releases/1.0.0/config.prod.yml` file like the following:

You will have to adjust the database `host`, `name`, `username` and `password`. Also the `api.key` and `app.port` but not required.

```yaml
# App configs
app:
  # App name
  name: ${APP_NAME:-mint}
  # Env mode (dev or prod)
  mode: ${APP_MODE:-prod}
  # HTTP port
  port: ${APP_PORT:-80}
  # Hostname
  hostname: ${APP_HOSTNAME:-127.0.0.1}
  # TLS configs
  tls:
    status: ${APP_TLS_STATUS:-off}
    crt_path: ${APP_TLS_PEMPATH:-cert/server.crt}
    key_path: ${APP_TLS_KEYPATH:-cert/server.key}

  # API Configs
  api:
      key: ${APP_API_KEY:-6c68b836-6f8e-465e-b59f-89c1db53afca}

  # Global timeout
  timeout: ${APP_TIMEOUT:-50}

  # Application Database
  database:
      # Database driver (sqlite3, mysql)
      driver: ${APP_DATABASE_DRIVER:-mysql}
      # Database Host
      host: ${APP_DATABASE_MYSQL_HOST:-localhost}
      # Database Port
      port: ${APP_DATABASE_MYSQL_PORT:-3306}
      # Database Name
      name: ${APP_DATABASE_MYSQL_DATABASE:-mint}
      # Database Username
      username: ${APP_DATABASE_MYSQL_USERNAME:-root}
      # Database Password
      password: ${APP_DATABASE_MYSQL_PASSWORD:-root}
      # Max DB Open Connections
      max_open_conns: ${APP_DATABASE_MAX_OPEN_CONNS:-5}
      # DB Connection Max Lifetime
      conns_max_lifetime: ${APP_DATABASE_CONNS_MAX_LIFETIME:-10}

  # Log configs
  log:
    # Log level, it can be debug, info, warn, error, panic, fatal
    level: ${APP_LOG_LEVEL:-info}
    # Output can be stdout or abs path to log file /var/logs/mint.log
    output: ${APP_LOG_OUTPUT:-/var/logs/mint.log}
    # Format can be json
    format: ${APP_LOG_FORMAT:-json}
```

6. Create a `systemd` service to run `mint`.

```zsh
$ echo "
[Unit]
Description=Mint Web Service
After=network.target

StartLimitIntervalSec=500
StartLimitBurst=5

[Service]
User=root
Group=www-data
WorkingDirectory=/srv/app/releases/1.0.0
Restart=on-failure
RestartSec=5s
ExecStart=/srv/app/releases/1.0.0/mint server -c /srv/app/releases/1.0.0/config.prod.yml
ExecReload=/bin/kill -s HUP $MAINPID
ExecStop=/bin/kill -s TERM $MAINPID
PrivateTmp=true

[Install]
WantedBy=multi-user.target" /etc/systemd/system/mint.service
```

7. Enable and start the service.

```zsh
$ systemctl daemon-reload

$ systemctl enable mint.service
$ systemctl start mint.service
```
