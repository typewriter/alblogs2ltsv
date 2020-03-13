package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"regexp"
)

func convert() {
	// see: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-access-logs.html#access-log-entry-syntax
	//      https://docs.aws.amazon.com/athena/latest/ug/application-load-balancer-logs.html#create-alb-table
	r := regexp.MustCompile("^(?P<type>[^ ]*) (?P<timestamp>[^ ]*) (?P<elb>[^ ]*) (?P<client>[^ ]*) (?P<target>[^ ]*) (?P<request_processing_time>[-.0-9]*) (?P<target_processing_time>[-.0-9]*) (?P<response_processing_time>[-.0-9]*) (?P<elb_status_code>|[-0-9]*) (?P<target_status_code>-|[-0-9]*) (?P<received_bytes>[-0-9]*) (?P<sent_bytes>[-0-9]*) \"(?P<request>[^\"]*)\" \"(?P<user_agent>[^\"]*)\" (?P<ssl_cipher>[A-Z0-9-]+) (?P<ssl_protocol>[A-Za-z0-9.-]*) (?P<target_group_arn>[^ ]*) \"(?P<trace_id>[^\"]*)\" \"(?P<domain_name>[^\"]*)\" \"(?P<chosen_cert_arn>[^\"]*)\" (?P<matched_rule_priority>[-.0-9]*) (?P<request_creation_time>[^ ]*) \"(?P<actions_executed>[^\"]*)\" \"(?P<redirect_url>[^\"]*)\" \"(?P<error_reason>[^\"]*)\" \"(?P<target_port_list>[^\"]*)\" \"(?P<target_status_code_list>[^\"]*)\"")

	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		match := r.FindStringSubmatch(text)

		if len(match) <= 1 {
			continue
		}

		for i, name := range r.SubexpNames() {
			if i == 1 {
				fmt.Printf("%s:%s", name, match[i])
			}
			if i >= 2 {
				fmt.Printf("\t%s:%s", name, match[i])
			}
		}
		fmt.Println()
	}
}

func main() {
	cmd := &cobra.Command{
		Use:   "alblogs2ltsv",
		Short: "AWS ALB (Application Load Balancer) access log converter to LTSV",
		Run: func(cmd *cobra.Command, args []string) {
			convert()
		},
	}

	cmd.Execute()
}
