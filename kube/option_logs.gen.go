// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var logsOptions = []prompt.Suggest{
	{Text: "--all-containers", Description: "Get all containers' logs in the pod(s)."},
	{Text: "-c", Description: "Print the logs of this container"},
	{Text: "--container", Description: "Print the logs of this container"},
	{Text: "-f", Description: "Specify if the logs should be streamed."},
	{Text: "--follow", Description: "Specify if the logs should be streamed."},
	{Text: "--ignore-errors", Description: "If watching / following pod logs, allow for any errors that occur to be non-fatal"},
	{Text: "--insecure-skip-tls-verify-backend", Description: "Skip verifying the identity of the kubelet that logs are requested from.  In theory, an attacker could provide invalid log content back. You might want to use this if your kubelet serving certificates have expired."},
	{Text: "--limit-bytes", Description: "Maximum bytes of logs to return. Defaults to no limit."},
	{Text: "--max-log-requests", Description: "Specify maximum number of concurrent logs to follow when using by a selector. Defaults to 5."},
	{Text: "--pod-running-timeout", Description: "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running"},
	{Text: "--prefix", Description: "Prefix each log line with the log source (pod name and container name)"},
	{Text: "-p", Description: "If true, print the logs for the previous instance of the container in a pod if it exists."},
	{Text: "--previous", Description: "If true, print the logs for the previous instance of the container in a pod if it exists."},
	{Text: "-l", Description: "Selector (label query) to filter on."},
	{Text: "--selector", Description: "Selector (label query) to filter on."},
	{Text: "--since", Description: "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used."},
	{Text: "--since-time", Description: "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used."},
	{Text: "--tail", Description: "Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwise 10, if a selector is provided."},
	{Text: "--timestamps", Description: "Include timestamps on each line in the log output"},
}
