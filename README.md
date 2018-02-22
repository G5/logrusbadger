# logrusbadger [![GoDoc](https://godoc.org/github.com/G5/logrusbadger?status.svg)](https://godoc.org/github.com/G5/logrusbadger)

[![Codeship Status for G5/logrusbadger](https://app.codeship.com/projects/d344e2c0-f9a0-0135-ba3b-36ac52e54289/status?branch=master)](https://app.codeship.com/projects/278592)
[![Maintainability](https://api.codeclimate.com/v1/badges/5e0e37d9c4ae1eaa72b5/maintainability)](https://codeclimate.com/github/G5/logrusbadger/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/5e0e37d9c4ae1eaa72b5/test_coverage)](https://codeclimate.com/github/G5/logrusbadger/test_coverage)

`logrusbadger` is a [logrus](https://github.com/sirupsen/logrus)-compatible hook that uses the official [Honeybadger go library](https://github.com/honeybadger-io/honeybadger-go) to notify Honeybadger.

Any fields included with the log message will be added to the context of the honeybadger notification, and the main message of the log will be used as it's type. Any error added using logrus's `WithError` function will be set as the exception's type. This encourages you to use generic error messages like `retrieving current user` as your message, while added specific details as additional fields.

### Usage

Just register the hook with logrus. All configuration (API keys, timeouts, etc) should happen via Honeybadger's official libraries.

```golang
log.AddHook(logrusbadger.NewHook())
```

### License

This library is MIT licensed. See the [LICENSE](https://raw.github.com/G5/logrusbadger/master/LICENSE) file in this repository for details.
