# fk-yew-camella

> An automated scheduled mailer to send recurring emails to companies that refuse to answer your follow-ups. Built to be the thorn in the side of corporate indifference.

## 💡 Why?

Because companies — especially some in the Philippines (hi, Camella 👋) — have a talent for ignoring emails. Instead of manually sending follow-up emails and pulling your hair out, let this little Go app automate the process for you.

Send as many emails as you want, until you get an answer you're satisfied with. Or until you feel petty enough to finally step on the brake. Up to you.

---

## 📦 Requirements

### ✅ [Go](https://golang.org/dl/)
If you want to customize the app or build it from source.

> **Q: Why Go? Why not JavaScript or Python?**  
> **A:** Because we don’t need a 50MB interpreter just to send an email. Go compiles to a single statically linked binary — often under 7MB; no bloat, no bullshit.

### ✅ [GitHub](https://github.com/)
You need GitHub to:
- Host this repository
- Let GitHub Actions schedule and run your mailer without your computer being online

### ✅ Gmail App Password
You’ll need to create a Gmail App Password (since normal passwords won't work with SMTP).  
Don’t know how? You're resourceful enough to find games and movies in torrent, you can Google this in under 5 minutes.

---

## 🚀 How to Use

### 1. Clone the repository

```bash
git clone https://github.com/jallenmanaloto/fk-yew-camella.git
cd fk-yew-camella
```

### 2. Fill out your preferences in `editme.json`
```bash
This is the only file you need to edit. Don’t touch anything else unless you know what you’re doing.
```
Here's a breakdown of what each field does:

| Field          | Type              | Description                                                                 |
|----------------|-------------------|-----------------------------------------------------------------------------|
| `email`        | string            | Your Gmail address (the sender)                                            |
| `password`     | string            | Your Gmail App Password |
| `enable`       | boolean           | Set to `true` to send, `false` to pause                                    |
| `subject`      | string            | Email subject line                       |
| `message_body` | string            | Body of the email — express your frustration here                                                            |
| `to`           | array of string   | Main recipients                                                             |
| `cc`           | array of string   | Optional — feel free to involve other people in your frustration                                     |
| `bcc`          | array of string   | Optional — idk why I put this, but use it if you wish to                              |
| `schedule`     | string            | One of: `daily`, `weekly`, `hourly`, or `custom`                           |
| `cron`         | string            | Used only if `schedule` is `custom`. Define your [cron expression here](https://crontab.guru) |


Example:
```json
{
  "email": "you@gmail.com",
  "password": "abcd efgh ijkl mnop",
  "enable": true,
  "subject": "Still Waiting for Your Reply",
  "message_body": "Hey, just following up again...",
  "to": ["support@shitcompany.com"],
  "cc": ["lawyer@example.com"],
  "bcc": [],
  "schedule": "daily",
  "cron": ""
}
```
### 3. Generate the scheduler workflow
This step creates the .github/workflows/scheduler.yml file based on your config:
```bash
go run . --generate
```
### 4. Commit the generated workflow
GitHub won't run the scheduled mailer unless the workflow exists in the repo.

## ⚙️ How It Works
When the app runs, it reads your `editme.json`.

It generates a GitHub Actions workflow based on your selected schedule.

It sets up your Go environment and sends the email using Gmail SMTP.

The email will continue sending until you set `"enable": false`.

You don't need your PC to be online. Let this devil deliver your wrath consistently

##  🐢 Why This Exists
Aren't we all tired of manually sending follow-ups? So now, they get reminders from us... every day. Until we get what we want. This app was born out of that pain and pettiness. You're welcome to fork it for your own war.

## 🧰 Want to Customize It?
If you're a developer and want to:

- Customize the build

- Add attachments

- Use other email providers

- Change the email template

Feel free to poke around the code. It's written in Go, modular, and tested.

## 🧼 Disclaimers
Don’t use this to spam people. Be civil (even if you're being a savage).

This isn’t a mass-mailing tool. It’s a polite-aggressive follow-up bot.

I take no responsibility for how you use it.
