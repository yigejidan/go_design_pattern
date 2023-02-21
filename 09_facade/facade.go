package _9_facade

/*
外观模式：为子系统中的一组接口提供一个一致的界面。Facade模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。
基本思想：如果客户端要跟许多子系统打交道，那么客户端需要了解各个子系统的接口，比较麻烦。
		如果有一个统一的“中介”，让客户端只跟中介打交道，中介再去跟各个子系统打交道，对客户端来说就比较简单。所以Facade就相当于搞了一个中介。
*/

// User 用户
type User struct {
	Name string
}

type IUser interface {
	Login(phone, code uint64) (*User, error)
	Register(phone, code uint64) (*User, error)
}

type IUserFacade interface {
	LoginOrRegister(phone, code uint64) (*User, error)
}

type loginService struct {
}

func (l *loginService) Login(phone, code uint64) (*User, error) {
	// 校验操作 ...
	return &User{Name: "test login"}, nil
}

type registerService struct {
}

func (r *registerService) Register(phone, code uint64) (*User, error) {
	// 校验操作 ...
	// 创建用户
	return &User{Name: "test register"}, nil
}

type UserService struct {
	loginService    loginService
	registerService registerService
}

func (u *UserService) LoginOrRegister(phone, code uint64) (*User, error) {
	user, err := u.loginService.Login(phone, code)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	return u.registerService.Register(phone, code)
}
