# GCP Client

It is code for GCP Cloud Functions.

## How to Deploy

```
git clone https://github.com/maito1201/hatebucord.git
cd ./hatebucord/gcp
gcloud functions deploy `your_function_name` --runtime go116 --trigger-http --entry-point=PostHatebuCommand --set-env-vars DISCORD_WEBHOOK=`your_webhook_url`
```