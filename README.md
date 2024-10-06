# Sign-Task

<p align="center">
<img src="https://img.shields.io/github/actions/workflow/status/starudream/sign-task/golang.yml?style=for-the-badge&logo=github&label=golang" alt="golang">
<img src="https://img.shields.io/github/v/release/starudream/sign-task?style=for-the-badge" alt="release">
<img src="https://img.shields.io/github/license/starudream/sign-task?style=for-the-badge" alt="license">
<br><br>
<img src="https://socialify.git.ci/starudream/sign-task/image?font=Inter&forks=1&issues=1&language=1&name=1&owner=1&pattern=Circuit%20Board&pulls=1&stargazers=1&theme=Auto" alt="project">
</p>

## Config

<details>

<summary>通用配置</summary>

```yaml
# 日志
#  https://pkg.go.dev/github.com/starudream/go-lib/core/v2/config/global#Config
log:
  console:
    format: text
    level: INFO
  file:
    enabled: true
    format: text
    level: DEBUG
    filename: ""
    max_size: 100
    max_backups: 10
    daily_rotate: true
# 通知
#  https://pkg.go.dev/github.com/starudream/go-lib/ntfy/v2#Config
ntfy:
```

</details>

<details>

<summary>完整配置</summary>

```yaml
# 打码
geetest:
  cron:
    spec: 0 10 0 * * *
    startup: true
  rr:
    key: "rrocr.com"
  tt:
    key: "ttocr.com"
# 斗鱼
douyu:
  cron:
    spec: 0 1 0 * * *
    jitter: 3
  accounts:
    - phone: ""
      did: ""
      ltp0: ""
      room: ""
      assigns:
        - count: 1
        - room: 9999
          all: true
      ignore_expired_check: false
# 库街区
kuro:
  accounts:
    - phone: ""
      dev_code: ""
      token: ""
# 米游社
miyoushe:
  cron:
    spec: 0 5 0 * * *
    jitter: 3
  accounts:
    - phone: ""
      device:
        id: ""
        type: ""
        name: ""
        model: ""
        version: ""
        channel: ""
      mid: ""
      stoken: ""
      uid: ""
      ctoken: ""
      sign_game_ids:
        - "6"
# 森空岛
skland:
  cron:
    spec: 0 3 0 * * *
    jitter: 3
  accounts:
    - phone: ""
      cred: ""
      token: ""
# 百度贴吧
tieba:
  cron:
    spec: 0 2 0 * * *
    jitter: 3
  accounts:
    - phone: ""
      bduss: ""
# 阿里云
aliyun:
  cron:
    spec: 10 0 0 * * *
  accounts:
    - id: ""
      secret: ""
```

</details>

## Usage

### Docker

```shell
mkdir sign && touch sign/app.yaml
docker run -it --rm -v $(pwd)/sign:/sign -e DEBUG=true starudream/sign-task /sign-task -c /sign/app.yaml --help
```

### Docker Compose

```yaml
version: "3"

services:
  sign:
    image: starudream/sign-task
    container_name: sign
    restart: always
    command: /sign-task -c /sign/app.yaml cron
    volumes:
      - "./sign/:/sign"
    environment:
      DEBUG: "true"
      app.log.console.level: "info"
      app.log.file.enabled: "true"
      app.log.file.level: "debug"
      app.log.file.filename: "/sign/app.log"
```

## [License](./LICENSE)
