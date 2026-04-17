# Can I paddle the Tagliamento?

Finally, a web app that answers the age old question:
Is there enough water in the Tagliamento river so I can paddle it with my kayak?

Recommend water measure (0.4m) researched by the awesome [Steve der Flusswanderer](https://www.flusswandern.at/tagliamento/).

## Deploy with Fly.io

1. [Install the Fly CLI](https://fly.io/docs/hands-on/install-flyctl/) and log in:
   ```bash
   brew install flyctl
   fly auth login
   ```

2. Update the `app` name in `fly.toml` to something unique, then launch:
   ```bash
   fly launch --no-deploy
   fly deploy
   ```

3. Your app is live at `https://<app-name>.fly.dev`.

The included `fly.toml` runs a single shared-CPU machine with 256 MB RAM, which is sufficient for this app and fits within Fly.io's free tier.
