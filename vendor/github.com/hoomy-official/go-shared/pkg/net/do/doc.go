// Package do provides a customizable HTTP request execution framework
// allowing for pre-request and post-request hooks. It employs a functional
// options approach for flexible request configuration.
//
// Use the Do function to perform an HTTP request with a specific URL and
// options, applying pre-defined or custom behaviors via handlers for actions
// such as logging or manipulating requests and responses. Default configured
// options set the HTTP method to GET, utilize the default HTTP client, and
// employ a no-operation logger. Users may override these by supplying their
// own options when invoking Do.
package do
