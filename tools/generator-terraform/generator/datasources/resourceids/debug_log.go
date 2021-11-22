package resourceids

import "log"

func (i DataSourceForResourceIdGenerator) log(format string, v ...interface{}) {
	if !i.Debug {
		return
	}

	log.Printf(format, v...)
}
