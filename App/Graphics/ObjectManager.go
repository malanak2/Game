package Graphics

type ObjectManagerT struct {
	Objects []*Renderable
}

var ObjectManager ObjectManagerT

func InitObjectManager() {
	ObjectManager = ObjectManagerT{Objects: make([]*Renderable, 0)}
}

func (o *ObjectManagerT) PushObject(object *Renderable) {
	o.Objects = append(o.Objects, object)
}
