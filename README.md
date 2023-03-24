# CheckSum Calculator for FIX Protocol Messages

Fix Protocol CheckSum Calculator

## How to use it

Used as reference: https://stackoverflow.com/questions/32708068/how-to-calculate-checksum-in-fix-manually

- [node.js] v19.7.0

```sh
node checksum.js "8=FIX.4.2|9=49|35=5|34=1|49=ARCA|52=20150916-04:14:05.306|56=TW|"
```

OutPut

```sh
┌─────────┬───────────────────────────────────────────────────────────────────────────┐
│ (index) │                                  Values                                   │
├─────────┼───────────────────────────────────────────────────────────────────────────┤
│  Input  │    '8=FIX.4.2|9=49|35=5|34=1|49=ARCA|52=20150916-04:14:05.306|56=TW|'     │
│ Output  │ '8=FIX.4.2|9=49|35=5|34=1|49=ARCA|52=20150916-04:14:05.306|56=TW|10=157|' │
└─────────┴───────────────────────────────────────────────────────────────────────────┘
```

- [GoLang] go1.20rc3

Update the file messages.txt with the messages you want to calculate. One message per line.

```sh
go run checksum.go
```

OutPut

```sh
8=FIX.4.2|9=49|35=5|34=1|49=ARCA|52=20150916-04:14:05.306|56=TW|10=157
8=FIX.4.4|9=49|35=5|34=1|49=ARCA|52=20150916-04:14:05.306|56=TW|10=159
8=FIX.5.0|9=49|35=5|34=1|49=ARCA|52=20150916-04:14:05.306|56=TW|10=156
```
