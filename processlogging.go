package supervisord

type (
	// A LogSegment represents a "tail" of a log
	LogSegment struct {
		Payload  string `xmlrpc:"string"`
		Offset   int    `xmlrpc:"offset"`
		Overflow bool   `xmlrpc:"overflow"`
	}
)

func (c *Client) logCall(method string, args ...interface{}) ([]LogSegment, error) {
	var logSegments []LogSegment
	err := c.Call(method, args, &logSegments)
	if err != nil {
		return nil, err
	}

	return logSegments, nil
}

// Read length bytes from name’s stdout log starting at offset.
func (c *Client) ReadProcessStdoutLog(name string, offset int, length int) (string, error) {
	return c.stringCall("supervisor.readProcessStdoutLog", name, offset, length)
}

// Read length bytes from name’s stderr log starting at offset.
func (c *Client) ReadProcessStderrLog(name string, offset int, length int) (string, error) {
	return c.stringCall("supervisor.readProcessStderrLog", name, offset, length)
}

// This is not implemented yet.
func (c *Client) TailProcessStdoutLog(name string, offset int, length int) ([]LogSegment, error) {
	return c.logCall("supervisor.tailProcessStdoutLog", name, offset, length)
}

// This is not implemented yet.
func (c *Client) TailProcessStderrLog(name string, offset int, length int) ([]LogSegment, error) {
	return nil, FIXMENotImplementedError
}

// Clear the stdout and stderr logs for the process name and reopen them.
func (c *Client) ClearProcessLogs(name string) error {
	return c.boolCall("supervisor.clearProcessLogs", name)
}

// Clear all process log files.
func (c *Client) ClearAllProcessLogs() error {
	return c.boolCall("supervisor.clearAllProcessLogs")
}
