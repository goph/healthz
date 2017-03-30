# Change Log


## Unreleased

### Added

- TCP Checker
- HTTP Timeout option
- TCP timeout option


## 0.5.0 - 2017-03-09

### Changed

- Replace DB Checker with Ping Checker


## 0.4.0 - 2017-02-16

### Added

- http.Handler constructor
- Status enum with `Healthy` and `Unhealthy` values
- Checker collection implementing the `Checker` interface
- `CheckFunc` to allow any func implementing a specific signature become a Checker

### Changed

- HealthService members are not exported anymore
- Rename `Ping` to `Check`
- Rename `ErrHealthCheckFailed` to `ErrCheckFailed`
- Rename `HealthChecker` to `Checker`
- Improved comments
- Moved tests to a separate package (see [here](https://medium.com/@povilasve/go-advanced-tips-tricks-a872503ac859#.ii6f9mgjv) why)

### Removed

- `HealthService.RegisterHandlers` method
- `Type` method from `Checker` interface
- `Probe` in favor of `Checker` interface and collection
- Mocked Checker
- Test framework


## 0.3.0 - 2017-02-09

### Changed

- Renamed `HealthService.RegisterMux` to `HealthService.RegisterHandlers`


## 0.2.0 - 2017-02-09

### Added

- DB (database/sql) and HTTP Health Checkers


## 0.1.0 - 2017-02-08

- Initial release
