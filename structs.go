package main

// AWS docs found here https://docs.aws.amazon.com/guardduty/latest/APIReference/API_GetFindings.html#API_GetFindings_ResponseSyntax

// GuardDutyFindingDetails contains the actual findings from GuardDuty
type GuardDutyFindingDetails struct {
	AccountID   string `json:"accountId"`
	ARN         string `json:"arn"`
	Confidence  int    `json:"confidence"`
	CreatedAt   string `json:"createdAt"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Partition   string `json:"partition"`
	Region      string `json:"region"`
	// Resource
	SchemaVersion string                   `json:"schemaVersion"`
	Resource      GuardDutyResourceActions `json:"resource"`
	Service       GuardDutyServiceActions  `json:"service"`
	Severity      int                      `json:"severity"`
	Title         string                   `json:"title"`
	Type          string                   `json:"type"`
	UpdatedAt     string                   `json:"updatedAt"`
}

type GuardDutyResourceActions struct {
	AccessKeyDetails struct {
		AccessKeyID string `json:"accessKeyId"`
		PrincipalID string `json:"principalId"`
		Username    string `json:"username"`
		UserType    string `json:"userType"`
	} `json:"accessKeyDetails"`
	InstanceDetails struct {
		AZ                 string `json:"availabilityZone"`
		IAMInstanceProfile struct {
			ARN string `json:"arn"`
			ID  string `json:"id"`
		} `json:"iamInstanceProfile"`
		ImageDescription  string                       `json:"imageDescription"`
		ImageID           string                       `json:"imageId"`
		InstanceID        string                       `json:"instanceID"`
		InstanceState     string                       `json:"instanceState"`
		InstanceType      string                       `json:"instanceType"`
		LaunchTime        string                       `json:"launchTime"`
		NetworkInterfaces []GuardDutyNetworkInterfaces `json:"networkInterfaces"`
		OutpostARN        string                       `json:"outpostArn"`
		Platform          string                       `json:"platform"`
		ProductCodes      []struct {
			Code        string `json:"code"`
			ProductType string `json:"productType"`
		} `json:"productCodes"`
		Tags []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"tags"`
		ResourceGroup string `json:"resourceGroup"`
	} `json:"instanceDetails"`
	ResourceType string `json:"resourceType"`
}

type GuardDutyServiceActions struct {
	Action struct {
		ActionType       string `json:"actionType"`
		AWSAPICallAction struct {
			API           string `json:"api"`
			CallerType    string `json:"callerType"`
			ServiceName   string `json:"serviceName"`
			DomainDetails struct {
				Domain string `json:"domain"`
			} `json:"domainDetails"`
			RemoteIPDetails GuardDutyRemoteIPDetails `json:"remoteIpDetails"`
		} `json:"awsApiCallAction"`
		DNSRequestAction struct {
			Domain string `json:"domain"`
		} `json:"dnsRequestAction"`
		NetworkConnectionAction struct {
			Blocked             bool   `json:"blocked"`
			ConnectionDirection string `json:"connectionDirection"`
			LocalIPDetails      struct {
				IPAddressV4 string `json:"ipAddressV4"`
			} `json:"localIpDetails"`
			LocalPortDetails  GuardDutyPortDetails     `json:"localPortDetails"`
			Protocol          string                   `json:"protocol"`
			RemoteIPDetails   GuardDutyRemoteIPDetails `json:"remoteIpDetails"`
			RemotePortDetails GuardDutyPortDetails     `json:"remotePortDetails"`
		} `json:"networkConnectionAction"`
		PortProbeAction struct {
			Blocked          bool                        `json:"blocked"`
			PortProbeDetails []GuardDutyPortProbeDetails `json:"portProbeDetails"`
		}
	} `json:"action"`
	Archived       bool   `json:"archived"`
	Count          int    `json:"count"`
	DetectorID     string `json:"detectorId"`
	EventFirstSeen string `json:"eventFirstSeen"`
	EventLastSeen  string `json:"eventLastSeen"`
	Evidence       struct {
		ThreatIntelDetails []struct {
			ThreatListName string   `json:"threatListName"`
			ThreatNames    []string `json:"threatNames"`
		} `json:"threatIntelligenceDetails"`
	} `json:"evidence"`
	ResourceRole string `json:"resourceRole"`
	ServiceName  string `json:"serviceName"`
	UserFeedback string `json:"userFeedback"`
}

// GuardDutyRemoteIPDetails provides details about the remote IP address involved in the alert
type GuardDutyRemoteIPDetails struct {
	City struct {
		CityName string `json:"cityName"`
	} `json:"city"`
	Country struct {
		CountryCode string `json:"countryCode"`
		CountryName string `json:"countryName"`
	} `json:"country"`
	GeoLocation struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"geoLocation"`
	IPAddressV4  string `json:"ipAddressV4"`
	Organization struct {
		ASN    string `json:"asn"`
		ASNOrg string `json:"asnorg"`
		ISP    string `json:"isp"`
		Org    string `json:"org"`
	} `json:"organization"`
}

// GuardDutyPortDetails provides information about the ports involved in the alert
type GuardDutyPortDetails struct {
	Port     int    `json:"port"`
	PortName string `json:"portName"`
}

// GuardDutyPortProbeDetails provides information about a port probe event
type GuardDutyPortProbeDetails struct {
	LocalIPDetails struct {
		IPAddressV4 string `json:"ipAddressV4"`
	} `json:"localIpDetails"`
	LocalPortDetails GuardDutyPortDetails     `json:"localPortDetails"`
	RemoteIPDetails  GuardDutyRemoteIPDetails `json:"remoteIpDetails"`
}

// GuardDutyNetworkInterfaces provides information about the network interface related to an alert
type GuardDutyNetworkInterfaces struct {
	IPv6Addresses      []string `json:"ipv6Addresses"`
	NetworkInterfaceID string   `json:"networkInterfaceId"`
	PrivateDNSName     string   `json:"privateDnsName"`
	PrivateIPAddress   string   `json:"privateIpAddress"`
	PrivateIPAddresses []struct {
		PrivateDNSName  string `json:"privateDnsName"`
		PrivateIPAddres string `json:"privateIpAddress"`
	} `json:"privateIpAddresses"`
	PublicDNSName  string `json:"publicDnsName"`
	PublicIP       string `json:"publicIp"`
	SecurityGroups []struct {
		GroupID   string `json:"groupId"`
		GroupName string `json:"groupName"`
	} `json:"securityGroups"`
	SubnetID string `json:"subnetId"`
	VPCID    string `json:"vpcId"`
}

// SlackMessage is the message we send to Slack
type SlackMessage struct {
	Text        string             `json:"text"`
	Attachments []SlackAttachments `json:"attachments"`
}

// SlackAttachments is the attachments we send in our message
type SlackAttachments struct {
	Text   string        `json:"text"`
	Color  string        `json:"color"`
	Title  string        `json:"title"`
	Fields []SlackFields `json:"fields"`
}

// SlackFields are additional fields in the attachment
type SlackFields struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}
