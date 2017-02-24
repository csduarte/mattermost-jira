# JIRA Mattermost Webhook Connector

## For Mattermost administrator
You need incoming webhook URL in `http://<mattermost_server>/hooks/<web_hook_id>` format. Can be copied from Mattermost config.

**Mattermost Config**
1. Go to `...` menu
2. `System console`
3. `Integrations` → `Custom Integrations`
4. Enable Settings
    - `Enable Incoming Webhooks`: `true`
    - `Enable integrations to override usernames`: `true`
    - `Enable integrations to override profile picture icons`: `true`

**Create Incoming Webhook**  
1. Go to `...` menu
2. `Integrations` → `Incoming Webhooks`
3. `Add Incoming Webhook`
 
## For JIRA administrator
1. JIRA Administration → System
2. ADVANCED → WebHooks
3. Create a WebHook:
    - URL:  https://**yourserver**?mattermost_hook_url=_**mattermost_hook_url**&channel=**channel_name**
    - Issue:
        - created: true
        - updated: true
        - deleted: true

## Build Binary
`go build`

## Run
`./mattermost-jira` map flag optional if you want room mapping
`./mattermost-jira 2&1 >> data.log &`  pipe stderr and stdout to file and disown process

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
        