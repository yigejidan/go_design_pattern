package _0_composite

/*
组合模式：经常用于树形结构，为了简化代码，使用Composite可以把一个叶子节点与一个父节点统一起来处理。
公司的人员组织就是一个典型的树状的结构，现在假设我们现在有部分，和员工，两种角色，一个部门下面可以存在子部门和员工，员工下面不能再包含其他节点。
*/

// IOrganization 组织接口，都实现统计人数的功能
type IOrganization interface {
	Count() int
}

// Employee 员工
type Employee struct {
	Name string
}

// Count 人数统计
func (e *Employee) Count() int {
	return 1
}

type Department struct {
	Name string

	SubOrganizations []IOrganization
}

func (d *Department) Count() int {
	res := 0
	for _, sub := range d.SubOrganizations {
		res += sub.Count()
	}
	return res
}

func (d *Department) AddSub(org IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, org)
}

func NewOrganization() IOrganization {
	root := &Department{Name: "futu"}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Department{
			"sub",
			[]IOrganization{&Employee{}},
		})
	}
	return root
}
