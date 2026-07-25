# Deploying to EC2 (GitHub Actions → GHCR → Docker → Nginx)

This pipeline builds a Docker image on every push to `main`, pushes it to the
GitHub Container Registry (GHCR), then SSHes into your EC2 instance to pull and
restart the container. Nginx on the host terminates HTTPS and proxies to the
container on `127.0.0.1:8080`.

## 1. Push the repo to GitHub

This folder isn't a git repo yet:

```bash
git init
git add .
git commit -m "chore: initial portfolio + deploy pipeline"
git branch -M master
git remote add origin git@github.com:<you>/<repo>.git
git push -u origin master
```

## 2. Prepare the EC2 instance (one-time)

Use Ubuntu 22.04+ (or Amazon Linux 2023). SSH in and install Docker + Nginx:

```bash
# Docker Engine + compose plugin
curl -fsSL https://get.docker.com | sh
sudo usermod -aG docker "$USER"   # log out/in afterwards

# Nginx + Certbot
sudo apt-get update
sudo apt-get install -y nginx certbot python3-certbot-nginx
```

Security group: allow inbound **80** and **443** (and **22** from your IP).
The app port **8080 is NOT exposed publicly** — it's bound to loopback and only
reachable through Nginx.

## 3. Configure Nginx + TLS (one-time)

```bash
sudo cp deploy/nginx/portfolio.conf /etc/nginx/sites-available/portfolio.conf
# edit server_name to your real domain
sudo ln -s /etc/nginx/sites-available/portfolio.conf /etc/nginx/sites-enabled/
sudo rm -f /etc/nginx/sites-enabled/default
sudo nginx -t && sudo systemctl reload nginx

# Point your domain's A record at the EC2 public IP first, then:
sudo certbot --nginx -d moshcore.com.ng -d www.moshcore.com.ng
```

Certbot rewrites the config to add the `443` server block and sets up auto-renewal.

## 4. Add GitHub repo secrets

`Settings → Secrets and variables → Actions → New repository secret`:

| Secret         | Value                                                            |
| -------------- | --------------------------------------------------------------- |
| `EC2_HOST`     | EC2 public IP or DNS name                                       |
| `EC2_USER`     | SSH user (e.g. `ubuntu` or `ec2-user`)                         |
| `EC2_SSH_KEY`  | The **private** key (full PEM contents) for that user           |
| `EC2_SSH_PORT` | *(optional)* SSH port, defaults to `22`                        |

`GITHUB_TOKEN` is provided automatically — no secret needed for GHCR.

> The workflow logs into GHCR on the server using `GITHUB_TOKEN`, so pulling a
> **private** image works out of the box. If you prefer, make the GHCR package
> public under the repo's Packages settings.

## 5. Deploy

Push to `master` (or run the workflow manually via "Run workflow"). The pipeline:

1. Builds the image and tags it `:<short-sha>` and `:latest`.
2. Pushes to `ghcr.io/<owner>/<repo>`.
3. Copies `deploy/docker-compose.yml` to `~/portfolio` on the server.
4. Pulls the new image and runs `docker compose up -d`.

Verify:

```bash
curl -s https://example.com/healthz
docker ps
docker logs portfolio
```

## Notes

- Rollback: `PORTFOLIO_IMAGE=ghcr.io/<owner>/<repo>:<old-sha> docker compose up -d`
  from `~/portfolio`.
- The résumé PDF is excluded from the image via `.dockerignore`.
- Consider **not** committing the résumé PDF to a public repo (it contains
  personal contact details) — add it to `.gitignore` if the repo is public.
