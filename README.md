# Elastic -> Discord Webhook Proxy

[![Go Report Card](https://goreportcard.com/badge/github.com/captainGeech42/elastic-discord-webhook-proxy)](https://goreportcard.com/report/github.com/captainGeech42/elastic-discord-webhook-proxy) [![Build](https://github.com/captainGeech42/elastic-discord-webhook-proxy/workflows/Build/badge.svg)](https://github.com/captainGeech42/elastic-discord-webhook-proxy/actions?query=workflow%3ABuild) [![Docker Hub Publish](https://github.com/captainGeech42/elastic-discord-webhook-proxy/workflows/Docker%20Hub%20Publish/badge.svg)](https://github.com/captainGeech42/elastic-discord-webhook-proxy/actions?query=workflow%3A%22Docker+Hub+Publish%22) [![Docker Hub Image](https://img.shields.io/docker/v/captaingeech/elastic-discord-webhook-proxy?color=blue)](https://hub.docker.com/repository/docker/captaingeech/elastic-discord-webhook-proxy/general)

_(see also [my webhook proxy for Terraform](https://github.com/captainGeech42/tf-discord-webhook-proxy))_

Have you ever wanted to use the webhook alert connector in Kibana to notify a Discord channel of alerts in your environment, only to realize that they can't natively talk to each other? Well not anymore, because here's the tool you've been searching for!

_Elastic webhook comes in, Discord webhook goes out, profit._

By default, an embedded rich message will be sent to Discord but this can be disabled in the config file.

Please note that at this time, only uptime monitor alerts are supported.

Sample embed message:

![Rich Message](https://i.imgur.com/m0oVJBb.png)

The color of the message (on the left side) will be either red if the alert is triggered, or green if not.

## Usage

1. Download: `go get github.com/captainGeech42/elastic-discord-webhook-proxy`
2. Copy the `config.ex.json` file to your current directory as `config.json`
3. Update the `WebhookURL` field with a Discord webhook URL ([Discord docs on webhooks](https://support.discord.com/hc/en-us/articles/228383668))
4. Run it: `./elastic-discord-webhook-proxy`

The proxy will be available at `http://host:8080/webhook`. Create a new webhook alert connector in Kibana pointing to that URL, and use the following as the body of the alert webhook action (for an uptime monitor):

```json
{
    "alertId": "{{alertId}}",
    "alertName": "{{alertName}}",
    "spaceId": "{{spaceId}}",
    "tags": "{{tags}}",
    "alertInstanceId": "{{alertInstanceId}}",
    "context": {
        "message": "{{context.message}}",
        "downMonitorsWithGeo": "{{context.downMonitorsWithGeo}}"
    },
    "state": {
        "firstCheckedAt": "{{state.firstCheckedAt}}",
        "firstTriggeredAt": "{{state.firstTriggeredAt}}",
        "currentTriggerStarted": "{{state.currentTriggerStarted}}",
        "isTriggered": "{{state.isTriggered}}",
        "lastCheckedAt": "{{state.lastCheckedAt}}",
        "lastResolvedAt": "{{state.lastResolvedAt}}",
        "lastTriggeredAt": "{{state.lastTriggeredAt}}"
    }
}
```

## Docker Image

This tool is also available via a Docker image on Docker Hub ([`captaingeech/elastic-discord-webhook-proxy`](https://hub.docker.com/repository/docker/captaingeech/elastic-discord-webhook-proxy)). When running via the Docker image, you can either use this image as a base image to `COPY` your `config.json` into `/app`, or set the following environment variables instead:

* `ELASTIC_PROXY_ENV=YES` (without this, a `config.json` will be looked for)
* `ELASTIC_PROXY_WEBHOOK_URL="https://discordapp.com/api/webhooks/xxxxxxxx/yyyyyyyyyyyyy"`
* `ELASTIC_PROXY_RICH_MESSAGES=YES` (optional, disable rich messages by setting to `NO`)

The Docker image will always have the proxy running on port 8080 in the container, you can choose to forward this outside of the container to whatever port you need.

Example execution of container:

```
docker run --rm -it -p8080:8080 \
           -e ELASTIC_PROXY_ENV=YES \
           -e ELASTIC_PROXY_WEBHOOK_URL="https://discordapp.com/api/webhooks/xxxxxxxx/yyyyyyyyyyyyy" \
           captaingeech/elastic-discord-webhook-proxy:latest
```
