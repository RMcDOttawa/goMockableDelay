# goMockableDelay

This module is a simple pass-through interface to the delay service in the standard go Time library.  It exposes two simple functions:

- Delay for a given number of seconds, and
- Delay until a given time

By using this service instead of calling the "time" package functions directly, you have the option of replacing your service with a mock for testing.

So, for example, you can write unit tests for a service that, in real life, includes delays, but mock those delays to be zero-length for testing purposes.

Functions

| Function        | Arguments        | Purpose                                                                               |
| --------------- | ---------------- | ------------------------------------------------------------------------------------- |
| NewDelayService | verbosity, debug | Creates a new delay service object, returning a pointer.                              |
| SetDebug        | boolean          | Sets the "debug" flag for the service                                                 |
| SetVerbosity    | int              | Sets the verbosity level, from 0 to 5                                                 |
| DelayDuration   | seconds int      | Delays for the given number of seconds. Returns the seconds delayed and an error code |
| DelayUntil      | time             | Delays until (at least) the given Time. Returns an error code.                        |

Create and use a MockDelayService using the normal mocking framework and inject it into your code under test for testing purposes.

e.g.,

````
mockDelayService := goMockableDelay.NewMockDelayService(ctrl)
mockDelayService.EXPECT().DelayDuration(10).Return(10, nil)
````