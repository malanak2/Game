package Graphics

type ObjectManagerT struct {
	objects []Renderable
}

var ObjectManager ObjectManagerT

func InitObjectManager() {
	ObjectManager = ObjectManagerT{objects: make([]Renderable, 0)}
}

func (o *ObjectManagerT) PushObject(object Renderable) {
	o.objects = append(o.objects, object)
}
