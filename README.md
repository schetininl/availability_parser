# Aviailability parser

Makes a GET request to the specified page and searches for the specified text there, if it finds it, it sends a link via the Telegram bot to the specified chatID.
The pause between requests is random according to the minimum and maximum values in minutes.

## Env parameters

- `BOT_TOKEN` Required
- `TG_CHAT_ID` Required
- `TARGET_URL` Required
- `TARGET_TEXT` Required
- `MIN_DURATION` Default:5
- `MIN_DURATION` Default:15

## Build

`make build`
