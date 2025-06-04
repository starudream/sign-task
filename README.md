# Sign-Task

<p align="center">
<img src="https://img.shields.io/github/actions/workflow/status/starudream/sign-task/golang.yml?style=for-the-badge&logo=github&label=golang" alt="golang">
<img src="https://img.shields.io/github/v/release/starudream/sign-task?style=for-the-badge" alt="release">
<img src="https://img.shields.io/github/license/starudream/sign-task?style=for-the-badge" alt="license">
<br><br>
<img src="https://socialify.git.ci/starudream/sign-task/image?font=Inter&forks=1&issues=1&language=1&name=1&owner=1&pattern=Circuit%20Board&pulls=1&stargazers=1&theme=Auto" alt="project">
</p>

## Feature

- [x] 斗鱼
  - [x] 领取荧光棒
  - [x] 赠送荧光棒续粉丝牌
- [x] 库街区
  - [x] 游戏签到（鸣潮）
  - [x] 论坛签到
- [x] 米游社
  - [x] 游戏签到（原神 崩坏 星铁 绝区零 等）
  - [x] 论坛签到
- [x] 森空岛
  - [x] 游戏签到（明日方舟）
- [x] 百度贴吧
  - [x] 客户端签到
- [x] 阿里云
  - [x] 余额提醒
  - [x] 历史账单详细
- [x] 火山引擎
  - [x] 余额提醒
  - [x] 历史账单详细
- [x] 腾讯云
  - [x] 余额提醒

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
#
# https://github.com/starudream/sign-task
#
# 打码
geetest:
  cron:
    disable: true
    spec: 0 0 12 * * *
    startup: false
    jitter: 10
# 斗鱼
douyu:
  cron:
    disable: true
    spec: 0 0 12 * * *
    startup: false
    jitter: 10
  accounts:
    - phone: douyu_phone
      did: douyu_did
      ltp0: douyu_ltp0
      room: 9999
      assigns:
        - count: 1
        - room: 9999
          all: true
      ignore_expired_check: false
# 库街区
kuro:
  cron:
    disable: true
    spec: 0 0 12 * * *
    startup: false
    jitter: 10
  accounts:
    - phone: kuro_phone
      dev_code: kuro_dev_code
      token: kuro_token
# 米游社
miyoushe:
  cron:
    disable: true
    spec: 0 0 12 * * *
    startup: false
    jitter: 10
  accounts:
    - phone: miyoushe_phone
      device:
        id: device_id
        type: device_type
        name: device_name
        model: device_model
        version: device_version
        channel: device_channel
      mid: miyoushe_mid
      stoken: miyoushe_stoken
      uid: miyoushe_uid
      ctoken: miyoushe_ctoken
      sign_game_ids:
        - "6"
# 森空岛
skland:
  cron:
    disable: true
    spec: 0 0 12 * * *
    startup: false
    jitter: 10
  accounts:
    - phone: skland_phone
      cred: skland_cred
      token: skland_token
# 百度贴吧
tieba:
  cron:
    disable: true
    spec: 0 0 12 * * *
    startup: false
    jitter: 10
  accounts:
    - phone: tieba_phone
      bduss: tieba_bduss
# 阿里云
aliyun:
  cron:
    disable: true
    spec: 0 0 12 * * *
    startup: false
    jitter: 10
  accounts:
    - id: aliyun_id
      secret: aliyun_secret
# 火山引擎
volcengine:
  cron:
    disable: true
    spec: 0 0 12 * * *
    startup: false
    jitter: 10
  accounts:
    - id: volcengine_id
      secret: volcengine_secret
# 腾讯云
qcloud:
  cron:
    disable: true
    spec: 0 0 12 * * *
    startup: false
    jitter: 10
  accounts:
    - id: qcloud_id
      key: qcloud_key
```

</details>

<details>

<summary>阿里云权限策略</summary>

```json
{
  "Version": "1",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "bss:DescribeAcccount",
        "bss:QueryAccountBill"
      ],
      "Resource": "*"
    }
  ]
}
```

</details>

<details>

<summary>火山引擎权限策略</summary>

```json
{
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "billing:QueryBalanceAcct",
        "billing:ListBillDetail"
      ],
      "Resource": [
        "*"
      ]
    }
  ]
}
```

</details>

<details>

<summary>腾讯云权限策略</summary>

```json
{
  "statement": [
    {
      "action": [
        "finance:DescribeAccountBalance"
      ],
      "effect": "allow",
      "resource": [
        "*"
      ]
    }
  ],
  "version": "2.0"
}
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
