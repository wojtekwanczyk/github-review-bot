# github-review-bot

Send notifications to reviewers about pull requests updates.

## Env vars

Kept in git ignored `secrets` file

```bash
export WEBEX_BOT_TOKEN="bot-token"
export WEBEX_ROOM_MAPPING_PATH="room_mapping.json"
```

## Room mapping

Example of format of `room_mapping.json` file

```bash
{
    "user1": "room-id1",
    "user2": "room-id2"
}
```
