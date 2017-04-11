package cli

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	admin_api "github.com/hexdecteam/easegateway-go-client/rest/1.0/admin/v1"
	health_api "github.com/hexdecteam/easegateway-go-client/rest/1.0/health/v1"
	stat_api "github.com/hexdecteam/easegateway-go-client/rest/1.0/statistics/v1"
)

var (
	serviceAddress = "localhost:9090"
)

func SetGatewayServerAddress(address string) error {
	if !strings.Contains(address, ":") {
		address = fmt.Sprintf("%s:%d", address, 9090)
	}

	host, portStr, err := net.SplitHostPort(address)
	port, err1 := strconv.Atoi(portStr)
	if err != nil || err1 != nil || port > 65535 {
		return fmt.Errorf("invalid port of gateway service address")
	}

	addr := net.ParseIP(host)
	if addr == nil {
		addresses, err := net.LookupHost(host)
		if err != nil || len(addresses) == 0 {
			return fmt.Errorf("invalid hostname or IP of gateway service address")
		}
	}

	serviceAddress = address

	return nil
}

func GatewayServerAddress() string {
	return serviceAddress
}

////

type multipleErr struct {
	errs []error
}

func (e *multipleErr) append(err error) {
	e.errs = append(e.errs, err)
}

func (e *multipleErr) String() string {
	if e.errs == nil {
		return "<nil>"
	}

	var s string
	for _, err := range e.errs {
		s = fmt.Sprintf("%s%s\n", s, err.Error())
	}
	return s
}

func (e *multipleErr) Error() string {
	return e.String()
}

// supply the interface gap
func (e *multipleErr) Return() error {
	if len(e.errs) == 0 {
		return nil
	}
	return e
}

////

func adminApi() *admin_api.AdminApi {
	return admin_api.NewAdminApi(serviceAddress)
}

func statisticsApi() *stat_api.StatisticsApi {
	return stat_api.NewStatisticsApi(serviceAddress)
}

func healthApi() *health_api.HealthApi {
	return health_api.NewHealthApi(serviceAddress)
}
