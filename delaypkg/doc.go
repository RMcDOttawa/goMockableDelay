package delaypkg

//	goMockableDelay module provides a simple pass-through to the built-in time.Sleep
//	delay functions.  This is done so that the delay functions can be mocked in tests.
//	Rather than having a unit test wait for a real delay, structure your code to use this
//	delay service, then mock the delay service in tests to return immediately.

//	func  NewDelayService(debug bool, verbosity int)
//		creates an instance of the delay service.  You can set debug and verbosity levels
//		to create output for testing.  verbosity is from 0 (silent) to 5 (verbose).
//		These can be changed later with SetDebug and SetVerbosity.

//	DelayDuration(seconds int) (int, error)
//		Delays for the given number of seconds
//		Returns the number of seconds actually delayed (useful for keeping a tally), or an error

//	DelayUntil(target time.Time) error
//		Delays until the given time.Time is reached
//			This is approximate.  It won't return before the target time, but may return after by a small amount
