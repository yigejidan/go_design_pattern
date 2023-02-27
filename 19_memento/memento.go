package _9_memento

/*
备忘录模式：在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。
备忘录模式（Memento），主要用于捕获一个对象的内部状态，以便在将来的某个时候恢复此状态。
*/

// InputText 用于保存数据
type InputText struct {
	content string
}

// Append 追加数据
func (in *InputText) Append(content string) {
	in.content += content
}

// GetText 获取数据
func (in *InputText) GetText() string {
	return in.content
}

// Snapshot 创建快照
func (in *InputText) Snapshot() *Snapshot {
	return &Snapshot{content: in.content}
}

func (in *InputText) Restore(s *Snapshot) {
	in.content = s.GetText()
}

type Snapshot struct {
	content string
}

func (s *Snapshot) GetText() string {
	return s.content
}
