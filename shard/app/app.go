package app;

import(
	Service "github.com/PxnPub/PxnGoCommon/service"
);



type AppShard struct {
	Version string
}



func New(version string) Service.App {
	return &AppShard{
		Version: version,
	};
}

func (app *AppShard) Main() {
	service := Service.New();
	service.Start();

print("test works!\n");

	service.WaitUntilEnd();
}
