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
`./mattermost-jira -addr=127.0.0.1 -port=5002 -log=./test.log`
Addr defaults to `0.0.0.0`
Port defaults to `5000`
Log defaults to `./mattermost-jira.log`

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

In general, teams work best when Jira alerts are customized to specific projects, issue types and events, for example:
 
- **New bugs created in Project X** - For a developer looking for feedback prior to release 
- **All newly opened and updated security issues in any project** - For a channel monitoring system security 
- **Resolved issues in Project A** - For a quality assurance team ready to test and close a resolved issue. 
- **All deleted issues** - For a channel monitoring any inappropriately deleted tickets
