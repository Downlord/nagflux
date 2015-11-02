package spoolfile

import (
	"fmt"
	"github.com/griesbacher/nagflux/helper"
)

type PerformanceData struct {
	hostname         string
	service          string
	awsstage         string
	awsservice       string
	awsserviceid     string
	awsinstanceid    string
	command          string
	performanceLabel string
	performanceType  string
	unit             string
	time             string
	value            string
	fieldseperator   string
	tags             map[string]string
	mvals            map[string]string
}

func (p *PerformanceData) String() string {
	if p.awsinstanceid == "" {
		p.awsinstanceid = "empty"
	}
	tableName := fmt.Sprintf(`%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s`, p.service, p.fieldseperator ,`command=`, p.command, p.fieldseperator ,`awsinstanceid=`, p.awsinstanceid, p.fieldseperator ,`awsservice=`, p.awsservice, p.fieldseperator ,`awsserviceid=`, p.awsserviceid, p.fieldseperator ,`awsstage=`, p.awsstage, p.fieldseperator, `host=` , p.hostname )
	if p.unit != "" {
		p.tags["unit"] = p.unit
	}

	p.tags["metric"] = p.performanceLabel
	
	if len(p.tags) > 0 {
		tableName += fmt.Sprintf(`,%s`, helper.PrintMapAsString(p.tags, ",", "="))
	}

	if len(p.mvals) > 0 {
		tableName += fmt.Sprintf(` %s`, helper.PrintMapAsString(p.mvals, ",", "="))
	}

   tableName += fmt.Sprintf(` %s`, p.time)
	return tableName
}
