package configs

import "rateBastion/enums"

type store struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

type Config struct {
	Strategy                       enums.Strategy  `json:"strategy"`
	MaxRequestsAllowedInTimeWindow int             `json:"max_requests_allowed_in_time_window"`
	TimeWindowInSeconds            int             `json:"time_window_in_seconds"`
	CacheType                      enums.CacheType `json:"cache_type"`
	CacheStore                     store           `json:"cache_store"`
}
