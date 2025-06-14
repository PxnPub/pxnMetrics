package heartbeat;

import(
	Context  "context"
	UtilsRPC "github.com/PxnPub/PxnGoCommon/rpc"
	GRPC     "google.golang.org/grpc"
	GStatus  "google.golang.org/grpc/status"
	GCodes   "google.golang.org/grpc/codes"
);



const KeyUserPerms = "user-perms";



type UserManager struct {
	AllowIPs map[string]string
	Users    map[string]*User
}

type User struct {
	AllowedShards []uint8
	AllowWebCalls bool
}



func NewUserManager() *UserManager {
	return &UserManager{
		AllowIPs: make(map[string]string),
		Users:    make(map[string]*User),
	};
}



func (man *UserManager) AllowIP(ip string, username string) *UserManager {
	man.AllowIPs[ip] = username;
	return man;
}



func (man *UserManager) GetNewUser(username string) *User {
	user, ok := man.Users[username];
	if !ok {
		user = &User{};
		man.Users[username] = user;
	}
	return user;
}

func (man *UserManager) AddPermWeb(username string) *UserManager {
	man.GetNewUser(username).AllowWebCalls = true;
	return man;
}

func (man *UserManager) AddPermShard(username string, index uint8) *UserManager {
	user := man.GetNewUser(username);
	user.AllowedShards = append(user.AllowedShards, index);
	return man;
}



func (man *UserManager) NewInterceptor() GRPC.UnaryServerInterceptor {
	return func(ctx Context.Context, req any, info *GRPC.UnaryServerInfo,
			handler GRPC.UnaryHandler) (any, error) {
		username, ok := ctx.Value(UtilsRPC.KeyUsername).(string);
		if !ok {
			return nil, GStatus.Errorf(
				GCodes.PermissionDenied,
				"Unable to find username",
			);
		}
		user, ok := man.Users[username];
		if !ok {
			return nil, GStatus.Errorf(
				GCodes.PermissionDenied,
				"Unable to find rpc user info",
			);
		}
		ctx = Context.WithValue(ctx, KeyUserPerms, user);
		return handler(ctx, req);
	};
}
