# Parking service tickets scraper

Small command line utility to give information about parking tickets based
on licence number for Belgrade, based on information provided by [Parking service](https://parking-servis.co.rs).

This works by scraping HTML provided by parking service. It has been tested for
entire 3 minutes. 

## Usage
Install:
```bash
$ go get github.com/delicb/parking-service-tickets-scraper
``` 

Use:
```bash
$ parking-service-tickets-scraper BG1234AB
```
