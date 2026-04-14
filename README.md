# cron

<p align="left">
  [![Go CI](https://github.com/verystar/crondesc/actions/workflows/go.yml/badge.svg)](https://github.com/verystar/crondesc/actions/workflows/go.yml)
  <a href="https://godoc.org/github.com/verystar/crondesc" title="GoDoc Reference" rel="nofollow"><img src="https://img.shields.io/badge/go-documentation-blue.svg?style=flat" alt="GoDoc Reference"></a>
  <a href="https://goreportcard.com/report/github.com/verystar/crondesc"><img src="https://goreportcard.com/badge/github.com/verystar/crondesc" alt="Code Status" /></a>
</p>

cron is a Go library that parses a cron expression and outputs a human readable description of the cron schedule.  
For example, given the expression `*/5 * * * *` it will output `Every 5 minutes`.

Translated to Go from [cron-expression-descriptor](https://github.com/bradymholt/cron-expression-descriptor) (C#) via [cRonstrue](https://github.com/bradymholt/cRonstrue) (Javascript).  
Original Author & Credit: Brady Holt (http://www.geekytidbits.com).

## Features

- Zero dependencies
- Supports all cron expression special characters including `* / , - ? L W #`
- Supports 5, 6 (w/ seconds or year), or 7 (w/ seconds and year) part cron expressions
- Supports [Quartz Job Scheduler](http://www.quartz-scheduler.org/) cron expressions
- i18n support with 26 locales.

## Installation

`crondesc` module can be used with both Go module (>= 1.11) and earlier Go versions.

```
go get -u -v github.com/verystar/crondesc
```

## Usage

```go
// Init with default EN locale
exprDesc, _ := crondesc.NewDescriptor()

desc, _ := exprDesc.ToDescription("* * * * *", crondesc.Locale_en)
// "Every minute"

desc, _ := exprDesc.ToDescription("0 23 ? * MON-FRI", crondesc.Locale_en)
// "At 11:00 PM, Monday through Friday"

desc, _ := exprDesc.ToDescription("23 14 * * SUN#2", crondesc.Locale_en)
// "At 02:23 PM, on the second Sunday of the month"

// Init with custom configs
exprDesc, _ := crondesc.NewDescriptor(
    crondesc.Use24HourTimeFormat(true),
    crondesc.DayOfWeekStartsAtOne(true),
    crondesc.Verbose(true),
    crondesc.SetLogger(log.New(os.Stdout, "cron: ", 0)),
    crondesc.SetLocales(crondesc.Locale_en, crondesc.Locale_fr),
)
```

For more usage examples, including a demonstration of how cron can handle some very complex cron expressions, you can reference [the unit tests](https://github.com/verystar/crondesc/blob/develop/locale_en_test.go) or [the example codes](https://github.com/verystar/crondesc/tree/develop/examples).

## i18n

To use the i18n support, you must configure the locales when create a new `ExpressionDescriptor` via `SetLocales()` option.

```go
exprDesc, _ := crondesc.NewDescriptor(
    crondesc.SetLocales(crondesc.Locale_en, crondesc.Locale_es, crondesc.Locale_fr),
)
// or load all crondesc.LocaleAll
exprDesc, _ := crondesc.NewDescriptor(crondesc.SetLocales(crondesc.LocaleAll))

desc, _ := exprDesc.ToDescription("* * * * *", crondesc.Locale_fr)
// Toutes les minutes
```

By default, `ExpressionDescriptor` always load the `Locale_en`. If you pass an unregistered locale into `ToDescription()` function, the result will be returned in English.

### Supported Locales

| Locale Code | Language              | Contributors                                             |
| ----------- | --------------------- | -------------------------------------------------------- |
| cs          | Czech                 | [hanbar](https://github.com/hanbar)                      |
| da          | Danish                | [Rasmus Melchior Jacobsen](https://github.com/rmja)      |
| de          | German                | [Michael Schuler](https://github.com/mschuler)           |
| en          | English               | [Brady Holt](https://github.com/bradymholt)              |
| es          | Spanish               | [Ivan Santos](https://github.com/ivansg)                 |
| fa          | Farsi                 | [A. Bahrami](https://github.com/alirezakoo)              |
| fi          | Finnish               | [Mikael Rosenberg](https://github.com/MR77FI)            |
| fr          | French                | [Arnaud TAMAILLON](https://github.com/Greybird)          |
| he          | Hebrew                | [Ilan Firsov](https://github.com/IlanF)                  |
| it          | Italian               | [rinaldihno](https://github.com/rinaldihno)              |
| ja          | Japanese              | [Alin Sarivan](https://github.com/asarivan)              |
| ko          | Korean                | [Ion Mincu](https://github.com/ionmincu)                 |
| nb          | Norwegian             | [Siarhei Khalipski](https://github.com/KhalipskiSiarhei) |
| nl          | Dutch                 | [TotalMace](https://github.com/TotalMace)                |
| pl          | Polish                | [foka](https://github.com/foka)                          |
| pt_BR       | Portuguese (Brazil)   | [Renato Lima](https://github.com/natenho)                |
| ro          | Romanian              | [Illegitimis](https://github.com/illegitimis)            |
| ru          | Russian               | [LbISS](https://github.com/LbISS)                        |
| sk          | Slovakian             | [hanbar](https://github.com/hanbar)                      |
| sl          | Slovenian             | [Jani Bevk](https://github.com/jenzy)                    |
| sv          | Swedish               | [roobin](https://github.com/roobin)                      |
| sw          | Swahili               | [Leylow Lujuo](https://github.com/leyluj)                |
| tr          | Turkish               | [Mustafa SADEDİL](https://github.com/sadedil)            |
| uk          | Ukrainian             | [Taras](https://github.com/tbudurovych)                  |
| zh_CN       | Chinese (Simplified)  | [Star Peng](https://github.com/starpeng)                 |
| zh_TW       | Chinese (Traditional) | [Ricky Chiang](https://github.com/metavige)              |
