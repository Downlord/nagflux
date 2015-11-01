package spoolfile

import (
	"fmt"
	"github.com/griesbacher/nagflux/helper"
)

type PerformanceData struct {
	hostname         string
	service          string
	command          string
	performanceLabel string
	performanceType  string
	unit             string
	time             string
	value            string
	fieldseperator   string
	tags             map[string]string
}

func (p *PerformanceData) String() string {
//	tableName := fmt.Sprintf(`%s%s%s%s%s%s%s%s%s`,
//		p.hostname, p.fieldseperator,
//		p.service, p.fieldseperator,
//		p.command, p.fieldseperator,
//		p.performanceLabel, p.fieldseperator,
//		p.performanceType)
	tableName := fmt.Sprintf(`%s%s%s%s%s%s%s`, p.service, p.fieldseperator ,`command=`, p.command, p.fieldseperator , `host=` , p.hostname )
	//if p.performanceType != "value" {
	//	//tableName += fmt.Sprintf(`%s%s%s%s`, p.fieldseperator , p.performanceLabel, `=`, p.performanceType)
	//	tableName += fmt.Sprintf(`%s%s%s`, p.fieldseperator , `perftype=`, p.performanceType)
	//}
	if p.unit != "" {
		tableName += fmt.Sprintf(`%s%s%s`, p.fieldseperator , `unit=`, p.unit)
	}

	p.tags["type"] = p.performanceType
	if len(p.tags) > 0 {
		tableName += fmt.Sprintf(`,%s`, helper.PrintMapAsString(p.tags, ",", "="))
	}

	//if p.performanceType == "value" {
		tableName += fmt.Sprintf(` %s=%s %s`, p.performanceLabel, p.value, p.time)
	//}else{
		//tableName += fmt.Sprintf(` value=%s %s`, p.value, p.time)
	//}
	return tableName
}
