package play

import (
	"regexp"
	"strconv"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

var (
	reClasic      *regexp.Regexp
	matcherClasic []string
)

func init() {
	//reClasic = regexp.MustCompile(`(?P<date>[^Z]+Z) (?P<elb>[^\s]+) (?P<ipclient>[^:]+?):[0-9]+ (?P<ipnode>[^:]+?):[0-9]+ (?P<reqtime>[0-9\.]+) (?P<backtime>[0-9\.]+) (?P<restime>[0-9\.]+) (?P<elbcode>[0-9]{3}) (?P<backcode>[0-9]{3}) (?P<lenght1>[0-9]+) (?P<lenght2>[0-9]+) "(?P<Method>[A-Z]+) (?P<URL>[^ ]+) HTTP/[0-9\.]+" "(?P<useragent>[^"]*)" .*`)
	reClasic = regexp.MustCompile(`^[\S]* - - [\[\]0-9\/A-z: +]*"(\S*) (\S*) HTTP\/1.1" (\S*) [0-9]* [\S]* "([\S ]*)"`)
	matcherClasic = reClasic.SubexpNames()
}

// parse the raw record line
func parse(line string) (*logLine, error) {
	matches := reClasic.FindAllStringSubmatch(line, -1)
	if matches == nil {
		return nil, errors.Errorf("Failed to parse the line %s", line)
	}

	match := matches[0]
	statusCode, err := strconv.Atoi(match[3])
	if err != nil {
		log.Errorf("Failed to parse status code (%s) from line %s", match[3], line)
	}

	return &logLine{
		statusCode: statusCode,
		method:     match[1],
		url:        match[2],
		userAgent:  match[4],
	}, nil
}
