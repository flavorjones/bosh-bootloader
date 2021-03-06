package templates

type BOSHEIPTemplateBuilder struct{}

func NewBOSHEIPTemplateBuilder() BOSHEIPTemplateBuilder {
	return BOSHEIPTemplateBuilder{}
}

func (t BOSHEIPTemplateBuilder) BOSHEIP() Template {
	return Template{
		Resources: map[string]Resource{
			"BOSHEIP": Resource{
				DependsOn: "VPCGatewayAttachment",
				Type:      "AWS::EC2::EIP",
				Properties: EIP{
					Domain: "vpc",
				},
			},
		},
		Outputs: map[string]Output{
			"BOSHEIP": {Value: Ref{"BOSHEIP"}},
			"BOSHURL": {
				Value: FnJoin{
					Delimeter: "",
					Values: []interface{}{
						"https://",
						Ref{"BOSHEIP"},
						":25555",
					},
				},
			},
		},
	}
}
