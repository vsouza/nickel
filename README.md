
# nickel
<img src="https://upload.wikimedia.org/wikipedia/commons/7/72/Jefferson-Nickel-Unc-Obv.jpg" width="240" align="right"/>
A Golang AWS Billing CSV Processor.

About AWS CSV Billing reports:

Billing reports provide information about your usage of AWS resources and estimated costs for that usage. You can have AWS generate billing reports that break down your estimated costs in different ways:
 * By the hour, day, or month
 * By each account in your organization
 * By product or product resource
 * By tags that you define yourself
You can create tags for your AWS resources to add your own labels to nearly every line item in your reports. 
AWS delivers the reports in CSV format to an Amazon S3 bucket that you specify, and updates the reports at least once a day. 


## How it works?

## Setup

Install Godeps

`make deps`

## Run

`make run` or `./nickel -c nickel.conf`

## Tests

`make tests`


## Credits

* Thanks for proTip, this project is a customized version of [github.com/ProTip/aws-elk-billing](https://github.com/ProTip/aws-elk-billing)

## License

[MIT License](http://vsouza.mit-license.org/) © Vinicius Souza
