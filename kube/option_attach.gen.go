// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/c-bata/go-prompt"
)

var attachOptions = []prompt.Suggest{
	{Text: "-c", Description: "Container name. If omitted, use the kubectl.kubernetes.io/default-container annotation for selecting the container to be attached or the first container in the pod will be chosen"},
	{Text: "--container", Description: "Container name. If omitted, use the kubectl.kubernetes.io/default-container annotation for selecting the container to be attached or the first container in the pod will be chosen"},
	{Text: "--pod-running-timeout", Description: "The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running"},
	{Text: "-q", Description: "Only print output from the remote session"},
	{Text: "--quiet", Description: "Only print output from the remote session"},
	{Text: "-i", Description: "Pass stdin to the container"},
	{Text: "--stdin", Description: "Pass stdin to the container"},
	{Text: "-t", Description: "Stdin is a TTY"},
	{Text: "--tty", Description: "Stdin is a TTY"},
}
