# JIRA Mattermost Webhook Connector

Enables Mattermost channels to receive notifications from one or more Jira servers, filtered by JQL query and event type, via webhooks. 

## For Mattermost administrator
You need a **Mattermost incoming webhook URL** in `http://<mattermost_server>/hooks/<web_hook_id>` format. Can be copied from Mattermost config.

**Mattermost Config**
1. In Team Site, as System Administrator, go to main menu
2. Go to `System Console`
3. Go to `Integrations` → `Custom Integrations`
4. Enable Settings
    - `Enable Incoming Webhooks`: `true`
    - `Enable integrations to override usernames`: `true`
    - `Enable integrations to override profile picture icons`: `true`

**Create Incoming Webhook**  
1. In Team Site, as System Administrator, go to main menu
2. Go to `Integrations` → `Incoming Webhooks`
3. Go to `Add Incoming Webhook`
 
## For JIRA administrator
1. Go to `JIRA Administration` → `System`
2. Go to `ADVANCED` → `WebHooks`
3. Create a WebHook:
    - URL:  `https://<mattermost-jira-server>?mattermost_hook_url=_<mattermost_incoming_webhook_URL>&channel=<channel_name>`
    - Issue:
        - created: true
        - updated: true
        - deleted: true

## Build Binary
`go build`

## Run
`./mattermost-jira -addr=127.0.0.1 -port=5002 -log=./test.log` 

- Addr defaults to `0.0.0.0` 
- Port defaults to `5000` 
- Log defaults to `./mattermost-jira.log`

## Test
While server is running in background or different session, execute:

Simple Test:
```bash
curl -X POST -H "Content-Type: application/json" --data @sample_hook.json "localhost:5000?mattermost_hook_url=http://localhost:8065/hooks/67qhmgccxffaunr886gfewoqfo&channel=off-topic"
```

Simple Repeated Test:
```bash
while sleep 0.2 
do 
	(curl -X POST -H "Content-Type: application/json" --data @sample_hook.json "localhost:5000?mattermost_hook_url=http://localhost:8065/hooks/67qhmgccxffaunr886gfewoqfo&channel=town-square") &
done
```

## Best Practices 

Teams work best when Jira notification are customized by query and events to meet the needs of specific channels to which notifications appear. For example: 
 
- **New bugs created in Project X** - For a developer looking for feedback prior to release 
- **All newly opened and updated security issues in any project** - For a channel monitoring system security 
- **Resolved issues in Project A** - For a quality assurance team ready to test and close a resolved issue. 
- **All deleted issues** - For a channel monitoring any inappropriately deleted tickets

The Jira Administrator can configure these alerts using the Webhooks user interface available at `https://<jira_server>/plugins/servlet/webhooks`

![image](https://user-images.githubusercontent.com/177788/28011452-4914d76a-6517-11e7-9538-a156fa7eadb9.png)
