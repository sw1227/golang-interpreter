package object

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// identifierの文字列とそれに紐づいたオウジェクトを関連づけるmapをwrap
// 関数呼び出し時に、元の環境に新しい束縛を追加しつつ、元の環境は上書きせずに済むようにする（環境の拡張）
type Environment struct {
	store map[string]Object
	outer *Environment // 拡張元の環境
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	// 見つからなければ拡張元の環境を参照
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// 書き込み時にはouterに影響を与えない
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
