# Change Log


## UNRELEASED

### Added

- http.Handler constructor

### Changed

- HealthService members are not exported anymore
- Rename `Ping` to `Check`
- Rename `ErrHealthCheckFailed` to `ErrCheckFailed`
- Rename `HealthChecker` to `Checker`

### Removed

- `HealthService.RegisterHandlers` method


## 0.3.0 - 2017-02-09

### Changed

- Renamed `HealthService.RegisterMux` to `HealthService.RegisterHandlers`


## 0.2.0 - 2017-02-09

### Added

- DB (database/sql) and HTTP Health Checkers


## 0.1.0 - 2017-02-08

- Initial release
