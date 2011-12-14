package ec2

import (
	"fmt"
	"launchpad.net/goamz/ec2"
	"launchpad.net/juju/go/juju"
)

func init() {
	juju.RegisterProvider("ec2", environProvider{})
}

type environProvider struct{}

var _ juju.EnvironProvider = environProvider{}

type environ struct {
	name   string
	config *providerConfig
	ec2    *ec2.EC2
}

var _ juju.Environ = (*environ)(nil)

type instance struct {
	*ec2.Instance
}

var _ juju.Instance = (*instance)(nil)

func (inst *instance) Id() string {
	return inst.InstanceId
}

func (inst *instance) DNSName() string {
	return inst.Instance.DNSName
}

func (environProvider) Open(name string, config interface{}) (e juju.Environ, err error) {
	cfg := config.(*providerConfig)
	return &environ{
		name:   name,
		config: cfg,
		ec2:    ec2.New(cfg.auth, Regions[cfg.region]),
	}, nil
}

func (e *environ) StartInstance(machineId int) (juju.Instance, error) {
	image, err := FindImageSpec(DefaultImageConstraint)
	if err != nil {
		return nil, fmt.Errorf("cannot find image: %v", err)
	}
	instances, err := e.ec2.RunInstances(&ec2.RunInstances{
		ImageId:      image.ImageId,
		MinCount:     1,
		MaxCount:     1,
		UserData:     nil,
		InstanceType: "m1.small",
	})
	if err != nil {
		return nil, fmt.Errorf("cannot run instances: %v", err)
	}
	if len(instances.Instances) != 1 {
		return nil, fmt.Errorf("expected 1 started instance, got %d", len(instances.Instances))
	}
	return &instance{&instances.Instances[0]}, nil
}

func (e *environ) StopInstances(insts []juju.Instance) error {
	if len(insts) == 0 {
		return nil
	}
	names := make([]string, len(insts))
	for i, inst := range insts {
		names[i] = inst.(*instance).InstanceId
	}
	_, err := e.ec2.TerminateInstances(names)
	return err
}

func (e *environ) Instances() ([]juju.Instance, error) {
	filter := ec2.NewFilter()
	filter.Add("instance-state-name", "pending", "running")

	resp, err := e.ec2.Instances(nil, filter)
	if err != nil {
		return nil, err
	}
	var insts []juju.Instance
	for i := range resp.Reservations {
		r := &resp.Reservations[i]
		for j := range r.Instances {
			insts = append(insts, &instance{&r.Instances[j]})
		}
	}
	return insts, nil
}

func (e *environ) Destroy() error {
	insts, err := e.Instances()
	if err != nil {
		return err
	}
	return e.StopInstances(insts)
}
