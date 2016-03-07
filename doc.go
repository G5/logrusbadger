// Package logrusbadger is a logrus-compatible hook that uses the official
// Honeybadger go library to notify Honeybadger. Any fields included with the
// log message will be added to the context of the honeybadger notification,
// and the main message of the log will be used as it's type. Any error added
// using logrus's WithError function will be set as the exception's type. This
// encourages you to use generic error messages like "retrieving current user"
// as your message, while added specific details as additional fields.
//
// Control over timeouts, API keys, and other options should be done through
// the Honeybadger's library. This uses its Notify package-level function, so
// any changes to that DefaultClient will take effect for this library as well.
//
// logrusbadger is influenced greatly by
// github.com/agonzalezro/logrus_honeybadger.
package logrusbadger
