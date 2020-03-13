# alblogs2ltsv

AWS ALB (Application Load Balancer) access log converter to LTSV

## Installation

Pre-build binaries available on [Releases page](https://github.com/typewriter/alblogs2ltsv/releases).

## Usage

```bash
$ zcat AWSACCOUNTID_elasticloadbalancing_REGION_LOADBALANCERID_TIME_IPADDRESS_RANDOM.log.gz | ./alblogs2ltsv
type:http	timestamp:2020-03-11T12:36:00.610235Z	elb:LOADBALANCERID	client:172.31.40.205:57096	target:172.31.40.210:80	request_processing_time:0.001	target_processing_time:0.000	response_processing_time:0.000	elb_status_code:200	target_status_code:200	received_bytes:493	sent_bytes:2499	request:GET http://LOADBALANCER.REGION.elb.amazonaws.com:80/example.html HTTP/1.1	user_agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36	ssl_cipher:-	ssl_protocol:-	target_group_arn:arn:aws:elasticloadbalancing:REGION:AWSACCOUNTID:targetgroup/TARGETGROUPID	trace_id:Root=TRACEID	domain_name:-	chosen_cert_arn:-	matched_rule_priority:0	request_creation_time:2020-03-11T12:36:00.609000Z	actions_executed:forward	redirect_url:-	error_reason:-	target_port_list:172.31.40.210:80	target_status_code_list:200
type:http	timestamp:2020-03-11T12:36:00.880434Z	elb:LOADBALANCERID	client:172.31.40.205:57096	target:172.31.40.210:80	request_processing_time:0.000	target_processing_time:0.000	response_processing_time:0.000	elb_status_code:404	target_status_code:404	received_bytes:458	sent_bytes:710	request:GET http://LOADBALANCER.REGION.elb.amazonaws.com:80/favicon.ico HTTP/1.1	user_agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.122 Safari/537.36	ssl_cipher:-	ssl_protocol:-	target_group_arn:arn:aws:elasticloadbalancing:REGION:AWSACCOUNTID:targetgroup/TARGETGROUPID	trace_id:Root=TRACEID	domain_name:-	chosen_cert_arn:-	matched_rule_priority:0	request_creation_time:2020-03-11T12:36:00.879000Z	actions_executed:forward	redirect_url:-	error_reason:-	target_port_list:172.31.40.210:80	target_status_code_list:404
```

## See

Amazon Athena makes it easy to query logs. This tool may be useful for more advanced processing or conversion.

[Querying Application Load Balancer Logs - Amazon Athena](https://docs.aws.amazon.com/athena/latest/ug/application-load-balancer-logs.html)

## License

[MIT](https://choosealicense.com/licenses/mit/)

