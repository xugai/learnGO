package rpcdemo

import "errors"

//Service.method

type DemoService struct {

}

type Args struct {
	Arg1, Arg2 int
}

func (DemoService) Div(args Args, result *float64) error {
	if args.Arg2 == 0 {
		return errors.New("division by zero")
	}
	*result = float64(args.Arg1) / float64(args.Arg2)
	return nil
}
