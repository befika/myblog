package permission

//operation
var (
	Create = "create"
	Update = "update"
	Read   = "read"
	Delete = "delete"
)

// permission name
var (
	CreateUser = "create user"
	UpdateUser = "update user"
	DeleteUser = "delete user"
	GetUsers   = "get users"
	GetUser    = "get user"

	CreateRole        = "create role"
	UpdateRole        = "update role"
	DeleteRole        = "delete role"
	GetRoles          = "get roles"
	GetRole           = "get role"
	AssignRole        = "assign role"
	DeleteRoleForUser = "delete role from user"
	GetUserRoles      = "get user roles"

	GetAllPermissions    = "get all permissions"
	GetRolePermissions   = "get role permissions"
	AddPermissionToRole  = "assign permission to role"
	DeleteRolePermission = "delete role permission"

	CreateBlog = "create blog"
	UpdateBlog = "update blog"
	DeleteBlog = "delete blog"
	GetBlogs   = "get blogs"
	GetBlog    = "get blog"

	AddSubscription  = "add subscription"
	GetSubscriptions = "get subscriptions"
	GetSubscription  = "get subscription"
)

var PermissionObjects = map[string]string{
	CreateUser: Object(User),
	UpdateUser: Object(User),
	DeleteUser: Object(User),
	GetUsers:   Object(User),
	GetUser:    Object(User),

	CreateRole:        Object(Role),
	UpdateRole:        Object(Role),
	DeleteRole:        Object(Role),
	GetRoles:          Object(Role),
	GetRole:           Object(Role),
	AssignRole:        Object(Role),
	DeleteRoleForUser: Object(Role),
	GetUserRoles:      Object(Role),

	GetAllPermissions:    Object(Permission),
	GetRolePermissions:   Object(Permission),
	DeleteRolePermission: Object(Permission),
	AddPermissionToRole:  Object(Permission),

	CreateBlog: Object(Blog),
	UpdateBlog: Object(Blog),
	DeleteBlog: Object(Blog),
	GetBlogs:   Object(Blog),
	GetBlog:    Object(Blog),

	AddSubscription:  Object(Subscription),
	GetSubscriptions: Object(Subscription),
	GetSubscription:  Object(Subscription),
}

var PermissionActions = map[string]string{
	CreateUser: Action(User, Create),
	UpdateUser: Action(User, Update),
	DeleteUser: Action(User, Delete),
	GetUser:    Action(User, Read),
	GetUsers:   Action(User, Read),

	CreateRole:        Action(Role, Create),
	UpdateRole:        Action(Role, Update),
	DeleteRole:        Action(Role, Delete),
	GetRole:           Action(Role, Read),
	GetRoles:          Action(Role, Read),
	DeleteRoleForUser: Action(Role, Delete),
	AssignRole:        Action(Role, Create),
	GetUserRoles:      Action(Role, Read),

	GetAllPermissions:    Action(Permission, Read),
	GetRolePermissions:   Action(Permission, Read),
	DeleteRolePermission: Action(Permission, Delete),
	AddPermissionToRole:  Action(Permission, Create),

	CreateBlog: Action(Blog, Create),
	UpdateBlog: Action(Blog, Update),
	DeleteBlog: Action(Blog, Delete),
	GetBlog:    Action(Blog, Read),
	GetBlogs:   Action(Blog, Read),

	AddSubscription:  Action(Blog, Create),
	GetSubscriptions: Action(Blog, Read),
	GetSubscription:  Action(Blog, Read),
}
