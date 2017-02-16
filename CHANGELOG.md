# Change Log


## UNRELEASED

### Added

- http.Handler constructor
- Status enum with `Healthy` and `Unhealthy` values
- Checker collection implementing the `Checker` interface

### Changed

- HealthService members are not exported anymore
- Rename `Ping` to `Check`
- Rename `ErrHealthCheckFailed` to `ErrCheckFailed`
- Rename `HealthChecker` to `Checker`
- Improved comments

### Removed

- `HealthService.RegisterHandlers` method
- `Type` method from `Checker` interface
- `Probe` in favor of `Checker` interface and collection
- Mocked Checker


## 0.3.0 - 2017-02-09

### Changed

- Renamed `HealthService.RegisterMux` to `HealthService.RegisterHandlers`


## 0.2.0 - 2017-02-09

### Added

- DB (database/sql) and HTTP Health Checkers


## 0.1.0 - 2017-02-08

- Initial release
