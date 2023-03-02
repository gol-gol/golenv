
## golenv

> facade for regular used Go Environment Variable calls
>
> extracted from an older pandora box of such packages at [abhishekkr/gol](https://github.com/abhishekkr/gol)

### Public Functions

* `OverrideIfEnv(envVar string, defaultValue string) string`

* `OverrideIfEnvBool(envVar string, defaultValue bool) bool`

* `OverrideIfEnvInt(envVar string, defaultValue int) int`

* `OverrideIfEnvUint64(envVar string, defaultValue uint64) uint64`

* `HasEnv(envVar string) bool`

* `EnvMap() map[string]string`

---
