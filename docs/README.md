# Setup local development

In your browser, navigate to <https://smee.io/> then setup new proxy

Install CLI tool

```sh
npm install --global smee-client
```

Then expose it to local

```sh
export WEBHOOK_PROXY_URL=https://smee.io/<id>
smee --url $WEBHOOK_PROXY_URL --path /webhook --port 3000
```

