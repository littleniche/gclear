## History Manager

### What was the Problem?

`HISTSIZE=10000`

```
$ history | grep clear | wc -l
1973
```

## Installation

```
$ go install github.com/littleniche/gclear@latest
```

### Shells supported currently

- `bash`
- `zsh`
- `fish`

### Usage

```
$ gclear
```

```
$ gclear -w make
```

### Setup cron job

- Create a daily cron job for gclear using `crontab`

```bash

#Edit crontab and add a job
$ crontab -e

# Add this line 
30 12 * * * /usr/local/bin/gclear

# Above cronjob will run daily at 12:30
```
- Test if cronjob is running
```bash
$ sudo grep –a “bash.sh” /var/log/syslog

Feb  2 11:01:01 tanmay-lenovo CRON[183156]: (tanmay) CMD (/usr/local/bin/gclear)
```